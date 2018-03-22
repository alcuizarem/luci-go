// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/datacenters.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "go.chromium.org/luci/machine-db/api/common/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A datacenter in the database.
type Datacenter struct {
	// The name of this datacenter. Uniquely identifies this datacenter.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this datacenter.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The state of this datacenter.
	State common.State `protobuf:"varint,3,opt,name=state,enum=common.State" json:"state,omitempty"`
}

func (m *Datacenter) Reset()                    { *m = Datacenter{} }
func (m *Datacenter) String() string            { return proto.CompactTextString(m) }
func (*Datacenter) ProtoMessage()               {}
func (*Datacenter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Datacenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Datacenter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Datacenter) GetState() common.State {
	if m != nil {
		return m.State
	}
	return common.State_STATE_UNSPECIFIED
}

// A request to list datacenters in the database.
type ListDatacentersRequest struct {
	// The names of datacenters to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *ListDatacentersRequest) Reset()                    { *m = ListDatacentersRequest{} }
func (m *ListDatacentersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDatacentersRequest) ProtoMessage()               {}
func (*ListDatacentersRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ListDatacentersRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// A response containing a list of datacenters in the database.
type ListDatacentersResponse struct {
	// The datacenters matching the request.
	Datacenters []*Datacenter `protobuf:"bytes,1,rep,name=datacenters" json:"datacenters,omitempty"`
}

func (m *ListDatacentersResponse) Reset()                    { *m = ListDatacentersResponse{} }
func (m *ListDatacentersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDatacentersResponse) ProtoMessage()               {}
func (*ListDatacentersResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ListDatacentersResponse) GetDatacenters() []*Datacenter {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

func init() {
	proto.RegisterType((*Datacenter)(nil), "crimson.Datacenter")
	proto.RegisterType((*ListDatacentersRequest)(nil), "crimson.ListDatacentersRequest")
	proto.RegisterType((*ListDatacentersResponse)(nil), "crimson.ListDatacentersResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/datacenters.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4e, 0xf4, 0x30,
	0x10, 0x84, 0x95, 0xff, 0xfe, 0x03, 0xdd, 0x46, 0x50, 0x18, 0x04, 0x11, 0x55, 0x14, 0x9a, 0x34,
	0xd8, 0xe2, 0x10, 0x0d, 0x15, 0x05, 0x25, 0x05, 0x32, 0x4f, 0xe0, 0x73, 0x56, 0x39, 0x4b, 0xd8,
	0x1b, 0xbc, 0x0e, 0xcf, 0x8f, 0xce, 0x89, 0xb8, 0x48, 0x34, 0x74, 0xf6, 0xce, 0xce, 0x37, 0xda,
	0x81, 0xe7, 0x9e, 0xa4, 0xdd, 0x47, 0xf2, 0x6e, 0xf4, 0x92, 0x62, 0xaf, 0x3e, 0x46, 0xeb, 0x94,
	0x37, 0x76, 0xef, 0x02, 0xde, 0x75, 0x3b, 0x65, 0x06, 0xa7, 0x6c, 0x74, 0x9e, 0x29, 0xa8, 0xaf,
	0x7b, 0xd5, 0x99, 0x64, 0x2c, 0x86, 0x84, 0x91, 0xe5, 0x10, 0x29, 0x91, 0x38, 0x9d, 0xd5, 0x9b,
	0xa7, 0x3f, 0xa1, 0xc8, 0xfb, 0x89, 0xc4, 0xc9, 0x24, 0x9c, 0x21, 0x4d, 0x0f, 0xf0, 0xf2, 0x43,
	0x16, 0x02, 0xfe, 0x07, 0xe3, 0xb1, 0x2a, 0xea, 0xa2, 0xdd, 0xe8, 0xfc, 0x16, 0x35, 0x94, 0x1d,
	0xb2, 0x8d, 0x6e, 0x48, 0x8e, 0x42, 0xf5, 0x2f, 0x4b, 0xcb, 0x91, 0xb8, 0x85, 0x75, 0x66, 0x56,
	0xab, 0xba, 0x68, 0xcf, 0xb7, 0x67, 0x72, 0xca, 0x92, 0xef, 0x87, 0xa1, 0x9e, 0xb4, 0x46, 0xc2,
	0xd5, 0xab, 0xe3, 0x74, 0x0c, 0x63, 0x8d, 0x9f, 0x23, 0x72, 0x12, 0x97, 0xb0, 0x3e, 0x04, 0x71,
	0x55, 0xd4, 0xab, 0x76, 0xa3, 0xa7, 0x4f, 0xf3, 0x06, 0xd7, 0xbf, 0xf6, 0x79, 0xa0, 0xc0, 0x28,
	0x1e, 0xa1, 0x5c, 0xb4, 0x91, 0x6d, 0xe5, 0xf6, 0x42, 0xce, 0x75, 0xc8, 0xa3, 0x45, 0x2f, 0xf7,
	0x76, 0x27, 0xf9, 0xe2, 0x87, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xc7, 0x46, 0x34, 0x7a,
	0x01, 0x00, 0x00,
}
