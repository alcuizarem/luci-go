// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/internal/db.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/scheduler/appengine/internal/db.proto
	go.chromium.org/luci/scheduler/appengine/internal/timers.proto
	go.chromium.org/luci/scheduler/appengine/internal/tq.proto
	go.chromium.org/luci/scheduler/appengine/internal/triggers.proto

It has these top-level messages:
	FinishedInvocation
	FinishedInvocationList
	Timer
	TimerList
	ReadProjectConfigTask
	LaunchInvocationTask
	LaunchInvocationsBatchTask
	TriageJobStateTask
	KickTriageTask
	InvocationFinishedTask
	FanOutTriggersTask
	EnqueueTriggersTask
	ScheduleTimersTask
	TimerTask
	CronTickTask
	Trigger
	TriggerList
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// FinishedInvocation represents a recently finished invocation of a job.
//
// It is stored as part of Job entity inside FinishedInvocationsRaw field.
type FinishedInvocation struct {
	InvocationId int64                      `protobuf:"varint,1,opt,name=invocation_id,json=invocationId" json:"invocation_id,omitempty"`
	Finished     *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=finished" json:"finished,omitempty"`
}

func (m *FinishedInvocation) Reset()                    { *m = FinishedInvocation{} }
func (m *FinishedInvocation) String() string            { return proto.CompactTextString(m) }
func (*FinishedInvocation) ProtoMessage()               {}
func (*FinishedInvocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FinishedInvocation) GetInvocationId() int64 {
	if m != nil {
		return m.InvocationId
	}
	return 0
}

func (m *FinishedInvocation) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

// FinishedInvocationList is stored in Job entities as FinishedInvocationsRaw.
type FinishedInvocationList struct {
	Invocations []*FinishedInvocation `protobuf:"bytes,1,rep,name=invocations" json:"invocations,omitempty"`
}

func (m *FinishedInvocationList) Reset()                    { *m = FinishedInvocationList{} }
func (m *FinishedInvocationList) String() string            { return proto.CompactTextString(m) }
func (*FinishedInvocationList) ProtoMessage()               {}
func (*FinishedInvocationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FinishedInvocationList) GetInvocations() []*FinishedInvocation {
	if m != nil {
		return m.Invocations
	}
	return nil
}

func init() {
	proto.RegisterType((*FinishedInvocation)(nil), "internal.db.FinishedInvocation")
	proto.RegisterType((*FinishedInvocationList)(nil), "internal.db.FinishedInvocationList")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/internal/db.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcf, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x05, 0x29, 0x1b, 0xbd, 0xec, 0x41, 0x42, 0x2f, 0x0d, 0xf5, 0x92, 0xd3, 0x2c,
	0x54, 0xf0, 0xe0, 0x4d, 0x0f, 0x42, 0xc1, 0x53, 0xf0, 0xa4, 0x07, 0x49, 0xb2, 0xd3, 0xcd, 0x40,
	0xb2, 0x13, 0xf7, 0x87, 0x7f, 0xbf, 0xd0, 0x76, 0xdb, 0x42, 0x8e, 0xf3, 0xe6, 0x7d, 0x7c, 0x33,
	0xe2, 0xc5, 0x30, 0x74, 0xbd, 0xe3, 0x91, 0xe2, 0x08, 0xec, 0x8c, 0x1a, 0x62, 0x47, 0xca, 0x77,
	0x3d, 0xea, 0x38, 0xa0, 0x53, 0xcd, 0x34, 0xa1, 0x35, 0x64, 0x51, 0x91, 0x0d, 0xe8, 0x6c, 0x33,
	0x28, 0xdd, 0xc2, 0xe4, 0x38, 0xb0, 0xcc, 0x53, 0x04, 0xba, 0x5d, 0xad, 0x0d, 0xb3, 0x19, 0x50,
	0x1d, 0x56, 0x6d, 0xdc, 0xab, 0x40, 0x23, 0xfa, 0xd0, 0x8c, 0xd3, 0xb1, 0xbd, 0xf9, 0x15, 0xf2,
	0x9d, 0x2c, 0xf9, 0x1e, 0xf5, 0xce, 0xfe, 0x71, 0xd7, 0x04, 0x62, 0x2b, 0x1f, 0xc5, 0x3d, 0x9d,
	0xa7, 0x1f, 0xd2, 0x45, 0x56, 0x66, 0xd5, 0xa2, 0xbe, 0xbb, 0x84, 0x3b, 0x2d, 0x9f, 0xc5, 0x72,
	0x7f, 0x42, 0x8b, 0x9b, 0x32, 0xab, 0xf2, 0xed, 0x0a, 0x8e, 0x3a, 0x48, 0x3a, 0xf8, 0x4c, 0xba,
	0xfa, 0xdc, 0xdd, 0x7c, 0x8b, 0x87, 0xb9, 0xf2, 0x83, 0x7c, 0x90, 0xaf, 0x22, 0xbf, 0x18, 0x7c,
	0x91, 0x95, 0x8b, 0x2a, 0xdf, 0xae, 0xe1, 0xea, 0x21, 0x98, 0x93, 0xf5, 0x35, 0xf3, 0x26, 0xbe,
	0x96, 0xa9, 0xde, 0xde, 0x1e, 0xce, 0x78, 0xfa, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x83, 0x7c,
	0xda, 0x4e, 0x01, 0x00, 0x00,
}
