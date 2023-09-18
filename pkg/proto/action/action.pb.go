// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: pkg/proto/action/action.proto

package action

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	player_info "mafia/pkg/proto/player_info"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActionType int32

const (
	ActionType_INIT         ActionType = 0
	ActionType_VOTE         ActionType = 1
	ActionType_CHECK_ROLE   ActionType = 2
	ActionType_PUBLISH_ROLE ActionType = 3
	ActionType_KILL         ActionType = 4
	ActionType_SEND_MESSAGE ActionType = 5
	ActionType_GO_TO_SLEEP  ActionType = 6
	ActionType_DISCONNECT   ActionType = 7
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "INIT",
		1: "VOTE",
		2: "CHECK_ROLE",
		3: "PUBLISH_ROLE",
		4: "KILL",
		5: "SEND_MESSAGE",
		6: "GO_TO_SLEEP",
		7: "DISCONNECT",
	}
	ActionType_value = map[string]int32{
		"INIT":         0,
		"VOTE":         1,
		"CHECK_ROLE":   2,
		"PUBLISH_ROLE": 3,
		"KILL":         4,
		"SEND_MESSAGE": 5,
		"GO_TO_SLEEP":  6,
		"DISCONNECT":   7,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_action_action_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_pkg_proto_action_action_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{0}
}

type InitAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *player_info.PlayerInfo `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *InitAction) Reset() {
	*x = InitAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_action_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitAction) ProtoMessage() {}

func (x *InitAction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_action_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitAction.ProtoReflect.Descriptor instead.
func (*InitAction) Descriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{0}
}

func (x *InitAction) GetPlayer() *player_info.PlayerInfo {
	if x != nil {
		return x.Player
	}
	return nil
}

type VoteAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *player_info.PlayerInfo `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *VoteAction) Reset() {
	*x = VoteAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_action_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteAction) ProtoMessage() {}

func (x *VoteAction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_action_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteAction.ProtoReflect.Descriptor instead.
func (*VoteAction) Descriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{1}
}

func (x *VoteAction) GetPlayer() *player_info.PlayerInfo {
	if x != nil {
		return x.Player
	}
	return nil
}

type CheckRoleAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *player_info.PlayerInfo `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *CheckRoleAction) Reset() {
	*x = CheckRoleAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_action_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRoleAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRoleAction) ProtoMessage() {}

func (x *CheckRoleAction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_action_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRoleAction.ProtoReflect.Descriptor instead.
func (*CheckRoleAction) Descriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{2}
}

func (x *CheckRoleAction) GetPlayer() *player_info.PlayerInfo {
	if x != nil {
		return x.Player
	}
	return nil
}

type PublishRoleAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *player_info.PlayerInfo `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *PublishRoleAction) Reset() {
	*x = PublishRoleAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_action_action_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRoleAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRoleAction) ProtoMessage() {}

func (x *PublishRoleAction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_action_action_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRoleAction.ProtoReflect.Descriptor instead.
func (*PublishRoleAction) Descriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{3}
}

func (x *PublishRoleAction) GetPlayer() *player_info.PlayerInfo {
	if x != nil {
		return x.Player
	}
	return nil
}

type KillAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *player_info.PlayerInfo `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *KillAction) Reset() {
	*x = KillAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_action_action_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillAction) ProtoMessage() {}

func (x *KillAction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_action_action_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillAction.ProtoReflect.Descriptor instead.
func (*KillAction) Descriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{4}
}

func (x *KillAction) GetPlayer() *player_info.PlayerInfo {
	if x != nil {
		return x.Player
	}
	return nil
}

type SendMessageAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageAction) Reset() {
	*x = SendMessageAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_action_action_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageAction) ProtoMessage() {}

func (x *SendMessageAction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_action_action_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageAction.ProtoReflect.Descriptor instead.
func (*SendMessageAction) Descriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{5}
}

func (x *SendMessageAction) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GoToSleepAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoToSleepAction) Reset() {
	*x = GoToSleepAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_action_action_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoToSleepAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoToSleepAction) ProtoMessage() {}

func (x *GoToSleepAction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_action_action_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoToSleepAction.ProtoReflect.Descriptor instead.
func (*GoToSleepAction) Descriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{6}
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ActionType `protobuf:"varint,1,opt,name=type,proto3,enum=action.ActionType" json:"type,omitempty"`
	// Types that are assignable to Actions:
	//
	//	*Action_InitAction
	//	*Action_VoteAction
	//	*Action_CheckAction
	//	*Action_PublishRoleAction
	//	*Action_KillAction
	//	*Action_SendMessageAction
	//	*Action_GoToSleepAction
	Actions isAction_Actions `protobuf_oneof:"actions"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_action_action_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_action_action_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_pkg_proto_action_action_proto_rawDescGZIP(), []int{7}
}

func (x *Action) GetType() ActionType {
	if x != nil {
		return x.Type
	}
	return ActionType_INIT
}

func (m *Action) GetActions() isAction_Actions {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (x *Action) GetInitAction() *InitAction {
	if x, ok := x.GetActions().(*Action_InitAction); ok {
		return x.InitAction
	}
	return nil
}

func (x *Action) GetVoteAction() *VoteAction {
	if x, ok := x.GetActions().(*Action_VoteAction); ok {
		return x.VoteAction
	}
	return nil
}

func (x *Action) GetCheckAction() *CheckRoleAction {
	if x, ok := x.GetActions().(*Action_CheckAction); ok {
		return x.CheckAction
	}
	return nil
}

func (x *Action) GetPublishRoleAction() *PublishRoleAction {
	if x, ok := x.GetActions().(*Action_PublishRoleAction); ok {
		return x.PublishRoleAction
	}
	return nil
}

func (x *Action) GetKillAction() *KillAction {
	if x, ok := x.GetActions().(*Action_KillAction); ok {
		return x.KillAction
	}
	return nil
}

func (x *Action) GetSendMessageAction() *SendMessageAction {
	if x, ok := x.GetActions().(*Action_SendMessageAction); ok {
		return x.SendMessageAction
	}
	return nil
}

func (x *Action) GetGoToSleepAction() *GoToSleepAction {
	if x, ok := x.GetActions().(*Action_GoToSleepAction); ok {
		return x.GoToSleepAction
	}
	return nil
}

type isAction_Actions interface {
	isAction_Actions()
}

type Action_InitAction struct {
	InitAction *InitAction `protobuf:"bytes,2,opt,name=init_action,json=initAction,proto3,oneof"`
}

type Action_VoteAction struct {
	VoteAction *VoteAction `protobuf:"bytes,3,opt,name=vote_action,json=voteAction,proto3,oneof"`
}

type Action_CheckAction struct {
	CheckAction *CheckRoleAction `protobuf:"bytes,4,opt,name=check_action,json=checkAction,proto3,oneof"`
}

type Action_PublishRoleAction struct {
	PublishRoleAction *PublishRoleAction `protobuf:"bytes,5,opt,name=publish_role_action,json=publishRoleAction,proto3,oneof"`
}

type Action_KillAction struct {
	KillAction *KillAction `protobuf:"bytes,6,opt,name=kill_action,json=killAction,proto3,oneof"`
}

type Action_SendMessageAction struct {
	SendMessageAction *SendMessageAction `protobuf:"bytes,7,opt,name=send_message_action,json=sendMessageAction,proto3,oneof"`
}

type Action_GoToSleepAction struct {
	GoToSleepAction *GoToSleepAction `protobuf:"bytes,8,opt,name=go_to_sleep_action,json=goToSleepAction,proto3,oneof"`
}

func (*Action_InitAction) isAction_Actions() {}

func (*Action_VoteAction) isAction_Actions() {}

func (*Action_CheckAction) isAction_Actions() {}

func (*Action_PublishRoleAction) isAction_Actions() {}

func (*Action_KillAction) isAction_Actions() {}

func (*Action_SendMessageAction) isAction_Actions() {}

func (*Action_GoToSleepAction) isAction_Actions() {}

var File_pkg_proto_action_action_proto protoreflect.FileDescriptor

var file_pkg_proto_action_action_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3d, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f,
	0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22,
	0x3d, 0x0a, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x42,
	0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x22, 0x44, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x6f, 0x6c,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x0a, 0x4b, 0x69, 0x6c, 0x6c,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x6f, 0x54, 0x6f, 0x53, 0x6c,
	0x65, 0x65, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x04, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x0b,
	0x69, 0x6e, 0x69, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a,
	0x76, 0x6f, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0c, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x6f, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x13, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x0a, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x13,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x12, 0x67, 0x6f, 0x5f,
	0x74, 0x6f, 0x5f, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x6f, 0x54, 0x6f, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x0f, 0x67, 0x6f, 0x54, 0x6f, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x09, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x7f, 0x0a, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e,
	0x49, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4f, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0x03,
	0x12, 0x08, 0x0a, 0x04, 0x4b, 0x49, 0x4c, 0x4c, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45,
	0x4e, 0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b,
	0x47, 0x4f, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x4c, 0x45, 0x45, 0x50, 0x10, 0x06, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x07, 0x42, 0x18, 0x5a,
	0x16, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_action_action_proto_rawDescOnce sync.Once
	file_pkg_proto_action_action_proto_rawDescData = file_pkg_proto_action_action_proto_rawDesc
)

func file_pkg_proto_action_action_proto_rawDescGZIP() []byte {
	file_pkg_proto_action_action_proto_rawDescOnce.Do(func() {
		file_pkg_proto_action_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_action_action_proto_rawDescData)
	})
	return file_pkg_proto_action_action_proto_rawDescData
}

var file_pkg_proto_action_action_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_action_action_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_proto_action_action_proto_goTypes = []interface{}{
	(ActionType)(0),                // 0: action.ActionType
	(*InitAction)(nil),             // 1: action.InitAction
	(*VoteAction)(nil),             // 2: action.VoteAction
	(*CheckRoleAction)(nil),        // 3: action.CheckRoleAction
	(*PublishRoleAction)(nil),      // 4: action.PublishRoleAction
	(*KillAction)(nil),             // 5: action.KillAction
	(*SendMessageAction)(nil),      // 6: action.SendMessageAction
	(*GoToSleepAction)(nil),        // 7: action.GoToSleepAction
	(*Action)(nil),                 // 8: action.Action
	(*player_info.PlayerInfo)(nil), // 9: player_info.PlayerInfo
}
var file_pkg_proto_action_action_proto_depIdxs = []int32{
	9,  // 0: action.InitAction.player:type_name -> player_info.PlayerInfo
	9,  // 1: action.VoteAction.player:type_name -> player_info.PlayerInfo
	9,  // 2: action.CheckRoleAction.player:type_name -> player_info.PlayerInfo
	9,  // 3: action.PublishRoleAction.player:type_name -> player_info.PlayerInfo
	9,  // 4: action.KillAction.player:type_name -> player_info.PlayerInfo
	0,  // 5: action.Action.type:type_name -> action.ActionType
	1,  // 6: action.Action.init_action:type_name -> action.InitAction
	2,  // 7: action.Action.vote_action:type_name -> action.VoteAction
	3,  // 8: action.Action.check_action:type_name -> action.CheckRoleAction
	4,  // 9: action.Action.publish_role_action:type_name -> action.PublishRoleAction
	5,  // 10: action.Action.kill_action:type_name -> action.KillAction
	6,  // 11: action.Action.send_message_action:type_name -> action.SendMessageAction
	7,  // 12: action.Action.go_to_sleep_action:type_name -> action.GoToSleepAction
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_pkg_proto_action_action_proto_init() }
func file_pkg_proto_action_action_proto_init() {
	if File_pkg_proto_action_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_action_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_action_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_action_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRoleAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_action_action_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRoleAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_action_action_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_action_action_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_action_action_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoToSleepAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_action_action_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_proto_action_action_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Action_InitAction)(nil),
		(*Action_VoteAction)(nil),
		(*Action_CheckAction)(nil),
		(*Action_PublishRoleAction)(nil),
		(*Action_KillAction)(nil),
		(*Action_SendMessageAction)(nil),
		(*Action_GoToSleepAction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_action_action_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_action_action_proto_goTypes,
		DependencyIndexes: file_pkg_proto_action_action_proto_depIdxs,
		EnumInfos:         file_pkg_proto_action_action_proto_enumTypes,
		MessageInfos:      file_pkg_proto_action_action_proto_msgTypes,
	}.Build()
	File_pkg_proto_action_action_proto = out.File
	file_pkg_proto_action_action_proto_rawDesc = nil
	file_pkg_proto_action_action_proto_goTypes = nil
	file_pkg_proto_action_action_proto_depIdxs = nil
}