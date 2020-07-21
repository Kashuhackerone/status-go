package push_notification_server

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"

	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/eth-node/crypto/ecies"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/protocol/common"
	"github.com/status-im/status-go/protocol/protobuf"
	"go.uber.org/zap"
)

const encryptedPayloadKeyLength = 16
const defaultGorushURL = "https://gorush.status.im"

type Config struct {
	// Identity is our identity key
	Identity *ecdsa.PrivateKey
	// GorushUrl is the url for the gorush service
	GorushURL string

	Logger *zap.Logger
}

type Server struct {
	persistence      Persistence
	config           *Config
	messageProcessor *common.MessageProcessor
}

func New(config *Config, persistence Persistence, messageProcessor *common.MessageProcessor) *Server {
	if len(config.GorushURL) == 0 {
		config.GorushURL = defaultGorushURL

	}
	return &Server{persistence: persistence, config: config, messageProcessor: messageProcessor}
}

func (p *Server) generateSharedKey(publicKey *ecdsa.PublicKey) ([]byte, error) {
	return ecies.ImportECDSA(p.config.Identity).GenerateShared(
		ecies.ImportECDSAPublic(publicKey),
		encryptedPayloadKeyLength,
		encryptedPayloadKeyLength,
	)
}

func (p *Server) validateUUID(u string) error {
	if len(u) == 0 {
		return errors.New("empty uuid")
	}
	_, err := uuid.Parse(u)
	return err
}

func (p *Server) decryptRegistration(publicKey *ecdsa.PublicKey, payload []byte) ([]byte, error) {
	sharedKey, err := p.generateSharedKey(publicKey)
	if err != nil {
		return nil, err
	}

	return common.Decrypt(payload, sharedKey)
}

// ValidateRegistration validates a new message against the last one received for a given installationID and and public key
// and return the decrypted message
func (s *Server) ValidateRegistration(publicKey *ecdsa.PublicKey, payload []byte) (*protobuf.PushNotificationRegistration, error) {
	if payload == nil {
		return nil, ErrEmptyPushNotificationRegistrationPayload
	}

	if publicKey == nil {
		return nil, ErrEmptyPushNotificationRegistrationPublicKey
	}

	decryptedPayload, err := s.decryptRegistration(publicKey, payload)
	if err != nil {
		return nil, err
	}

	registration := &protobuf.PushNotificationRegistration{}

	if err := proto.Unmarshal(decryptedPayload, registration); err != nil {
		return nil, ErrCouldNotUnmarshalPushNotificationRegistration
	}

	if registration.Version < 1 {
		return nil, ErrInvalidPushNotificationRegistrationVersion
	}

	if err := s.validateUUID(registration.InstallationId); err != nil {
		return nil, ErrMalformedPushNotificationRegistrationInstallationID
	}

	previousRegistration, err := s.persistence.GetPushNotificationRegistrationByPublicKeyAndInstallationID(common.HashPublicKey(publicKey), registration.InstallationId)
	if err != nil {
		return nil, err
	}

	if previousRegistration != nil && registration.Version <= previousRegistration.Version {
		return nil, ErrInvalidPushNotificationRegistrationVersion
	}

	// Unregistering message
	if registration.Unregister {
		return registration, nil
	}

	if err := s.validateUUID(registration.AccessToken); err != nil {
		return nil, ErrMalformedPushNotificationRegistrationAccessToken
	}

	if len(registration.Grant) == 0 {
		return nil, ErrMalformedPushNotificationRegistrationGrant
	}

	if err := s.verifyGrantSignature(publicKey, registration.AccessToken, registration.Grant); err != nil {

		s.config.Logger.Error("failed to verify grant", zap.Error(err))
		return nil, ErrMalformedPushNotificationRegistrationGrant
	}

	if len(registration.Token) == 0 {
		return nil, ErrMalformedPushNotificationRegistrationDeviceToken
	}

	if registration.TokenType == protobuf.PushNotificationRegistration_UNKNOWN_TOKEN_TYPE {
		return nil, ErrUnknownPushNotificationRegistrationTokenType
	}

	return registration, nil
}

func (s *Server) HandlePushNotificationQuery(query *protobuf.PushNotificationQuery) *protobuf.PushNotificationQueryResponse {

	s.config.Logger.Info("handling push notification query")
	response := &protobuf.PushNotificationQueryResponse{}
	if query == nil || len(query.PublicKeys) == 0 {
		return response
	}

	registrations, err := s.persistence.GetPushNotificationRegistrationByPublicKeys(query.PublicKeys)
	if err != nil {
		s.config.Logger.Error("failed to retrieve registration", zap.Error(err))
		return response
	}

	for _, idAndResponse := range registrations {

		registration := idAndResponse.Registration
		info := &protobuf.PushNotificationQueryInfo{
			PublicKey:      idAndResponse.ID,
			Grant:          registration.Grant,
			Version:        registration.Version,
			InstallationId: registration.InstallationId,
		}

		if registration.AllowFromContactsOnly {
			info.AllowedUserList = registration.AllowedUserList
		} else {
			info.AccessToken = registration.AccessToken
		}
		response.Info = append(response.Info, info)
	}

	response.Success = true
	return response
}

func (s *Server) HandlePushNotificationRequest(request *protobuf.PushNotificationRequest) *protobuf.PushNotificationResponse {
	s.config.Logger.Info("handling pn request")
	response := &protobuf.PushNotificationResponse{}
	// We don't even send a response in this case
	if request == nil || len(request.MessageId) == 0 {
		s.config.Logger.Warn("empty message id")
		return nil
	}

	response.MessageId = request.MessageId

	// TODO: filter by chat id
	// Collect successful requests & registrations
	var requestAndRegistrations []*RequestAndRegistration

	for _, pn := range request.Requests {
		registration, err := s.persistence.GetPushNotificationRegistrationByPublicKeyAndInstallationID(pn.PublicKey, pn.InstallationId)
		report := &protobuf.PushNotificationReport{
			PublicKey:      pn.PublicKey,
			InstallationId: pn.InstallationId,
		}

		if err != nil {
			s.config.Logger.Error("failed to retrieve registration", zap.Error(err))
			report.Error = protobuf.PushNotificationReport_UNKNOWN_ERROR_TYPE
		} else if registration == nil {
			s.config.Logger.Warn("empty registration")
			report.Error = protobuf.PushNotificationReport_NOT_REGISTERED
		} else if registration.AccessToken != pn.AccessToken {
			s.config.Logger.Warn("invalid access token")
			report.Error = protobuf.PushNotificationReport_WRONG_TOKEN
		} else {
			// For now we just assume that the notification will be successful
			requestAndRegistrations = append(requestAndRegistrations, &RequestAndRegistration{
				Request:      pn,
				Registration: registration,
			})
			report.Success = true
		}

		response.Reports = append(response.Reports, report)
	}

	s.config.Logger.Info("built pn request")
	if len(requestAndRegistrations) == 0 {
		s.config.Logger.Warn("no request and registration")
		return response
	}

	// This can be done asynchronously
	goRushRequest := PushNotificationRegistrationToGoRushRequest(requestAndRegistrations)
	//TODO: REMOVE ME
	s.config.Logger.Info("REQUEST", zap.Any("REQUEST", goRushRequest))
	err := sendGoRushNotification(goRushRequest, s.config.GorushURL)
	if err != nil {
		s.config.Logger.Error("failed to send go rush notification", zap.Error(err))
		// TODO: handle this error?
	}

	return response
}

func (s *Server) HandlePushNotificationRegistration(publicKey *ecdsa.PublicKey, payload []byte) *protobuf.PushNotificationRegistrationResponse {

	s.config.Logger.Info("handling push notification registration")
	response := &protobuf.PushNotificationRegistrationResponse{
		RequestId: common.Shake256(payload),
	}

	registration, err := s.ValidateRegistration(publicKey, payload)

	if err != nil {
		if err == ErrInvalidPushNotificationRegistrationVersion {
			response.Error = protobuf.PushNotificationRegistrationResponse_VERSION_MISMATCH
		} else {
			response.Error = protobuf.PushNotificationRegistrationResponse_MALFORMED_MESSAGE
		}
		s.config.Logger.Warn("registration did not validate", zap.Error(err))
		return response
	}

	if registration.Unregister {
		// We save an empty registration, only keeping version and installation-id
		emptyRegistration := &protobuf.PushNotificationRegistration{
			Version:        registration.Version,
			InstallationId: registration.InstallationId,
		}
		if err := s.persistence.SavePushNotificationRegistration(common.HashPublicKey(publicKey), emptyRegistration); err != nil {
			response.Error = protobuf.PushNotificationRegistrationResponse_INTERNAL_ERROR
			s.config.Logger.Error("failed to unregister ", zap.Error(err))
			return response
		}

	} else if err := s.persistence.SavePushNotificationRegistration(common.HashPublicKey(publicKey), registration); err != nil {
		response.Error = protobuf.PushNotificationRegistrationResponse_INTERNAL_ERROR
		s.config.Logger.Error("failed to save registration", zap.Error(err))
		return response
	}

	if err := s.listenToPublicKeyQueryTopic(common.HashPublicKey(publicKey)); err != nil {
		response.Error = protobuf.PushNotificationRegistrationResponse_INTERNAL_ERROR
		s.config.Logger.Error("failed to listen to topic", zap.Error(err))
		return response

	}
	response.Success = true

	s.config.Logger.Info("handled push notification registration successfully")

	return response
}

func (s *Server) Start() error {
	s.config.Logger.Info("starting push notification server")
	if s.config.Identity == nil {
		s.config.Logger.Info("Identity nil")
		// Pull identity from database
		identity, err := s.persistence.GetIdentity()
		if err != nil {
			return err
		}
		if identity == nil {
			identity, err = crypto.GenerateKey()
			if err != nil {
				return err
			}
			if err := s.persistence.SaveIdentity(identity); err != nil {
				return err
			}
		}
		s.config.Identity = identity
	}

	pks, err := s.persistence.GetPushNotificationRegistrationPublicKeys()
	if err != nil {
		return err
	}
	for _, pk := range pks {
		if err := s.listenToPublicKeyQueryTopic(pk); err != nil {
			return err
		}
	}

	s.config.Logger.Info("started push notification server", zap.String("identity", types.EncodeHex(crypto.FromECDSAPub(&s.config.Identity.PublicKey))))

	return nil
}

func (s *Server) listenToPublicKeyQueryTopic(hashedPublicKey []byte) error {
	if s.messageProcessor == nil {
		return nil
	}
	encodedPublicKey := hex.EncodeToString(hashedPublicKey)
	return s.messageProcessor.JoinPublic(encodedPublicKey)
}

func (p *Server) HandlePushNotificationRegistration2(publicKey *ecdsa.PublicKey, payload []byte) error {
	response := p.HandlePushNotificationRegistration(publicKey, payload)
	if response == nil {
		return nil
	}
	encodedMessage, err := proto.Marshal(response)
	if err != nil {
		return err
	}

	rawMessage := &common.RawMessage{
		Payload:     encodedMessage,
		MessageType: protobuf.ApplicationMetadataMessage_PUSH_NOTIFICATION_REGISTRATION_RESPONSE,
	}

	_, err = p.messageProcessor.SendPrivate(context.Background(), publicKey, rawMessage)
	return err
}

func (p *Server) HandlePushNotificationQuery2(publicKey *ecdsa.PublicKey, messageID []byte, query protobuf.PushNotificationQuery) error {
	response := p.HandlePushNotificationQuery(&query)
	if response == nil {
		return nil
	}
	response.MessageId = messageID
	encodedMessage, err := proto.Marshal(response)
	if err != nil {
		return err
	}

	rawMessage := &common.RawMessage{
		Payload:         encodedMessage,
		MessageType:     protobuf.ApplicationMetadataMessage_PUSH_NOTIFICATION_QUERY_RESPONSE,
		SkipNegotiation: true,
	}

	_, err = p.messageProcessor.SendPrivate(context.Background(), publicKey, rawMessage)
	return err

}

func (s *Server) HandlePushNotificationRequest2(publicKey *ecdsa.PublicKey,
	request protobuf.PushNotificationRequest) error {
	s.config.Logger.Info("handling pn request")
	response := s.HandlePushNotificationRequest(&request)
	if response == nil {
		return nil
	}
	encodedMessage, err := proto.Marshal(response)
	if err != nil {
		return err
	}

	rawMessage := &common.RawMessage{
		Payload:         encodedMessage,
		MessageType:     protobuf.ApplicationMetadataMessage_PUSH_NOTIFICATION_RESPONSE,
		SkipNegotiation: true,
	}

	_, err = s.messageProcessor.SendPrivate(context.Background(), publicKey, rawMessage)
	return err
}

// buildGrantSignatureMaterial builds a grant for a specific server.
// We use 3 components:
// 1) The client public key. Not sure this applies to our signature scheme, but best to be conservative. https://crypto.stackexchange.com/questions/15538/given-a-message-and-signature-find-a-public-key-that-makes-the-signature-valid
// 2) The server public key
// 3) The access token
// By verifying this signature, a client can trust the server was instructed to store this access token.

func (s *Server) buildGrantSignatureMaterial(clientPublicKey *ecdsa.PublicKey, serverPublicKey *ecdsa.PublicKey, accessToken string) []byte {
	var signatureMaterial []byte
	signatureMaterial = append(signatureMaterial, crypto.CompressPubkey(clientPublicKey)...)
	signatureMaterial = append(signatureMaterial, crypto.CompressPubkey(serverPublicKey)...)
	signatureMaterial = append(signatureMaterial, []byte(accessToken)...)
	a := crypto.Keccak256(signatureMaterial)
	return a
}

func (s *Server) verifyGrantSignature(clientPublicKey *ecdsa.PublicKey, accessToken string, grant []byte) error {
	signatureMaterial := s.buildGrantSignatureMaterial(clientPublicKey, &s.config.Identity.PublicKey, accessToken)
	recoveredPublicKey, err := crypto.SigToPub(signatureMaterial, grant)
	if err != nil {
		return err
	}

	if !common.IsPubKeyEqual(recoveredPublicKey, clientPublicKey) {
		return errors.New("pubkey mismatch")
	}
	return nil

}
