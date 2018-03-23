// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/cipd/api/cipd/v1/repo.proto

package api

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Roles used in package prefix ACLs.
//
// A user can have one or more such roles for a package prefix. They get
// inherited by all subprefixes.
type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	// Readers can fetch package instances and package metadata (e.g. list of
	// instances, all tags, all refs), but not prefix metadata (e.g. ACLs).
	Role_READER Role = 1
	// Writers can do everything that readers can, plus create new packages,
	// upload package instances, attach tags, move refs.
	Role_WRITER Role = 2
	// Owners can do everything that writers can, plus read prefix metadata for
	// all parent prefixes and all subprefixes, and modify prefix metadata for
	// all subprefixes.
	Role_OWNER Role = 3
)

var Role_name = map[int32]string{
	0: "ROLE_UNSPECIFIED",
	1: "READER",
	2: "WRITER",
	3: "OWNER",
}
var Role_value = map[string]int32{
	"ROLE_UNSPECIFIED": 0,
	"READER":           1,
	"WRITER":           2,
	"OWNER":            3,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}
func (Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PrefixRequest struct {
	// A prefix within the repository, e.g. "a/b/c".
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *PrefixRequest) Reset()                    { *m = PrefixRequest{} }
func (m *PrefixRequest) String() string            { return proto.CompactTextString(m) }
func (*PrefixRequest) ProtoMessage()               {}
func (*PrefixRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *PrefixRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// PrefixMetadata is metadata defined at some concrete package prefix.
//
// It applies to this prefix and all subprefixes, recursively.
type PrefixMetadata struct {
	// Prefix this metadata is defined at, e.g. "a/b/c".
	//
	// Note: there's no metadata at the root, so prefix must never be "".
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	// An opaque string that identifies a particular version of this metadata.
	//
	// Used by UpdatePrefixMetadata to prevent an accidental overwrite of changes.
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
	// When the metadata was modified the last time.
	//
	// Managed by the server, ignored when passed to UpdatePrefixMetadata.
	UpdateTime *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	// Identity string of whoever modified the metadata the last time.
	//
	// Managed by the server, ignored when passed to UpdatePrefixMetadata.
	UpdateUser string `protobuf:"bytes,4,opt,name=update_user,json=updateUser" json:"update_user,omitempty"`
	// ACLs that apply to this prefix and all subprefixes, as a mapping from
	// a role to a list of users and groups that have it.
	Acls []*PrefixMetadata_ACL `protobuf:"bytes,5,rep,name=acls" json:"acls,omitempty"`
}

func (m *PrefixMetadata) Reset()                    { *m = PrefixMetadata{} }
func (m *PrefixMetadata) String() string            { return proto.CompactTextString(m) }
func (*PrefixMetadata) ProtoMessage()               {}
func (*PrefixMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PrefixMetadata) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *PrefixMetadata) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *PrefixMetadata) GetUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *PrefixMetadata) GetUpdateUser() string {
	if m != nil {
		return m.UpdateUser
	}
	return ""
}

func (m *PrefixMetadata) GetAcls() []*PrefixMetadata_ACL {
	if m != nil {
		return m.Acls
	}
	return nil
}

type PrefixMetadata_ACL struct {
	// Role that this ACL describes.
	Role Role `protobuf:"varint,1,opt,name=role,enum=cipd.Role" json:"role,omitempty"`
	// Users and groups that have the specified role.
	//
	// Each entry has a form "<kind>:<value>", e.g "group:..." or "user:...".
	Principals []string `protobuf:"bytes,2,rep,name=principals" json:"principals,omitempty"`
}

func (m *PrefixMetadata_ACL) Reset()                    { *m = PrefixMetadata_ACL{} }
func (m *PrefixMetadata_ACL) String() string            { return proto.CompactTextString(m) }
func (*PrefixMetadata_ACL) ProtoMessage()               {}
func (*PrefixMetadata_ACL) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *PrefixMetadata_ACL) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (m *PrefixMetadata_ACL) GetPrincipals() []string {
	if m != nil {
		return m.Principals
	}
	return nil
}

type InheritedPrefixMetadata struct {
	// Per-prefix metadata that applies to a prefix, ordered by prefix length.
	//
	// For example, when requesting metadata for prefix "a/b/c/d" the reply may
	// contain entries for "a", "a/b", "a/b/c/d" (in that order, with "a/b/c"
	// skipped in this example as not having any metadata attached).
	PerPrefixMetadata []*PrefixMetadata `protobuf:"bytes,1,rep,name=per_prefix_metadata,json=perPrefixMetadata" json:"per_prefix_metadata,omitempty"`
}

func (m *InheritedPrefixMetadata) Reset()                    { *m = InheritedPrefixMetadata{} }
func (m *InheritedPrefixMetadata) String() string            { return proto.CompactTextString(m) }
func (*InheritedPrefixMetadata) ProtoMessage()               {}
func (*InheritedPrefixMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *InheritedPrefixMetadata) GetPerPrefixMetadata() []*PrefixMetadata {
	if m != nil {
		return m.PerPrefixMetadata
	}
	return nil
}

func init() {
	proto.RegisterType((*PrefixRequest)(nil), "cipd.PrefixRequest")
	proto.RegisterType((*PrefixMetadata)(nil), "cipd.PrefixMetadata")
	proto.RegisterType((*PrefixMetadata_ACL)(nil), "cipd.PrefixMetadata.ACL")
	proto.RegisterType((*InheritedPrefixMetadata)(nil), "cipd.InheritedPrefixMetadata")
	proto.RegisterEnum("cipd.Role", Role_name, Role_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Repository service

type RepositoryClient interface {
	// Returns metadata associated with the given prefix.
	//
	// Requires the caller to have OWNER role for the requested prefix or any of
	// parent prefixes, otherwise the call fails with PERMISSION_DENIED error.
	//
	// If the caller has OWNER permission in any of parent prefixes, but the
	// requested prefix has no metadata associated with it, the call fails with
	// NOT_FOUND error.
	GetPrefixMetadata(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*PrefixMetadata, error)
	// Returns metadata associated with the given prefix and all parent prefixes.
	//
	// Requires the caller to have OWNER role for the requested prefix or any of
	// parent prefixes, otherwise the call fails with PERMISSION_DENIED error.
	//
	// Note that if the caller has permission to see the metadata for the
	// requested prefix, they will also see metadata for all parent prefixes,
	// since it is needed to assemble the final metadata for the prefix (it
	// includes inherited properties from all parent prefixes).
	GetInheritedPrefixMetadata(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*InheritedPrefixMetadata, error)
	// Updates or creates metadata associated with the given prefix.
	//
	// Requires the caller to have OWNER role for the requested prefix or any of
	// parent prefixes, otherwise the call fails with PERMISSION_DENIED error.
	//
	// This method checks 'fingerprint' field of the PrefixMetadata object. If the
	// metadata for the given prefix already exists, and the fingerprint in the
	// request doesn't match the current fingerprint, the request fails with
	// FAILED_PRECONDITION error.
	//
	// If the metadata doesn't exist yet, its fingerprint is assumed to be empty
	// string. So pass empty fingerprint when creating initial metadata objects.
	//
	// If the caller passes empty fingerprint, but the metadata already exists,
	// the request fails with ALREADY_EXISTS error.
	//
	// Note that there's no way to delete metadata once it was created. Passing
	// empty PrefixMetadata object is the best that can be done.
	//
	// On success returns PrefixMetadata object with the updated fingerprint.
	UpdatePrefixMetadata(ctx context.Context, in *PrefixMetadata, opts ...grpc.CallOption) (*PrefixMetadata, error)
}
type repositoryPRPCClient struct {
	client *prpc.Client
}

func NewRepositoryPRPCClient(client *prpc.Client) RepositoryClient {
	return &repositoryPRPCClient{client}
}

func (c *repositoryPRPCClient) GetPrefixMetadata(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*PrefixMetadata, error) {
	out := new(PrefixMetadata)
	err := c.client.Call(ctx, "cipd.Repository", "GetPrefixMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryPRPCClient) GetInheritedPrefixMetadata(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*InheritedPrefixMetadata, error) {
	out := new(InheritedPrefixMetadata)
	err := c.client.Call(ctx, "cipd.Repository", "GetInheritedPrefixMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryPRPCClient) UpdatePrefixMetadata(ctx context.Context, in *PrefixMetadata, opts ...grpc.CallOption) (*PrefixMetadata, error) {
	out := new(PrefixMetadata)
	err := c.client.Call(ctx, "cipd.Repository", "UpdatePrefixMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type repositoryClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryClient(cc *grpc.ClientConn) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) GetPrefixMetadata(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*PrefixMetadata, error) {
	out := new(PrefixMetadata)
	err := grpc.Invoke(ctx, "/cipd.Repository/GetPrefixMetadata", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) GetInheritedPrefixMetadata(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*InheritedPrefixMetadata, error) {
	out := new(InheritedPrefixMetadata)
	err := grpc.Invoke(ctx, "/cipd.Repository/GetInheritedPrefixMetadata", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) UpdatePrefixMetadata(ctx context.Context, in *PrefixMetadata, opts ...grpc.CallOption) (*PrefixMetadata, error) {
	out := new(PrefixMetadata)
	err := grpc.Invoke(ctx, "/cipd.Repository/UpdatePrefixMetadata", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Repository service

type RepositoryServer interface {
	// Returns metadata associated with the given prefix.
	//
	// Requires the caller to have OWNER role for the requested prefix or any of
	// parent prefixes, otherwise the call fails with PERMISSION_DENIED error.
	//
	// If the caller has OWNER permission in any of parent prefixes, but the
	// requested prefix has no metadata associated with it, the call fails with
	// NOT_FOUND error.
	GetPrefixMetadata(context.Context, *PrefixRequest) (*PrefixMetadata, error)
	// Returns metadata associated with the given prefix and all parent prefixes.
	//
	// Requires the caller to have OWNER role for the requested prefix or any of
	// parent prefixes, otherwise the call fails with PERMISSION_DENIED error.
	//
	// Note that if the caller has permission to see the metadata for the
	// requested prefix, they will also see metadata for all parent prefixes,
	// since it is needed to assemble the final metadata for the prefix (it
	// includes inherited properties from all parent prefixes).
	GetInheritedPrefixMetadata(context.Context, *PrefixRequest) (*InheritedPrefixMetadata, error)
	// Updates or creates metadata associated with the given prefix.
	//
	// Requires the caller to have OWNER role for the requested prefix or any of
	// parent prefixes, otherwise the call fails with PERMISSION_DENIED error.
	//
	// This method checks 'fingerprint' field of the PrefixMetadata object. If the
	// metadata for the given prefix already exists, and the fingerprint in the
	// request doesn't match the current fingerprint, the request fails with
	// FAILED_PRECONDITION error.
	//
	// If the metadata doesn't exist yet, its fingerprint is assumed to be empty
	// string. So pass empty fingerprint when creating initial metadata objects.
	//
	// If the caller passes empty fingerprint, but the metadata already exists,
	// the request fails with ALREADY_EXISTS error.
	//
	// Note that there's no way to delete metadata once it was created. Passing
	// empty PrefixMetadata object is the best that can be done.
	//
	// On success returns PrefixMetadata object with the updated fingerprint.
	UpdatePrefixMetadata(context.Context, *PrefixMetadata) (*PrefixMetadata, error)
}

func RegisterRepositoryServer(s prpc.Registrar, srv RepositoryServer) {
	s.RegisterService(&_Repository_serviceDesc, srv)
}

func _Repository_GetPrefixMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).GetPrefixMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Repository/GetPrefixMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).GetPrefixMetadata(ctx, req.(*PrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_GetInheritedPrefixMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).GetInheritedPrefixMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Repository/GetInheritedPrefixMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).GetInheritedPrefixMetadata(ctx, req.(*PrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_UpdatePrefixMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefixMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).UpdatePrefixMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Repository/UpdatePrefixMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).UpdatePrefixMetadata(ctx, req.(*PrefixMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

var _Repository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cipd.Repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrefixMetadata",
			Handler:    _Repository_GetPrefixMetadata_Handler,
		},
		{
			MethodName: "GetInheritedPrefixMetadata",
			Handler:    _Repository_GetInheritedPrefixMetadata_Handler,
		},
		{
			MethodName: "UpdatePrefixMetadata",
			Handler:    _Repository_UpdatePrefixMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/cipd/api/cipd/v1/repo.proto",
}

func init() { proto.RegisterFile("go.chromium.org/luci/cipd/api/cipd/v1/repo.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x86, 0xe7, 0xd8, 0x09, 0xf4, 0x84, 0x15, 0x57, 0x0d, 0x9b, 0x31, 0xac, 0x35, 0xb9, 0x59,
	0x18, 0xc3, 0xde, 0xb2, 0xcb, 0xc1, 0x46, 0x9a, 0x78, 0x25, 0x90, 0xb5, 0x41, 0x6b, 0x28, 0xec,
	0xc6, 0xb8, 0xce, 0x89, 0x2b, 0xb0, 0x23, 0x4d, 0x96, 0xc7, 0xf6, 0x30, 0x7b, 0xb9, 0x3d, 0xc9,
	0xb0, 0xec, 0xae, 0x4d, 0x70, 0xee, 0xe4, 0xcf, 0xff, 0x91, 0x3e, 0x1d, 0x1d, 0x78, 0x97, 0x72,
	0x3f, 0xb9, 0x97, 0x3c, 0x67, 0x65, 0xee, 0x73, 0x99, 0x06, 0x59, 0x99, 0xb0, 0x20, 0x61, 0x62,
	0x1d, 0xc4, 0xa2, 0x59, 0xfc, 0x7c, 0x1f, 0x48, 0x14, 0xdc, 0x17, 0x92, 0x2b, 0x4e, 0xac, 0x8a,
	0xb9, 0xe7, 0x29, 0xe7, 0x69, 0x86, 0x81, 0x66, 0x77, 0xe5, 0x26, 0x50, 0x2c, 0xc7, 0x42, 0xc5,
	0xb9, 0xa8, 0x63, 0xc3, 0xd7, 0xf0, 0x7c, 0x29, 0x71, 0xc3, 0x7e, 0x51, 0xfc, 0x51, 0x62, 0xa1,
	0xc8, 0x0b, 0xe8, 0x09, 0x0d, 0x1c, 0xc3, 0x33, 0x46, 0x47, 0xb4, 0xf9, 0x1a, 0xfe, 0xe9, 0xc0,
	0x71, 0x9d, 0xfc, 0x8a, 0x2a, 0x5e, 0xc7, 0x2a, 0x3e, 0x14, 0x25, 0x1e, 0xf4, 0x37, 0x6c, 0x9b,
	0xa2, 0x14, 0x92, 0x6d, 0x95, 0xd3, 0xd1, 0x3f, 0x9f, 0x22, 0xf2, 0x11, 0xfa, 0xa5, 0x58, 0xc7,
	0x0a, 0xa3, 0xca, 0xc7, 0x31, 0x3d, 0x63, 0xd4, 0x1f, 0xbb, 0x7e, 0x2d, 0xeb, 0x3f, 0xc8, 0xfa,
	0x37, 0x0f, 0xb2, 0x14, 0xea, 0x78, 0x05, 0xc8, 0xf9, 0xff, 0xe2, 0xb2, 0x40, 0xe9, 0x58, 0x7a,
	0xfb, 0x26, 0xb0, 0x2a, 0x50, 0x92, 0xb7, 0x60, 0xc5, 0x49, 0x56, 0x38, 0x5d, 0xcf, 0x1c, 0xf5,
	0xc7, 0x8e, 0x5f, 0x75, 0xc2, 0xdf, 0x75, 0xf7, 0x27, 0xd3, 0x05, 0xd5, 0x29, 0x37, 0x04, 0x73,
	0x32, 0x5d, 0x90, 0x33, 0xb0, 0x24, 0xcf, 0x50, 0x5f, 0xe5, 0x78, 0x0c, 0x75, 0x11, 0xe5, 0x19,
	0x52, 0xcd, 0xc9, 0x19, 0x40, 0xe5, 0x9e, 0x30, 0x11, 0x67, 0x85, 0xd3, 0xf1, 0xcc, 0xea, 0xd0,
	0x47, 0x32, 0x8c, 0xe0, 0xe5, 0x7c, 0x7b, 0x8f, 0x92, 0x29, 0x5c, 0xef, 0xf5, 0x69, 0x06, 0xa7,
	0x02, 0x65, 0x54, 0x77, 0x27, 0xca, 0x1b, 0xec, 0x18, 0x5a, 0x6f, 0xd0, 0xa6, 0x47, 0x4f, 0x04,
	0xca, 0x5d, 0xf4, 0xe6, 0x33, 0x58, 0x95, 0x0e, 0x19, 0x80, 0x4d, 0xaf, 0x17, 0x61, 0xb4, 0xba,
	0xfa, 0xb6, 0x0c, 0xa7, 0xf3, 0x2f, 0xf3, 0x70, 0x66, 0x3f, 0x23, 0x00, 0x3d, 0x1a, 0x4e, 0x66,
	0x21, 0xb5, 0x8d, 0x6a, 0x7d, 0x4b, 0xe7, 0x37, 0x21, 0xb5, 0x3b, 0xe4, 0x08, 0xba, 0xd7, 0xb7,
	0x57, 0x21, 0xb5, 0xcd, 0xf1, 0x5f, 0x03, 0x80, 0xa2, 0xe0, 0x05, 0x53, 0x5c, 0xfe, 0x26, 0x9f,
	0xe0, 0xe4, 0x12, 0xd5, 0x9e, 0xea, 0xe9, 0x53, 0x9b, 0x66, 0x24, 0xdc, 0x56, 0x45, 0xb2, 0x04,
	0xf7, 0x12, 0xd5, 0xa1, 0x3b, 0xb7, 0x6e, 0xf4, 0xaa, 0x86, 0x87, 0x6a, 0x2e, 0x60, 0xb0, 0xd2,
	0xaf, 0xb8, 0xc7, 0x5b, 0xcf, 0x6f, 0xb7, 0xba, 0xe8, 0x7e, 0x37, 0x63, 0xc1, 0xee, 0x7a, 0x7a,
	0x86, 0x3e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x84, 0x5c, 0x39, 0xa3, 0x38, 0x03, 0x00, 0x00,
}
