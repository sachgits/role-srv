// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roles.proto

/*
Package chremoas_roles is a generated protocol buffer package.

It is generated from these files:
	roles.proto

It has these top-level messages:
	NilMessage
	GetDiscordUserRequest
	GetDiscordUserResponse
	SyncRequest
	StringList
	Role
	UpdateInfo
	GetRolesResponse
	FilterList
	Filter
	Members
	MemberList
*/
package chremoas_roles

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NilMessage struct {
}

func (m *NilMessage) Reset()                    { *m = NilMessage{} }
func (m *NilMessage) String() string            { return proto.CompactTextString(m) }
func (*NilMessage) ProtoMessage()               {}
func (*NilMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetDiscordUserRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *GetDiscordUserRequest) Reset()                    { *m = GetDiscordUserRequest{} }
func (m *GetDiscordUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDiscordUserRequest) ProtoMessage()               {}
func (*GetDiscordUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetDiscordUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetDiscordUserResponse struct {
	Username      string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Discriminator string `protobuf:"bytes,2,opt,name=Discriminator" json:"Discriminator,omitempty"`
	Avatar        string `protobuf:"bytes,3,opt,name=Avatar" json:"Avatar,omitempty"`
	Bot           bool   `protobuf:"varint,4,opt,name=Bot" json:"Bot,omitempty"`
	MfaEnabled    bool   `protobuf:"varint,5,opt,name=MfaEnabled" json:"MfaEnabled,omitempty"`
	Verified      bool   `protobuf:"varint,6,opt,name=Verified" json:"Verified,omitempty"`
	Email         string `protobuf:"bytes,7,opt,name=Email" json:"Email,omitempty"`
}

func (m *GetDiscordUserResponse) Reset()                    { *m = GetDiscordUserResponse{} }
func (m *GetDiscordUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDiscordUserResponse) ProtoMessage()               {}
func (*GetDiscordUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetDiscordUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetDiscordUserResponse) GetDiscriminator() string {
	if m != nil {
		return m.Discriminator
	}
	return ""
}

func (m *GetDiscordUserResponse) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *GetDiscordUserResponse) GetBot() bool {
	if m != nil {
		return m.Bot
	}
	return false
}

func (m *GetDiscordUserResponse) GetMfaEnabled() bool {
	if m != nil {
		return m.MfaEnabled
	}
	return false
}

func (m *GetDiscordUserResponse) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *GetDiscordUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SyncRequest struct {
	ChannelId   string `protobuf:"bytes,1,opt,name=ChannelId" json:"ChannelId,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=UserId" json:"UserId,omitempty"`
	SendMessage bool   `protobuf:"varint,3,opt,name=SendMessage" json:"SendMessage,omitempty"`
}

func (m *SyncRequest) Reset()                    { *m = SyncRequest{} }
func (m *SyncRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()               {}
func (*SyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SyncRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SyncRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SyncRequest) GetSendMessage() bool {
	if m != nil {
		return m.SendMessage
	}
	return false
}

type StringList struct {
	Value []string `protobuf:"bytes,1,rep,name=Value" json:"Value,omitempty"`
}

func (m *StringList) Reset()                    { *m = StringList{} }
func (m *StringList) String() string            { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()               {}
func (*StringList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type Role struct {
	Type      string `protobuf:"bytes,1,opt,name=Type" json:"Type,omitempty"`
	ShortName string `protobuf:"bytes,2,opt,name=ShortName" json:"ShortName,omitempty"`
	FilterA   string `protobuf:"bytes,3,opt,name=FilterA" json:"FilterA,omitempty"`
	FilterB   string `protobuf:"bytes,4,opt,name=FilterB" json:"FilterB,omitempty"`
	Sig       bool   `protobuf:"varint,5,opt,name=Sig" json:"Sig,omitempty"`
	Joinable  bool   `protobuf:"varint,6,opt,name=Joinable" json:"Joinable,omitempty"`
	// Discord
	Name        string `protobuf:"bytes,20,opt,name=Name" json:"Name,omitempty"`
	Color       int32  `protobuf:"varint,21,opt,name=Color" json:"Color,omitempty"`
	Hoist       bool   `protobuf:"varint,22,opt,name=Hoist" json:"Hoist,omitempty"`
	Position    int32  `protobuf:"varint,23,opt,name=Position" json:"Position,omitempty"`
	Permissions int32  `protobuf:"varint,24,opt,name=Permissions" json:"Permissions,omitempty"`
	Managed     bool   `protobuf:"varint,25,opt,name=Managed" json:"Managed,omitempty"`
	Mentionable bool   `protobuf:"varint,26,opt,name=Mentionable" json:"Mentionable,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Role) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Role) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *Role) GetFilterA() string {
	if m != nil {
		return m.FilterA
	}
	return ""
}

func (m *Role) GetFilterB() string {
	if m != nil {
		return m.FilterB
	}
	return ""
}

func (m *Role) GetSig() bool {
	if m != nil {
		return m.Sig
	}
	return false
}

func (m *Role) GetJoinable() bool {
	if m != nil {
		return m.Joinable
	}
	return false
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetColor() int32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *Role) GetHoist() bool {
	if m != nil {
		return m.Hoist
	}
	return false
}

func (m *Role) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Role) GetPermissions() int32 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

func (m *Role) GetManaged() bool {
	if m != nil {
		return m.Managed
	}
	return false
}

func (m *Role) GetMentionable() bool {
	if m != nil {
		return m.Mentionable
	}
	return false
}

type UpdateInfo struct {
	Name  string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=Value" json:"Value,omitempty"`
}

func (m *UpdateInfo) Reset()                    { *m = UpdateInfo{} }
func (m *UpdateInfo) String() string            { return proto.CompactTextString(m) }
func (*UpdateInfo) ProtoMessage()               {}
func (*UpdateInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpdateInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetRolesResponse struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=Roles" json:"Roles,omitempty"`
}

func (m *GetRolesResponse) Reset()                    { *m = GetRolesResponse{} }
func (m *GetRolesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRolesResponse) ProtoMessage()               {}
func (*GetRolesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetRolesResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type FilterList struct {
	FilterList []*Filter `protobuf:"bytes,1,rep,name=FilterList" json:"FilterList,omitempty"`
}

func (m *FilterList) Reset()                    { *m = FilterList{} }
func (m *FilterList) String() string            { return proto.CompactTextString(m) }
func (*FilterList) ProtoMessage()               {}
func (*FilterList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FilterList) GetFilterList() []*Filter {
	if m != nil {
		return m.FilterList
	}
	return nil
}

type Filter struct {
	Name        string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Members struct {
	Name   []string `protobuf:"bytes,1,rep,name=Name" json:"Name,omitempty"`
	Filter string   `protobuf:"bytes,2,opt,name=Filter" json:"Filter,omitempty"`
}

func (m *Members) Reset()                    { *m = Members{} }
func (m *Members) String() string            { return proto.CompactTextString(m) }
func (*Members) ProtoMessage()               {}
func (*Members) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Members) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Members) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type MemberList struct {
	Members []string `protobuf:"bytes,1,rep,name=Members" json:"Members,omitempty"`
}

func (m *MemberList) Reset()                    { *m = MemberList{} }
func (m *MemberList) String() string            { return proto.CompactTextString(m) }
func (*MemberList) ProtoMessage()               {}
func (*MemberList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MemberList) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*NilMessage)(nil), "chremoas.roles.NilMessage")
	proto.RegisterType((*GetDiscordUserRequest)(nil), "chremoas.roles.GetDiscordUserRequest")
	proto.RegisterType((*GetDiscordUserResponse)(nil), "chremoas.roles.GetDiscordUserResponse")
	proto.RegisterType((*SyncRequest)(nil), "chremoas.roles.SyncRequest")
	proto.RegisterType((*StringList)(nil), "chremoas.roles.StringList")
	proto.RegisterType((*Role)(nil), "chremoas.roles.Role")
	proto.RegisterType((*UpdateInfo)(nil), "chremoas.roles.UpdateInfo")
	proto.RegisterType((*GetRolesResponse)(nil), "chremoas.roles.GetRolesResponse")
	proto.RegisterType((*FilterList)(nil), "chremoas.roles.FilterList")
	proto.RegisterType((*Filter)(nil), "chremoas.roles.Filter")
	proto.RegisterType((*Members)(nil), "chremoas.roles.Members")
	proto.RegisterType((*MemberList)(nil), "chremoas.roles.MemberList")
}

func init() { proto.RegisterFile("roles.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdf, 0x4f, 0xdb, 0x48,
	0x10, 0x4e, 0x08, 0xf9, 0x35, 0x01, 0x84, 0x56, 0x21, 0xf8, 0x72, 0xa7, 0x53, 0xb4, 0xba, 0x43,
	0xe8, 0x1e, 0x72, 0x12, 0x55, 0xdb, 0xa7, 0xa2, 0x26, 0x04, 0xc2, 0x8f, 0x06, 0x21, 0xa7, 0xf0,
	0x6e, 0xf0, 0x10, 0x56, 0x72, 0xbc, 0xa9, 0x77, 0x41, 0xca, 0xff, 0x59, 0xf5, 0xdf, 0x69, 0xb5,
	0xbb, 0x5e, 0xdb, 0x24, 0x81, 0xb4, 0xcd, 0x9b, 0xbf, 0xd9, 0x99, 0x6f, 0x67, 0xbe, 0x99, 0x59,
	0x19, 0x6a, 0x11, 0x0f, 0x50, 0xb4, 0x27, 0x11, 0x97, 0x9c, 0x6c, 0xdd, 0x3d, 0x44, 0x38, 0xe6,
	0x9e, 0x68, 0x6b, 0x2b, 0xdd, 0x00, 0xb8, 0x64, 0xc1, 0x00, 0x85, 0xf0, 0x46, 0x48, 0xff, 0x87,
	0x9d, 0x3e, 0xca, 0x1e, 0x13, 0x77, 0x3c, 0xf2, 0xaf, 0x05, 0x46, 0x2e, 0x7e, 0x79, 0x44, 0x21,
	0x49, 0x03, 0x4a, 0x0a, 0x9e, 0xf9, 0x4e, 0xbe, 0x95, 0xdf, 0xaf, 0xba, 0x31, 0xa2, 0xdf, 0xf2,
	0xd0, 0x98, 0x8d, 0x10, 0x13, 0x1e, 0x0a, 0x24, 0x4d, 0xa8, 0x28, 0x1c, 0x7a, 0x63, 0x8c, 0x83,
	0x12, 0x4c, 0xfe, 0x81, 0x4d, 0x15, 0x12, 0xb1, 0x31, 0x0b, 0x3d, 0xc9, 0x23, 0x67, 0x4d, 0x3b,
	0x3c, 0x37, 0xaa, 0x4b, 0x3b, 0x4f, 0x9e, 0xf4, 0x22, 0xa7, 0x60, 0x2e, 0x35, 0x88, 0x6c, 0x43,
	0xa1, 0xcb, 0xa5, 0xb3, 0xde, 0xca, 0xef, 0x57, 0x5c, 0xf5, 0x49, 0xfe, 0x06, 0x18, 0xdc, 0x7b,
	0xc7, 0xa1, 0x77, 0x1b, 0xa0, 0xef, 0x14, 0xf5, 0x41, 0xc6, 0xa2, 0x72, 0xb9, 0xc1, 0x88, 0xdd,
	0x33, 0xf4, 0x9d, 0x92, 0x3e, 0x4d, 0x30, 0xa9, 0x43, 0xf1, 0x78, 0xec, 0xb1, 0xc0, 0x29, 0xeb,
	0x4b, 0x0c, 0xa0, 0x08, 0xb5, 0xe1, 0x34, 0xbc, 0xb3, 0xf5, 0xff, 0x05, 0xd5, 0xa3, 0x07, 0x2f,
	0x0c, 0x31, 0x48, 0x24, 0x48, 0x0d, 0x19, 0x75, 0xd6, 0xb2, 0xea, 0x90, 0x16, 0xd4, 0x86, 0x18,
	0xfa, 0xb1, 0xba, 0xba, 0x8a, 0x8a, 0x9b, 0x35, 0x51, 0x0a, 0x30, 0x94, 0x11, 0x0b, 0x47, 0x9f,
	0x98, 0x90, 0x2a, 0x95, 0x1b, 0x2f, 0x78, 0x54, 0x7a, 0x15, 0x54, 0x2a, 0x1a, 0xd0, 0xaf, 0x6b,
	0xb0, 0xee, 0xf2, 0x00, 0x09, 0x81, 0xf5, 0xcf, 0xd3, 0x89, 0x55, 0x53, 0x7f, 0xab, 0xc4, 0x86,
	0x0f, 0x3c, 0x92, 0x97, 0x4a, 0x66, 0x73, 0x7b, 0x6a, 0x20, 0x0e, 0x94, 0x4f, 0x58, 0x20, 0x31,
	0xea, 0xc4, 0x12, 0x5a, 0x98, 0x9e, 0x74, 0xb5, 0x8e, 0xc9, 0x49, 0x57, 0xa9, 0x3b, 0x64, 0xa3,
	0x58, 0x44, 0xf5, 0xa9, 0xd4, 0x3b, 0xe7, 0x4c, 0x4b, 0x69, 0xd5, 0xb3, 0x58, 0xe5, 0xa4, 0xaf,
	0xae, 0x9b, 0x9c, 0xf4, 0xad, 0x75, 0x28, 0x1e, 0xf1, 0x80, 0x47, 0xce, 0x4e, 0x2b, 0xbf, 0x5f,
	0x74, 0x0d, 0x50, 0xd6, 0x53, 0xce, 0x84, 0x74, 0x1a, 0x9a, 0xc2, 0x00, 0xc5, 0x7d, 0xc5, 0x05,
	0x93, 0x8c, 0x87, 0xce, 0xae, 0x76, 0x4f, 0xb0, 0x92, 0xef, 0x0a, 0xa3, 0x31, 0x13, 0x82, 0xf1,
	0x50, 0x38, 0x8e, 0x3e, 0xce, 0x9a, 0x54, 0x15, 0x03, 0x2f, 0xf4, 0x46, 0xe8, 0x3b, 0x7f, 0x68,
	0x56, 0x0b, 0x55, 0xec, 0x00, 0x43, 0x45, 0xa3, 0xd3, 0x6e, 0x1a, 0xe9, 0x33, 0x26, 0x7a, 0x0a,
	0x70, 0x3d, 0xf1, 0x3d, 0x89, 0x67, 0xe1, 0x3d, 0x4f, 0xea, 0xc8, 0x67, 0xea, 0xd8, 0x86, 0xc2,
	0x05, 0x4e, 0x63, 0x55, 0xd5, 0x67, 0xda, 0x20, 0xa3, 0x66, 0xdc, 0xa0, 0x43, 0xd8, 0xee, 0xa3,
	0x54, 0x2d, 0x12, 0xc9, 0xf4, 0xff, 0x07, 0x45, 0x6d, 0xd0, 0xad, 0xac, 0x1d, 0xd4, 0xdb, 0xcf,
	0xf7, 0xae, 0xad, 0x0e, 0x5d, 0xe3, 0x42, 0x7b, 0x00, 0x46, 0x7c, 0x3d, 0x04, 0xef, 0xb2, 0x28,
	0x0e, 0x6f, 0xcc, 0x86, 0x1b, 0x0f, 0x37, 0xe3, 0x49, 0x0f, 0xa1, 0x64, 0xd0, 0xc2, 0x5a, 0x5a,
	0x50, 0xeb, 0xa1, 0x5a, 0xae, 0x89, 0x96, 0xda, 0xd4, 0x94, 0x35, 0xd1, 0xb7, 0x50, 0x1e, 0xe0,
	0xf8, 0x16, 0x23, 0x91, 0x21, 0x28, 0x24, 0x04, 0x0d, 0x4b, 0x6f, 0x67, 0xdc, 0x20, 0xba, 0x07,
	0x60, 0xc2, 0x74, 0xf2, 0x4e, 0x42, 0x12, 0x07, 0x5b, 0x78, 0xf0, 0xbd, 0x12, 0x2b, 0x42, 0x3e,
	0x40, 0xb9, 0xe3, 0xfb, 0x7a, 0xa2, 0x17, 0xca, 0xd2, 0x6c, 0xce, 0x5a, 0x33, 0x2f, 0x54, 0x8e,
	0x9c, 0xd8, 0xbe, 0x69, 0x86, 0x39, 0xdf, 0xb4, 0xa7, 0x4b, 0x78, 0x3e, 0x02, 0xb8, 0x38, 0xe6,
	0x4f, 0xf8, 0xdb, 0x99, 0x9c, 0x43, 0xc5, 0xf6, 0x9d, 0xbc, 0xe2, 0xd9, 0x6c, 0xcd, 0x9e, 0xcd,
	0x4e, 0x0b, 0xcd, 0x91, 0xf7, 0x50, 0x8e, 0xad, 0x2f, 0xa4, 0xb2, 0xd0, 0x4a, 0x73, 0xa4, 0x0f,
	0xb5, 0x38, 0xf0, 0x02, 0xa7, 0xaf, 0xe7, 0x31, 0x77, 0x96, 0x3e, 0x3d, 0x34, 0x47, 0x4e, 0x61,
	0x23, 0x26, 0x52, 0x0f, 0xcb, 0x6a, 0x4c, 0xfa, 0xed, 0xb4, 0xd3, 0xf4, 0xe7, 0x9c, 0x73, 0xfa,
	0xb0, 0x2e, 0xed, 0x75, 0x55, 0x3b, 0x6b, 0x89, 0x57, 0xe2, 0x81, 0x3e, 0x4a, 0x33, 0xb1, 0xbf,
	0x58, 0x59, 0x66, 0xc3, 0x72, 0xa4, 0x03, 0xd5, 0x8e, 0xef, 0xc7, 0x6b, 0xf6, 0xc2, 0x52, 0x2e,
	0x49, 0xa5, 0x07, 0x1b, 0x66, 0xec, 0x56, 0x62, 0xe9, 0xea, 0x82, 0xac, 0xc2, 0x3f, 0xcd, 0x91,
	0x6e, 0x2a, 0xcd, 0x91, 0x23, 0x80, 0x8e, 0xef, 0x5b, 0x8e, 0xdd, 0xc5, 0xbe, 0x62, 0xa9, 0xb2,
	0x9b, 0xa6, 0x9c, 0x15, 0x79, 0x3c, 0xd8, 0x7a, 0xfe, 0x1f, 0x41, 0xfe, 0x5d, 0xb0, 0x35, 0xf3,
	0x7f, 0x26, 0xcd, 0xbd, 0x65, 0x6e, 0x76, 0xc5, 0x6e, 0x4b, 0xfa, 0x0f, 0xe8, 0xcd, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x56, 0x34, 0xe0, 0xfe, 0x10, 0x09, 0x00, 0x00,
}
