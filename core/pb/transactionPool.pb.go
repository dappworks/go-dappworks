// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/core/pb/transactionPool.proto

package corepb

import (
	fmt "fmt"
	pb "github.com/dappley/go-dappley/core/transaction/pb"
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

type TransactionNode struct {
	Children             map[string]*pb.Transaction `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value                *pb.Transaction            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Size                 int64                      `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TransactionNode) Reset()         { *m = TransactionNode{} }
func (m *TransactionNode) String() string { return proto.CompactTextString(m) }
func (*TransactionNode) ProtoMessage()    {}
func (*TransactionNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_5febfe3e1645e55a, []int{0}
}

func (m *TransactionNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionNode.Unmarshal(m, b)
}
func (m *TransactionNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionNode.Marshal(b, m, deterministic)
}
func (m *TransactionNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionNode.Merge(m, src)
}
func (m *TransactionNode) XXX_Size() int {
	return xxx_messageInfo_TransactionNode.Size(m)
}
func (m *TransactionNode) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionNode.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionNode proto.InternalMessageInfo

func (m *TransactionNode) GetChildren() map[string]*pb.Transaction {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *TransactionNode) GetValue() *pb.Transaction {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TransactionNode) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type TransactionPool struct {
	Txs                  map[string]*TransactionNode `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TipOrder             []string                    `protobuf:"bytes,2,rep,name=tipOrder,proto3" json:"tipOrder,omitempty"`
	CurrSize             uint32                      `protobuf:"varint,3,opt,name=currSize,proto3" json:"currSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TransactionPool) Reset()         { *m = TransactionPool{} }
func (m *TransactionPool) String() string { return proto.CompactTextString(m) }
func (*TransactionPool) ProtoMessage()    {}
func (*TransactionPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_5febfe3e1645e55a, []int{1}
}

func (m *TransactionPool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionPool.Unmarshal(m, b)
}
func (m *TransactionPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionPool.Marshal(b, m, deterministic)
}
func (m *TransactionPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionPool.Merge(m, src)
}
func (m *TransactionPool) XXX_Size() int {
	return xxx_messageInfo_TransactionPool.Size(m)
}
func (m *TransactionPool) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionPool.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionPool proto.InternalMessageInfo

func (m *TransactionPool) GetTxs() map[string]*TransactionNode {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *TransactionPool) GetTipOrder() []string {
	if m != nil {
		return m.TipOrder
	}
	return nil
}

func (m *TransactionPool) GetCurrSize() uint32 {
	if m != nil {
		return m.CurrSize
	}
	return 0
}

func init() {
	proto.RegisterType((*TransactionNode)(nil), "corepb.TransactionNode")
	proto.RegisterMapType((map[string]*pb.Transaction)(nil), "corepb.TransactionNode.ChildrenEntry")
	proto.RegisterType((*TransactionPool)(nil), "corepb.TransactionPool")
	proto.RegisterMapType((map[string]*TransactionNode)(nil), "corepb.TransactionPool.TxsEntry")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/core/pb/transactionPool.proto", fileDescriptor_5febfe3e1645e55a)
}

var fileDescriptor_5febfe3e1645e55a = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xd9, 0xae, 0x96, 0x74, 0x4a, 0x51, 0xf6, 0x62, 0xc8, 0x69, 0x29, 0x08, 0xb9, 0x74,
	0x23, 0xf1, 0x22, 0x7a, 0x92, 0xe2, 0xd5, 0xca, 0x5a, 0xf0, 0x9c, 0x3f, 0x4b, 0x1b, 0x8c, 0xd9,
	0x65, 0xb3, 0x91, 0xc6, 0x87, 0xf4, 0x65, 0x7c, 0x01, 0x49, 0xd2, 0xc4, 0xa4, 0x54, 0xd0, 0xdb,
	0xcc, 0xce, 0xb7, 0xdf, 0xc7, 0x6f, 0x06, 0xee, 0x36, 0x89, 0xd9, 0x16, 0x21, 0x8b, 0xe4, 0x9b,
	0x17, 0x07, 0x4a, 0xa5, 0xa2, 0xf4, 0x36, 0x72, 0xd1, 0x96, 0x91, 0xd4, 0xc2, 0x53, 0xa1, 0x67,
	0x74, 0x90, 0xe5, 0x41, 0x64, 0x12, 0x99, 0x3d, 0x49, 0x99, 0x32, 0xa5, 0xa5, 0x91, 0x64, 0x5c,
	0x8d, 0x55, 0xe8, 0x2c, 0xff, 0x60, 0xd2, 0x73, 0x38, 0x30, 0x6c, 0xcc, 0xe6, 0x5f, 0x08, 0xce,
	0xd6, 0x3f, 0xaf, 0x8f, 0x32, 0x16, 0xe4, 0x1e, 0xac, 0x68, 0x9b, 0xa4, 0xb1, 0x16, 0x99, 0x8d,
	0x28, 0x76, 0xa7, 0xfe, 0x25, 0x6b, 0x32, 0xd9, 0x81, 0x94, 0x2d, 0xf7, 0xba, 0x87, 0xcc, 0xe8,
	0x92, 0x77, 0xdf, 0xc8, 0x15, 0x9c, 0xbe, 0x07, 0x69, 0x21, 0xec, 0x11, 0x45, 0xee, 0xd4, 0x77,
	0x58, 0x2f, 0x79, 0x68, 0xc3, 0x1b, 0x21, 0x21, 0x70, 0x92, 0x27, 0x1f, 0xc2, 0xc6, 0x14, 0xb9,
	0x98, 0xd7, 0xb5, 0xf3, 0x02, 0xb3, 0x41, 0x00, 0x39, 0x07, 0xfc, 0x2a, 0x4a, 0x1b, 0x51, 0xe4,
	0x4e, 0x78, 0x55, 0xfe, 0x3f, 0xe8, 0x76, 0x74, 0x83, 0xe6, 0x9f, 0x43, 0xea, 0x6a, 0xb9, 0xc4,
	0x07, 0x6c, 0x76, 0xf9, 0x1e, 0x98, 0x1e, 0x01, 0xae, 0x4f, 0xb0, 0xde, 0xe5, 0x0d, 0x6b, 0x25,
	0x26, 0x0e, 0x58, 0x26, 0x51, 0x2b, 0x1d, 0x0b, 0x6d, 0x8f, 0x28, 0x76, 0x27, 0xbc, 0xeb, 0xab,
	0x59, 0x54, 0x68, 0xfd, 0xdc, 0x42, 0xcd, 0x78, 0xd7, 0x3b, 0x2b, 0xb0, 0x5a, 0xa3, 0x23, 0x4c,
	0x8b, 0x21, 0xd3, 0xc5, 0x2f, 0xcb, 0xef, 0x01, 0x85, 0xe3, 0xfa, 0x9a, 0xd7, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x96, 0x5d, 0x0b, 0xee, 0x59, 0x02, 0x00, 0x00,
}
