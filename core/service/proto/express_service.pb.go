// Code generated by protoc-gen-go. DO NOT EDIT.
// source: express_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type ExpressProviderListResponse struct {
	Value                []*SExpressProvider `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExpressProviderListResponse) Reset()         { *m = ExpressProviderListResponse{} }
func (m *ExpressProviderListResponse) String() string { return proto.CompactTextString(m) }
func (*ExpressProviderListResponse) ProtoMessage()    {}
func (*ExpressProviderListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{0}
}
func (m *ExpressProviderListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressProviderListResponse.Unmarshal(m, b)
}
func (m *ExpressProviderListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressProviderListResponse.Marshal(b, m, deterministic)
}
func (dst *ExpressProviderListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressProviderListResponse.Merge(dst, src)
}
func (m *ExpressProviderListResponse) XXX_Size() int {
	return xxx_messageInfo_ExpressProviderListResponse.Size(m)
}
func (m *ExpressProviderListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressProviderListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressProviderListResponse proto.InternalMessageInfo

func (m *ExpressProviderListResponse) GetValue() []*SExpressProvider {
	if m != nil {
		return m.Value
	}
	return nil
}

// 快递服务商
type SExpressProvider struct {
	// 快递公司编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 快递名称
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 首字母，用于索引分组
	Letter string `protobuf:"bytes,3,opt,name=Letter,proto3" json:"Letter,omitempty"`
	// 分组,多个组,用","隔开
	GroupFlag string `protobuf:"bytes,4,opt,name=GroupFlag,proto3" json:"GroupFlag,omitempty"`
	// 快递公司编码
	Code string `protobuf:"bytes,5,opt,name=Code,proto3" json:"Code,omitempty"`
	// 接口编码
	ApiCode              string   `protobuf:"bytes,6,opt,name=ApiCode,proto3" json:"ApiCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SExpressProvider) Reset()         { *m = SExpressProvider{} }
func (m *SExpressProvider) String() string { return proto.CompactTextString(m) }
func (*SExpressProvider) ProtoMessage()    {}
func (*SExpressProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{1}
}
func (m *SExpressProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SExpressProvider.Unmarshal(m, b)
}
func (m *SExpressProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SExpressProvider.Marshal(b, m, deterministic)
}
func (dst *SExpressProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SExpressProvider.Merge(dst, src)
}
func (m *SExpressProvider) XXX_Size() int {
	return xxx_messageInfo_SExpressProvider.Size(m)
}
func (m *SExpressProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_SExpressProvider.DiscardUnknown(m)
}

var xxx_messageInfo_SExpressProvider proto.InternalMessageInfo

func (m *SExpressProvider) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SExpressProvider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SExpressProvider) GetLetter() string {
	if m != nil {
		return m.Letter
	}
	return ""
}

func (m *SExpressProvider) GetGroupFlag() string {
	if m != nil {
		return m.GroupFlag
	}
	return ""
}

func (m *SExpressProvider) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SExpressProvider) GetApiCode() string {
	if m != nil {
		return m.ApiCode
	}
	return ""
}

// 快递模板
type SExpressTemplate struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 运营商编号
	SellerId int64 `protobuf:"varint,2,opt,name=SellerId,proto3" json:"SellerId,omitempty"`
	// 运费模板名称
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	// 是否卖价承担运费
	IsFree bool `protobuf:"varint,4,opt,name=IsFree,proto3" json:"IsFree,omitempty"`
	// 运费计价依据
	Basis int32 `protobuf:"varint,5,opt,name=Basis,proto3" json:"Basis,omitempty"`
	// 首次计价单位,如首重为2kg
	FirstUnit int32 `protobuf:"varint,6,opt,name=FirstUnit,proto3" json:"FirstUnit,omitempty"`
	// 首次计价单价,如续重1kg
	FirstFee float64 `protobuf:"fixed64,7,opt,name=FirstFee,proto3" json:"FirstFee,omitempty"`
	// 超过首次计价计算单位,如续重1kg
	AddUnit int32 `protobuf:"varint,8,opt,name=AddUnit,proto3" json:"AddUnit,omitempty"`
	// 超过首次计价单价，如续重1kg
	AddFee float64 `protobuf:"fixed64,9,opt,name=AddFee,proto3" json:"AddFee,omitempty"`
	// 是否启用
	Enabled              bool     `protobuf:"varint,10,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SExpressTemplate) Reset()         { *m = SExpressTemplate{} }
func (m *SExpressTemplate) String() string { return proto.CompactTextString(m) }
func (*SExpressTemplate) ProtoMessage()    {}
func (*SExpressTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{2}
}
func (m *SExpressTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SExpressTemplate.Unmarshal(m, b)
}
func (m *SExpressTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SExpressTemplate.Marshal(b, m, deterministic)
}
func (dst *SExpressTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SExpressTemplate.Merge(dst, src)
}
func (m *SExpressTemplate) XXX_Size() int {
	return xxx_messageInfo_SExpressTemplate.Size(m)
}
func (m *SExpressTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_SExpressTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_SExpressTemplate proto.InternalMessageInfo

func (m *SExpressTemplate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SExpressTemplate) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *SExpressTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SExpressTemplate) GetIsFree() bool {
	if m != nil {
		return m.IsFree
	}
	return false
}

func (m *SExpressTemplate) GetBasis() int32 {
	if m != nil {
		return m.Basis
	}
	return 0
}

func (m *SExpressTemplate) GetFirstUnit() int32 {
	if m != nil {
		return m.FirstUnit
	}
	return 0
}

func (m *SExpressTemplate) GetFirstFee() float64 {
	if m != nil {
		return m.FirstFee
	}
	return 0
}

func (m *SExpressTemplate) GetAddUnit() int32 {
	if m != nil {
		return m.AddUnit
	}
	return 0
}

func (m *SExpressTemplate) GetAddFee() float64 {
	if m != nil {
		return m.AddFee
	}
	return 0
}

func (m *SExpressTemplate) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type ExpressTemplateId struct {
	SellerId             int64    `protobuf:"varint,1,opt,name=SellerId,proto3" json:"SellerId,omitempty"`
	TemplateId           int64    `protobuf:"varint,2,opt,name=TemplateId,proto3" json:"TemplateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpressTemplateId) Reset()         { *m = ExpressTemplateId{} }
func (m *ExpressTemplateId) String() string { return proto.CompactTextString(m) }
func (*ExpressTemplateId) ProtoMessage()    {}
func (*ExpressTemplateId) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{3}
}
func (m *ExpressTemplateId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressTemplateId.Unmarshal(m, b)
}
func (m *ExpressTemplateId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressTemplateId.Marshal(b, m, deterministic)
}
func (dst *ExpressTemplateId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressTemplateId.Merge(dst, src)
}
func (m *ExpressTemplateId) XXX_Size() int {
	return xxx_messageInfo_ExpressTemplateId.Size(m)
}
func (m *ExpressTemplateId) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressTemplateId.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressTemplateId proto.InternalMessageInfo

func (m *ExpressTemplateId) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *ExpressTemplateId) GetTemplateId() int64 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

type GetTemplatesRequest struct {
	SellerId int64 `protobuf:"varint,1,opt,name=SellerId,proto3" json:"SellerId,omitempty"`
	// 仅返回已启用的模板
	OnlyEnabled          bool     `protobuf:"varint,2,opt,name=OnlyEnabled,proto3" json:"OnlyEnabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTemplatesRequest) Reset()         { *m = GetTemplatesRequest{} }
func (m *GetTemplatesRequest) String() string { return proto.CompactTextString(m) }
func (*GetTemplatesRequest) ProtoMessage()    {}
func (*GetTemplatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{4}
}
func (m *GetTemplatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTemplatesRequest.Unmarshal(m, b)
}
func (m *GetTemplatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTemplatesRequest.Marshal(b, m, deterministic)
}
func (dst *GetTemplatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTemplatesRequest.Merge(dst, src)
}
func (m *GetTemplatesRequest) XXX_Size() int {
	return xxx_messageInfo_GetTemplatesRequest.Size(m)
}
func (m *GetTemplatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTemplatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTemplatesRequest proto.InternalMessageInfo

func (m *GetTemplatesRequest) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *GetTemplatesRequest) GetOnlyEnabled() bool {
	if m != nil {
		return m.OnlyEnabled
	}
	return false
}

type ExpressTemplateListResponse struct {
	Value                []*SExpressTemplate `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExpressTemplateListResponse) Reset()         { *m = ExpressTemplateListResponse{} }
func (m *ExpressTemplateListResponse) String() string { return proto.CompactTextString(m) }
func (*ExpressTemplateListResponse) ProtoMessage()    {}
func (*ExpressTemplateListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{5}
}
func (m *ExpressTemplateListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressTemplateListResponse.Unmarshal(m, b)
}
func (m *ExpressTemplateListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressTemplateListResponse.Marshal(b, m, deterministic)
}
func (dst *ExpressTemplateListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressTemplateListResponse.Merge(dst, src)
}
func (m *ExpressTemplateListResponse) XXX_Size() int {
	return xxx_messageInfo_ExpressTemplateListResponse.Size(m)
}
func (m *ExpressTemplateListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressTemplateListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressTemplateListResponse proto.InternalMessageInfo

func (m *ExpressTemplateListResponse) GetValue() []*SExpressTemplate {
	if m != nil {
		return m.Value
	}
	return nil
}

// 快递地区模板
type SExpressAreaTemplate struct {
	// 模板编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 地区编号列表，通常精确到省即可
	CodeList string `protobuf:"bytes,2,opt,name=CodeList,proto3" json:"CodeList,omitempty"`
	// 地区名称列表
	NameList string `protobuf:"bytes,3,opt,name=NameList,proto3" json:"NameList,omitempty"`
	// 首次数值，如 首重为2kg
	FirstUnit int32 `protobuf:"varint,4,opt,name=FirstUnit,proto3" json:"FirstUnit,omitempty"`
	// 首次金额，如首重10元
	FirstFee float64 `protobuf:"fixed64,5,opt,name=FirstFee,proto3" json:"FirstFee,omitempty"`
	// 增加数值，如续重1kg
	AddUnit int32 `protobuf:"varint,6,opt,name=AddUnit,proto3" json:"AddUnit,omitempty"`
	// 增加产生费用，如续重1kg 10元
	AddFee               float64  `protobuf:"fixed64,7,opt,name=AddFee,proto3" json:"AddFee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SExpressAreaTemplate) Reset()         { *m = SExpressAreaTemplate{} }
func (m *SExpressAreaTemplate) String() string { return proto.CompactTextString(m) }
func (*SExpressAreaTemplate) ProtoMessage()    {}
func (*SExpressAreaTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{6}
}
func (m *SExpressAreaTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SExpressAreaTemplate.Unmarshal(m, b)
}
func (m *SExpressAreaTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SExpressAreaTemplate.Marshal(b, m, deterministic)
}
func (dst *SExpressAreaTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SExpressAreaTemplate.Merge(dst, src)
}
func (m *SExpressAreaTemplate) XXX_Size() int {
	return xxx_messageInfo_SExpressAreaTemplate.Size(m)
}
func (m *SExpressAreaTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_SExpressAreaTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_SExpressAreaTemplate proto.InternalMessageInfo

func (m *SExpressAreaTemplate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SExpressAreaTemplate) GetCodeList() string {
	if m != nil {
		return m.CodeList
	}
	return ""
}

func (m *SExpressAreaTemplate) GetNameList() string {
	if m != nil {
		return m.NameList
	}
	return ""
}

func (m *SExpressAreaTemplate) GetFirstUnit() int32 {
	if m != nil {
		return m.FirstUnit
	}
	return 0
}

func (m *SExpressAreaTemplate) GetFirstFee() float64 {
	if m != nil {
		return m.FirstFee
	}
	return 0
}

func (m *SExpressAreaTemplate) GetAddUnit() int32 {
	if m != nil {
		return m.AddUnit
	}
	return 0
}

func (m *SExpressAreaTemplate) GetAddFee() float64 {
	if m != nil {
		return m.AddFee
	}
	return 0
}

type SaveAreaExpTemplateRequest struct {
	SellerId             int64                 `protobuf:"varint,1,opt,name=SellerId,proto3" json:"SellerId,omitempty"`
	TemplateId           int64                 `protobuf:"varint,2,opt,name=TemplateId,proto3" json:"TemplateId,omitempty"`
	Value                *SExpressAreaTemplate `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SaveAreaExpTemplateRequest) Reset()         { *m = SaveAreaExpTemplateRequest{} }
func (m *SaveAreaExpTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*SaveAreaExpTemplateRequest) ProtoMessage()    {}
func (*SaveAreaExpTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{7}
}
func (m *SaveAreaExpTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveAreaExpTemplateRequest.Unmarshal(m, b)
}
func (m *SaveAreaExpTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveAreaExpTemplateRequest.Marshal(b, m, deterministic)
}
func (dst *SaveAreaExpTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveAreaExpTemplateRequest.Merge(dst, src)
}
func (m *SaveAreaExpTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_SaveAreaExpTemplateRequest.Size(m)
}
func (m *SaveAreaExpTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveAreaExpTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveAreaExpTemplateRequest proto.InternalMessageInfo

func (m *SaveAreaExpTemplateRequest) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *SaveAreaExpTemplateRequest) GetTemplateId() int64 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *SaveAreaExpTemplateRequest) GetValue() *SExpressAreaTemplate {
	if m != nil {
		return m.Value
	}
	return nil
}

type AreaTemplateId struct {
	SellerId             int64    `protobuf:"varint,1,opt,name=SellerId,proto3" json:"SellerId,omitempty"`
	TemplateId           int64    `protobuf:"varint,2,opt,name=TemplateId,proto3" json:"TemplateId,omitempty"`
	AreaTemplateId       int64    `protobuf:"varint,3,opt,name=AreaTemplateId,proto3" json:"AreaTemplateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AreaTemplateId) Reset()         { *m = AreaTemplateId{} }
func (m *AreaTemplateId) String() string { return proto.CompactTextString(m) }
func (*AreaTemplateId) ProtoMessage()    {}
func (*AreaTemplateId) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{8}
}
func (m *AreaTemplateId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaTemplateId.Unmarshal(m, b)
}
func (m *AreaTemplateId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaTemplateId.Marshal(b, m, deterministic)
}
func (dst *AreaTemplateId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaTemplateId.Merge(dst, src)
}
func (m *AreaTemplateId) XXX_Size() int {
	return xxx_messageInfo_AreaTemplateId.Size(m)
}
func (m *AreaTemplateId) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaTemplateId.DiscardUnknown(m)
}

var xxx_messageInfo_AreaTemplateId proto.InternalMessageInfo

func (m *AreaTemplateId) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *AreaTemplateId) GetTemplateId() int64 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *AreaTemplateId) GetAreaTemplateId() int64 {
	if m != nil {
		return m.AreaTemplateId
	}
	return 0
}

type SaveTemplateResponse struct {
	ErrCode              int64    `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	TemplateId           int64    `protobuf:"varint,3,opt,name=TemplateId,proto3" json:"TemplateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveTemplateResponse) Reset()         { *m = SaveTemplateResponse{} }
func (m *SaveTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*SaveTemplateResponse) ProtoMessage()    {}
func (*SaveTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_fd7d2b5e604c3a7f, []int{9}
}
func (m *SaveTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTemplateResponse.Unmarshal(m, b)
}
func (m *SaveTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTemplateResponse.Marshal(b, m, deterministic)
}
func (dst *SaveTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTemplateResponse.Merge(dst, src)
}
func (m *SaveTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_SaveTemplateResponse.Size(m)
}
func (m *SaveTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTemplateResponse proto.InternalMessageInfo

func (m *SaveTemplateResponse) GetErrCode() int64 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SaveTemplateResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *SaveTemplateResponse) GetTemplateId() int64 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func init() {
	proto.RegisterType((*ExpressProviderListResponse)(nil), "ExpressProviderListResponse")
	proto.RegisterType((*SExpressProvider)(nil), "SExpressProvider")
	proto.RegisterType((*SExpressTemplate)(nil), "SExpressTemplate")
	proto.RegisterType((*ExpressTemplateId)(nil), "ExpressTemplateId")
	proto.RegisterType((*GetTemplatesRequest)(nil), "GetTemplatesRequest")
	proto.RegisterType((*ExpressTemplateListResponse)(nil), "ExpressTemplateListResponse")
	proto.RegisterType((*SExpressAreaTemplate)(nil), "SExpressAreaTemplate")
	proto.RegisterType((*SaveAreaExpTemplateRequest)(nil), "SaveAreaExpTemplateRequest")
	proto.RegisterType((*AreaTemplateId)(nil), "AreaTemplateId")
	proto.RegisterType((*SaveTemplateResponse)(nil), "SaveTemplateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExpressServiceClient is the client API for ExpressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExpressServiceClient interface {
	// 获取快递公司
	GetExpressProvider(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SExpressProvider, error)
	// 保存快递公司
	SaveExpressProvider(ctx context.Context, in *SExpressProvider, opts ...grpc.CallOption) (*Result, error)
	// 获取可用的快递公司
	GetProviders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExpressProviderListResponse, error)
	// 保存快递模板
	SaveTemplate(ctx context.Context, in *SExpressTemplate, opts ...grpc.CallOption) (*SaveTemplateResponse, error)
	// 获取单个快递模板
	GetTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*SExpressTemplate, error)
	// 获取卖家的快递模板
	GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*ExpressTemplateListResponse, error)
	// 删除模板
	DeleteTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*Result, error)
	// 保存地区快递模板
	SaveAreaTemplate(ctx context.Context, in *SaveAreaExpTemplateRequest, opts ...grpc.CallOption) (*Result, error)
	// 删除模板地区设定
	DeleteAreaTemplate(ctx context.Context, in *AreaTemplateId, opts ...grpc.CallOption) (*Result, error)
}

type expressServiceClient struct {
	cc *grpc.ClientConn
}

func NewExpressServiceClient(cc *grpc.ClientConn) ExpressServiceClient {
	return &expressServiceClient{cc}
}

func (c *expressServiceClient) GetExpressProvider(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SExpressProvider, error) {
	out := new(SExpressProvider)
	err := c.cc.Invoke(ctx, "/ExpressService/GetExpressProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveExpressProvider(ctx context.Context, in *SExpressProvider, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveExpressProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetProviders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExpressProviderListResponse, error) {
	out := new(ExpressProviderListResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/GetProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveTemplate(ctx context.Context, in *SExpressTemplate, opts ...grpc.CallOption) (*SaveTemplateResponse, error) {
	out := new(SaveTemplateResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*SExpressTemplate, error) {
	out := new(SExpressTemplate)
	err := c.cc.Invoke(ctx, "/ExpressService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*ExpressTemplateListResponse, error) {
	out := new(ExpressTemplateListResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/GetTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) DeleteTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveAreaTemplate(ctx context.Context, in *SaveAreaExpTemplateRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveAreaTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) DeleteAreaTemplate(ctx context.Context, in *AreaTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/DeleteAreaTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpressServiceServer is the server API for ExpressService service.
type ExpressServiceServer interface {
	// 获取快递公司
	GetExpressProvider(context.Context, *IdOrName) (*SExpressProvider, error)
	// 保存快递公司
	SaveExpressProvider(context.Context, *SExpressProvider) (*Result, error)
	// 获取可用的快递公司
	GetProviders(context.Context, *Empty) (*ExpressProviderListResponse, error)
	// 保存快递模板
	SaveTemplate(context.Context, *SExpressTemplate) (*SaveTemplateResponse, error)
	// 获取单个快递模板
	GetTemplate(context.Context, *ExpressTemplateId) (*SExpressTemplate, error)
	// 获取卖家的快递模板
	GetTemplates(context.Context, *GetTemplatesRequest) (*ExpressTemplateListResponse, error)
	// 删除模板
	DeleteTemplate(context.Context, *ExpressTemplateId) (*Result, error)
	// 保存地区快递模板
	SaveAreaTemplate(context.Context, *SaveAreaExpTemplateRequest) (*Result, error)
	// 删除模板地区设定
	DeleteAreaTemplate(context.Context, *AreaTemplateId) (*Result, error)
}

func RegisterExpressServiceServer(s *grpc.Server, srv ExpressServiceServer) {
	s.RegisterService(&_ExpressService_serviceDesc, srv)
}

func _ExpressService_GetExpressProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetExpressProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetExpressProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetExpressProvider(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveExpressProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SExpressProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveExpressProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveExpressProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveExpressProvider(ctx, req.(*SExpressProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetProviders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SExpressTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveTemplate(ctx, req.(*SExpressTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetTemplate(ctx, req.(*ExpressTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetTemplates(ctx, req.(*GetTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).DeleteTemplate(ctx, req.(*ExpressTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveAreaTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAreaExpTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveAreaTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveAreaTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveAreaTemplate(ctx, req.(*SaveAreaExpTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_DeleteAreaTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).DeleteAreaTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/DeleteAreaTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).DeleteAreaTemplate(ctx, req.(*AreaTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExpressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExpressService",
	HandlerType: (*ExpressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExpressProvider",
			Handler:    _ExpressService_GetExpressProvider_Handler,
		},
		{
			MethodName: "SaveExpressProvider",
			Handler:    _ExpressService_SaveExpressProvider_Handler,
		},
		{
			MethodName: "GetProviders",
			Handler:    _ExpressService_GetProviders_Handler,
		},
		{
			MethodName: "SaveTemplate",
			Handler:    _ExpressService_SaveTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _ExpressService_GetTemplate_Handler,
		},
		{
			MethodName: "GetTemplates",
			Handler:    _ExpressService_GetTemplates_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _ExpressService_DeleteTemplate_Handler,
		},
		{
			MethodName: "SaveAreaTemplate",
			Handler:    _ExpressService_SaveAreaTemplate_Handler,
		},
		{
			MethodName: "DeleteAreaTemplate",
			Handler:    _ExpressService_DeleteAreaTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "express_service.proto",
}

func init() {
	proto.RegisterFile("express_service.proto", fileDescriptor_express_service_fd7d2b5e604c3a7f)
}

var fileDescriptor_express_service_fd7d2b5e604c3a7f = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xda, 0x4a,
	0x10, 0xc6, 0x98, 0x9f, 0x30, 0x44, 0x9c, 0x64, 0x43, 0x8e, 0x2c, 0x27, 0x8a, 0x90, 0x2f, 0xce,
	0x41, 0xaa, 0xb4, 0xad, 0x48, 0xd5, 0x8b, 0xf4, 0x2a, 0x69, 0x00, 0x21, 0xa5, 0x4d, 0x65, 0xda,
	0x5e, 0xf4, 0xa6, 0x32, 0xf1, 0x88, 0x58, 0x32, 0xd8, 0xdd, 0x5d, 0xa2, 0xe4, 0x01, 0xda, 0xab,
	0x3e, 0x41, 0x1f, 0xa9, 0x4f, 0x55, 0xed, 0xda, 0x6b, 0x8c, 0x21, 0x34, 0x52, 0xaf, 0xf0, 0x37,
	0xb3, 0x33, 0xf3, 0xcd, 0x2f, 0x70, 0x88, 0xf7, 0x31, 0x43, 0xce, 0xbf, 0x70, 0x64, 0x77, 0xc1,
	0x0d, 0xd2, 0x98, 0x45, 0x22, 0xb2, 0x77, 0xa7, 0x61, 0x34, 0xf1, 0xc2, 0x04, 0x39, 0x03, 0x38,
	0xea, 0x27, 0xcf, 0xde, 0xb3, 0xe8, 0x2e, 0xf0, 0x91, 0x5d, 0x05, 0x5c, 0xb8, 0xc8, 0xe3, 0x68,
	0xce, 0x91, 0xfc, 0x0f, 0xd5, 0x4f, 0x5e, 0xb8, 0x40, 0xcb, 0xe8, 0x98, 0xdd, 0x66, 0x6f, 0x9f,
	0x8e, 0x0b, 0xaf, 0xdd, 0x44, 0xef, 0xfc, 0x34, 0x60, 0xaf, 0xa8, 0x23, 0x2d, 0x28, 0x8f, 0x7c,
	0xcb, 0xe8, 0x18, 0x5d, 0xd3, 0x2d, 0x8f, 0x7c, 0x42, 0xa0, 0xf2, 0xce, 0x9b, 0xa1, 0x55, 0xee,
	0x18, 0xdd, 0x86, 0xab, 0xbe, 0xc9, 0xbf, 0x50, 0xbb, 0x42, 0x21, 0x90, 0x59, 0xa6, 0x92, 0xa6,
	0x88, 0x1c, 0x43, 0x63, 0xc8, 0xa2, 0x45, 0x3c, 0x08, 0xbd, 0xa9, 0x55, 0x51, 0xaa, 0xa5, 0x40,
	0x7a, 0x7a, 0x13, 0xf9, 0x68, 0x55, 0x13, 0x4f, 0xf2, 0x9b, 0x58, 0x50, 0x3f, 0x8f, 0x03, 0x25,
	0xae, 0x29, 0xb1, 0x86, 0xce, 0xf7, 0xf2, 0x92, 0xdc, 0x07, 0x9c, 0xc5, 0xa1, 0x27, 0x70, 0x8d,
	0x9c, 0x0d, 0x3b, 0x63, 0x0c, 0x43, 0x64, 0x23, 0x5f, 0x11, 0x34, 0xdd, 0x0c, 0x67, 0xc4, 0xcd,
	0x55, 0xe2, 0x23, 0x3e, 0x60, 0x88, 0x8a, 0xdd, 0x8e, 0x9b, 0x22, 0xd2, 0x86, 0xea, 0x85, 0xc7,
	0x03, 0xae, 0xb8, 0x55, 0xdd, 0x04, 0xc8, 0x74, 0x06, 0x01, 0xe3, 0xe2, 0xe3, 0x3c, 0x10, 0x8a,
	0x5e, 0xd5, 0x5d, 0x0a, 0x64, 0x6c, 0x05, 0x06, 0x88, 0x56, 0xbd, 0x63, 0x74, 0x0d, 0x37, 0xc3,
	0x2a, 0x2d, 0xdf, 0x57, 0x76, 0x3b, 0xca, 0x4e, 0x43, 0xc9, 0xe0, 0xdc, 0xf7, 0xa5, 0x4d, 0x43,
	0xd9, 0xa4, 0x48, 0x5a, 0xf4, 0xe7, 0xde, 0x24, 0x44, 0xdf, 0x02, 0x45, 0x4d, 0x43, 0xe7, 0x1a,
	0xf6, 0x0b, 0x65, 0x28, 0x24, 0x6e, 0x14, 0x12, 0x3f, 0x01, 0x58, 0xbe, 0x4c, 0xcb, 0x92, 0x93,
	0x38, 0x63, 0x38, 0x18, 0xa2, 0xd0, 0x02, 0xee, 0xe2, 0xd7, 0x05, 0x72, 0xb1, 0xd5, 0x65, 0x07,
	0x9a, 0xd7, 0xf3, 0xf0, 0x41, 0x33, 0x2c, 0x2b, 0x86, 0x79, 0x51, 0x6e, 0x26, 0xb5, 0xe3, 0xa7,
	0xcd, 0xa4, 0x7e, 0xad, 0x67, 0xf2, 0x97, 0x01, 0x6d, 0xad, 0x3b, 0x67, 0xe8, 0x6d, 0x6b, 0xbd,
	0x9c, 0x13, 0x19, 0x25, 0x9d, 0xcd, 0x0c, 0x4b, 0x9d, 0x6c, 0xb7, 0xd2, 0x25, 0xed, 0xcf, 0xf0,
	0x6a, 0x53, 0x2b, 0xdb, 0x9a, 0x5a, 0x7d, 0xbc, 0xa9, 0xb5, 0xc7, 0x9a, 0x5a, 0xcf, 0x37, 0xd5,
	0xf9, 0x66, 0x80, 0x3d, 0xf6, 0xee, 0x50, 0x26, 0xd2, 0xbf, 0x8f, 0xb3, 0x5c, 0x9f, 0x50, 0xf1,
	0x3f, 0x34, 0x91, 0x3c, 0xd3, 0x05, 0x95, 0xf9, 0x35, 0x7b, 0x87, 0x74, 0x53, 0xd1, 0x74, 0x51,
	0x05, 0xb4, 0xf2, 0xe2, 0xbf, 0x9b, 0x1f, 0xf2, 0x5f, 0xd1, 0x9b, 0xe2, 0x60, 0xba, 0x05, 0xa9,
	0x73, 0x0b, 0x6d, 0x99, 0xfc, 0x32, 0xeb, 0x74, 0x16, 0xe4, 0xa8, 0x33, 0xa6, 0x76, 0x3e, 0x09,
	0xad, 0xa1, 0xac, 0x63, 0x9f, 0xb1, 0xb7, 0x7c, 0x9a, 0x76, 0x34, 0x45, 0x05, 0x46, 0x66, 0x91,
	0x51, 0xef, 0x47, 0x05, 0x5a, 0x69, 0xfa, 0xe3, 0xe4, 0x6e, 0x92, 0x97, 0x40, 0x86, 0x28, 0x8a,
	0xc7, 0xad, 0x41, 0x47, 0xfe, 0x35, 0x93, 0xb3, 0x60, 0xaf, 0x9f, 0x45, 0xa7, 0x44, 0x4e, 0xe1,
	0x40, 0x52, 0x2e, 0x9a, 0xad, 0xbf, 0xb5, 0xeb, 0xd4, 0x45, 0xbe, 0x08, 0x85, 0x53, 0x22, 0xaf,
	0x60, 0x77, 0x88, 0x42, 0x6b, 0x38, 0xa9, 0xd1, 0xfe, 0x2c, 0x16, 0x0f, 0xf6, 0x31, 0xdd, 0x72,
	0xa5, 0x9d, 0x12, 0x39, 0x83, 0xdd, 0x7c, 0x7d, 0xc8, 0xfa, 0x52, 0xd8, 0x87, 0x74, 0x53, 0x05,
	0x55, 0xcc, 0x66, 0x6e, 0x87, 0x09, 0xa1, 0x6b, 0x27, 0xc2, 0x5e, 0x77, 0xe7, 0x94, 0xc8, 0xa5,
	0xe2, 0x9a, 0xed, 0x3e, 0x69, 0xd3, 0x0d, 0xa7, 0x60, 0xc9, 0x7c, 0xd3, 0x2e, 0x3b, 0x25, 0xf2,
	0x1c, 0x5a, 0x97, 0x18, 0xa2, 0xc0, 0xad, 0x04, 0x72, 0x25, 0x3a, 0x83, 0x3d, 0xbd, 0x07, 0x99,
	0xc9, 0x11, 0x7d, 0x7c, 0x35, 0xf2, 0xb6, 0x2f, 0x80, 0x24, 0xc1, 0x56, 0xac, 0xff, 0xa1, 0xab,
	0xd3, 0x96, 0xb3, 0xb8, 0x38, 0x81, 0x83, 0x9b, 0x68, 0x46, 0xa7, 0x81, 0xb8, 0x5d, 0x4c, 0xe8,
	0x34, 0xea, 0x45, 0x94, 0xc5, 0x37, 0x9f, 0xeb, 0xf4, 0xb5, 0xfa, 0xff, 0x9c, 0xd4, 0xd4, 0xcf,
	0xe9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0xea, 0x91, 0xb2, 0x6d, 0x07, 0x00, 0x00,
}