// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/dm/api/distributor/swarming/v1/isolate_ref.proto
// DO NOT EDIT!

package swarmingV1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IsolatedRef struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Server string `protobuf:"bytes,2,opt,name=server" json:"server,omitempty"`
}

func (m *IsolatedRef) Reset()                    { *m = IsolatedRef{} }
func (m *IsolatedRef) String() string            { return proto.CompactTextString(m) }
func (*IsolatedRef) ProtoMessage()               {}
func (*IsolatedRef) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterType((*IsolatedRef)(nil), "swarmingV1.IsolatedRef")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/dm/api/distributor/swarming/v1/isolate_ref.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xf2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x29, 0x4d, 0xce, 0x04, 0x13, 0xba, 0xe9, 0xf9,
	0xfa, 0x29, 0xb9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x29, 0x99, 0xc5, 0x25, 0x45, 0x99, 0x49, 0xa5,
	0x25, 0xf9, 0x45, 0xfa, 0xc5, 0xe5, 0x89, 0x45, 0xb9, 0x99, 0x79, 0xe9, 0xfa, 0x65, 0x86, 0xfa,
	0x99, 0xc5, 0xf9, 0x39, 0x89, 0x25, 0xa9, 0xf1, 0x45, 0xa9, 0x69, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x5c, 0x30, 0xe9, 0x30, 0x43, 0x25, 0x53, 0x2e, 0x6e, 0x4f, 0x88, 0x82, 0x94, 0xa0,
	0xd4, 0x34, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x20,
	0x4b, 0x48, 0x8c, 0x8b, 0xad, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x48, 0x82, 0x09, 0x2c, 0x06, 0xe5,
	0x25, 0xb1, 0x81, 0x4d, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x27, 0x70, 0xaa, 0x76, 0x95,
	0x00, 0x00, 0x00,
}
