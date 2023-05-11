// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/core/transaction/pb/transaction.proto

package transactionpb

import (
	fmt "fmt"
	pb "github.com/dappley/go-dappley/core/transactionbase/pb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Transactions struct {
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Transactions) Reset()         { *m = Transactions{} }
func (m *Transactions) String() string { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()    {}
func (*Transactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_4138a4cf34c3b76a, []int{0}
}

func (m *Transactions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transactions.Unmarshal(m, b)
}
func (m *Transactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transactions.Marshal(b, m, deterministic)
}
func (m *Transactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transactions.Merge(m, src)
}
func (m *Transactions) XXX_Size() int {
	return xxx_messageInfo_Transactions.Size(m)
}
func (m *Transactions) XXX_DiscardUnknown() {
	xxx_messageInfo_Transactions.DiscardUnknown(m)
}

var xxx_messageInfo_Transactions proto.InternalMessageInfo

func (m *Transactions) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type NonceTransactions struct {
	Transactions         []*NonceTransaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NonceTransactions) Reset()         { *m = NonceTransactions{} }
func (m *NonceTransactions) String() string { return proto.CompactTextString(m) }
func (*NonceTransactions) ProtoMessage()    {}
func (*NonceTransactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_4138a4cf34c3b76a, []int{1}
}

func (m *NonceTransactions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonceTransactions.Unmarshal(m, b)
}
func (m *NonceTransactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonceTransactions.Marshal(b, m, deterministic)
}
func (m *NonceTransactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonceTransactions.Merge(m, src)
}
func (m *NonceTransactions) XXX_Size() int {
	return xxx_messageInfo_NonceTransactions.Size(m)
}
func (m *NonceTransactions) XXX_DiscardUnknown() {
	xxx_messageInfo_NonceTransactions.DiscardUnknown(m)
}

var xxx_messageInfo_NonceTransactions proto.InternalMessageInfo

func (m *NonceTransactions) GetTransactions() []*NonceTransaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type NonceTransaction struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Nonce                uint64       `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NonceTransaction) Reset()         { *m = NonceTransaction{} }
func (m *NonceTransaction) String() string { return proto.CompactTextString(m) }
func (*NonceTransaction) ProtoMessage()    {}
func (*NonceTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4138a4cf34c3b76a, []int{2}
}

func (m *NonceTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonceTransaction.Unmarshal(m, b)
}
func (m *NonceTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonceTransaction.Marshal(b, m, deterministic)
}
func (m *NonceTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonceTransaction.Merge(m, src)
}
func (m *NonceTransaction) XXX_Size() int {
	return xxx_messageInfo_NonceTransaction.Size(m)
}
func (m *NonceTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_NonceTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_NonceTransaction proto.InternalMessageInfo

func (m *NonceTransaction) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *NonceTransaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type Transaction struct {
	Id                   []byte         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vin                  []*pb.TXInput  `protobuf:"bytes,2,rep,name=vin,proto3" json:"vin,omitempty"`
	Vout                 []*pb.TXOutput `protobuf:"bytes,3,rep,name=vout,proto3" json:"vout,omitempty"`
	Tip                  []byte         `protobuf:"bytes,4,opt,name=tip,proto3" json:"tip,omitempty"`
	GasLimit             []byte         `protobuf:"bytes,5,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice             []byte         `protobuf:"bytes,6,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	Type                 int32          `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4138a4cf34c3b76a, []int{3}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Transaction) GetVin() []*pb.TXInput {
	if m != nil {
		return m.Vin
	}
	return nil
}

func (m *Transaction) GetVout() []*pb.TXOutput {
	if m != nil {
		return m.Vout
	}
	return nil
}

func (m *Transaction) GetTip() []byte {
	if m != nil {
		return m.Tip
	}
	return nil
}

func (m *Transaction) GetGasLimit() []byte {
	if m != nil {
		return m.GasLimit
	}
	return nil
}

func (m *Transaction) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *Transaction) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type TransactionNode struct {
	Children             map[string]*Transaction `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value                *Transaction            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Size                 int64                   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TransactionNode) Reset()         { *m = TransactionNode{} }
func (m *TransactionNode) String() string { return proto.CompactTextString(m) }
func (*TransactionNode) ProtoMessage()    {}
func (*TransactionNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_4138a4cf34c3b76a, []int{4}
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

func (m *TransactionNode) GetChildren() map[string]*Transaction {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *TransactionNode) GetValue() *Transaction {
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
	return fileDescriptor_4138a4cf34c3b76a, []int{5}
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
	proto.RegisterType((*Transactions)(nil), "transactionpb.Transactions")
	proto.RegisterType((*NonceTransactions)(nil), "transactionpb.NonceTransactions")
	proto.RegisterType((*NonceTransaction)(nil), "transactionpb.NonceTransaction")
	proto.RegisterType((*Transaction)(nil), "transactionpb.Transaction")
	proto.RegisterType((*TransactionNode)(nil), "transactionpb.TransactionNode")
	proto.RegisterMapType((map[string]*Transaction)(nil), "transactionpb.TransactionNode.ChildrenEntry")
	proto.RegisterType((*TransactionJournal)(nil), "transactionpb.TransactionJournal")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/core/transaction/pb/transaction.proto", fileDescriptor_4138a4cf34c3b76a)
}

var fileDescriptor_4138a4cf34c3b76a = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xe5, 0x64, 0xb7, 0xb4, 0xb3, 0x5b, 0x28, 0x16, 0x07, 0x6b, 0x7b, 0x20, 0xca, 0x69,
	0x0f, 0x25, 0x41, 0x70, 0x41, 0x08, 0x71, 0x60, 0x55, 0x89, 0x7f, 0x5a, 0x90, 0x85, 0x44, 0x6f,
	0xc8, 0x49, 0xcc, 0xd6, 0x22, 0xb5, 0xad, 0xd8, 0x59, 0x29, 0x3c, 0x2b, 0x4f, 0xc2, 0x09, 0xd9,
	0x69, 0x17, 0x27, 0x14, 0x68, 0x6f, 0x33, 0x9e, 0xdf, 0x7c, 0x33, 0x99, 0x4f, 0x81, 0xd5, 0x46,
	0xd8, 0xf3, 0xb6, 0xc8, 0x4a, 0x75, 0x91, 0x57, 0x4c, 0xeb, 0x9a, 0x77, 0xf9, 0x46, 0x3d, 0xba,
	0x0a, 0x4b, 0xd5, 0xf0, 0xdc, 0x36, 0x4c, 0x1a, 0x56, 0x5a, 0xa1, 0x64, 0xae, 0x8b, 0x30, 0xcd,
	0x74, 0xa3, 0xac, 0xc2, 0x87, 0xc1, 0x93, 0x2e, 0x16, 0xef, 0x6e, 0xa7, 0x59, 0x30, 0xc3, 0x47,
	0xba, 0xaf, 0x98, 0xe1, 0xbd, 0x76, 0xba, 0x86, 0xf9, 0xa7, 0xdf, 0x05, 0x83, 0x5f, 0xc2, 0x3c,
	0x00, 0x0d, 0x41, 0x49, 0xbc, 0x9c, 0x3d, 0x59, 0x64, 0x83, 0x15, 0xb2, 0xa0, 0x85, 0x0e, 0xf8,
	0xf4, 0x0c, 0xee, 0xaf, 0x95, 0x2c, 0xf9, 0x40, 0x74, 0x75, 0xad, 0xe8, 0xc3, 0x91, 0xe8, 0xb8,
	0x6f, 0xa4, 0xfc, 0x15, 0x8e, 0xc6, 0x04, 0x7e, 0x01, 0xb3, 0x80, 0x21, 0x28, 0x41, 0xff, 0x59,
	0x36, 0xc4, 0xf1, 0x03, 0x98, 0x4a, 0xa7, 0x48, 0xa2, 0x04, 0x2d, 0x27, 0xb4, 0x4f, 0xd2, 0x1f,
	0x08, 0x66, 0xe1, 0x8c, 0xbb, 0x10, 0x89, 0xca, 0x4b, 0xcf, 0x69, 0x24, 0x2a, 0x7c, 0x02, 0xf1,
	0x56, 0x48, 0x12, 0xfd, 0x79, 0x18, 0x77, 0x69, 0x37, 0xef, 0xec, 0x8d, 0xd4, 0xad, 0xa5, 0x0e,
	0xc3, 0x39, 0x4c, 0xb6, 0xaa, 0xb5, 0x24, 0xf6, 0xf8, 0xf1, 0xb5, 0xf8, 0x87, 0xd6, 0x3a, 0xde,
	0x83, 0xf8, 0x08, 0x62, 0x2b, 0x34, 0x99, 0xf8, 0x79, 0x2e, 0xc4, 0xc7, 0x70, 0xb0, 0x61, 0xe6,
	0x4b, 0x2d, 0x2e, 0x84, 0x25, 0x53, 0xff, 0xbe, 0xbf, 0x61, 0xe6, 0xbd, 0xcb, 0xaf, 0x8a, 0xba,
	0x11, 0x25, 0x27, 0x7b, 0xbb, 0xe2, 0x47, 0x97, 0x63, 0x0c, 0x13, 0xdb, 0x69, 0x4e, 0xee, 0x24,
	0x68, 0x39, 0xa5, 0x3e, 0x4e, 0x7f, 0x22, 0xb8, 0x17, 0x7c, 0xde, 0x5a, 0x55, 0x1c, 0xbf, 0x86,
	0xfd, 0xf2, 0x5c, 0xd4, 0x55, 0xc3, 0xe5, 0xa5, 0x37, 0x27, 0x7f, 0xbf, 0xa1, 0xeb, 0xc8, 0x56,
	0x97, 0xf8, 0xa9, 0xb4, 0x4d, 0x47, 0x77, 0xdd, 0xf8, 0x31, 0x4c, 0xb7, 0xac, 0x6e, 0xfb, 0x93,
	0xfe, 0xdb, 0x8a, 0x1e, 0x74, 0x3b, 0x1a, 0xf1, 0x9d, 0x93, 0x38, 0x41, 0xcb, 0x98, 0xfa, 0x78,
	0xf1, 0x19, 0x0e, 0x07, 0x03, 0xdc, 0x51, 0xbe, 0xf1, 0xce, 0x9b, 0x70, 0x40, 0x5d, 0x78, 0xfb,
	0x41, 0xcf, 0xa3, 0x67, 0x28, 0x3d, 0x05, 0x1c, 0x54, 0xde, 0xaa, 0xb6, 0x91, 0xac, 0xde, 0x79,
	0x84, 0x6e, 0xe8, 0x51, 0xb1, 0xe7, 0xff, 0x9d, 0xa7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0x86, 0xa8, 0xfd, 0xde, 0x03, 0x00, 0x00,
}
