// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/config/v1/vlans.proto

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "go.chromium.org/luci/machine-db/api/common/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A VLAN to store in the database.
type VLAN struct {
	// The ID of this VLAN. Must be unique.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// An alias for this VLAN.
	Alias string `protobuf:"bytes,2,opt,name=alias" json:"alias,omitempty"`
	// The block of IPv4 addresses belonging to this VLAN.
	CidrBlock string `protobuf:"bytes,3,opt,name=cidr_block,json=cidrBlock" json:"cidr_block,omitempty"`
	// The state of this VLAN.
	State common.State `protobuf:"varint,4,opt,name=state,enum=common.State" json:"state,omitempty"`
}

func (m *VLAN) Reset()                    { *m = VLAN{} }
func (m *VLAN) String() string            { return proto.CompactTextString(m) }
func (*VLAN) ProtoMessage()               {}
func (*VLAN) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *VLAN) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VLAN) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *VLAN) GetCidrBlock() string {
	if m != nil {
		return m.CidrBlock
	}
	return ""
}

func (m *VLAN) GetState() common.State {
	if m != nil {
		return m.State
	}
	return common.State_STATE_UNSPECIFIED
}

// A list of VLANs.
type VLANs struct {
	// A list of VLANs.
	Vlan []*VLAN `protobuf:"bytes,1,rep,name=vlan" json:"vlan,omitempty"`
}

func (m *VLANs) Reset()                    { *m = VLANs{} }
func (m *VLANs) String() string            { return proto.CompactTextString(m) }
func (*VLANs) ProtoMessage()               {}
func (*VLANs) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *VLANs) GetVlan() []*VLAN {
	if m != nil {
		return m.Vlan
	}
	return nil
}

func init() {
	proto.RegisterType((*VLAN)(nil), "config.VLAN")
	proto.RegisterType((*VLANs)(nil), "config.VLANs")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/config/v1/vlans.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0x47, 0xc9, 0xfe, 0x39, 0xb8, 0x51, 0xaf, 0x08, 0x16, 0x41, 0x10, 0xc2, 0xd9, 0xac, 0x85,
	0x09, 0x9e, 0x95, 0x76, 0x5a, 0x8b, 0xc5, 0x0a, 0xb6, 0x92, 0xcd, 0x9e, 0x7b, 0x83, 0x9b, 0x9d,
	0x25, 0xd9, 0xbb, 0xcf, 0x2f, 0x49, 0xfc, 0x00, 0x57, 0xe6, 0x3d, 0x78, 0x93, 0x1f, 0x3c, 0x0f,
	0xa4, 0xec, 0xc1, 0x93, 0xc3, 0xa3, 0x53, 0xe4, 0x07, 0x3d, 0x1e, 0x2d, 0x6a, 0x67, 0xec, 0x01,
	0xa7, 0xfd, 0x43, 0xdf, 0x69, 0x33, 0xa3, 0xb6, 0x34, 0xfd, 0xe0, 0xa0, 0x4f, 0x8f, 0xfa, 0x34,
	0x9a, 0x29, 0xa8, 0xd9, 0xd3, 0x42, 0x7c, 0x95, 0xf1, 0xcd, 0xcb, 0x79, 0x09, 0xe7, 0x68, 0x8a,
	0x89, 0xb0, 0x98, 0x65, 0xff, 0xdf, 0xd8, 0xce, 0x50, 0x7d, 0xbd, 0xbf, 0x7e, 0xf0, 0x0d, 0x14,
	0xd8, 0x0b, 0x26, 0x59, 0x53, 0xb6, 0x05, 0xf6, 0xfc, 0x1a, 0x6a, 0x33, 0xa2, 0x09, 0xa2, 0x90,
	0xac, 0x59, 0xb7, 0xf9, 0xc1, 0x6f, 0x01, 0x2c, 0xf6, 0xfe, 0xbb, 0x1b, 0xc9, 0xfe, 0x8a, 0x32,
	0xa9, 0x75, 0x24, 0x6f, 0x11, 0xf0, 0x3b, 0xa8, 0x53, 0x5c, 0x54, 0x92, 0x35, 0x9b, 0xdd, 0x95,
	0xca, 0x47, 0xd5, 0x67, 0x84, 0x6d, 0x76, 0xdb, 0x7b, 0xa8, 0xe3, 0xc5, 0xc0, 0x25, 0x54, 0x71,
	0x8d, 0x60, 0xb2, 0x6c, 0x2e, 0x76, 0x97, 0x2a, 0xaf, 0x51, 0x51, 0xb6, 0xc9, 0x74, 0xab, 0xf4,
	0xc7, 0xa7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x2c, 0x11, 0x99, 0x24, 0x01, 0x00, 0x00,
}
