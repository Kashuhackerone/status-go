package push_notification_client

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/sha3"

	"github.com/google/uuid"

	"github.com/status-im/status-go/eth-node/crypto/ecies"
	"github.com/status-im/status-go/protocol/encryption"
	"github.com/status-im/status-go/protocol/protobuf"
)

const accessTokenKeyLength = 16

type PushNotificationServer struct {
	key        *ecdsa.PublicKey
	registered bool
}

type Config struct {
	// Identity is our identity key
	Identity *ecdsa.PrivateKey
	// SendEnabled indicates whether we should be sending push notifications
	SendEnabled bool
	// RemoteNotificationsEnabled is whether we should register with a remote server for push notifications
	RemoteNotificationsEnabled bool

	// AllowOnlyFromContacts indicates whether we should be receiving push notifications
	// only from contacts
	AllowOnlyFromContacts bool
	// ContactIDs is the public keys for each contact that we allow notifications from
	ContactIDs []*ecdsa.PublicKey
	// MutedChatIDs is the IDs of the chats we don't want to receive notifications from
	MutedChatIDs []string
	// PushNotificationServers is an array of push notification servers we want to register with
	PushNotificationServers []*PushNotificationServer
	// InstallationID is the installation-id for this device
	InstallationID string
}

type Client struct {
	persistence *Persistence
	config      *Config

	// lastPushNotificationVersion is the latest known push notification version
	lastPushNotificationVersion uint64

	// AccessToken is the access token that is currently being used
	AccessToken string
	// DeviceToken is the device token for this device
	DeviceToken string

	// randomReader only used for testing so we have deterministic encryption
	reader io.Reader
}

func New(persistence *Persistence) *Client {
	return &Client{persistence: persistence, reader: rand.Reader}
}

// This likely will return a channel as it's an asynchrous operation
func fetchNotificationInfoFor(publicKey *ecdsa.PublicKey) error {
	return nil
}

// Sends an actual push notification, where do we get the chatID?
func sendPushNotificationTo(publicKey *ecdsa.PublicKey, chatID string) error {
	return nil
}

// This should schedule:
// 1) Check we have reasonably fresh push notifications info
// 2) Otherwise it should fetch them
// 3) Send a push notification to the devices in question
func (p *Client) HandleMessageSent(publicKey *ecdsa.PublicKey, spec *encryption.ProtocolMessageSpec, messageIDs [][]byte) error {
	return nil
}

func (p *Client) NotifyOnMessageID(messageID []byte) error {
	return nil
}

func (p *Client) mutedChatIDsHashes() [][]byte {
	var mutedChatListHashes [][]byte

	for _, chatID := range p.config.MutedChatIDs {
		mutedChatListHashes = append(mutedChatListHashes, shake256(chatID))
	}

	return mutedChatListHashes
}

func (p *Client) encryptToken(publicKey *ecdsa.PublicKey, token []byte) ([]byte, error) {
	sharedKey, err := ecies.ImportECDSA(p.config.Identity).GenerateShared(
		ecies.ImportECDSAPublic(publicKey),
		accessTokenKeyLength,
		accessTokenKeyLength,
	)
	if err != nil {
		return nil, err
	}
	encryptedToken, err := encryptAccessToken(token, sharedKey, p.reader)
	if err != nil {
		return nil, err
	}
	return encryptedToken, nil
}

func (p *Client) allowedUserList(token []byte) ([][]byte, error) {
	var encryptedTokens [][]byte
	for _, publicKey := range p.config.ContactIDs {
		encryptedToken, err := p.encryptToken(publicKey, token)
		if err != nil {
			return nil, err
		}

		encryptedTokens = append(encryptedTokens, encryptedToken)

	}
	return encryptedTokens, nil
}

func (p *Client) buildPushNotificationRegistrationMessage() (*protobuf.PushNotificationRegistration, error) {
	token := uuid.New().String()
	allowedUserList, err := p.allowedUserList([]byte(token))
	if err != nil {
		return nil, err
	}

	options := &protobuf.PushNotificationRegistration{
		AccessToken:     token,
		Version:         p.lastPushNotificationVersion + 1,
		InstallationId:  p.config.InstallationID,
		Token:           p.DeviceToken,
		Enabled:         p.config.RemoteNotificationsEnabled,
		BlockedChatList: p.mutedChatIDsHashes(),
		AllowedUserList: allowedUserList,
	}
	return options, nil
}

func (p *Client) Register(deviceToken string) error {
	return nil
}

// HandlePushNotificationRegistrationResponse should check whether the response was successful or not, retry if necessary otherwise store the result in the database
func (p *Client) HandlePushNotificationRegistrationResponse(response *protobuf.PushNotificationRegistrationResponse) error {
	return nil
}

// HandlePushNotificationAdvertisement should store any info related to push notifications
func (p *Client) HandlePushNotificationAdvertisement(info *protobuf.PushNotificationAdvertisementInfo) error {
	return nil
}

// HandlePushNotificationQueryResponse should update the data in the database for a given user
func (p *Client) HandlePushNotificationQueryResponse(response *protobuf.PushNotificationQueryResponse) error {
	return nil
}

// HandlePushNotificationAcknowledgement should set the request as processed
func (p *Client) HandlePushNotificationAcknowledgement(ack *protobuf.PushNotificationAcknowledgement) error {
	return nil
}

func (p *Client) SetContactIDs(contactIDs []*ecdsa.PublicKey) error {
	p.config.ContactIDs = contactIDs
	// Update or schedule update
	return nil
}

func (p *Client) SetMutedChatIDs(chatIDs []string) error {
	p.config.MutedChatIDs = chatIDs
	// Update or schedule update
	return nil
}

func encryptAccessToken(plaintext []byte, key []byte, reader io.Reader) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func shake256(input string) []byte {
	buf := []byte(input)
	h := make([]byte, 64)
	sha3.ShakeSum256(h, buf)
	return h
}