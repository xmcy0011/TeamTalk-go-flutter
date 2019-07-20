// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IM.Group.proto

package IM_Group

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

type IMNormalGroupListReq struct {
	//cmd id:			0x0401
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMNormalGroupListReq) Reset()         { *m = IMNormalGroupListReq{} }
func (m *IMNormalGroupListReq) String() string { return proto.CompactTextString(m) }
func (*IMNormalGroupListReq) ProtoMessage()    {}
func (*IMNormalGroupListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{0}
}

func (m *IMNormalGroupListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMNormalGroupListReq.Unmarshal(m, b)
}
func (m *IMNormalGroupListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMNormalGroupListReq.Marshal(b, m, deterministic)
}
func (m *IMNormalGroupListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMNormalGroupListReq.Merge(m, src)
}
func (m *IMNormalGroupListReq) XXX_Size() int {
	return xxx_messageInfo_IMNormalGroupListReq.Size(m)
}
func (m *IMNormalGroupListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMNormalGroupListReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMNormalGroupListReq proto.InternalMessageInfo

func (m *IMNormalGroupListReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMNormalGroupListReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMNormalGroupListRsp struct {
	//cmd id:			0x0402
	UserId               uint32              `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupVersionList     []*GroupVersionInfo `protobuf:"bytes,2,rep,name=group_version_list,json=groupVersionList,proto3" json:"group_version_list,omitempty"`
	AttachData           []byte              `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IMNormalGroupListRsp) Reset()         { *m = IMNormalGroupListRsp{} }
func (m *IMNormalGroupListRsp) String() string { return proto.CompactTextString(m) }
func (*IMNormalGroupListRsp) ProtoMessage()    {}
func (*IMNormalGroupListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{1}
}

func (m *IMNormalGroupListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMNormalGroupListRsp.Unmarshal(m, b)
}
func (m *IMNormalGroupListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMNormalGroupListRsp.Marshal(b, m, deterministic)
}
func (m *IMNormalGroupListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMNormalGroupListRsp.Merge(m, src)
}
func (m *IMNormalGroupListRsp) XXX_Size() int {
	return xxx_messageInfo_IMNormalGroupListRsp.Size(m)
}
func (m *IMNormalGroupListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMNormalGroupListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMNormalGroupListRsp proto.InternalMessageInfo

func (m *IMNormalGroupListRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMNormalGroupListRsp) GetGroupVersionList() []*GroupVersionInfo {
	if m != nil {
		return m.GroupVersionList
	}
	return nil
}

func (m *IMNormalGroupListRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupInfoListReq struct {
	//cmd id:			0x0403
	UserId               uint32              `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupVersionList     []*GroupVersionInfo `protobuf:"bytes,2,rep,name=group_version_list,json=groupVersionList,proto3" json:"group_version_list,omitempty"`
	AttachData           []byte              `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IMGroupInfoListReq) Reset()         { *m = IMGroupInfoListReq{} }
func (m *IMGroupInfoListReq) String() string { return proto.CompactTextString(m) }
func (*IMGroupInfoListReq) ProtoMessage()    {}
func (*IMGroupInfoListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{2}
}

func (m *IMGroupInfoListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupInfoListReq.Unmarshal(m, b)
}
func (m *IMGroupInfoListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupInfoListReq.Marshal(b, m, deterministic)
}
func (m *IMGroupInfoListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupInfoListReq.Merge(m, src)
}
func (m *IMGroupInfoListReq) XXX_Size() int {
	return xxx_messageInfo_IMGroupInfoListReq.Size(m)
}
func (m *IMGroupInfoListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupInfoListReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupInfoListReq proto.InternalMessageInfo

func (m *IMGroupInfoListReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupInfoListReq) GetGroupVersionList() []*GroupVersionInfo {
	if m != nil {
		return m.GroupVersionList
	}
	return nil
}

func (m *IMGroupInfoListReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupInfoListRsp struct {
	//cmd id:			0x0404
	UserId               uint32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupInfoList        []*GroupInfo `protobuf:"bytes,2,rep,name=group_info_list,json=groupInfoList,proto3" json:"group_info_list,omitempty"`
	AttachData           []byte       `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *IMGroupInfoListRsp) Reset()         { *m = IMGroupInfoListRsp{} }
func (m *IMGroupInfoListRsp) String() string { return proto.CompactTextString(m) }
func (*IMGroupInfoListRsp) ProtoMessage()    {}
func (*IMGroupInfoListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{3}
}

func (m *IMGroupInfoListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupInfoListRsp.Unmarshal(m, b)
}
func (m *IMGroupInfoListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupInfoListRsp.Marshal(b, m, deterministic)
}
func (m *IMGroupInfoListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupInfoListRsp.Merge(m, src)
}
func (m *IMGroupInfoListRsp) XXX_Size() int {
	return xxx_messageInfo_IMGroupInfoListRsp.Size(m)
}
func (m *IMGroupInfoListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupInfoListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupInfoListRsp proto.InternalMessageInfo

func (m *IMGroupInfoListRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupInfoListRsp) GetGroupInfoList() []*GroupInfo {
	if m != nil {
		return m.GroupInfoList
	}
	return nil
}

func (m *IMGroupInfoListRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupCreateReq struct {
	//cmd id:			0x0405
	UserId               uint32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupType            GroupType `protobuf:"varint,2,opt,name=group_type,json=groupType,proto3,enum=IM.BaseDefine.GroupType" json:"group_type,omitempty"`
	GroupName            string    `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	GroupAvatar          string    `protobuf:"bytes,4,opt,name=group_avatar,json=groupAvatar,proto3" json:"group_avatar,omitempty"`
	MemberIdList         []uint32  `protobuf:"varint,5,rep,packed,name=member_id_list,json=memberIdList,proto3" json:"member_id_list,omitempty"`
	AttachData           []byte    `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IMGroupCreateReq) Reset()         { *m = IMGroupCreateReq{} }
func (m *IMGroupCreateReq) String() string { return proto.CompactTextString(m) }
func (*IMGroupCreateReq) ProtoMessage()    {}
func (*IMGroupCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{4}
}

func (m *IMGroupCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupCreateReq.Unmarshal(m, b)
}
func (m *IMGroupCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupCreateReq.Marshal(b, m, deterministic)
}
func (m *IMGroupCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupCreateReq.Merge(m, src)
}
func (m *IMGroupCreateReq) XXX_Size() int {
	return xxx_messageInfo_IMGroupCreateReq.Size(m)
}
func (m *IMGroupCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupCreateReq proto.InternalMessageInfo

func (m *IMGroupCreateReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupCreateReq) GetGroupType() GroupType {
	if m != nil {
		return m.GroupType
	}
	return GroupType_GROUP_TYPE_NONE
}

func (m *IMGroupCreateReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *IMGroupCreateReq) GetGroupAvatar() string {
	if m != nil {
		return m.GroupAvatar
	}
	return ""
}

func (m *IMGroupCreateReq) GetMemberIdList() []uint32 {
	if m != nil {
		return m.MemberIdList
	}
	return nil
}

func (m *IMGroupCreateReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupCreateRsp struct {
	//cmd id:			0x0406
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResultCode           uint32   `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	GroupId              uint32   `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	GroupName            string   `protobuf:"bytes,4,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	UserIdList           []uint32 `protobuf:"varint,5,rep,packed,name=user_id_list,json=userIdList,proto3" json:"user_id_list,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMGroupCreateRsp) Reset()         { *m = IMGroupCreateRsp{} }
func (m *IMGroupCreateRsp) String() string { return proto.CompactTextString(m) }
func (*IMGroupCreateRsp) ProtoMessage()    {}
func (*IMGroupCreateRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{5}
}

func (m *IMGroupCreateRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupCreateRsp.Unmarshal(m, b)
}
func (m *IMGroupCreateRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupCreateRsp.Marshal(b, m, deterministic)
}
func (m *IMGroupCreateRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupCreateRsp.Merge(m, src)
}
func (m *IMGroupCreateRsp) XXX_Size() int {
	return xxx_messageInfo_IMGroupCreateRsp.Size(m)
}
func (m *IMGroupCreateRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupCreateRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupCreateRsp proto.InternalMessageInfo

func (m *IMGroupCreateRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupCreateRsp) GetResultCode() uint32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *IMGroupCreateRsp) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IMGroupCreateRsp) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *IMGroupCreateRsp) GetUserIdList() []uint32 {
	if m != nil {
		return m.UserIdList
	}
	return nil
}

func (m *IMGroupCreateRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupChangeMemberReq struct {
	//cmd id:			0x0407
	UserId               uint32          `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChangeType           GroupModifyType `protobuf:"varint,2,opt,name=change_type,json=changeType,proto3,enum=IM.BaseDefine.GroupModifyType" json:"change_type,omitempty"`
	GroupId              uint32          `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	MemberIdList         []uint32        `protobuf:"varint,4,rep,packed,name=member_id_list,json=memberIdList,proto3" json:"member_id_list,omitempty"`
	AttachData           []byte          `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IMGroupChangeMemberReq) Reset()         { *m = IMGroupChangeMemberReq{} }
func (m *IMGroupChangeMemberReq) String() string { return proto.CompactTextString(m) }
func (*IMGroupChangeMemberReq) ProtoMessage()    {}
func (*IMGroupChangeMemberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{6}
}

func (m *IMGroupChangeMemberReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupChangeMemberReq.Unmarshal(m, b)
}
func (m *IMGroupChangeMemberReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupChangeMemberReq.Marshal(b, m, deterministic)
}
func (m *IMGroupChangeMemberReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupChangeMemberReq.Merge(m, src)
}
func (m *IMGroupChangeMemberReq) XXX_Size() int {
	return xxx_messageInfo_IMGroupChangeMemberReq.Size(m)
}
func (m *IMGroupChangeMemberReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupChangeMemberReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupChangeMemberReq proto.InternalMessageInfo

func (m *IMGroupChangeMemberReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupChangeMemberReq) GetChangeType() GroupModifyType {
	if m != nil {
		return m.ChangeType
	}
	return GroupModifyType_GROUP_MODIFY_TYPE_NONE
}

func (m *IMGroupChangeMemberReq) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IMGroupChangeMemberReq) GetMemberIdList() []uint32 {
	if m != nil {
		return m.MemberIdList
	}
	return nil
}

func (m *IMGroupChangeMemberReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupChangeMemberRsp struct {
	//cmd id:			0x0408
	UserId               uint32          `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChangeType           GroupModifyType `protobuf:"varint,2,opt,name=change_type,json=changeType,proto3,enum=IM.BaseDefine.GroupModifyType" json:"change_type,omitempty"`
	ResultCode           uint32          `protobuf:"varint,3,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	GroupId              uint32          `protobuf:"varint,4,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	CurUserIdList        []uint32        `protobuf:"varint,5,rep,packed,name=cur_user_id_list,json=curUserIdList,proto3" json:"cur_user_id_list,omitempty"`
	ChgUserIdList        []uint32        `protobuf:"varint,6,rep,packed,name=chg_user_id_list,json=chgUserIdList,proto3" json:"chg_user_id_list,omitempty"`
	AttachData           []byte          `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IMGroupChangeMemberRsp) Reset()         { *m = IMGroupChangeMemberRsp{} }
func (m *IMGroupChangeMemberRsp) String() string { return proto.CompactTextString(m) }
func (*IMGroupChangeMemberRsp) ProtoMessage()    {}
func (*IMGroupChangeMemberRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{7}
}

func (m *IMGroupChangeMemberRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupChangeMemberRsp.Unmarshal(m, b)
}
func (m *IMGroupChangeMemberRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupChangeMemberRsp.Marshal(b, m, deterministic)
}
func (m *IMGroupChangeMemberRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupChangeMemberRsp.Merge(m, src)
}
func (m *IMGroupChangeMemberRsp) XXX_Size() int {
	return xxx_messageInfo_IMGroupChangeMemberRsp.Size(m)
}
func (m *IMGroupChangeMemberRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupChangeMemberRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupChangeMemberRsp proto.InternalMessageInfo

func (m *IMGroupChangeMemberRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupChangeMemberRsp) GetChangeType() GroupModifyType {
	if m != nil {
		return m.ChangeType
	}
	return GroupModifyType_GROUP_MODIFY_TYPE_NONE
}

func (m *IMGroupChangeMemberRsp) GetResultCode() uint32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *IMGroupChangeMemberRsp) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IMGroupChangeMemberRsp) GetCurUserIdList() []uint32 {
	if m != nil {
		return m.CurUserIdList
	}
	return nil
}

func (m *IMGroupChangeMemberRsp) GetChgUserIdList() []uint32 {
	if m != nil {
		return m.ChgUserIdList
	}
	return nil
}

func (m *IMGroupChangeMemberRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupShieldReq struct {
	//cmd id:			0x0409
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId              uint32   `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ShieldStatus         uint32   `protobuf:"varint,3,opt,name=shield_status,json=shieldStatus,proto3" json:"shield_status,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMGroupShieldReq) Reset()         { *m = IMGroupShieldReq{} }
func (m *IMGroupShieldReq) String() string { return proto.CompactTextString(m) }
func (*IMGroupShieldReq) ProtoMessage()    {}
func (*IMGroupShieldReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{8}
}

func (m *IMGroupShieldReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupShieldReq.Unmarshal(m, b)
}
func (m *IMGroupShieldReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupShieldReq.Marshal(b, m, deterministic)
}
func (m *IMGroupShieldReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupShieldReq.Merge(m, src)
}
func (m *IMGroupShieldReq) XXX_Size() int {
	return xxx_messageInfo_IMGroupShieldReq.Size(m)
}
func (m *IMGroupShieldReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupShieldReq.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupShieldReq proto.InternalMessageInfo

func (m *IMGroupShieldReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupShieldReq) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IMGroupShieldReq) GetShieldStatus() uint32 {
	if m != nil {
		return m.ShieldStatus
	}
	return 0
}

func (m *IMGroupShieldReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupShieldRsp struct {
	//cmd id:			0x040a
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId              uint32   `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ResultCode           uint32   `protobuf:"varint,3,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	AttachData           []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData,proto3" json:"attach_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IMGroupShieldRsp) Reset()         { *m = IMGroupShieldRsp{} }
func (m *IMGroupShieldRsp) String() string { return proto.CompactTextString(m) }
func (*IMGroupShieldRsp) ProtoMessage()    {}
func (*IMGroupShieldRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{9}
}

func (m *IMGroupShieldRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupShieldRsp.Unmarshal(m, b)
}
func (m *IMGroupShieldRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupShieldRsp.Marshal(b, m, deterministic)
}
func (m *IMGroupShieldRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupShieldRsp.Merge(m, src)
}
func (m *IMGroupShieldRsp) XXX_Size() int {
	return xxx_messageInfo_IMGroupShieldRsp.Size(m)
}
func (m *IMGroupShieldRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupShieldRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupShieldRsp proto.InternalMessageInfo

func (m *IMGroupShieldRsp) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupShieldRsp) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IMGroupShieldRsp) GetResultCode() uint32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *IMGroupShieldRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupChangeMemberNotify struct {
	//cmd id: 			0x040b
	UserId               uint32          `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChangeType           GroupModifyType `protobuf:"varint,2,opt,name=change_type,json=changeType,proto3,enum=IM.BaseDefine.GroupModifyType" json:"change_type,omitempty"`
	GroupId              uint32          `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	CurUserIdList        []uint32        `protobuf:"varint,4,rep,packed,name=cur_user_id_list,json=curUserIdList,proto3" json:"cur_user_id_list,omitempty"`
	ChgUserIdList        []uint32        `protobuf:"varint,5,rep,packed,name=chg_user_id_list,json=chgUserIdList,proto3" json:"chg_user_id_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IMGroupChangeMemberNotify) Reset()         { *m = IMGroupChangeMemberNotify{} }
func (m *IMGroupChangeMemberNotify) String() string { return proto.CompactTextString(m) }
func (*IMGroupChangeMemberNotify) ProtoMessage()    {}
func (*IMGroupChangeMemberNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c8c45c9c95134d, []int{10}
}

func (m *IMGroupChangeMemberNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IMGroupChangeMemberNotify.Unmarshal(m, b)
}
func (m *IMGroupChangeMemberNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IMGroupChangeMemberNotify.Marshal(b, m, deterministic)
}
func (m *IMGroupChangeMemberNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IMGroupChangeMemberNotify.Merge(m, src)
}
func (m *IMGroupChangeMemberNotify) XXX_Size() int {
	return xxx_messageInfo_IMGroupChangeMemberNotify.Size(m)
}
func (m *IMGroupChangeMemberNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_IMGroupChangeMemberNotify.DiscardUnknown(m)
}

var xxx_messageInfo_IMGroupChangeMemberNotify proto.InternalMessageInfo

func (m *IMGroupChangeMemberNotify) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IMGroupChangeMemberNotify) GetChangeType() GroupModifyType {
	if m != nil {
		return m.ChangeType
	}
	return GroupModifyType_GROUP_MODIFY_TYPE_NONE
}

func (m *IMGroupChangeMemberNotify) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IMGroupChangeMemberNotify) GetCurUserIdList() []uint32 {
	if m != nil {
		return m.CurUserIdList
	}
	return nil
}

func (m *IMGroupChangeMemberNotify) GetChgUserIdList() []uint32 {
	if m != nil {
		return m.ChgUserIdList
	}
	return nil
}

func init() {
	proto.RegisterType((*IMNormalGroupListReq)(nil), "IM.Group.IMNormalGroupListReq")
	proto.RegisterType((*IMNormalGroupListRsp)(nil), "IM.Group.IMNormalGroupListRsp")
	proto.RegisterType((*IMGroupInfoListReq)(nil), "IM.Group.IMGroupInfoListReq")
	proto.RegisterType((*IMGroupInfoListRsp)(nil), "IM.Group.IMGroupInfoListRsp")
	proto.RegisterType((*IMGroupCreateReq)(nil), "IM.Group.IMGroupCreateReq")
	proto.RegisterType((*IMGroupCreateRsp)(nil), "IM.Group.IMGroupCreateRsp")
	proto.RegisterType((*IMGroupChangeMemberReq)(nil), "IM.Group.IMGroupChangeMemberReq")
	proto.RegisterType((*IMGroupChangeMemberRsp)(nil), "IM.Group.IMGroupChangeMemberRsp")
	proto.RegisterType((*IMGroupShieldReq)(nil), "IM.Group.IMGroupShieldReq")
	proto.RegisterType((*IMGroupShieldRsp)(nil), "IM.Group.IMGroupShieldRsp")
	proto.RegisterType((*IMGroupChangeMemberNotify)(nil), "IM.Group.IMGroupChangeMemberNotify")
}

func init() { proto.RegisterFile("IM.Group.proto", fileDescriptor_72c8c45c9c95134d) }

var fileDescriptor_72c8c45c9c95134d = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0xc1, 0x8e, 0x93, 0x40,
	0x18, 0xc7, 0x43, 0xdb, 0xed, 0xee, 0x7e, 0x2d, 0xb5, 0xc1, 0x8d, 0x4b, 0x4d, 0xb4, 0xb8, 0x9a,
	0xc8, 0x89, 0x83, 0x1e, 0x3c, 0xaa, 0xbb, 0x9b, 0x18, 0x12, 0x69, 0x0c, 0xeb, 0x7a, 0x25, 0x53,
	0x18, 0x28, 0xa6, 0x30, 0x15, 0x86, 0x4d, 0xfa, 0x06, 0x1e, 0x3c, 0xf8, 0x02, 0xc6, 0xf8, 0x38,
	0xde, 0x7d, 0x02, 0xef, 0xbe, 0x83, 0x61, 0x86, 0x66, 0x69, 0x3b, 0x0e, 0x55, 0x13, 0xbd, 0xc1,
	0x37, 0x1f, 0xdf, 0xfc, 0xfe, 0xf3, 0x9b, 0xa6, 0x30, 0xb0, 0x1d, 0xeb, 0x45, 0x46, 0x8a, 0x85,
	0xb5, 0xc8, 0x08, 0x25, 0xda, 0xc1, 0xea, 0xfd, 0xf6, 0x4d, 0xdb, 0xb1, 0x4e, 0x51, 0x8e, 0xcf,
	0x71, 0x18, 0xa7, 0x98, 0x2f, 0x9f, 0xbc, 0x82, 0x23, 0xdb, 0x99, 0x90, 0x2c, 0x41, 0x73, 0xd6,
	0xf5, 0x32, 0xce, 0xa9, 0x8b, 0xdf, 0x69, 0xc7, 0xb0, 0x5f, 0xe4, 0x38, 0xf3, 0xe2, 0x40, 0x57,
	0x0c, 0xc5, 0x54, 0xdd, 0x6e, 0xf9, 0x6a, 0x07, 0xda, 0x18, 0x7a, 0x88, 0x52, 0xe4, 0xcf, 0xbc,
	0x00, 0x51, 0xa4, 0x1f, 0x19, 0x8a, 0xd9, 0x77, 0x81, 0x97, 0xce, 0x11, 0x45, 0x27, 0x9f, 0x15,
	0xd1, 0xc8, 0x7c, 0xf1, 0xeb, 0x91, 0x0e, 0x68, 0x51, 0xd9, 0xe8, 0x5d, 0xe1, 0x2c, 0x8f, 0x49,
	0xea, 0xcd, 0xe3, 0x9c, 0xea, 0x2d, 0xa3, 0x6d, 0xf6, 0x1e, 0x8d, 0xad, 0x75, 0x6a, 0x36, 0xf1,
	0x0d, 0xef, 0xb3, 0xd3, 0x90, 0xb8, 0xc3, 0xa8, 0x56, 0x29, 0xb7, 0x6a, 0x26, 0xfc, 0xa4, 0x80,
	0x66, 0x3b, 0x6c, 0x52, 0x39, 0xa2, 0x31, 0xf2, 0xbf, 0xe6, 0xfb, 0x28, 0xe0, 0x93, 0x9d, 0xdf,
	0x33, 0xb8, 0xc1, 0xf9, 0xe2, 0x34, 0x24, 0x75, 0x38, 0x5d, 0x04, 0xc7, 0xa8, 0xd4, 0xa8, 0x3e,
	0xbd, 0x19, 0xe9, 0x87, 0x02, 0xc3, 0x0a, 0xe9, 0x2c, 0xc3, 0x88, 0x62, 0xe9, 0x81, 0x3d, 0x01,
	0xe0, 0x40, 0x74, 0xb9, 0xc0, 0x7a, 0xcb, 0x50, 0xcc, 0x81, 0x98, 0xe5, 0xf5, 0x72, 0x81, 0xdd,
	0xc3, 0x68, 0xf5, 0xa8, 0xdd, 0x59, 0x7d, 0x98, 0xa2, 0x04, 0xeb, 0x6d, 0x43, 0x31, 0x0f, 0xab,
	0xe5, 0x09, 0x4a, 0xb0, 0x76, 0x0f, 0xfa, 0x7c, 0x19, 0x5d, 0x21, 0x8a, 0x32, 0xbd, 0xc3, 0x1a,
	0x7a, 0xac, 0xf6, 0x9c, 0x95, 0xb4, 0x07, 0x30, 0x48, 0x70, 0x32, 0x65, 0x54, 0xfc, 0x28, 0xf6,
	0x8c, 0xb6, 0xa9, 0xba, 0x7d, 0x5e, 0xb5, 0x83, 0xdd, 0xf2, 0x7e, 0xdd, 0xca, 0x2b, 0x13, 0x30,
	0x86, 0x5e, 0x86, 0xf3, 0x62, 0x4e, 0x3d, 0x9f, 0x04, 0x3c, 0xb0, 0xea, 0x02, 0x2f, 0x9d, 0x91,
	0x00, 0x6b, 0x23, 0x38, 0xa8, 0x0c, 0x05, 0x2c, 0x95, 0xea, 0xee, 0x73, 0x01, 0xc1, 0x46, 0xe4,
	0xce, 0x66, 0x64, 0x03, 0xfa, 0xd5, 0x9e, 0xf5, 0x34, 0xc0, 0x37, 0xde, 0x2d, 0xcb, 0x37, 0x05,
	0x6e, 0xad, 0xb2, 0xcc, 0x50, 0x1a, 0x61, 0x87, 0x1d, 0x85, 0xd4, 0xe0, 0x53, 0xe8, 0xf9, 0xac,
	0xb7, 0xae, 0xf0, 0xae, 0x48, 0xa1, 0x43, 0x82, 0x38, 0x5c, 0x32, 0x91, 0xc0, 0x3f, 0x61, 0x26,
	0x25, 0x89, 0xb7, 0x15, 0x75, 0xfe, 0x44, 0xd1, 0x97, 0x96, 0x38, 0x96, 0x4c, 0xd4, 0x5f, 0xc7,
	0xda, 0x30, 0xdd, 0x96, 0x9a, 0xee, 0xac, 0xe7, 0x7e, 0x08, 0x43, 0xbf, 0xc8, 0x3c, 0x81, 0x4e,
	0xd5, 0x2f, 0xb2, 0xcb, 0x6b, 0xa3, 0x65, 0xe3, 0x2c, 0x5a, 0x6f, 0xec, 0x56, 0x8d, 0xb3, 0xe8,
	0xf2, 0x37, 0xd4, 0x7f, 0xb8, 0xbe, 0xc6, 0x17, 0xb3, 0x18, 0xcf, 0x03, 0xa9, 0xf4, 0x3a, 0x7b,
	0x6b, 0x9d, 0xfd, 0x3e, 0xa8, 0x39, 0x1b, 0xe0, 0xe5, 0x14, 0xd1, 0x22, 0xaf, 0x92, 0xf7, 0x79,
	0xf1, 0x82, 0xd5, 0x9a, 0x71, 0xde, 0x6f, 0xe1, 0xc8, 0x64, 0x49, 0x70, 0x1a, 0x35, 0x34, 0xa2,
	0x7c, 0x57, 0x60, 0x24, 0xb8, 0x3d, 0x13, 0x42, 0xe3, 0x70, 0xf9, 0x7f, 0x7e, 0x17, 0xa2, 0xfb,
	0xd1, 0xd9, 0xf5, 0x7e, 0xec, 0x09, 0xee, 0xc7, 0xe9, 0x08, 0x8e, 0x7d, 0x92, 0x58, 0x09, 0x89,
	0x8a, 0xb7, 0x31, 0xb6, 0x28, 0xe5, 0x7f, 0xfa, 0xd3, 0x22, 0x9c, 0x76, 0xd9, 0xd3, 0xe3, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xd3, 0x11, 0xae, 0x2f, 0x08, 0x00, 0x00,
}
