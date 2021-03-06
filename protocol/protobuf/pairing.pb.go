// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pairing.proto

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

type PairInstallation struct {
	Clock                uint64   `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	InstallationId       string   `protobuf:"bytes,2,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	DeviceType           string   `protobuf:"bytes,3,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairInstallation) Reset()         { *m = PairInstallation{} }
func (m *PairInstallation) String() string { return proto.CompactTextString(m) }
func (*PairInstallation) ProtoMessage()    {}
func (*PairInstallation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61ab7221f0b5518, []int{0}
}

func (m *PairInstallation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairInstallation.Unmarshal(m, b)
}
func (m *PairInstallation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairInstallation.Marshal(b, m, deterministic)
}
func (m *PairInstallation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairInstallation.Merge(m, src)
}
func (m *PairInstallation) XXX_Size() int {
	return xxx_messageInfo_PairInstallation.Size(m)
}
func (m *PairInstallation) XXX_DiscardUnknown() {
	xxx_messageInfo_PairInstallation.DiscardUnknown(m)
}

var xxx_messageInfo_PairInstallation proto.InternalMessageInfo

func (m *PairInstallation) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *PairInstallation) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PairInstallation) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *PairInstallation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SyncInstallationContact struct {
	Clock                uint64   `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ProfileImage         string   `protobuf:"bytes,3,opt,name=profile_image,json=profileImage,proto3" json:"profile_image,omitempty"`
	EnsName              string   `protobuf:"bytes,4,opt,name=ens_name,json=ensName,proto3" json:"ens_name,omitempty"`
	LastUpdated          uint64   `protobuf:"varint,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	SystemTags           []string `protobuf:"bytes,6,rep,name=system_tags,json=systemTags,proto3" json:"system_tags,omitempty"`
	LocalNickname        string   `protobuf:"bytes,7,opt,name=local_nickname,json=localNickname,proto3" json:"local_nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncInstallationContact) Reset()         { *m = SyncInstallationContact{} }
func (m *SyncInstallationContact) String() string { return proto.CompactTextString(m) }
func (*SyncInstallationContact) ProtoMessage()    {}
func (*SyncInstallationContact) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61ab7221f0b5518, []int{1}
}

func (m *SyncInstallationContact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncInstallationContact.Unmarshal(m, b)
}
func (m *SyncInstallationContact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncInstallationContact.Marshal(b, m, deterministic)
}
func (m *SyncInstallationContact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncInstallationContact.Merge(m, src)
}
func (m *SyncInstallationContact) XXX_Size() int {
	return xxx_messageInfo_SyncInstallationContact.Size(m)
}
func (m *SyncInstallationContact) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncInstallationContact.DiscardUnknown(m)
}

var xxx_messageInfo_SyncInstallationContact proto.InternalMessageInfo

func (m *SyncInstallationContact) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *SyncInstallationContact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SyncInstallationContact) GetProfileImage() string {
	if m != nil {
		return m.ProfileImage
	}
	return ""
}

func (m *SyncInstallationContact) GetEnsName() string {
	if m != nil {
		return m.EnsName
	}
	return ""
}

func (m *SyncInstallationContact) GetLastUpdated() uint64 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

func (m *SyncInstallationContact) GetSystemTags() []string {
	if m != nil {
		return m.SystemTags
	}
	return nil
}

func (m *SyncInstallationContact) GetLocalNickname() string {
	if m != nil {
		return m.LocalNickname
	}
	return ""
}

type SyncInstallationAccount struct {
	Clock                uint64   `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	ProfileImage         string   `protobuf:"bytes,2,opt,name=profile_image,json=profileImage,proto3" json:"profile_image,omitempty"`
	LastUpdated          uint64   `protobuf:"varint,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncInstallationAccount) Reset()         { *m = SyncInstallationAccount{} }
func (m *SyncInstallationAccount) String() string { return proto.CompactTextString(m) }
func (*SyncInstallationAccount) ProtoMessage()    {}
func (*SyncInstallationAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61ab7221f0b5518, []int{2}
}

func (m *SyncInstallationAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncInstallationAccount.Unmarshal(m, b)
}
func (m *SyncInstallationAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncInstallationAccount.Marshal(b, m, deterministic)
}
func (m *SyncInstallationAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncInstallationAccount.Merge(m, src)
}
func (m *SyncInstallationAccount) XXX_Size() int {
	return xxx_messageInfo_SyncInstallationAccount.Size(m)
}
func (m *SyncInstallationAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncInstallationAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SyncInstallationAccount proto.InternalMessageInfo

func (m *SyncInstallationAccount) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *SyncInstallationAccount) GetProfileImage() string {
	if m != nil {
		return m.ProfileImage
	}
	return ""
}

func (m *SyncInstallationAccount) GetLastUpdated() uint64 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

type SyncInstallationPublicChat struct {
	Clock                uint64   `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncInstallationPublicChat) Reset()         { *m = SyncInstallationPublicChat{} }
func (m *SyncInstallationPublicChat) String() string { return proto.CompactTextString(m) }
func (*SyncInstallationPublicChat) ProtoMessage()    {}
func (*SyncInstallationPublicChat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61ab7221f0b5518, []int{3}
}

func (m *SyncInstallationPublicChat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncInstallationPublicChat.Unmarshal(m, b)
}
func (m *SyncInstallationPublicChat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncInstallationPublicChat.Marshal(b, m, deterministic)
}
func (m *SyncInstallationPublicChat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncInstallationPublicChat.Merge(m, src)
}
func (m *SyncInstallationPublicChat) XXX_Size() int {
	return xxx_messageInfo_SyncInstallationPublicChat.Size(m)
}
func (m *SyncInstallationPublicChat) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncInstallationPublicChat.DiscardUnknown(m)
}

var xxx_messageInfo_SyncInstallationPublicChat proto.InternalMessageInfo

func (m *SyncInstallationPublicChat) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *SyncInstallationPublicChat) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SyncInstallation struct {
	Contacts             []*SyncInstallationContact    `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	PublicChats          []*SyncInstallationPublicChat `protobuf:"bytes,2,rep,name=public_chats,json=publicChats,proto3" json:"public_chats,omitempty"`
	Account              *SyncInstallationAccount      `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SyncInstallation) Reset()         { *m = SyncInstallation{} }
func (m *SyncInstallation) String() string { return proto.CompactTextString(m) }
func (*SyncInstallation) ProtoMessage()    {}
func (*SyncInstallation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61ab7221f0b5518, []int{4}
}

func (m *SyncInstallation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncInstallation.Unmarshal(m, b)
}
func (m *SyncInstallation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncInstallation.Marshal(b, m, deterministic)
}
func (m *SyncInstallation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncInstallation.Merge(m, src)
}
func (m *SyncInstallation) XXX_Size() int {
	return xxx_messageInfo_SyncInstallation.Size(m)
}
func (m *SyncInstallation) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncInstallation.DiscardUnknown(m)
}

var xxx_messageInfo_SyncInstallation proto.InternalMessageInfo

func (m *SyncInstallation) GetContacts() []*SyncInstallationContact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *SyncInstallation) GetPublicChats() []*SyncInstallationPublicChat {
	if m != nil {
		return m.PublicChats
	}
	return nil
}

func (m *SyncInstallation) GetAccount() *SyncInstallationAccount {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*PairInstallation)(nil), "protobuf.PairInstallation")
	proto.RegisterType((*SyncInstallationContact)(nil), "protobuf.SyncInstallationContact")
	proto.RegisterType((*SyncInstallationAccount)(nil), "protobuf.SyncInstallationAccount")
	proto.RegisterType((*SyncInstallationPublicChat)(nil), "protobuf.SyncInstallationPublicChat")
	proto.RegisterType((*SyncInstallation)(nil), "protobuf.SyncInstallation")
}

func init() { proto.RegisterFile("pairing.proto", fileDescriptor_d61ab7221f0b5518) }

var fileDescriptor_d61ab7221f0b5518 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0xb4, 0xbb, 0x2d, 0xd3, 0x3f, 0xac, 0x2c, 0x24, 0x0c, 0x17, 0xb2, 0x01, 0x44,
	0x4f, 0x3d, 0xc0, 0x11, 0x71, 0x80, 0x3d, 0xa0, 0x5e, 0x56, 0xab, 0xb0, 0x9c, 0xad, 0xa9, 0xe3,
	0xcd, 0x5a, 0xeb, 0xd8, 0x56, 0xec, 0x80, 0xf2, 0x02, 0xbc, 0x22, 0xef, 0xc0, 0x53, 0xa0, 0x38,
	0x6d, 0x89, 0x5a, 0x82, 0x38, 0xc5, 0xf9, 0x32, 0x9e, 0xef, 0xfb, 0xcd, 0x04, 0x16, 0x16, 0x65,
	0x25, 0x75, 0xb1, 0xb6, 0x95, 0xf1, 0x86, 0x4c, 0xc3, 0x63, 0x5b, 0xdf, 0xa5, 0x3f, 0x22, 0xb8,
	0xb8, 0x41, 0x59, 0x6d, 0xb4, 0xf3, 0xa8, 0x14, 0x7a, 0x69, 0x34, 0x79, 0x02, 0x67, 0x5c, 0x19,
	0xfe, 0x40, 0xa3, 0x24, 0x5a, 0x8d, 0xb3, 0xee, 0x85, 0xbc, 0x81, 0xc7, 0xb2, 0x57, 0xc5, 0x64,
	0x4e, 0xe3, 0x24, 0x5a, 0x3d, 0xca, 0x96, 0x7d, 0x79, 0x93, 0x93, 0x17, 0x30, 0xcb, 0xc5, 0x37,
	0xc9, 0x05, 0xf3, 0x8d, 0x15, 0x74, 0x14, 0x8a, 0xa0, 0x93, 0x6e, 0x1b, 0x2b, 0x08, 0x81, 0xb1,
	0xc6, 0x52, 0xd0, 0x71, 0xf8, 0x12, 0xce, 0xe9, 0xaf, 0x08, 0x9e, 0x7e, 0x69, 0x34, 0xef, 0x07,
	0xb9, 0x32, 0xda, 0x23, 0xf7, 0x03, 0x79, 0x96, 0x10, 0x1f, 0x22, 0xc4, 0x32, 0x27, 0x2f, 0x61,
	0x61, 0x2b, 0x73, 0x27, 0x95, 0x60, 0xb2, 0xc4, 0x62, 0x6f, 0x3c, 0xdf, 0x89, 0x9b, 0x56, 0x23,
	0xcf, 0x60, 0x2a, 0xb4, 0x63, 0x3d, 0xfb, 0x89, 0xd0, 0xee, 0x1a, 0x4b, 0x41, 0x2e, 0x61, 0xae,
	0xd0, 0x79, 0x56, 0xdb, 0x1c, 0xbd, 0xc8, 0xe9, 0x59, 0x30, 0x9b, 0xb5, 0xda, 0xd7, 0x4e, 0x6a,
	0xc9, 0x5c, 0xe3, 0xbc, 0x28, 0x99, 0xc7, 0xc2, 0xd1, 0xf3, 0x64, 0xd4, 0x92, 0x75, 0xd2, 0x2d,
	0x16, 0x8e, 0xbc, 0x86, 0xa5, 0x32, 0x1c, 0x15, 0xd3, 0x92, 0x3f, 0x04, 0x93, 0x49, 0x30, 0x59,
	0x04, 0xf5, 0x7a, 0x27, 0xa6, 0xdf, 0x4f, 0x59, 0x3f, 0x72, 0x6e, 0x6a, 0x3d, 0xc4, 0x7a, 0xc2,
	0x16, 0xff, 0x85, 0xed, 0x18, 0x60, 0x74, 0x02, 0x90, 0x7e, 0x82, 0xe7, 0xc7, 0xc6, 0x37, 0xf5,
	0x56, 0x49, 0x7e, 0x75, 0x8f, 0xff, 0x39, 0xe7, 0xf4, 0x67, 0x04, 0x17, 0xc7, 0x4d, 0xc8, 0x07,
	0x98, 0xf2, 0x6e, 0x5b, 0x8e, 0x46, 0xc9, 0x68, 0x35, 0x7b, 0x7b, 0xb9, 0xde, 0xff, 0x64, 0xeb,
	0x81, 0xbd, 0x66, 0x87, 0x2b, 0xe4, 0x33, 0xcc, 0x6d, 0xc8, 0xc1, 0xf8, 0x3d, 0x7a, 0x47, 0xe3,
	0xd0, 0xe2, 0xd5, 0x70, 0x8b, 0x3f, 0xa9, 0xb3, 0x99, 0x3d, 0x9c, 0x1d, 0x79, 0x0f, 0x13, 0xec,
	0x26, 0x19, 0xf0, 0xff, 0x19, 0x63, 0x37, 0xf2, 0x6c, 0x7f, 0x63, 0x7b, 0x1e, 0x4a, 0xdf, 0xfd,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x5c, 0xf1, 0xac, 0x2e, 0x03, 0x00, 0x00,
}
