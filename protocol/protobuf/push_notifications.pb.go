// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push_notifications.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PushNotificationRegistration_TokenType int32

const (
	PushNotificationRegistration_UNKNOWN_TOKEN_TYPE PushNotificationRegistration_TokenType = 0
	PushNotificationRegistration_APN_TOKEN          PushNotificationRegistration_TokenType = 1
	PushNotificationRegistration_FIREBASE_TOKEN     PushNotificationRegistration_TokenType = 2
)

var PushNotificationRegistration_TokenType_name = map[int32]string{
	0: "UNKNOWN_TOKEN_TYPE",
	1: "APN_TOKEN",
	2: "FIREBASE_TOKEN",
}

var PushNotificationRegistration_TokenType_value = map[string]int32{
	"UNKNOWN_TOKEN_TYPE": 0,
	"APN_TOKEN":          1,
	"FIREBASE_TOKEN":     2,
}

func (x PushNotificationRegistration_TokenType) String() string {
	return proto.EnumName(PushNotificationRegistration_TokenType_name, int32(x))
}

func (PushNotificationRegistration_TokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{0, 0}
}

type PushNotificationRegistrationResponse_ErrorType int32

const (
	PushNotificationRegistrationResponse_UNKNOWN_ERROR_TYPE     PushNotificationRegistrationResponse_ErrorType = 0
	PushNotificationRegistrationResponse_MALFORMED_MESSAGE      PushNotificationRegistrationResponse_ErrorType = 1
	PushNotificationRegistrationResponse_VERSION_MISMATCH       PushNotificationRegistrationResponse_ErrorType = 2
	PushNotificationRegistrationResponse_UNSUPPORTED_TOKEN_TYPE PushNotificationRegistrationResponse_ErrorType = 3
	PushNotificationRegistrationResponse_INTERNAL_ERROR         PushNotificationRegistrationResponse_ErrorType = 4
)

var PushNotificationRegistrationResponse_ErrorType_name = map[int32]string{
	0: "UNKNOWN_ERROR_TYPE",
	1: "MALFORMED_MESSAGE",
	2: "VERSION_MISMATCH",
	3: "UNSUPPORTED_TOKEN_TYPE",
	4: "INTERNAL_ERROR",
}

var PushNotificationRegistrationResponse_ErrorType_value = map[string]int32{
	"UNKNOWN_ERROR_TYPE":     0,
	"MALFORMED_MESSAGE":      1,
	"VERSION_MISMATCH":       2,
	"UNSUPPORTED_TOKEN_TYPE": 3,
	"INTERNAL_ERROR":         4,
}

func (x PushNotificationRegistrationResponse_ErrorType) String() string {
	return proto.EnumName(PushNotificationRegistrationResponse_ErrorType_name, int32(x))
}

func (PushNotificationRegistrationResponse_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{1, 0}
}

type PushNotificationRegistration struct {
	TokenType            PushNotificationRegistration_TokenType `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3,enum=protobuf.PushNotificationRegistration_TokenType" json:"token_type,omitempty"`
	Token                string                                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	InstallationId       string                                 `protobuf:"bytes,3,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	AccessToken          string                                 `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Enabled              bool                                   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Version              uint64                                 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	AllowedUserList      [][]byte                               `protobuf:"bytes,7,rep,name=allowed_user_list,json=allowedUserList,proto3" json:"allowed_user_list,omitempty"`
	BlockedChatList      [][]byte                               `protobuf:"bytes,8,rep,name=blocked_chat_list,json=blockedChatList,proto3" json:"blocked_chat_list,omitempty"`
	Unregister           bool                                   `protobuf:"varint,9,opt,name=unregister,proto3" json:"unregister,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *PushNotificationRegistration) Reset()         { *m = PushNotificationRegistration{} }
func (m *PushNotificationRegistration) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRegistration) ProtoMessage()    {}
func (*PushNotificationRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{0}
}

func (m *PushNotificationRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRegistration.Unmarshal(m, b)
}
func (m *PushNotificationRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRegistration.Marshal(b, m, deterministic)
}
func (m *PushNotificationRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRegistration.Merge(m, src)
}
func (m *PushNotificationRegistration) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRegistration.Size(m)
}
func (m *PushNotificationRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRegistration proto.InternalMessageInfo

func (m *PushNotificationRegistration) GetTokenType() PushNotificationRegistration_TokenType {
	if m != nil {
		return m.TokenType
	}
	return PushNotificationRegistration_UNKNOWN_TOKEN_TYPE
}

func (m *PushNotificationRegistration) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PushNotificationRegistration) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PushNotificationRegistration) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotificationRegistration) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PushNotificationRegistration) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PushNotificationRegistration) GetAllowedUserList() [][]byte {
	if m != nil {
		return m.AllowedUserList
	}
	return nil
}

func (m *PushNotificationRegistration) GetBlockedChatList() [][]byte {
	if m != nil {
		return m.BlockedChatList
	}
	return nil
}

func (m *PushNotificationRegistration) GetUnregister() bool {
	if m != nil {
		return m.Unregister
	}
	return false
}

type PushNotificationRegistrationResponse struct {
	Success              bool                                           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                PushNotificationRegistrationResponse_ErrorType `protobuf:"varint,2,opt,name=error,proto3,enum=protobuf.PushNotificationRegistrationResponse_ErrorType" json:"error,omitempty"`
	RequestId            []byte                                         `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *PushNotificationRegistrationResponse) Reset()         { *m = PushNotificationRegistrationResponse{} }
func (m *PushNotificationRegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRegistrationResponse) ProtoMessage()    {}
func (*PushNotificationRegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{1}
}

func (m *PushNotificationRegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Unmarshal(m, b)
}
func (m *PushNotificationRegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Marshal(b, m, deterministic)
}
func (m *PushNotificationRegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRegistrationResponse.Merge(m, src)
}
func (m *PushNotificationRegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Size(m)
}
func (m *PushNotificationRegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRegistrationResponse proto.InternalMessageInfo

func (m *PushNotificationRegistrationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PushNotificationRegistrationResponse) GetError() PushNotificationRegistrationResponse_ErrorType {
	if m != nil {
		return m.Error
	}
	return PushNotificationRegistrationResponse_UNKNOWN_ERROR_TYPE
}

func (m *PushNotificationRegistrationResponse) GetRequestId() []byte {
	if m != nil {
		return m.RequestId
	}
	return nil
}

type PushNotificationAdvertisementInfo struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	AccessToken          string   `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	InstallationId       string   `protobuf:"bytes,3,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationAdvertisementInfo) Reset()         { *m = PushNotificationAdvertisementInfo{} }
func (m *PushNotificationAdvertisementInfo) String() string { return proto.CompactTextString(m) }
func (*PushNotificationAdvertisementInfo) ProtoMessage()    {}
func (*PushNotificationAdvertisementInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{2}
}

func (m *PushNotificationAdvertisementInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationAdvertisementInfo.Unmarshal(m, b)
}
func (m *PushNotificationAdvertisementInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationAdvertisementInfo.Marshal(b, m, deterministic)
}
func (m *PushNotificationAdvertisementInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationAdvertisementInfo.Merge(m, src)
}
func (m *PushNotificationAdvertisementInfo) XXX_Size() int {
	return xxx_messageInfo_PushNotificationAdvertisementInfo.Size(m)
}
func (m *PushNotificationAdvertisementInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationAdvertisementInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationAdvertisementInfo proto.InternalMessageInfo

func (m *PushNotificationAdvertisementInfo) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotificationAdvertisementInfo) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotificationAdvertisementInfo) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

type ContactCodeAdvertisement struct {
	PushNotificationInfo []*PushNotificationAdvertisementInfo `protobuf:"bytes,1,rep,name=push_notification_info,json=pushNotificationInfo,proto3" json:"push_notification_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ContactCodeAdvertisement) Reset()         { *m = ContactCodeAdvertisement{} }
func (m *ContactCodeAdvertisement) String() string { return proto.CompactTextString(m) }
func (*ContactCodeAdvertisement) ProtoMessage()    {}
func (*ContactCodeAdvertisement) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{3}
}

func (m *ContactCodeAdvertisement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactCodeAdvertisement.Unmarshal(m, b)
}
func (m *ContactCodeAdvertisement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactCodeAdvertisement.Marshal(b, m, deterministic)
}
func (m *ContactCodeAdvertisement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactCodeAdvertisement.Merge(m, src)
}
func (m *ContactCodeAdvertisement) XXX_Size() int {
	return xxx_messageInfo_ContactCodeAdvertisement.Size(m)
}
func (m *ContactCodeAdvertisement) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactCodeAdvertisement.DiscardUnknown(m)
}

var xxx_messageInfo_ContactCodeAdvertisement proto.InternalMessageInfo

func (m *ContactCodeAdvertisement) GetPushNotificationInfo() []*PushNotificationAdvertisementInfo {
	if m != nil {
		return m.PushNotificationInfo
	}
	return nil
}

type PushNotificationQuery struct {
	PublicKeys           [][]byte `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationQuery) Reset()         { *m = PushNotificationQuery{} }
func (m *PushNotificationQuery) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQuery) ProtoMessage()    {}
func (*PushNotificationQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{4}
}

func (m *PushNotificationQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQuery.Unmarshal(m, b)
}
func (m *PushNotificationQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQuery.Marshal(b, m, deterministic)
}
func (m *PushNotificationQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQuery.Merge(m, src)
}
func (m *PushNotificationQuery) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQuery.Size(m)
}
func (m *PushNotificationQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQuery proto.InternalMessageInfo

func (m *PushNotificationQuery) GetPublicKeys() [][]byte {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type PushNotificationQueryInfo struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	InstallationId       string   `protobuf:"bytes,2,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	AllowedUserList      [][]byte `protobuf:"bytes,4,rep,name=allowed_user_list,json=allowedUserList,proto3" json:"allowed_user_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationQueryInfo) Reset()         { *m = PushNotificationQueryInfo{} }
func (m *PushNotificationQueryInfo) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQueryInfo) ProtoMessage()    {}
func (*PushNotificationQueryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{5}
}

func (m *PushNotificationQueryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQueryInfo.Unmarshal(m, b)
}
func (m *PushNotificationQueryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQueryInfo.Marshal(b, m, deterministic)
}
func (m *PushNotificationQueryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQueryInfo.Merge(m, src)
}
func (m *PushNotificationQueryInfo) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQueryInfo.Size(m)
}
func (m *PushNotificationQueryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQueryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQueryInfo proto.InternalMessageInfo

func (m *PushNotificationQueryInfo) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotificationQueryInfo) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PushNotificationQueryInfo) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotificationQueryInfo) GetAllowedUserList() [][]byte {
	if m != nil {
		return m.AllowedUserList
	}
	return nil
}

type PushNotificationQueryResponse struct {
	Info                 []*PushNotificationQueryInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	Success              bool                         `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PushNotificationQueryResponse) Reset()         { *m = PushNotificationQueryResponse{} }
func (m *PushNotificationQueryResponse) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQueryResponse) ProtoMessage()    {}
func (*PushNotificationQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{6}
}

func (m *PushNotificationQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQueryResponse.Unmarshal(m, b)
}
func (m *PushNotificationQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQueryResponse.Marshal(b, m, deterministic)
}
func (m *PushNotificationQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQueryResponse.Merge(m, src)
}
func (m *PushNotificationQueryResponse) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQueryResponse.Size(m)
}
func (m *PushNotificationQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQueryResponse proto.InternalMessageInfo

func (m *PushNotificationQueryResponse) GetInfo() []*PushNotificationQueryInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PushNotificationQueryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type PushNotification struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ChatId               string   `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotification) Reset()         { *m = PushNotification{} }
func (m *PushNotification) String() string { return proto.CompactTextString(m) }
func (*PushNotification) ProtoMessage()    {}
func (*PushNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{7}
}

func (m *PushNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotification.Unmarshal(m, b)
}
func (m *PushNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotification.Marshal(b, m, deterministic)
}
func (m *PushNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotification.Merge(m, src)
}
func (m *PushNotification) XXX_Size() int {
	return xxx_messageInfo_PushNotification.Size(m)
}
func (m *PushNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotification.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotification proto.InternalMessageInfo

func (m *PushNotification) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotification) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

type PushNotificationRequest struct {
	Requests             []*PushNotification `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	Message              []byte              `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MessageId            string              `protobuf:"bytes,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	AckRequired          string              `protobuf:"bytes,4,opt,name=ack_required,json=ackRequired,proto3" json:"ack_required,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PushNotificationRequest) Reset()         { *m = PushNotificationRequest{} }
func (m *PushNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRequest) ProtoMessage()    {}
func (*PushNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{8}
}

func (m *PushNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRequest.Unmarshal(m, b)
}
func (m *PushNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRequest.Marshal(b, m, deterministic)
}
func (m *PushNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRequest.Merge(m, src)
}
func (m *PushNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRequest.Size(m)
}
func (m *PushNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRequest proto.InternalMessageInfo

func (m *PushNotificationRequest) GetRequests() []*PushNotification {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *PushNotificationRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PushNotificationRequest) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *PushNotificationRequest) GetAckRequired() string {
	if m != nil {
		return m.AckRequired
	}
	return ""
}

type PushNotificationAcknowledgement struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationAcknowledgement) Reset()         { *m = PushNotificationAcknowledgement{} }
func (m *PushNotificationAcknowledgement) String() string { return proto.CompactTextString(m) }
func (*PushNotificationAcknowledgement) ProtoMessage()    {}
func (*PushNotificationAcknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{9}
}

func (m *PushNotificationAcknowledgement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationAcknowledgement.Unmarshal(m, b)
}
func (m *PushNotificationAcknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationAcknowledgement.Marshal(b, m, deterministic)
}
func (m *PushNotificationAcknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationAcknowledgement.Merge(m, src)
}
func (m *PushNotificationAcknowledgement) XXX_Size() int {
	return xxx_messageInfo_PushNotificationAcknowledgement.Size(m)
}
func (m *PushNotificationAcknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationAcknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationAcknowledgement proto.InternalMessageInfo

func (m *PushNotificationAcknowledgement) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("protobuf.PushNotificationRegistration_TokenType", PushNotificationRegistration_TokenType_name, PushNotificationRegistration_TokenType_value)
	proto.RegisterEnum("protobuf.PushNotificationRegistrationResponse_ErrorType", PushNotificationRegistrationResponse_ErrorType_name, PushNotificationRegistrationResponse_ErrorType_value)
	proto.RegisterType((*PushNotificationRegistration)(nil), "protobuf.PushNotificationRegistration")
	proto.RegisterType((*PushNotificationRegistrationResponse)(nil), "protobuf.PushNotificationRegistrationResponse")
	proto.RegisterType((*PushNotificationAdvertisementInfo)(nil), "protobuf.PushNotificationAdvertisementInfo")
	proto.RegisterType((*ContactCodeAdvertisement)(nil), "protobuf.ContactCodeAdvertisement")
	proto.RegisterType((*PushNotificationQuery)(nil), "protobuf.PushNotificationQuery")
	proto.RegisterType((*PushNotificationQueryInfo)(nil), "protobuf.PushNotificationQueryInfo")
	proto.RegisterType((*PushNotificationQueryResponse)(nil), "protobuf.PushNotificationQueryResponse")
	proto.RegisterType((*PushNotification)(nil), "protobuf.PushNotification")
	proto.RegisterType((*PushNotificationRequest)(nil), "protobuf.PushNotificationRequest")
	proto.RegisterType((*PushNotificationAcknowledgement)(nil), "protobuf.PushNotificationAcknowledgement")
}

func init() { proto.RegisterFile("push_notifications.proto", fileDescriptor_200acd86044eaa5d) }

var fileDescriptor_200acd86044eaa5d = []byte{
	// 766 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xda, 0x48,
	0x14, 0x5e, 0x1b, 0x92, 0xe0, 0x13, 0x96, 0x90, 0x51, 0x7e, 0xbc, 0xd1, 0x66, 0x43, 0xbc, 0x95,
	0x8a, 0x5a, 0x09, 0xb5, 0xa9, 0xd4, 0xe6, 0x96, 0x12, 0xa7, 0xb5, 0x12, 0x0c, 0x1d, 0xa0, 0x55,
	0xaf, 0x2c, 0x63, 0x0f, 0xc1, 0xc2, 0xb1, 0xe9, 0xcc, 0x38, 0x11, 0x17, 0x95, 0xfa, 0x04, 0x7d,
	0x93, 0xde, 0xe4, 0x1d, 0xfa, 0x5e, 0x95, 0xc7, 0x36, 0x21, 0x40, 0x22, 0xae, 0xec, 0xf3, 0xcd,
	0xf9, 0x66, 0xe6, 0x7c, 0xe7, 0x7c, 0x03, 0xea, 0x38, 0x62, 0x43, 0x2b, 0x08, 0xb9, 0x37, 0xf0,
	0x1c, 0x9b, 0x7b, 0x61, 0xc0, 0x6a, 0x63, 0x1a, 0xf2, 0x10, 0x15, 0xc4, 0xa7, 0x1f, 0x0d, 0xb4,
	0xdf, 0x39, 0xf8, 0xb7, 0x1d, 0xb1, 0xa1, 0x39, 0x93, 0x85, 0xc9, 0x95, 0xc7, 0x38, 0x15, 0xff,
	0xa8, 0x05, 0xc0, 0xc3, 0x11, 0x09, 0x2c, 0x3e, 0x19, 0x13, 0x55, 0xaa, 0x48, 0xd5, 0xd2, 0xc9,
	0xab, 0x5a, 0xc6, 0xaf, 0x3d, 0xc5, 0xad, 0x75, 0x63, 0x62, 0x77, 0x32, 0x26, 0x58, 0xe1, 0xd9,
	0x2f, 0xda, 0x81, 0x35, 0x11, 0xa8, 0x72, 0x45, 0xaa, 0x2a, 0x38, 0x09, 0xd0, 0x73, 0xd8, 0xf2,
	0x02, 0xc6, 0x6d, 0xdf, 0x17, 0x54, 0xcb, 0x73, 0xd5, 0x9c, 0x58, 0x2f, 0xcd, 0xc2, 0x86, 0x8b,
	0x8e, 0xa1, 0x68, 0x3b, 0x0e, 0x61, 0xcc, 0x4a, 0x76, 0xc9, 0x8b, 0xac, 0xcd, 0x04, 0x13, 0x07,
	0x22, 0x15, 0x36, 0x48, 0x60, 0xf7, 0x7d, 0xe2, 0xaa, 0x6b, 0x15, 0xa9, 0x5a, 0xc0, 0x59, 0x18,
	0xaf, 0xdc, 0x10, 0xca, 0xbc, 0x30, 0x50, 0xd7, 0x2b, 0x52, 0x35, 0x8f, 0xb3, 0x10, 0xbd, 0x80,
	0x6d, 0xdb, 0xf7, 0xc3, 0x5b, 0xe2, 0x5a, 0x11, 0x23, 0xd4, 0xf2, 0x3d, 0xc6, 0xd5, 0x8d, 0x4a,
	0xae, 0x5a, 0xc4, 0x5b, 0xe9, 0x42, 0x8f, 0x11, 0x7a, 0xe9, 0x31, 0x1e, 0xe7, 0xf6, 0xfd, 0xd0,
	0x19, 0x11, 0xd7, 0x72, 0x86, 0x36, 0x4f, 0x72, 0x0b, 0x49, 0x6e, 0xba, 0xd0, 0x18, 0xda, 0x5c,
	0xe4, 0xfe, 0x07, 0x10, 0x05, 0x54, 0x88, 0x42, 0xa8, 0xaa, 0x88, 0xeb, 0xcc, 0x20, 0xda, 0x39,
	0x28, 0x53, 0x95, 0xd0, 0x1e, 0xa0, 0x9e, 0x79, 0x61, 0xb6, 0xbe, 0x98, 0x56, 0xb7, 0x75, 0xa1,
	0x9b, 0x56, 0xf7, 0x6b, 0x5b, 0x2f, 0xff, 0x85, 0xfe, 0x06, 0xa5, 0xde, 0x4e, 0xb1, 0xb2, 0x84,
	0x10, 0x94, 0xce, 0x0d, 0xac, 0xbf, 0xaf, 0x77, 0xf4, 0x14, 0x93, 0xb5, 0x3b, 0x19, 0x9e, 0x3d,
	0xd5, 0x0b, 0x4c, 0xd8, 0x38, 0x0c, 0x18, 0x89, 0x25, 0x60, 0x91, 0x10, 0x4b, 0x34, 0xb3, 0x80,
	0xb3, 0x10, 0x99, 0xb0, 0x46, 0x28, 0x0d, 0xa9, 0x68, 0x4c, 0xe9, 0xe4, 0x74, 0xb5, 0x26, 0x67,
	0x1b, 0xd7, 0xf4, 0x98, 0x2b, 0x9a, 0x9d, 0x6c, 0x83, 0x0e, 0x01, 0x28, 0xf9, 0x16, 0x11, 0xc6,
	0xb3, 0x6e, 0x16, 0xb1, 0x92, 0x22, 0x86, 0xab, 0xfd, 0x90, 0x40, 0x99, 0x72, 0x66, 0x4b, 0xd7,
	0x31, 0x6e, 0xe1, 0xac, 0xf4, 0x5d, 0xd8, 0x6e, 0xd6, 0x2f, 0xcf, 0x5b, 0xb8, 0xa9, 0x9f, 0x59,
	0x4d, 0xbd, 0xd3, 0xa9, 0x7f, 0xd0, 0xcb, 0x12, 0xda, 0x81, 0xf2, 0x67, 0x1d, 0x77, 0x8c, 0x96,
	0x69, 0x35, 0x8d, 0x4e, 0xb3, 0xde, 0x6d, 0x7c, 0x2c, 0xcb, 0xe8, 0x00, 0xf6, 0x7a, 0x66, 0xa7,
	0xd7, 0x6e, 0xb7, 0x70, 0x57, 0x3f, 0x9b, 0xd5, 0x30, 0x17, 0x8b, 0x66, 0x98, 0x5d, 0x1d, 0x9b,
	0xf5, 0xcb, 0xe4, 0x84, 0x72, 0x5e, 0xfb, 0x29, 0xc1, 0xf1, 0x7c, 0x6d, 0x75, 0xf7, 0x86, 0x50,
	0xee, 0x31, 0x72, 0x4d, 0x02, 0x6e, 0x04, 0x83, 0x30, 0xae, 0x63, 0x1c, 0xf5, 0x7d, 0xcf, 0xb1,
	0x46, 0x64, 0x22, 0x44, 0x2b, 0x62, 0x25, 0x41, 0x2e, 0xc8, 0x64, 0x61, 0x20, 0xe5, 0xc5, 0x81,
	0x5c, 0x75, 0xb8, 0xb5, 0xef, 0xa0, 0x36, 0xc2, 0x80, 0xdb, 0x0e, 0x6f, 0x84, 0x2e, 0x79, 0x70,
	0x15, 0x64, 0xc3, 0xde, 0x82, 0x9f, 0x2d, 0x2f, 0x18, 0x84, 0xaa, 0x54, 0xc9, 0x55, 0x37, 0x4f,
	0x5e, 0x3e, 0xde, 0xaf, 0x85, 0x9a, 0xf0, 0xce, 0x78, 0x2e, 0x25, 0x46, 0xb5, 0x53, 0xd8, 0x9d,
	0xa7, 0x7e, 0x8a, 0x08, 0x9d, 0xa0, 0x23, 0xd8, 0xbc, 0x97, 0x80, 0x89, 0x03, 0x8b, 0x18, 0xa6,
	0x1a, 0x30, 0xed, 0x4e, 0x82, 0x7f, 0x96, 0x52, 0x85, 0x82, 0xf3, 0x12, 0x49, 0x2b, 0x49, 0x24,
	0x2f, 0xf5, 0xff, 0xc3, 0x6e, 0xe4, 0xe6, 0xbb, 0xb1, 0xd4, 0xc7, 0xf9, 0xa5, 0x3e, 0xd6, 0x28,
	0x1c, 0x2e, 0xbd, 0xf3, 0xd4, 0x2b, 0xef, 0x20, 0x3f, 0x23, 0xf0, 0xff, 0x8f, 0x0b, 0x3c, 0x2d,
	0x15, 0x0b, 0xc2, 0xac, 0xc9, 0xe4, 0x07, 0x26, 0xd3, 0x4c, 0x28, 0xcf, 0x93, 0x57, 0x91, 0x67,
	0x1f, 0x36, 0xc4, 0x53, 0x33, 0x95, 0x65, 0x3d, 0x0e, 0x0d, 0x57, 0xfb, 0x25, 0xc1, 0xfe, 0xa2,
	0x3d, 0x85, 0xc7, 0xd0, 0x5b, 0x28, 0xa4, 0x76, 0x63, 0x69, 0x09, 0x07, 0x4f, 0x78, 0x7a, 0x9a,
	0x1b, 0xdf, 0xfe, 0x9a, 0x30, 0x66, 0x5f, 0x11, 0x71, 0x58, 0x11, 0x67, 0x61, 0x2c, 0x7e, 0xfa,
	0x7b, 0x3f, 0xc3, 0x4a, 0x8a, 0x64, 0x6f, 0xf3, 0xc8, 0x8a, 0x37, 0xf2, 0x28, 0x71, 0xef, 0xdf,
	0xe6, 0x11, 0x4e, 0x21, 0xed, 0x35, 0x1c, 0x2d, 0x4c, 0xa7, 0x33, 0x0a, 0xc2, 0x5b, 0x9f, 0xb8,
	0x57, 0xc9, 0xa0, 0x97, 0x40, 0xf6, 0xdc, 0x54, 0x04, 0xd9, 0x73, 0xfb, 0xeb, 0xe2, 0xce, 0x6f,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x16, 0x49, 0x40, 0xfd, 0xcf, 0x06, 0x00, 0x00,
}