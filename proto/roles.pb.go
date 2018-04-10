// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roles.proto

/*
Package chremoas_roles is a generated protocol buffer package.

It is generated from these files:
	roles.proto

It has these top-level messages:
	NilMessage
	StringList
	Role
	UpdateInfo
	GetRolesResponse
	RoleSyncResponse
	MemberSyncResult
	MemberSyncResponse
	FilterList
	Filter
	Members
	MemberList
*/
package chremoas_roles

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MemberSyncAction int32

const (
	MemberSyncAction_NOTHING MemberSyncAction = 0
	MemberSyncAction_ADDED   MemberSyncAction = 1
	MemberSyncAction_REMOVED MemberSyncAction = 2
)

var MemberSyncAction_name = map[int32]string{
	0: "NOTHING",
	1: "ADDED",
	2: "REMOVED",
}
var MemberSyncAction_value = map[string]int32{
	"NOTHING": 0,
	"ADDED":   1,
	"REMOVED": 2,
}

func (x MemberSyncAction) String() string {
	return proto.EnumName(MemberSyncAction_name, int32(x))
}
func (MemberSyncAction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NilMessage struct {
}

func (m *NilMessage) Reset()                    { *m = NilMessage{} }
func (m *NilMessage) String() string            { return proto.CompactTextString(m) }
func (*NilMessage) ProtoMessage()               {}
func (*NilMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StringList struct {
	Value []string `protobuf:"bytes,1,rep,name=Value" json:"Value,omitempty"`
}

func (m *StringList) Reset()                    { *m = StringList{} }
func (m *StringList) String() string            { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()               {}
func (*StringList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	// Discord
	Name        string `protobuf:"bytes,5,opt,name=Name" json:"Name,omitempty"`
	Color       int32  `protobuf:"varint,6,opt,name=Color" json:"Color,omitempty"`
	Hoist       bool   `protobuf:"varint,7,opt,name=Hoist" json:"Hoist,omitempty"`
	Position    int32  `protobuf:"varint,8,opt,name=Position" json:"Position,omitempty"`
	Permissions int32  `protobuf:"varint,9,opt,name=Permissions" json:"Permissions,omitempty"`
	Managed     bool   `protobuf:"varint,10,opt,name=Managed" json:"Managed,omitempty"`
	Mentionable bool   `protobuf:"varint,11,opt,name=Mentionable" json:"Mentionable,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*UpdateInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*GetRolesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetRolesResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type RoleSyncResponse struct {
	Added   []string `protobuf:"bytes,1,rep,name=Added" json:"Added,omitempty"`
	Removed []string `protobuf:"bytes,2,rep,name=Removed" json:"Removed,omitempty"`
}

func (m *RoleSyncResponse) Reset()                    { *m = RoleSyncResponse{} }
func (m *RoleSyncResponse) String() string            { return proto.CompactTextString(m) }
func (*RoleSyncResponse) ProtoMessage()               {}
func (*RoleSyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RoleSyncResponse) GetAdded() []string {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *RoleSyncResponse) GetRemoved() []string {
	if m != nil {
		return m.Removed
	}
	return nil
}

type MemberSyncResult struct {
	Action MemberSyncAction `protobuf:"varint,1,opt,name=Action,enum=chremoas.roles.MemberSyncAction" json:"Action,omitempty"`
	User   string           `protobuf:"bytes,2,opt,name=User" json:"User,omitempty"`
	Role   string           `protobuf:"bytes,3,opt,name=Role" json:"Role,omitempty"`
}

func (m *MemberSyncResult) Reset()                    { *m = MemberSyncResult{} }
func (m *MemberSyncResult) String() string            { return proto.CompactTextString(m) }
func (*MemberSyncResult) ProtoMessage()               {}
func (*MemberSyncResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MemberSyncResult) GetAction() MemberSyncAction {
	if m != nil {
		return m.Action
	}
	return MemberSyncAction_NOTHING
}

func (m *MemberSyncResult) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *MemberSyncResult) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type MemberSyncResponse struct {
	Results []*MemberSyncResult `protobuf:"bytes,1,rep,name=Results" json:"Results,omitempty"`
}

func (m *MemberSyncResponse) Reset()                    { *m = MemberSyncResponse{} }
func (m *MemberSyncResponse) String() string            { return proto.CompactTextString(m) }
func (*MemberSyncResponse) ProtoMessage()               {}
func (*MemberSyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MemberSyncResponse) GetResults() []*MemberSyncResult {
	if m != nil {
		return m.Results
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
	proto.RegisterType((*StringList)(nil), "chremoas.roles.StringList")
	proto.RegisterType((*Role)(nil), "chremoas.roles.Role")
	proto.RegisterType((*UpdateInfo)(nil), "chremoas.roles.UpdateInfo")
	proto.RegisterType((*GetRolesResponse)(nil), "chremoas.roles.GetRolesResponse")
	proto.RegisterType((*RoleSyncResponse)(nil), "chremoas.roles.RoleSyncResponse")
	proto.RegisterType((*MemberSyncResult)(nil), "chremoas.roles.MemberSyncResult")
	proto.RegisterType((*MemberSyncResponse)(nil), "chremoas.roles.MemberSyncResponse")
	proto.RegisterType((*FilterList)(nil), "chremoas.roles.FilterList")
	proto.RegisterType((*Filter)(nil), "chremoas.roles.Filter")
	proto.RegisterType((*Members)(nil), "chremoas.roles.Members")
	proto.RegisterType((*MemberList)(nil), "chremoas.roles.MemberList")
	proto.RegisterEnum("chremoas.roles.MemberSyncAction", MemberSyncAction_name, MemberSyncAction_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Roles service

type RolesClient interface {
	AddRole(ctx context.Context, in *Role, opts ...client.CallOption) (*NilMessage, error)
	UpdateRole(ctx context.Context, in *UpdateInfo, opts ...client.CallOption) (*NilMessage, error)
	RemoveRole(ctx context.Context, in *Role, opts ...client.CallOption) (*NilMessage, error)
	GetRoles(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*GetRolesResponse, error)
	GetRole(ctx context.Context, in *Role, opts ...client.CallOption) (*Role, error)
	SyncRoles(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*RoleSyncResponse, error)
	GetRoleKeys(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*StringList, error)
	GetRoleTypes(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*StringList, error)
	SyncMembers(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*MemberSyncResponse, error)
	GetFilters(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*FilterList, error)
	AddFilter(ctx context.Context, in *Filter, opts ...client.CallOption) (*NilMessage, error)
	RemoveFilter(ctx context.Context, in *Filter, opts ...client.CallOption) (*NilMessage, error)
	GetMembers(ctx context.Context, in *Filter, opts ...client.CallOption) (*MemberList, error)
	AddMembers(ctx context.Context, in *Members, opts ...client.CallOption) (*NilMessage, error)
	RemoveMembers(ctx context.Context, in *Members, opts ...client.CallOption) (*NilMessage, error)
}

type rolesClient struct {
	c           client.Client
	serviceName string
}

func NewRolesClient(serviceName string, c client.Client) RolesClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "chremoas.roles"
	}
	return &rolesClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *rolesClient) AddRole(ctx context.Context, in *Role, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.AddRole", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) UpdateRole(ctx context.Context, in *UpdateInfo, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.UpdateRole", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) RemoveRole(ctx context.Context, in *Role, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.RemoveRole", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetRoles(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*GetRolesResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.GetRoles", in)
	out := new(GetRolesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetRole(ctx context.Context, in *Role, opts ...client.CallOption) (*Role, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.GetRole", in)
	out := new(Role)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) SyncRoles(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*RoleSyncResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.SyncRoles", in)
	out := new(RoleSyncResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetRoleKeys(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*StringList, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.GetRoleKeys", in)
	out := new(StringList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetRoleTypes(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*StringList, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.GetRoleTypes", in)
	out := new(StringList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) SyncMembers(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*MemberSyncResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.SyncMembers", in)
	out := new(MemberSyncResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetFilters(ctx context.Context, in *NilMessage, opts ...client.CallOption) (*FilterList, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.GetFilters", in)
	out := new(FilterList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) AddFilter(ctx context.Context, in *Filter, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.AddFilter", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) RemoveFilter(ctx context.Context, in *Filter, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.RemoveFilter", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetMembers(ctx context.Context, in *Filter, opts ...client.CallOption) (*MemberList, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.GetMembers", in)
	out := new(MemberList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) AddMembers(ctx context.Context, in *Members, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.AddMembers", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) RemoveMembers(ctx context.Context, in *Members, opts ...client.CallOption) (*NilMessage, error) {
	req := c.c.NewRequest(c.serviceName, "Roles.RemoveMembers", in)
	out := new(NilMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Roles service

type RolesHandler interface {
	AddRole(context.Context, *Role, *NilMessage) error
	UpdateRole(context.Context, *UpdateInfo, *NilMessage) error
	RemoveRole(context.Context, *Role, *NilMessage) error
	GetRoles(context.Context, *NilMessage, *GetRolesResponse) error
	GetRole(context.Context, *Role, *Role) error
	SyncRoles(context.Context, *NilMessage, *RoleSyncResponse) error
	GetRoleKeys(context.Context, *NilMessage, *StringList) error
	GetRoleTypes(context.Context, *NilMessage, *StringList) error
	SyncMembers(context.Context, *NilMessage, *MemberSyncResponse) error
	GetFilters(context.Context, *NilMessage, *FilterList) error
	AddFilter(context.Context, *Filter, *NilMessage) error
	RemoveFilter(context.Context, *Filter, *NilMessage) error
	GetMembers(context.Context, *Filter, *MemberList) error
	AddMembers(context.Context, *Members, *NilMessage) error
	RemoveMembers(context.Context, *Members, *NilMessage) error
}

func RegisterRolesHandler(s server.Server, hdlr RolesHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Roles{hdlr}, opts...))
}

type Roles struct {
	RolesHandler
}

func (h *Roles) AddRole(ctx context.Context, in *Role, out *NilMessage) error {
	return h.RolesHandler.AddRole(ctx, in, out)
}

func (h *Roles) UpdateRole(ctx context.Context, in *UpdateInfo, out *NilMessage) error {
	return h.RolesHandler.UpdateRole(ctx, in, out)
}

func (h *Roles) RemoveRole(ctx context.Context, in *Role, out *NilMessage) error {
	return h.RolesHandler.RemoveRole(ctx, in, out)
}

func (h *Roles) GetRoles(ctx context.Context, in *NilMessage, out *GetRolesResponse) error {
	return h.RolesHandler.GetRoles(ctx, in, out)
}

func (h *Roles) GetRole(ctx context.Context, in *Role, out *Role) error {
	return h.RolesHandler.GetRole(ctx, in, out)
}

func (h *Roles) SyncRoles(ctx context.Context, in *NilMessage, out *RoleSyncResponse) error {
	return h.RolesHandler.SyncRoles(ctx, in, out)
}

func (h *Roles) GetRoleKeys(ctx context.Context, in *NilMessage, out *StringList) error {
	return h.RolesHandler.GetRoleKeys(ctx, in, out)
}

func (h *Roles) GetRoleTypes(ctx context.Context, in *NilMessage, out *StringList) error {
	return h.RolesHandler.GetRoleTypes(ctx, in, out)
}

func (h *Roles) SyncMembers(ctx context.Context, in *NilMessage, out *MemberSyncResponse) error {
	return h.RolesHandler.SyncMembers(ctx, in, out)
}

func (h *Roles) GetFilters(ctx context.Context, in *NilMessage, out *FilterList) error {
	return h.RolesHandler.GetFilters(ctx, in, out)
}

func (h *Roles) AddFilter(ctx context.Context, in *Filter, out *NilMessage) error {
	return h.RolesHandler.AddFilter(ctx, in, out)
}

func (h *Roles) RemoveFilter(ctx context.Context, in *Filter, out *NilMessage) error {
	return h.RolesHandler.RemoveFilter(ctx, in, out)
}

func (h *Roles) GetMembers(ctx context.Context, in *Filter, out *MemberList) error {
	return h.RolesHandler.GetMembers(ctx, in, out)
}

func (h *Roles) AddMembers(ctx context.Context, in *Members, out *NilMessage) error {
	return h.RolesHandler.AddMembers(ctx, in, out)
}

func (h *Roles) RemoveMembers(ctx context.Context, in *Members, out *NilMessage) error {
	return h.RolesHandler.RemoveMembers(ctx, in, out)
}

func init() { proto.RegisterFile("roles.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4d, 0x4f, 0xdb, 0x40,
	0x10, 0x75, 0x12, 0x9c, 0xc4, 0x63, 0x8a, 0xac, 0x55, 0x44, 0xad, 0xa8, 0x87, 0x68, 0x0f, 0x55,
	0xc4, 0x21, 0x07, 0xaa, 0x96, 0xaa, 0x52, 0x51, 0x1d, 0x0c, 0x09, 0xa5, 0x09, 0xc8, 0x7c, 0xdc,
	0x4d, 0xbc, 0x05, 0x4b, 0x8e, 0x37, 0xf2, 0x9a, 0x4a, 0xf9, 0x43, 0x55, 0x7f, 0x66, 0xb5, 0x5f,
	0xb1, 0x09, 0x21, 0x94, 0xe6, 0xb6, 0xf3, 0x66, 0xe6, 0xed, 0x9b, 0xd9, 0x19, 0xcb, 0x60, 0x67,
	0x34, 0x21, 0xac, 0x37, 0xcb, 0x68, 0x4e, 0xd1, 0xce, 0xe4, 0x3e, 0x23, 0x53, 0x1a, 0xb2, 0x9e,
	0x40, 0xf1, 0x36, 0xc0, 0x38, 0x4e, 0x46, 0x84, 0xb1, 0xf0, 0x8e, 0x60, 0x0c, 0x70, 0x99, 0x67,
	0x71, 0x7a, 0xf7, 0x23, 0x66, 0x39, 0x6a, 0x81, 0x79, 0x13, 0x26, 0x0f, 0xc4, 0xad, 0x74, 0x6a,
	0x5d, 0x2b, 0x90, 0x06, 0xfe, 0x5d, 0x85, 0xad, 0x80, 0x26, 0x04, 0x21, 0xd8, 0xba, 0x9a, 0xcf,
	0xb8, 0xb7, 0xd2, 0xb5, 0x02, 0x71, 0x46, 0xef, 0xc0, 0xba, 0xbc, 0xa7, 0x59, 0x3e, 0x0e, 0xa7,
	0xc4, 0xad, 0x0a, 0x47, 0x01, 0x20, 0x17, 0x1a, 0x27, 0x71, 0x92, 0x93, 0xcc, 0x73, 0x6b, 0xc2,
	0xa7, 0xcd, 0xc2, 0xd3, 0x77, 0xb7, 0xca, 0x9e, 0x3e, 0xbf, 0x45, 0x90, 0x99, 0xf2, 0x16, 0xc1,
	0xd3, 0x02, 0xf3, 0x88, 0x26, 0x34, 0x73, 0xeb, 0x9d, 0x4a, 0xd7, 0x0c, 0xa4, 0xc1, 0xd1, 0x21,
	0x8d, 0x59, 0xee, 0x36, 0x3a, 0x95, 0x6e, 0x33, 0x90, 0x06, 0x6a, 0x43, 0xf3, 0x82, 0xb2, 0x38,
	0x8f, 0x69, 0xea, 0x36, 0x45, 0xf8, 0xc2, 0x46, 0x1d, 0xb0, 0x2f, 0x48, 0x36, 0x8d, 0x19, 0x8b,
	0x69, 0xca, 0x5c, 0x4b, 0xb8, 0xcb, 0x10, 0xd7, 0x35, 0x0a, 0xd3, 0xf0, 0x8e, 0x44, 0x2e, 0x08,
	0x56, 0x6d, 0xf2, 0xdc, 0x11, 0x49, 0x39, 0x4d, 0x78, 0x9b, 0x10, 0xd7, 0x16, 0xde, 0x32, 0x84,
	0x87, 0x00, 0xd7, 0xb3, 0x28, 0xcc, 0xc9, 0x69, 0xfa, 0x93, 0x2e, 0xea, 0xa8, 0x94, 0xea, 0x70,
	0xa0, 0x76, 0x46, 0xe6, 0xaa, 0x4f, 0xfc, 0x58, 0xb4, 0x5c, 0xf6, 0x47, 0xb5, 0xfc, 0x10, 0x9c,
	0x01, 0xc9, 0x79, 0xd3, 0x59, 0x40, 0xd8, 0x8c, 0xa6, 0x8c, 0xa0, 0x3d, 0x30, 0x05, 0x20, 0x1e,
	0xc7, 0xde, 0x6f, 0xf5, 0x1e, 0x3f, 0x6c, 0x8f, 0x3b, 0x03, 0x19, 0x82, 0xfb, 0xe0, 0xf0, 0xc3,
	0xe5, 0x3c, 0x9d, 0x2c, 0xf2, 0x5b, 0x60, 0x7a, 0x51, 0x44, 0x22, 0xfd, 0xb8, 0xc2, 0xe0, 0xf5,
	0x06, 0x64, 0x4a, 0x7f, 0x91, 0xc8, 0xad, 0x0a, 0x5c, 0x9b, 0x38, 0x07, 0x67, 0x44, 0xa6, 0xb7,
	0x24, 0x53, 0x2c, 0x0f, 0x49, 0x8e, 0x3e, 0x43, 0xdd, 0x9b, 0x88, 0xce, 0xf2, 0xaa, 0x76, 0xf6,
	0x3b, 0xcb, 0x22, 0x8a, 0x0c, 0x19, 0x17, 0xa8, 0x78, 0xde, 0x8d, 0x6b, 0x46, 0x32, 0x55, 0xba,
	0x38, 0x73, 0x8c, 0xab, 0x54, 0xa5, 0x8b, 0x33, 0xbe, 0x00, 0xf4, 0xe8, 0x56, 0xa9, 0xfd, 0x0b,
	0x57, 0xc9, 0x15, 0xe8, 0xea, 0xd7, 0x5c, 0x2c, 0x03, 0x03, 0x9d, 0x80, 0x7d, 0x00, 0x39, 0x5a,
	0x62, 0xc4, 0x3f, 0x95, 0x2d, 0x45, 0xb6, 0xbb, 0x4c, 0x26, 0x23, 0x82, 0x52, 0x24, 0x3e, 0x84,
	0xba, 0xb4, 0x56, 0xbe, 0x6b, 0x07, 0x6c, 0x9f, 0xb0, 0x49, 0x16, 0xcf, 0x44, 0x73, 0x64, 0x91,
	0x65, 0x08, 0x7f, 0x84, 0x86, 0x94, 0xc8, 0x4a, 0x04, 0xb5, 0x05, 0xc1, 0xae, 0xa6, 0x57, 0xb9,
	0xca, 0xc2, 0xef, 0x01, 0x64, 0x9a, 0x10, 0xef, 0x2e, 0x48, 0x54, 0xb2, 0x36, 0xf7, 0x0e, 0xca,
	0x8f, 0xa5, 0x5a, 0x6e, 0x43, 0x63, 0x7c, 0x7e, 0x35, 0x3c, 0x1d, 0x0f, 0x1c, 0x03, 0x59, 0x60,
	0x7a, 0xbe, 0x7f, 0xec, 0x3b, 0x15, 0x8e, 0x07, 0xc7, 0xa3, 0xf3, 0x9b, 0x63, 0xdf, 0xa9, 0xee,
	0xff, 0x69, 0xaa, 0xb1, 0x42, 0x5f, 0xa1, 0xe1, 0x45, 0x91, 0x58, 0xf4, 0x95, 0xb3, 0xd5, 0x6e,
	0x2f, 0xa3, 0xa5, 0xef, 0x88, 0x81, 0x4e, 0xf4, 0xf0, 0x0b, 0x86, 0x27, 0xb1, 0xc5, 0x62, 0xbc,
	0xc0, 0xf3, 0x0d, 0x40, 0x4e, 0xe0, 0x7f, 0x2b, 0xf9, 0x0e, 0x4d, 0xbd, 0x3c, 0x68, 0x4d, 0x64,
	0xfb, 0xc9, 0x0c, 0x2d, 0xaf, 0x1c, 0x36, 0xd0, 0x01, 0x34, 0x14, 0xfa, 0x8c, 0x94, 0x95, 0x28,
	0x36, 0xd0, 0x19, 0x58, 0x62, 0x18, 0x5f, 0xaf, 0x62, 0x79, 0x71, 0xb1, 0x81, 0x06, 0x60, 0x2b,
	0x15, 0x67, 0x64, 0xbe, 0x9e, 0xee, 0x89, 0xaf, 0xf8, 0xbc, 0x63, 0x03, 0x0d, 0x61, 0x5b, 0x11,
	0xf1, 0x8f, 0xf7, 0x26, 0x4c, 0xe7, 0x60, 0x73, 0x91, 0x7a, 0xa6, 0xd7, 0x11, 0xe1, 0xb5, 0xbb,
	0xaa, 0x6b, 0x3c, 0x01, 0x18, 0x90, 0x5c, 0x8e, 0xfd, 0x2b, 0x85, 0x95, 0xd6, 0xd4, 0x40, 0x1e,
	0x58, 0x5e, 0x14, 0xa9, 0x5d, 0x7d, 0x66, 0xb3, 0x5f, 0x18, 0x20, 0x1f, 0xb6, 0xe5, 0x08, 0x6e,
	0xc4, 0xd2, 0x17, 0x05, 0xe9, 0x06, 0xfd, 0x33, 0x47, 0xb1, 0xee, 0xd8, 0x40, 0x47, 0x00, 0x5e,
	0x14, 0x69, 0x8e, 0xb7, 0xab, 0x63, 0xd9, 0x8b, 0x9b, 0xf9, 0x46, 0x96, 0xb3, 0x19, 0xcf, 0x6d,
	0x5d, 0xfc, 0x50, 0x7c, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x52, 0x98, 0xea, 0xd7, 0x5f, 0x08,
	0x00, 0x00,
}
