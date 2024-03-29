// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package api

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConfigureResponse_Status int32

const (
	ConfigureResponse_OK    ConfigureResponse_Status = 0
	ConfigureResponse_ERROR ConfigureResponse_Status = 1
)

var ConfigureResponse_Status_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}

var ConfigureResponse_Status_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x ConfigureResponse_Status) String() string {
	return proto.EnumName(ConfigureResponse_Status_name, int32(x))
}

func (ConfigureResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1, 0}
}

type MeasureResponse_Status int32

const (
	MeasureResponse_OK    MeasureResponse_Status = 0
	MeasureResponse_ERROR MeasureResponse_Status = 1
)

var MeasureResponse_Status_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}

var MeasureResponse_Status_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x MeasureResponse_Status) String() string {
	return proto.EnumName(MeasureResponse_Status_name, int32(x))
}

func (MeasureResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3, 0}
}

type ConfigureRequest struct {
	Msg                  string    `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	SrInfo               []*SRInfo `protobuf:"bytes,2,rep,name=sr_info,json=srInfo,proto3" json:"sr_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ConfigureRequest) Reset()         { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()    {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ConfigureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigureRequest.Unmarshal(m, b)
}
func (m *ConfigureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigureRequest.Marshal(b, m, deterministic)
}
func (m *ConfigureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureRequest.Merge(m, src)
}
func (m *ConfigureRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigureRequest.Size(m)
}
func (m *ConfigureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureRequest proto.InternalMessageInfo

func (m *ConfigureRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ConfigureRequest) GetSrInfo() []*SRInfo {
	if m != nil {
		return m.SrInfo
	}
	return nil
}

type ConfigureResponse struct {
	Status               ConfigureResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=api.ConfigureResponse_Status" json:"status,omitempty"`
	Msg                  string                   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ConfigureResponse) Reset()         { *m = ConfigureResponse{} }
func (m *ConfigureResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigureResponse) ProtoMessage()    {}
func (*ConfigureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ConfigureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigureResponse.Unmarshal(m, b)
}
func (m *ConfigureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigureResponse.Marshal(b, m, deterministic)
}
func (m *ConfigureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureResponse.Merge(m, src)
}
func (m *ConfigureResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigureResponse.Size(m)
}
func (m *ConfigureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureResponse proto.InternalMessageInfo

func (m *ConfigureResponse) GetStatus() ConfigureResponse_Status {
	if m != nil {
		return m.Status
	}
	return ConfigureResponse_OK
}

func (m *ConfigureResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type MeasureRequest struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Param                *Param   `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeasureRequest) Reset()         { *m = MeasureRequest{} }
func (m *MeasureRequest) String() string { return proto.CompactTextString(m) }
func (*MeasureRequest) ProtoMessage()    {}
func (*MeasureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *MeasureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeasureRequest.Unmarshal(m, b)
}
func (m *MeasureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeasureRequest.Marshal(b, m, deterministic)
}
func (m *MeasureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasureRequest.Merge(m, src)
}
func (m *MeasureRequest) XXX_Size() int {
	return xxx_messageInfo_MeasureRequest.Size(m)
}
func (m *MeasureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MeasureRequest proto.InternalMessageInfo

func (m *MeasureRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *MeasureRequest) GetParam() *Param {
	if m != nil {
		return m.Param
	}
	return nil
}

type MeasureResponse struct {
	Status               MeasureResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=api.MeasureResponse_Status" json:"status,omitempty"`
	Msg                  string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MeasureResponse) Reset()         { *m = MeasureResponse{} }
func (m *MeasureResponse) String() string { return proto.CompactTextString(m) }
func (*MeasureResponse) ProtoMessage()    {}
func (*MeasureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *MeasureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeasureResponse.Unmarshal(m, b)
}
func (m *MeasureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeasureResponse.Marshal(b, m, deterministic)
}
func (m *MeasureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasureResponse.Merge(m, src)
}
func (m *MeasureResponse) XXX_Size() int {
	return xxx_messageInfo_MeasureResponse.Size(m)
}
func (m *MeasureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MeasureResponse proto.InternalMessageInfo

func (m *MeasureResponse) GetStatus() MeasureResponse_Status {
	if m != nil {
		return m.Status
	}
	return MeasureResponse_OK
}

func (m *MeasureResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SRInfo struct {
	SrcAddr              string   `protobuf:"bytes,1,opt,name=src_addr,json=srcAddr,proto3" json:"src_addr,omitempty"`
	Vrf                  int32    `protobuf:"varint,2,opt,name=vrf,proto3" json:"vrf,omitempty"`
	DstAddr              string   `protobuf:"bytes,3,opt,name=dst_addr,json=dstAddr,proto3" json:"dst_addr,omitempty"`
	SidList              []string `protobuf:"bytes,4,rep,name=sid_list,json=sidList,proto3" json:"sid_list,omitempty"`
	TableName            string   `protobuf:"bytes,5,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SRInfo) Reset()         { *m = SRInfo{} }
func (m *SRInfo) String() string { return proto.CompactTextString(m) }
func (*SRInfo) ProtoMessage()    {}
func (*SRInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *SRInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SRInfo.Unmarshal(m, b)
}
func (m *SRInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SRInfo.Marshal(b, m, deterministic)
}
func (m *SRInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SRInfo.Merge(m, src)
}
func (m *SRInfo) XXX_Size() int {
	return xxx_messageInfo_SRInfo.Size(m)
}
func (m *SRInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SRInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SRInfo proto.InternalMessageInfo

func (m *SRInfo) GetSrcAddr() string {
	if m != nil {
		return m.SrcAddr
	}
	return ""
}

func (m *SRInfo) GetVrf() int32 {
	if m != nil {
		return m.Vrf
	}
	return 0
}

func (m *SRInfo) GetDstAddr() string {
	if m != nil {
		return m.DstAddr
	}
	return ""
}

func (m *SRInfo) GetSidList() []string {
	if m != nil {
		return m.SidList
	}
	return nil
}

func (m *SRInfo) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

type Param struct {
	PacketNum            int32    `protobuf:"varint,1,opt,name=packet_num,json=packetNum,proto3" json:"packet_num,omitempty"`
	PacketSize           int32    `protobuf:"varint,2,opt,name=packet_size,json=packetSize,proto3" json:"packet_size,omitempty"`
	RepeatNum            int32    `protobuf:"varint,3,opt,name=repeat_num,json=repeatNum,proto3" json:"repeat_num,omitempty"`
	MeasNum              int32    `protobuf:"varint,4,opt,name=meas_num,json=measNum,proto3" json:"meas_num,omitempty"`
	SmaInterval          int32    `protobuf:"varint,5,opt,name=sma_interval,json=smaInterval,proto3" json:"sma_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Param) Reset()         { *m = Param{} }
func (m *Param) String() string { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()    {}
func (*Param) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *Param) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Param.Unmarshal(m, b)
}
func (m *Param) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Param.Marshal(b, m, deterministic)
}
func (m *Param) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Param.Merge(m, src)
}
func (m *Param) XXX_Size() int {
	return xxx_messageInfo_Param.Size(m)
}
func (m *Param) XXX_DiscardUnknown() {
	xxx_messageInfo_Param.DiscardUnknown(m)
}

var xxx_messageInfo_Param proto.InternalMessageInfo

func (m *Param) GetPacketNum() int32 {
	if m != nil {
		return m.PacketNum
	}
	return 0
}

func (m *Param) GetPacketSize() int32 {
	if m != nil {
		return m.PacketSize
	}
	return 0
}

func (m *Param) GetRepeatNum() int32 {
	if m != nil {
		return m.RepeatNum
	}
	return 0
}

func (m *Param) GetMeasNum() int32 {
	if m != nil {
		return m.MeasNum
	}
	return 0
}

func (m *Param) GetSmaInterval() int32 {
	if m != nil {
		return m.SmaInterval
	}
	return 0
}

func init() {
	proto.RegisterEnum("api.ConfigureResponse_Status", ConfigureResponse_Status_name, ConfigureResponse_Status_value)
	proto.RegisterEnum("api.MeasureResponse_Status", MeasureResponse_Status_name, MeasureResponse_Status_value)
	proto.RegisterType((*ConfigureRequest)(nil), "api.ConfigureRequest")
	proto.RegisterType((*ConfigureResponse)(nil), "api.ConfigureResponse")
	proto.RegisterType((*MeasureRequest)(nil), "api.MeasureRequest")
	proto.RegisterType((*MeasureResponse)(nil), "api.MeasureResponse")
	proto.RegisterType((*SRInfo)(nil), "api.SRInfo")
	proto.RegisterType((*Param)(nil), "api.Param")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0x25, 0xae, 0x9d, 0x66, 0xf2, 0xa3, 0x84, 0x05, 0xaa, 0x94, 0xaa, 0x22, 0x58, 0x1c,
	0x7a, 0x69, 0x22, 0xa5, 0x70, 0xe1, 0x06, 0x88, 0x43, 0x0b, 0xb4, 0x68, 0x73, 0xe3, 0x62, 0x6d,
	0xec, 0x49, 0xba, 0x6a, 0xd6, 0x36, 0x3b, 0xeb, 0x48, 0x09, 0xdf, 0x80, 0xcf, 0xc1, 0x07, 0x45,
	0xfb, 0x87, 0xa8, 0x84, 0xde, 0x38, 0x79, 0x77, 0xe6, 0xbd, 0xa7, 0x37, 0x3b, 0xcf, 0xf0, 0x80,
	0x50, 0xaf, 0x64, 0x8e, 0xa3, 0x5a, 0x57, 0xa6, 0x62, 0x91, 0xa8, 0x65, 0x7a, 0x09, 0xfd, 0xf7,
	0x55, 0x39, 0x97, 0x8b, 0x46, 0x23, 0xc7, 0x6f, 0x0d, 0x92, 0x61, 0x7d, 0x88, 0x14, 0x2d, 0x06,
	0xad, 0x61, 0xeb, 0xb4, 0xcb, 0xed, 0x91, 0xbd, 0x84, 0x0e, 0xe9, 0x4c, 0x96, 0xf3, 0x6a, 0xd0,
	0x1e, 0x46, 0xa7, 0xbd, 0x49, 0x6f, 0x24, 0x6a, 0x39, 0x9a, 0xf2, 0x8b, 0x72, 0x5e, 0xf1, 0x84,
	0xb4, 0xfd, 0xa6, 0x6b, 0x78, 0x74, 0x47, 0x8b, 0xea, 0xaa, 0x24, 0x64, 0xaf, 0x21, 0x21, 0x23,
	0x4c, 0x43, 0x4e, 0xef, 0x60, 0x72, 0xe2, 0x98, 0x7f, 0xe1, 0x46, 0x53, 0x07, 0xe2, 0x01, 0xfc,
	0xdb, 0x43, 0x7b, 0xeb, 0x21, 0x3d, 0x86, 0xc4, 0x63, 0x58, 0x02, 0xed, 0xeb, 0x8f, 0xfd, 0xff,
	0x58, 0x17, 0xe2, 0x0f, 0x9c, 0x5f, 0xf3, 0x7e, 0x2b, 0xbd, 0x84, 0x83, 0xcf, 0x28, 0xe8, 0xce,
	0x10, 0x87, 0x90, 0x28, 0x34, 0x37, 0x55, 0x11, 0xe6, 0x08, 0x37, 0x36, 0x84, 0xb8, 0x16, 0x5a,
	0x28, 0x27, 0xdd, 0x9b, 0x80, 0xb3, 0xf3, 0xc5, 0x56, 0xb8, 0x6f, 0xa4, 0x0d, 0x3c, 0xdc, 0x6a,
	0x85, 0x21, 0xce, 0x77, 0x86, 0x38, 0x76, 0xac, 0x1d, 0xd4, 0x3f, 0x8e, 0xf0, 0xa3, 0x05, 0x89,
	0x7f, 0x50, 0x76, 0x04, 0xfb, 0xa4, 0xf3, 0x4c, 0x14, 0x85, 0x0e, 0xee, 0x3b, 0xa4, 0xf3, 0xb7,
	0x45, 0xa1, 0xad, 0xe8, 0x4a, 0xcf, 0x9d, 0x68, 0xcc, 0xed, 0xd1, 0x82, 0x0b, 0x32, 0x1e, 0x1c,
	0x79, 0x70, 0x41, 0xc6, 0x81, 0xad, 0x8e, 0x2c, 0xb2, 0xa5, 0x24, 0x33, 0xd8, 0x1b, 0x46, 0x4e,
	0x47, 0x16, 0x9f, 0x24, 0x19, 0x76, 0x02, 0x60, 0xc4, 0x6c, 0x89, 0x59, 0x29, 0x14, 0x0e, 0x62,
	0xc7, 0xeb, 0xba, 0xca, 0x95, 0x50, 0x98, 0xfe, 0x6c, 0x41, 0xec, 0x1e, 0xc5, 0x02, 0x6b, 0x91,
	0xdf, 0xa2, 0xc9, 0xca, 0x46, 0x39, 0x37, 0x31, 0xef, 0xfa, 0xca, 0x55, 0xa3, 0xd8, 0x73, 0xe8,
	0x85, 0x36, 0xc9, 0x0d, 0x06, 0x5f, 0x81, 0x31, 0x95, 0x1b, 0xb4, 0x7c, 0x8d, 0x35, 0x0a, 0xcf,
	0x8f, 0x3c, 0xdf, 0x57, 0x2c, 0xff, 0x08, 0xf6, 0x15, 0x0a, 0x72, 0xcd, 0x3d, 0xd7, 0xec, 0xd8,
	0xbb, 0x6d, 0xbd, 0x80, 0xff, 0x49, 0x89, 0x4c, 0x96, 0x06, 0xf5, 0x4a, 0x2c, 0x9d, 0xc9, 0x98,
	0xf7, 0x48, 0x89, 0x8b, 0x50, 0x9a, 0x7c, 0x87, 0xce, 0xd4, 0x67, 0x9a, 0xbd, 0x81, 0xee, 0x36,
	0x54, 0xec, 0xe9, 0x6e, 0xc8, 0x5c, 0x26, 0x9e, 0x1d, 0xde, 0x9f, 0x3d, 0xf6, 0x0a, 0x3a, 0x61,
	0x97, 0xec, 0xf1, 0x9f, 0x9b, 0xf5, 0xbc, 0x27, 0xf7, 0xad, 0xfb, 0x5d, 0xfa, 0x75, 0xb8, 0x90,
	0xe6, 0xa6, 0x99, 0x8d, 0xf2, 0x4a, 0x8d, 0xd7, 0x67, 0xb7, 0x1b, 0x35, 0xc6, 0x52, 0x17, 0x67,
	0xb4, 0x26, 0x83, 0x6a, 0x2c, 0x6a, 0x39, 0x4b, 0xdc, 0xaf, 0x76, 0xfe, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0xa5, 0x64, 0x77, 0x7b, 0x03, 0x00, 0x00,
}
