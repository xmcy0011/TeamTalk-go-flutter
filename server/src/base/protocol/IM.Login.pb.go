// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IM.Login.proto

package IM_Login

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

type IMMsgServReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMMsgServReq) Reset()         { *m = IMMsgServReq{} }
func (m *IMMsgServReq) String() string { return proto.CompactTextString(m) }
func (*IMMsgServReq) ProtoMessage()    {}
func (*IMMsgServReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{0}
}

func (m *IMMsgServReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMMsgServReq.Unmarshal(m, b)
}
func (m *IMMsgServReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMMsgServReq.Marshal(b, m, deterministic)
}
func (m *IMMsgServReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMMsgServReq.Merge(m, src)
}
func (m *IMMsgServReq) XXX_Size() int {
	return xxx_messageInfo_IMMsgServReq.Size(m)
}
func (m *IMMsgServReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMMsgServReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMMsgServReq proto.InternalMessageInfo

type IMMsgServRsp struct {
	//cmd id:		0x0102
	ResultCode           ResultType `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3,enum=IM.BaseDefine.ResultType" json:"result_code,omitempty"`
	PriorIp              string     `protobuf:"bytes,2,opt,name=prior_ip,json=priorIp,proto3" json:"prior_ip,omitempty"`
	BackipIp             string     `protobuf:"bytes,3,opt,name=backip_ip,json=backipIp,proto3" json:"backip_ip,omitempty"`
	Port                 uint32     `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IMMsgServRsp) Reset()         { *m = IMMsgServRsp{} }
func (m *IMMsgServRsp) String() string { return proto.CompactTextString(m) }
func (*IMMsgServRsp) ProtoMessage()    {}
func (*IMMsgServRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{1}
}

func (m *IMMsgServRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMMsgServRsp.Unmarshal(m, b)
}
func (m *IMMsgServRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMMsgServRsp.Marshal(b, m, deterministic)
}
func (m *IMMsgServRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMMsgServRsp.Merge(m, src)
}
func (m *IMMsgServRsp) XXX_Size() int {
	return xxx_messageInfo_IMMsgServRsp.Size(m)
}
func (m *IMMsgServRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMMsgServRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMMsgServRsp proto.InternalMessageInfo

func (m *IMMsgServRsp) GetResultCode() ResultType {
	if m != nil {
		return m.ResultCode
	}
	return ResultType_REFUSE_REASON_NONE
}

func (m *IMMsgServRsp) GetPriorIp() string {
	if m != nil {
		return m.PriorIp
	}
	return ""
}

func (m *IMMsgServRsp) GetBackipIp() string {
	if m != nil {
		return m.BackipIp
	}
	return ""
}

func (m *IMMsgServRsp) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type IMLoginReq struct {
	//cmd id:		0x0103
	UserName             string       `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password             string       `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	OnlineStatus         UserStatType `protobuf:"varint,3,opt,name=online_status,json=onlineStatus,proto3,enum=IM.BaseDefine.UserStatType" json:"online_status,omitempty"`
	ClientType           ClientType   `protobuf:"varint,4,opt,name=client_type,json=clientType,proto3,enum=IM.BaseDefine.ClientType" json:"client_type,omitempty"`
	ClientVersion        string       `protobuf:"bytes,5,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *IMLoginReq) Reset()         { *m = IMLoginReq{} }
func (m *IMLoginReq) String() string { return proto.CompactTextString(m) }
func (*IMLoginReq) ProtoMessage()    {}
func (*IMLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{2}
}

func (m *IMLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMLoginReq.Unmarshal(m, b)
}
func (m *IMLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMLoginReq.Marshal(b, m, deterministic)
}
func (m *IMLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMLoginReq.Merge(m, src)
}
func (m *IMLoginReq) XXX_Size() int {
	return xxx_messageInfo_IMLoginReq.Size(m)
}
func (m *IMLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMLoginReq proto.InternalMessageInfo

func (m *IMLoginReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *IMLoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *IMLoginReq) GetOnlineStatus() UserStatType {
	if m != nil {
		return m.OnlineStatus
	}
	return UserStatType_USER_STATUS_NONE
}

func (m *IMLoginReq) GetClientType() ClientType {
	if m != nil {
		return m.ClientType
	}
	return ClientType_CLIENT_TYPE_NONE
}

func (m *IMLoginReq) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

type IMLoginRes struct {
	//cmd id:		0x0104
	ServerTime           uint32       `protobuf:"varint,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	ResultCode           ResultType   `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=IM.BaseDefine.ResultType" json:"result_code,omitempty"`
	ResultString         string       `protobuf:"bytes,3,opt,name=result_string,json=resultString,proto3" json:"result_string,omitempty"`
	OnlineStatus         UserStatType `protobuf:"varint,4,opt,name=online_status,json=onlineStatus,proto3,enum=IM.BaseDefine.UserStatType" json:"online_status,omitempty"`
	UserInfo             *UserInfo    `protobuf:"bytes,5,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *IMLoginRes) Reset()         { *m = IMLoginRes{} }
func (m *IMLoginRes) String() string { return proto.CompactTextString(m) }
func (*IMLoginRes) ProtoMessage()    {}
func (*IMLoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{3}
}

func (m *IMLoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMLoginRes.Unmarshal(m, b)
}
func (m *IMLoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMLoginRes.Marshal(b, m, deterministic)
}
func (m *IMLoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMLoginRes.Merge(m, src)
}
func (m *IMLoginRes) XXX_Size() int {
	return xxx_messageInfo_IMLoginRes.Size(m)
}
func (m *IMLoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_IMLoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_IMLoginRes proto.InternalMessageInfo

func (m *IMLoginRes) GetServerTime() uint32 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

func (m *IMLoginRes) GetResultCode() ResultType {
	if m != nil {
		return m.ResultCode
	}
	return ResultType_REFUSE_REASON_NONE
}

func (m *IMLoginRes) GetResultString() string {
	if m != nil {
		return m.ResultString
	}
	return ""
}

func (m *IMLoginRes) GetOnlineStatus() UserStatType {
	if m != nil {
		return m.OnlineStatus
	}
	return UserStatType_USER_STATUS_NONE
}

func (m *IMLoginRes) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type IMLogoutReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMLogoutReq) Reset()         { *m = IMLogoutReq{} }
func (m *IMLogoutReq) String() string { return proto.CompactTextString(m) }
func (*IMLogoutReq) ProtoMessage()    {}
func (*IMLogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{4}
}

func (m *IMLogoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMLogoutReq.Unmarshal(m, b)
}
func (m *IMLogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMLogoutReq.Marshal(b, m, deterministic)
}
func (m *IMLogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMLogoutReq.Merge(m, src)
}
func (m *IMLogoutReq) XXX_Size() int {
	return xxx_messageInfo_IMLogoutReq.Size(m)
}
func (m *IMLogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMLogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMLogoutReq proto.InternalMessageInfo

type IMLogoutRsp struct {
	//cmd id:		0x0106
	ResultCode           uint32   `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMLogoutRsp) Reset()         { *m = IMLogoutRsp{} }
func (m *IMLogoutRsp) String() string { return proto.CompactTextString(m) }
func (*IMLogoutRsp) ProtoMessage()    {}
func (*IMLogoutRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{5}
}

func (m *IMLogoutRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMLogoutRsp.Unmarshal(m, b)
}
func (m *IMLogoutRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMLogoutRsp.Marshal(b, m, deterministic)
}
func (m *IMLogoutRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMLogoutRsp.Merge(m, src)
}
func (m *IMLogoutRsp) XXX_Size() int {
	return xxx_messageInfo_IMLogoutRsp.Size(m)
}
func (m *IMLogoutRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMLogoutRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMLogoutRsp proto.InternalMessageInfo

func (m *IMLogoutRsp) GetResultCode() uint32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

type IMKickUser struct {
	//cmd id:		0x0107
	UserId               uint32         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	KickReason           KickReasonType `protobuf:"varint,2,opt,name=kick_reason,json=kickReason,proto3,enum=IM.BaseDefine.KickReasonType" json:"kick_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IMKickUser) Reset()         { *m = IMKickUser{} }
func (m *IMKickUser) String() string { return proto.CompactTextString(m) }
func (*IMKickUser) ProtoMessage()    {}
func (*IMKickUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{6}
}

func (m *IMKickUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMKickUser.Unmarshal(m, b)
}
func (m *IMKickUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMKickUser.Marshal(b, m, deterministic)
}
func (m *IMKickUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMKickUser.Merge(m, src)
}
func (m *IMKickUser) XXX_Size() int {
	return xxx_messageInfo_IMKickUser.Size(m)
}
func (m *IMKickUser) XXX_DiscardUnknown() {
	xxx_messageInfo_IMKickUser.DiscardUnknown(m)
}

var xxx_messageInfo_IMKickUser proto.InternalMessageInfo

func (m *IMKickUser) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMKickUser) GetKickReason() KickReasonType {
	if m != nil {
		return m.KickReason
	}
	return KickReasonType_KICK_REASON_NONE
}

type IMDeviceTokenReq struct {
	//cmd id:		0x0108
	UserId               uint32     `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceToken          string     `protobuf:"bytes,2,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
	ClientType           ClientType `protobuf:"varint,3,opt,name=client_type,json=clientType,proto3,enum=IM.BaseDefine.ClientType" json:"client_type,omitempty"`
	AttachData           []byte     `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IMDeviceTokenReq) Reset()         { *m = IMDeviceTokenReq{} }
func (m *IMDeviceTokenReq) String() string { return proto.CompactTextString(m) }
func (*IMDeviceTokenReq) ProtoMessage()    {}
func (*IMDeviceTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{7}
}

func (m *IMDeviceTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMDeviceTokenReq.Unmarshal(m, b)
}
func (m *IMDeviceTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMDeviceTokenReq.Marshal(b, m, deterministic)
}
func (m *IMDeviceTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMDeviceTokenReq.Merge(m, src)
}
func (m *IMDeviceTokenReq) XXX_Size() int {
	return xxx_messageInfo_IMDeviceTokenReq.Size(m)
}
func (m *IMDeviceTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMDeviceTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMDeviceTokenReq proto.InternalMessageInfo

func (m *IMDeviceTokenReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMDeviceTokenReq) GetDeviceToken() string {
	if m != nil {
		return m.DeviceToken
	}
	return ""
}

func (m *IMDeviceTokenReq) GetClientType() ClientType {
	if m != nil {
		return m.ClientType
	}
	return ClientType_CLIENT_TYPE_NONE
}

func (m *IMDeviceTokenReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMDeviceTokenRsp struct {
	//cmd id: 		0x0109
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMDeviceTokenRsp) Reset()         { *m = IMDeviceTokenRsp{} }
func (m *IMDeviceTokenRsp) String() string { return proto.CompactTextString(m) }
func (*IMDeviceTokenRsp) ProtoMessage()    {}
func (*IMDeviceTokenRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{8}
}

func (m *IMDeviceTokenRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMDeviceTokenRsp.Unmarshal(m, b)
}
func (m *IMDeviceTokenRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMDeviceTokenRsp.Marshal(b, m, deterministic)
}
func (m *IMDeviceTokenRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMDeviceTokenRsp.Merge(m, src)
}
func (m *IMDeviceTokenRsp) XXX_Size() int {
	return xxx_messageInfo_IMDeviceTokenRsp.Size(m)
}
func (m *IMDeviceTokenRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMDeviceTokenRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMDeviceTokenRsp proto.InternalMessageInfo

func (m *IMDeviceTokenRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMDeviceTokenRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

//只给移动端请求
type IMKickPCClientReq struct {
	//cmd id:		0x010a
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMKickPCClientReq) Reset()         { *m = IMKickPCClientReq{} }
func (m *IMKickPCClientReq) String() string { return proto.CompactTextString(m) }
func (*IMKickPCClientReq) ProtoMessage()    {}
func (*IMKickPCClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{9}
}

func (m *IMKickPCClientReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMKickPCClientReq.Unmarshal(m, b)
}
func (m *IMKickPCClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMKickPCClientReq.Marshal(b, m, deterministic)
}
func (m *IMKickPCClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMKickPCClientReq.Merge(m, src)
}
func (m *IMKickPCClientReq) XXX_Size() int {
	return xxx_messageInfo_IMKickPCClientReq.Size(m)
}
func (m *IMKickPCClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMKickPCClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMKickPCClientReq proto.InternalMessageInfo

func (m *IMKickPCClientReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type IMKickPCClientRsp struct {
	//cmd id: 		0x010b
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResultCode           uint32   `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMKickPCClientRsp) Reset()         { *m = IMKickPCClientRsp{} }
func (m *IMKickPCClientRsp) String() string { return proto.CompactTextString(m) }
func (*IMKickPCClientRsp) ProtoMessage()    {}
func (*IMKickPCClientRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{10}
}

func (m *IMKickPCClientRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMKickPCClientRsp.Unmarshal(m, b)
}
func (m *IMKickPCClientRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMKickPCClientRsp.Marshal(b, m, deterministic)
}
func (m *IMKickPCClientRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMKickPCClientRsp.Merge(m, src)
}
func (m *IMKickPCClientRsp) XXX_Size() int {
	return xxx_messageInfo_IMKickPCClientRsp.Size(m)
}
func (m *IMKickPCClientRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMKickPCClientRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMKickPCClientRsp proto.InternalMessageInfo

func (m *IMKickPCClientRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMKickPCClientRsp) GetResultCode() uint32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

// 一旦设置以后，22:00 -- 07:00不发送
type IMPushShieldReq struct {
	//cmd id:			0x010c
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShieldStatus         uint32   `protobuf:"varint,2,opt,name=shield_status,json=shieldStatus,proto3" json:"shield_status,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMPushShieldReq) Reset()         { *m = IMPushShieldReq{} }
func (m *IMPushShieldReq) String() string { return proto.CompactTextString(m) }
func (*IMPushShieldReq) ProtoMessage()    {}
func (*IMPushShieldReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{11}
}

func (m *IMPushShieldReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMPushShieldReq.Unmarshal(m, b)
}
func (m *IMPushShieldReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMPushShieldReq.Marshal(b, m, deterministic)
}
func (m *IMPushShieldReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMPushShieldReq.Merge(m, src)
}
func (m *IMPushShieldReq) XXX_Size() int {
	return xxx_messageInfo_IMPushShieldReq.Size(m)
}
func (m *IMPushShieldReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMPushShieldReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMPushShieldReq proto.InternalMessageInfo

func (m *IMPushShieldReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMPushShieldReq) GetShieldStatus() uint32 {
	if m != nil {
		return m.ShieldStatus
	}
	return 0
}

func (m *IMPushShieldReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMPushShieldRsp struct {
	//cmd id:			0x010d
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResultCode           uint32   `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	ShieldStatus         uint32   `protobuf:"varint,3,opt,name=shield_status,json=shieldStatus,proto3" json:"shield_status,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMPushShieldRsp) Reset()         { *m = IMPushShieldRsp{} }
func (m *IMPushShieldRsp) String() string { return proto.CompactTextString(m) }
func (*IMPushShieldRsp) ProtoMessage()    {}
func (*IMPushShieldRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{12}
}

func (m *IMPushShieldRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMPushShieldRsp.Unmarshal(m, b)
}
func (m *IMPushShieldRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMPushShieldRsp.Marshal(b, m, deterministic)
}
func (m *IMPushShieldRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMPushShieldRsp.Merge(m, src)
}
func (m *IMPushShieldRsp) XXX_Size() int {
	return xxx_messageInfo_IMPushShieldRsp.Size(m)
}
func (m *IMPushShieldRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMPushShieldRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMPushShieldRsp proto.InternalMessageInfo

func (m *IMPushShieldRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMPushShieldRsp) GetResultCode() uint32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *IMPushShieldRsp) GetShieldStatus() uint32 {
	if m != nil {
		return m.ShieldStatus
	}
	return 0
}

func (m *IMPushShieldRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

// 如果用户重新安装app，第一次启动登录成功后，app主动查询
// 服务端返回IMQueryPushShieldRsp
type IMQueryPushShieldReq struct {
	//cmd id:			0x010e
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMQueryPushShieldReq) Reset()         { *m = IMQueryPushShieldReq{} }
func (m *IMQueryPushShieldReq) String() string { return proto.CompactTextString(m) }
func (*IMQueryPushShieldReq) ProtoMessage()    {}
func (*IMQueryPushShieldReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{13}
}

func (m *IMQueryPushShieldReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMQueryPushShieldReq.Unmarshal(m, b)
}
func (m *IMQueryPushShieldReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMQueryPushShieldReq.Marshal(b, m, deterministic)
}
func (m *IMQueryPushShieldReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMQueryPushShieldReq.Merge(m, src)
}
func (m *IMQueryPushShieldReq) XXX_Size() int {
	return xxx_messageInfo_IMQueryPushShieldReq.Size(m)
}
func (m *IMQueryPushShieldReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMQueryPushShieldReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMQueryPushShieldReq proto.InternalMessageInfo

func (m *IMQueryPushShieldReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMQueryPushShieldReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMQueryPushShieldRsp struct {
	//cmd id:			0x010f
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResultCode           uint32   `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	ShieldStatus         uint32   `protobuf:"varint,3,opt,name=shield_status,json=shieldStatus,proto3" json:"shield_status,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMQueryPushShieldRsp) Reset()         { *m = IMQueryPushShieldRsp{} }
func (m *IMQueryPushShieldRsp) String() string { return proto.CompactTextString(m) }
func (*IMQueryPushShieldRsp) ProtoMessage()    {}
func (*IMQueryPushShieldRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4474ae06df949ed, []int{14}
}

func (m *IMQueryPushShieldRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMQueryPushShieldRsp.Unmarshal(m, b)
}
func (m *IMQueryPushShieldRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMQueryPushShieldRsp.Marshal(b, m, deterministic)
}
func (m *IMQueryPushShieldRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMQueryPushShieldRsp.Merge(m, src)
}
func (m *IMQueryPushShieldRsp) XXX_Size() int {
	return xxx_messageInfo_IMQueryPushShieldRsp.Size(m)
}
func (m *IMQueryPushShieldRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMQueryPushShieldRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMQueryPushShieldRsp proto.InternalMessageInfo

func (m *IMQueryPushShieldRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMQueryPushShieldRsp) GetResultCode() uint32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *IMQueryPushShieldRsp) GetShieldStatus() uint32 {
	if m != nil {
		return m.ShieldStatus
	}
	return 0
}

func (m *IMQueryPushShieldRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

func init() {
	proto.RegisterType((*IMMsgServReq)(nil), "IM.Login.IMMsgServReq")
	proto.RegisterType((*IMMsgServRsp)(nil), "IM.Login.IMMsgServRsp")
	proto.RegisterType((*IMLoginReq)(nil), "IM.Login.IMLoginReq")
	proto.RegisterType((*IMLoginRes)(nil), "IM.Login.IMLoginRes")
	proto.RegisterType((*IMLogoutReq)(nil), "IM.Login.IMLogoutReq")
	proto.RegisterType((*IMLogoutRsp)(nil), "IM.Login.IMLogoutRsp")
	proto.RegisterType((*IMKickUser)(nil), "IM.Login.IMKickUser")
	proto.RegisterType((*IMDeviceTokenReq)(nil), "IM.Login.IMDeviceTokenReq")
	proto.RegisterType((*IMDeviceTokenRsp)(nil), "IM.Login.IMDeviceTokenRsp")
	proto.RegisterType((*IMKickPCClientReq)(nil), "IM.Login.IMKickPCClientReq")
	proto.RegisterType((*IMKickPCClientRsp)(nil), "IM.Login.IMKickPCClientRsp")
	proto.RegisterType((*IMPushShieldReq)(nil), "IM.Login.IMPushShieldReq")
	proto.RegisterType((*IMPushShieldRsp)(nil), "IM.Login.IMPushShieldRsp")
	proto.RegisterType((*IMQueryPushShieldReq)(nil), "IM.Login.IMQueryPushShieldReq")
	proto.RegisterType((*IMQueryPushShieldRsp)(nil), "IM.Login.IMQueryPushShieldRsp")
}

func init() { proto.RegisterFile("IM.Login.proto", fileDescriptor_e4474ae06df949ed) }

var fileDescriptor_e4474ae06df949ed = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x95, 0x81, 0x0f, 0xc2, 0x75, 0x9c, 0xaf, 0x9d, 0x22, 0x11, 0x40, 0x15, 0xa9, 0x51, 0xa5,
	0x2c, 0x2a, 0x2f, 0x68, 0x57, 0x5d, 0x54, 0x15, 0xb0, 0xb1, 0x8a, 0x2b, 0xea, 0xd0, 0x6e, 0xad,
	0xc1, 0xbe, 0x09, 0xd3, 0x24, 0x9e, 0x61, 0x66, 0x9c, 0x8a, 0x17, 0xe8, 0xbe, 0x1b, 0x5e, 0xa2,
	0x6f, 0xd6, 0xa7, 0xa8, 0x3c, 0x63, 0x68, 0x70, 0x80, 0x34, 0xed, 0xa2, 0xbb, 0xb9, 0xc7, 0xf7,
	0xe7, 0x9c, 0x73, 0xaf, 0x64, 0x68, 0x85, 0x51, 0x70, 0xcc, 0x07, 0x2c, 0x0f, 0x84, 0xe4, 0x9a,
	0x93, 0xc6, 0x75, 0xbc, 0xfd, 0x24, 0x8c, 0x82, 0x03, 0xaa, 0xf0, 0x08, 0xfb, 0x2c, 0x47, 0xfb,
	0xd9, 0x6f, 0x41, 0x33, 0x8c, 0x22, 0x35, 0xe8, 0xa1, 0x9c, 0xc4, 0x78, 0xe1, 0x5f, 0x39, 0xd3,
	0x80, 0x12, 0xe4, 0x35, 0xb8, 0x12, 0x55, 0x31, 0xd2, 0x49, 0xca, 0x33, 0x6c, 0x3b, 0x1d, 0xa7,
	0xdb, 0xda, 0xdf, 0x0a, 0x6e, 0xf7, 0x8a, 0x4d, 0xc6, 0xe9, 0xa5, 0xc0, 0x18, 0x6c, 0xf6, 0x21,
	0xcf, 0x90, 0x6c, 0x41, 0x43, 0x48, 0xc6, 0x65, 0xc2, 0x44, 0x7b, 0xa9, 0xe3, 0x74, 0xd7, 0xe3,
	0x35, 0x13, 0x87, 0x82, 0xec, 0xc0, 0xfa, 0x19, 0x4d, 0x87, 0x4c, 0x94, 0xdf, 0x96, 0xcd, 0xb7,
	0x86, 0x05, 0x42, 0x41, 0x08, 0xac, 0x08, 0x2e, 0x75, 0x7b, 0xa5, 0xe3, 0x74, 0xbd, 0xd8, 0xbc,
	0xfd, 0x1f, 0x0e, 0x40, 0x18, 0x19, 0x25, 0x31, 0x5e, 0x94, 0xf5, 0x85, 0x42, 0x99, 0xe4, 0x74,
	0x6c, 0x49, 0xad, 0xc7, 0x8d, 0x12, 0x78, 0x4f, 0xc7, 0x48, 0xb6, 0xa1, 0x21, 0xa8, 0x52, 0x5f,
	0xb8, 0xcc, 0xaa, 0xb9, 0x37, 0x31, 0x79, 0x0b, 0x1e, 0xcf, 0x47, 0x2c, 0xc7, 0x44, 0x69, 0xaa,
	0x0b, 0x65, 0x86, 0xb7, 0xf6, 0x77, 0x6a, 0x8a, 0x3e, 0x2a, 0x94, 0x3d, 0x4d, 0xad, 0xa6, 0xa6,
	0xad, 0xe8, 0x99, 0x82, 0xd2, 0x91, 0x74, 0xc4, 0x30, 0xd7, 0x89, 0xbe, 0x14, 0x68, 0x48, 0xce,
	0x3a, 0x72, 0x68, 0x32, 0xac, 0x23, 0xe9, 0xcd, 0x9b, 0x3c, 0x87, 0x56, 0x55, 0x3b, 0x41, 0xa9,
	0x18, 0xcf, 0xdb, 0xff, 0x19, 0x7e, 0x9e, 0x45, 0x3f, 0x59, 0xd0, 0xff, 0xba, 0x34, 0x25, 0x56,
	0x91, 0x5d, 0x70, 0x15, 0xca, 0x09, 0xca, 0x44, 0xb3, 0x4a, 0xae, 0x17, 0x83, 0x85, 0x4e, 0xd9,
	0x18, 0xeb, 0x4b, 0x5a, 0x5a, 0x64, 0x49, 0x7b, 0xe0, 0x55, 0xb5, 0x4a, 0x4b, 0x96, 0x0f, 0xaa,
	0x6d, 0x34, 0x2d, 0xd8, 0x33, 0xd8, 0xac, 0x6b, 0x2b, 0x8b, 0xba, 0xf6, 0xaa, 0x5a, 0x18, 0xcb,
	0xfb, 0xdc, 0x88, 0x76, 0xf7, 0x37, 0xef, 0xa8, 0x0e, 0xf3, 0x3e, 0xb7, 0x9b, 0x2c, 0x5f, 0xbe,
	0x07, 0xae, 0xf1, 0x81, 0x17, 0xba, 0xbc, 0xce, 0x60, 0x2a, 0x54, 0xa2, 0xf4, 0xa5, 0x7e, 0x9b,
	0xde, 0xb4, 0x36, 0x1f, 0x4b, 0x1b, 0xdf, 0xb1, 0x74, 0x58, 0xb6, 0x26, 0x9b, 0xb0, 0x66, 0x29,
	0x64, 0x55, 0xea, 0xaa, 0x99, 0x93, 0x91, 0x37, 0xe0, 0x0e, 0x59, 0x3a, 0x4c, 0x24, 0x52, 0xc5,
	0xf3, 0xca, 0xbe, 0xa7, 0x35, 0x76, 0x65, 0x9b, 0xd8, 0x24, 0x58, 0x0b, 0x87, 0x37, 0xb1, 0xff,
	0xdd, 0x81, 0x47, 0x61, 0x74, 0x84, 0x13, 0x96, 0xe2, 0x29, 0x1f, 0xa2, 0xb9, 0xd0, 0x7b, 0xa7,
	0x3d, 0x83, 0x66, 0x66, 0x52, 0x13, 0x5d, 0xe6, 0x56, 0x17, 0xea, 0x66, 0xbf, 0xca, 0xeb, 0x27,
	0xb6, 0xbc, 0xc8, 0x89, 0xed, 0x82, 0x4b, 0xb5, 0xa6, 0xe9, 0x79, 0x92, 0x51, 0x4d, 0xdb, 0x1b,
	0x1d, 0xa7, 0xdb, 0x8c, 0xc1, 0x42, 0x47, 0x54, 0x53, 0xff, 0xb8, 0x4e, 0x56, 0x89, 0xfb, 0xc9,
	0xce, 0xed, 0xf6, 0x02, 0x1e, 0x5b, 0x8b, 0x4f, 0x0e, 0x2d, 0xa1, 0x87, 0xb4, 0xfb, 0xd1, 0x4c,
	0xf6, 0x9c, 0xe1, 0xf5, 0xb3, 0xbe, 0xbd, 0x5f, 0x01, 0xff, 0x87, 0xd1, 0x49, 0xa1, 0xce, 0x7b,
	0xe7, 0x0c, 0x47, 0xd9, 0x83, 0xb6, 0xef, 0x81, 0xa7, 0x4c, 0xd6, 0xf5, 0x09, 0xdb, 0x76, 0x4d,
	0x0b, 0x56, 0x57, 0x3a, 0x57, 0xee, 0x37, 0xa7, 0x36, 0xf2, 0x6f, 0xf8, 0xcf, 0x72, 0x5a, 0xfe,
	0x13, 0x4e, 0x27, 0xb0, 0x11, 0x46, 0x1f, 0x0a, 0x94, 0x97, 0xbf, 0x69, 0xc5, 0xdc, 0x8e, 0x57,
	0xce, 0x5d, 0x2d, 0xff, 0xbd, 0xd4, 0x83, 0x2d, 0xd8, 0x4c, 0xf9, 0x38, 0x18, 0xf3, 0x41, 0xf1,
	0x99, 0x61, 0xa0, 0xb5, 0xfd, 0x8d, 0x9d, 0x15, 0xfd, 0xb3, 0x55, 0xf3, 0x7a, 0xf9, 0x33, 0x00,
	0x00, 0xff, 0xff, 0xff, 0xf7, 0xb8, 0x32, 0x01, 0x07, 0x00, 0x00,
}
