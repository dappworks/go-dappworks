// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/core/pb/transactionJournal.proto

package corepb

import (
	fmt "fmt"
	pb "github.com/dappley/go-dappley/core/transaction_base/pb"
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

type TransactionJournal struct {
	Vout                 []*pb.TXOutput `protobuf:"bytes,1,rep,name=vout,proto3" json:"vout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransactionJournal) Reset()         { *m = TransactionJournal{} }
func (m *TransactionJournal) String() string { return proto.CompactTextString(m) }
func (*TransactionJournal) ProtoMessage()    {}
func (*TransactionJournal) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a777b3b39f59290, []int{0}
}

func (m *TransactionJournal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionJournal.Unmarshal(m, b)
}
func (m *TransactionJournal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionJournal.Marshal(b, m, deterministic)
}
func (m *TransactionJournal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionJournal.Merge(m, src)
}
func (m *TransactionJournal) XXX_Size() int {
	return xxx_messageInfo_TransactionJournal.Size(m)
}
func (m *TransactionJournal) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionJournal.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionJournal proto.InternalMessageInfo

func (m *TransactionJournal) GetVout() []*pb.TXOutput {
	if m != nil {
		return m.Vout
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionJournal)(nil), "corepb.TransactionJournal")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/core/pb/transactionJournal.proto", fileDescriptor_3a777b3b39f59290)
}

var fileDescriptor_3a777b3b39f59290 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x49, 0x2c, 0x28, 0xc8, 0x49, 0xad, 0xd4, 0x4f,
	0xcf, 0xd7, 0x85, 0x31, 0x93, 0xf3, 0x8b, 0x52, 0xf5, 0x0b, 0x92, 0xf4, 0x4b, 0x8a, 0x12, 0xf3,
	0x8a, 0x13, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0xbc, 0xf2, 0x4b, 0x8b, 0xf2, 0x12, 0x73, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0x40, 0x2a, 0x0a, 0x92, 0xa4, 0x7c, 0x88, 0x30, 0x07, 0xc9,
	0x90, 0xf8, 0xa4, 0xc4, 0x62, 0x74, 0x83, 0x9d, 0x12, 0x8b, 0x53, 0x21, 0xa6, 0x2a, 0xb9, 0x72,
	0x09, 0x85, 0x60, 0xd8, 0x28, 0xa4, 0xcf, 0xc5, 0x52, 0x96, 0x5f, 0x5a, 0x22, 0xc1, 0xa8, 0xc0,
	0xac, 0xc1, 0x6d, 0x24, 0xad, 0x87, 0xa4, 0x17, 0x64, 0x5c, 0x41, 0x92, 0x5e, 0x48, 0x84, 0x7f,
	0x69, 0x49, 0x41, 0x69, 0x49, 0x10, 0x58, 0x61, 0x12, 0x1b, 0xd8, 0x34, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe6, 0x1f, 0x45, 0x5b, 0xe5, 0x00, 0x00, 0x00,
}
