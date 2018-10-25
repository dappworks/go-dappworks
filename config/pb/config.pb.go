// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/config/pb/config.proto

package configpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	ConsensusConfig      *ConsensusConfig `protobuf:"bytes,1,opt,name=consensusConfig,proto3" json:"consensusConfig,omitempty"`
	NodeConfig           *NodeConfig      `protobuf:"bytes,2,opt,name=nodeConfig,proto3" json:"nodeConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_b25824abde06a062, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetConsensusConfig() *ConsensusConfig {
	if m != nil {
		return m.ConsensusConfig
	}
	return nil
}

func (m *Config) GetNodeConfig() *NodeConfig {
	if m != nil {
		return m.NodeConfig
	}
	return nil
}

type ConsensusConfig struct {
	MinerAddr            string   `protobuf:"bytes,1,opt,name=minerAddr,proto3" json:"minerAddr,omitempty"`
	PrivKey              string   `protobuf:"bytes,2,opt,name=privKey,proto3" json:"privKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusConfig) Reset()         { *m = ConsensusConfig{} }
func (m *ConsensusConfig) String() string { return proto.CompactTextString(m) }
func (*ConsensusConfig) ProtoMessage()    {}
func (*ConsensusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_b25824abde06a062, []int{1}
}
func (m *ConsensusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusConfig.Unmarshal(m, b)
}
func (m *ConsensusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusConfig.Marshal(b, m, deterministic)
}
func (dst *ConsensusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusConfig.Merge(dst, src)
}
func (m *ConsensusConfig) XXX_Size() int {
	return xxx_messageInfo_ConsensusConfig.Size(m)
}
func (m *ConsensusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusConfig proto.InternalMessageInfo

func (m *ConsensusConfig) GetMinerAddr() string {
	if m != nil {
		return m.MinerAddr
	}
	return ""
}

func (m *ConsensusConfig) GetPrivKey() string {
	if m != nil {
		return m.PrivKey
	}
	return ""
}

type NodeConfig struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Seed                 string   `protobuf:"bytes,2,opt,name=seed,proto3" json:"seed,omitempty"`
	DbPath               string   `protobuf:"bytes,3,opt,name=dbPath,proto3" json:"dbPath,omitempty"`
	RpcPort              uint32   `protobuf:"varint,4,opt,name=rpcPort,proto3" json:"rpcPort,omitempty"`
	KeyPath              string   `protobuf:"bytes,5,opt,name=keyPath,proto3" json:"keyPath,omitempty"`
	TxPoolLimit          uint32   `protobuf:"varint,6,opt,name=txPoolLimit,proto3" json:"txPoolLimit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}
func (*NodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_b25824abde06a062, []int{2}
}
func (m *NodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig.Unmarshal(m, b)
}
func (m *NodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig.Marshal(b, m, deterministic)
}
func (dst *NodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig.Merge(dst, src)
}
func (m *NodeConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig.Size(m)
}
func (m *NodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig proto.InternalMessageInfo

func (m *NodeConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *NodeConfig) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *NodeConfig) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *NodeConfig) GetRpcPort() uint32 {
	if m != nil {
		return m.RpcPort
	}
	return 0
}

func (m *NodeConfig) GetKeyPath() string {
	if m != nil {
		return m.KeyPath
	}
	return ""
}

func (m *NodeConfig) GetTxPoolLimit() uint32 {
	if m != nil {
		return m.TxPoolLimit
	}
	return 0
}

type DynastyConfig struct {
	Producers            []string `protobuf:"bytes,1,rep,name=producers,proto3" json:"producers,omitempty"`
	MaxProducers         uint32   `protobuf:"varint,2,opt,name=maxProducers,proto3" json:"maxProducers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DynastyConfig) Reset()         { *m = DynastyConfig{} }
func (m *DynastyConfig) String() string { return proto.CompactTextString(m) }
func (*DynastyConfig) ProtoMessage()    {}
func (*DynastyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_b25824abde06a062, []int{3}
}
func (m *DynastyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynastyConfig.Unmarshal(m, b)
}
func (m *DynastyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynastyConfig.Marshal(b, m, deterministic)
}
func (dst *DynastyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynastyConfig.Merge(dst, src)
}
func (m *DynastyConfig) XXX_Size() int {
	return xxx_messageInfo_DynastyConfig.Size(m)
}
func (m *DynastyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynastyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynastyConfig proto.InternalMessageInfo

func (m *DynastyConfig) GetProducers() []string {
	if m != nil {
		return m.Producers
	}
	return nil
}

func (m *DynastyConfig) GetMaxProducers() uint32 {
	if m != nil {
		return m.MaxProducers
	}
	return 0
}

type CliConfig struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CliConfig) Reset()         { *m = CliConfig{} }
func (m *CliConfig) String() string { return proto.CompactTextString(m) }
func (*CliConfig) ProtoMessage()    {}
func (*CliConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_b25824abde06a062, []int{4}
}
func (m *CliConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CliConfig.Unmarshal(m, b)
}
func (m *CliConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CliConfig.Marshal(b, m, deterministic)
}
func (dst *CliConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CliConfig.Merge(dst, src)
}
func (m *CliConfig) XXX_Size() int {
	return xxx_messageInfo_CliConfig.Size(m)
}
func (m *CliConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CliConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CliConfig proto.InternalMessageInfo

func (m *CliConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *CliConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "configpb.Config")
	proto.RegisterType((*ConsensusConfig)(nil), "configpb.ConsensusConfig")
	proto.RegisterType((*NodeConfig)(nil), "configpb.NodeConfig")
	proto.RegisterType((*DynastyConfig)(nil), "configpb.DynastyConfig")
	proto.RegisterType((*CliConfig)(nil), "configpb.CliConfig")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/config/pb/config.proto", fileDescriptor_config_b25824abde06a062)
}

var fileDescriptor_config_b25824abde06a062 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x53, 0xc0, 0x4a, 0x07, 0x09, 0xc9, 0xc6, 0x98, 0x6a, 0x3c, 0x90, 0x3d, 0x79, 0xb1,
	0x24, 0xca, 0xcd, 0x93, 0xa9, 0x17, 0xa3, 0x31, 0xb5, 0x6f, 0xd0, 0x76, 0x57, 0xd8, 0x48, 0xbb,
	0x9b, 0xdd, 0x45, 0xe9, 0xd9, 0x37, 0xf1, 0x49, 0xa5, 0x43, 0xff, 0x00, 0x07, 0x6f, 0xf3, 0xfd,
	0x76, 0xbe, 0x6f, 0x66, 0xd2, 0xc2, 0x7c, 0x21, 0xec, 0x72, 0x9d, 0x06, 0x99, 0xcc, 0x67, 0x2c,
	0x51, 0x6a, 0xc5, 0xcb, 0xd9, 0x42, 0xde, 0x36, 0x65, 0x26, 0x8b, 0x0f, 0xb1, 0x98, 0xa9, 0xb4,
	0xae, 0x02, 0xa5, 0xa5, 0x95, 0x64, 0xb8, 0x53, 0x2a, 0xa5, 0x3f, 0x0e, 0xb8, 0x21, 0x0a, 0x12,
	0xc2, 0x64, 0x8b, 0x0d, 0x2f, 0xcc, 0xda, 0xec, 0x90, 0xef, 0x4c, 0x9d, 0x9b, 0xd1, 0xdd, 0x65,
	0xd0, 0xb4, 0x07, 0xe1, 0x61, 0x43, 0x7c, 0xec, 0x20, 0x73, 0x80, 0x42, 0x32, 0x5e, 0xfb, 0x7b,
	0xe8, 0x3f, 0xef, 0xfc, 0x6f, 0xed, 0x5b, 0xbc, 0xd7, 0x47, 0x9f, 0x61, 0x72, 0x94, 0x4c, 0xae,
	0xc1, 0xcb, 0x45, 0xc1, 0xf5, 0x23, 0x63, 0x1a, 0xf7, 0xf0, 0xe2, 0x0e, 0x10, 0x1f, 0x4e, 0x95,
	0x16, 0x5f, 0x2f, 0xbc, 0xc4, 0x19, 0x5e, 0xdc, 0x48, 0xfa, 0xeb, 0x00, 0x74, 0x53, 0x08, 0x81,
	0x81, 0x92, 0xda, 0x62, 0xc2, 0x38, 0xc6, 0xba, 0x62, 0x86, 0x73, 0x56, 0x3b, 0xb1, 0x26, 0x17,
	0xe0, 0xb2, 0x34, 0x4a, 0xec, 0xd2, 0xef, 0x23, 0xad, 0x55, 0x35, 0x48, 0xab, 0x2c, 0xaa, 0x22,
	0x06, 0x18, 0xd1, 0xc8, 0xea, 0xe5, 0x93, 0x97, 0x68, 0x39, 0xd9, 0xad, 0x50, 0x4b, 0x32, 0x85,
	0x91, 0xdd, 0x44, 0x52, 0xae, 0x5e, 0x45, 0x2e, 0xac, 0xef, 0xa2, 0x6f, 0x1f, 0xd1, 0x77, 0x18,
	0x3f, 0x95, 0x45, 0x62, 0x6c, 0xd9, 0x5d, 0xbb, 0xfd, 0x32, 0x6c, 0x9d, 0x71, 0x6d, 0xb6, 0xbb,
	0xf6, 0xab, 0x6b, 0x5b, 0x40, 0x28, 0x9c, 0xe5, 0xc9, 0x26, 0x6a, 0x1b, 0x7a, 0x98, 0x78, 0xc0,
	0xe8, 0x03, 0x78, 0xe1, 0x4a, 0xfc, 0x73, 0xf5, 0x15, 0x0c, 0x55, 0x62, 0xcc, 0xb7, 0xd4, 0xcd,
	0xe5, 0xad, 0x4e, 0x5d, 0xfc, 0x2d, 0xee, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x43, 0x9d, 0x10,
	0xc1, 0x4e, 0x02, 0x00, 0x00,
}
