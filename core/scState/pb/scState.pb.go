// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/core/scState/pb/scState.proto

package scstatepb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type State struct {
	State                map[string]string `protobuf:"bytes,1,rep,name=state,proto3" json:"state,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3509d4354c0467, []int{0}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetState() map[string]string {
	if m != nil {
		return m.State
	}
	return nil
}

type ScState struct {
	States               map[string]*State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ScState) Reset()         { *m = ScState{} }
func (m *ScState) String() string { return proto.CompactTextString(m) }
func (*ScState) ProtoMessage()    {}
func (*ScState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3509d4354c0467, []int{1}
}

func (m *ScState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScState.Unmarshal(m, b)
}
func (m *ScState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScState.Marshal(b, m, deterministic)
}
func (m *ScState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScState.Merge(m, src)
}
func (m *ScState) XXX_Size() int {
	return xxx_messageInfo_ScState.Size(m)
}
func (m *ScState) XXX_DiscardUnknown() {
	xxx_messageInfo_ScState.DiscardUnknown(m)
}

var xxx_messageInfo_ScState proto.InternalMessageInfo

func (m *ScState) GetStates() map[string]*State {
	if m != nil {
		return m.States
	}
	return nil
}

type Log struct {
	Log                  map[string]string `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3509d4354c0467, []int{2}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetLog() map[string]string {
	if m != nil {
		return m.Log
	}
	return nil
}

type ChangeLog struct {
	Log                  map[string]*Log `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ChangeLog) Reset()         { *m = ChangeLog{} }
func (m *ChangeLog) String() string { return proto.CompactTextString(m) }
func (*ChangeLog) ProtoMessage()    {}
func (*ChangeLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3509d4354c0467, []int{3}
}

func (m *ChangeLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeLog.Unmarshal(m, b)
}
func (m *ChangeLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeLog.Marshal(b, m, deterministic)
}
func (m *ChangeLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeLog.Merge(m, src)
}
func (m *ChangeLog) XXX_Size() int {
	return xxx_messageInfo_ChangeLog.Size(m)
}
func (m *ChangeLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeLog.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeLog proto.InternalMessageInfo

func (m *ChangeLog) GetLog() map[string]*Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func init() {
	proto.RegisterType((*State)(nil), "scstatepb.State")
	proto.RegisterMapType((map[string]string)(nil), "scstatepb.State.StateEntry")
	proto.RegisterType((*ScState)(nil), "scstatepb.ScState")
	proto.RegisterMapType((map[string]*State)(nil), "scstatepb.ScState.StatesEntry")
	proto.RegisterType((*Log)(nil), "scstatepb.Log")
	proto.RegisterMapType((map[string]string)(nil), "scstatepb.Log.LogEntry")
	proto.RegisterType((*ChangeLog)(nil), "scstatepb.ChangeLog")
	proto.RegisterMapType((map[string]*Log)(nil), "scstatepb.ChangeLog.LogEntry")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/core/scState/pb/scState.proto", fileDescriptor_1f3509d4354c0467)
}

var fileDescriptor_1f3509d4354c0467 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x49, 0x2c, 0x28, 0xc8, 0x49, 0xad, 0xd4, 0x4f,
	0xcf, 0xd7, 0x85, 0x31, 0x93, 0xf3, 0x8b, 0x52, 0xf5, 0x8b, 0x93, 0x83, 0x4b, 0x12, 0x4b, 0x52,
	0xf5, 0x0b, 0x92, 0x60, 0x4c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xce, 0xe2, 0xe4, 0x62,
	0x10, 0xb7, 0x20, 0x49, 0xa9, 0x84, 0x8b, 0x15, 0x2c, 0x23, 0x64, 0xc8, 0xc5, 0x0a, 0x16, 0x93,
	0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd6, 0x83, 0xab, 0xd1, 0x83, 0x68, 0x05, 0x93, 0xae,
	0x79, 0x25, 0x45, 0x95, 0x41, 0x10, 0x95, 0x52, 0x16, 0x5c, 0x5c, 0x08, 0x41, 0x21, 0x01, 0x2e,
	0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x84, 0x8b,
	0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0x2c, 0x06, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x2a,
	0xf5, 0x31, 0x72, 0xb1, 0x07, 0x43, 0x9c, 0x24, 0x64, 0xc6, 0xc5, 0x06, 0x36, 0xae, 0x18, 0x6a,
	0xb3, 0x1c, 0xb2, 0xcd, 0xc9, 0x48, 0x76, 0x17, 0x43, 0x2c, 0x87, 0xaa, 0x96, 0xf2, 0xe6, 0xe2,
	0x46, 0x12, 0xc6, 0x62, 0xbd, 0x1a, 0xb2, 0xf5, 0xdc, 0x46, 0x02, 0xe8, 0x3e, 0x42, 0x76, 0x50,
	0x06, 0x17, 0xb3, 0x4f, 0x7e, 0xba, 0x90, 0x26, 0x17, 0x73, 0x4e, 0x7e, 0x3a, 0xd4, 0x21, 0xe2,
	0x48, 0x1a, 0x7c, 0xf2, 0xd3, 0x41, 0x18, 0xe2, 0x02, 0x90, 0x1a, 0x29, 0x33, 0x2e, 0x0e, 0x98,
	0x00, 0x49, 0x5e, 0x6f, 0x61, 0xe4, 0xe2, 0x74, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x05, 0x59, 0xa8,
	0x8f, 0x6c, 0xa1, 0x2c, 0x92, 0x85, 0x70, 0x25, 0x68, 0xd6, 0xba, 0xe1, 0xb5, 0x56, 0x05, 0xd5,
	0xcb, 0x7c, 0xa8, 0x3e, 0x40, 0x72, 0x46, 0x12, 0x1b, 0x38, 0x25, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xbc, 0x50, 0x96, 0x17, 0x48, 0x02, 0x00, 0x00,
}
