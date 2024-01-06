// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc_master.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event int32

const (
	Event_EVENT_UNSPECIFIED     Event = 0
	Event_EVENT_REGISTER_CLIENT Event = 1
	Event_EVENT_REGISTER_SERVER Event = 2
	Event_EVENT_ERROR           Event = 3
	Event_EVENT_DATA            Event = 4
	Event_EVENT_UPDATE_FRPC     Event = 5
	Event_EVENT_REMOVE_FRPC     Event = 6
	Event_EVENT_UPDATE_FRPS     Event = 7
	Event_EVENT_REMOVE_FRPS     Event = 8
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "EVENT_REGISTER_CLIENT",
		2: "EVENT_REGISTER_SERVER",
		3: "EVENT_ERROR",
		4: "EVENT_DATA",
		5: "EVENT_UPDATE_FRPC",
		6: "EVENT_REMOVE_FRPC",
		7: "EVENT_UPDATE_FRPS",
		8: "EVENT_REMOVE_FRPS",
	}
	Event_value = map[string]int32{
		"EVENT_UNSPECIFIED":     0,
		"EVENT_REGISTER_CLIENT": 1,
		"EVENT_REGISTER_SERVER": 2,
		"EVENT_ERROR":           3,
		"EVENT_DATA":            4,
		"EVENT_UPDATE_FRPC":     5,
		"EVENT_REMOVE_FRPC":     6,
		"EVENT_UPDATE_FRPS":     7,
		"EVENT_REMOVE_FRPS":     8,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_master_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_rpc_master_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{0}
}

type ServerBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId     string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ServerSecret string `protobuf:"bytes,2,opt,name=server_secret,json=serverSecret,proto3" json:"server_secret,omitempty"`
}

func (x *ServerBase) Reset() {
	*x = ServerBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerBase) ProtoMessage() {}

func (x *ServerBase) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerBase.ProtoReflect.Descriptor instead.
func (*ServerBase) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{0}
}

func (x *ServerBase) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *ServerBase) GetServerSecret() string {
	if x != nil {
		return x.ServerSecret
	}
	return ""
}

type ClientBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *ClientBase) Reset() {
	*x = ClientBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientBase) ProtoMessage() {}

func (x *ClientBase) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientBase.ProtoReflect.Descriptor instead.
func (*ClientBase) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{1}
}

func (x *ClientBase) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientBase) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type ServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event     Event  `protobuf:"varint,1,opt,name=event,proto3,enum=master.Event" json:"event,omitempty"`
	ClientId  string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SessionId string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ServerMessage) Reset() {
	*x = ServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMessage) ProtoMessage() {}

func (x *ServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMessage.ProtoReflect.Descriptor instead.
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{2}
}

func (x *ServerMessage) GetEvent() Event {
	if x != nil {
		return x.Event
	}
	return Event_EVENT_UNSPECIFIED
}

func (x *ServerMessage) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ServerMessage) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ServerMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event     Event  `protobuf:"varint,1,opt,name=event,proto3,enum=master.Event" json:"event,omitempty"`
	ClientId  string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SessionId string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Secret    string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	Data      []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ClientMessage) Reset() {
	*x = ClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessage) ProtoMessage() {}

func (x *ClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessage.ProtoReflect.Descriptor instead.
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{3}
}

func (x *ClientMessage) GetEvent() Event {
	if x != nil {
		return x.Event
	}
	return Event_EVENT_UNSPECIFIED
}

func (x *ClientMessage) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientMessage) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ClientMessage) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *ClientMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PullClientConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *ClientBase `protobuf:"bytes,255,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *PullClientConfigReq) Reset() {
	*x = PullClientConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullClientConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullClientConfigReq) ProtoMessage() {}

func (x *PullClientConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullClientConfigReq.ProtoReflect.Descriptor instead.
func (*PullClientConfigReq) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{4}
}

func (x *PullClientConfigReq) GetBase() *ClientBase {
	if x != nil {
		return x.Base
	}
	return nil
}

type PullClientConfigResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Client *Client `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *PullClientConfigResp) Reset() {
	*x = PullClientConfigResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullClientConfigResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullClientConfigResp) ProtoMessage() {}

func (x *PullClientConfigResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullClientConfigResp.ProtoReflect.Descriptor instead.
func (*PullClientConfigResp) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{5}
}

func (x *PullClientConfigResp) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *PullClientConfigResp) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type PullServerConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *ServerBase `protobuf:"bytes,255,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *PullServerConfigReq) Reset() {
	*x = PullServerConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullServerConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullServerConfigReq) ProtoMessage() {}

func (x *PullServerConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullServerConfigReq.ProtoReflect.Descriptor instead.
func (*PullServerConfigReq) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{6}
}

func (x *PullServerConfigReq) GetBase() *ServerBase {
	if x != nil {
		return x.Base
	}
	return nil
}

type PullServerConfigResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Server *Server `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *PullServerConfigResp) Reset() {
	*x = PullServerConfigResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullServerConfigResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullServerConfigResp) ProtoMessage() {}

func (x *PullServerConfigResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullServerConfigResp.ProtoReflect.Descriptor instead.
func (*PullServerConfigResp) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{7}
}

func (x *PullServerConfigResp) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *PullServerConfigResp) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type FRPAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  string      `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token string      `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Base  *ServerBase `protobuf:"bytes,255,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *FRPAuthRequest) Reset() {
	*x = FRPAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FRPAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FRPAuthRequest) ProtoMessage() {}

func (x *FRPAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FRPAuthRequest.ProtoReflect.Descriptor instead.
func (*FRPAuthRequest) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{8}
}

func (x *FRPAuthRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *FRPAuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FRPAuthRequest) GetBase() *ServerBase {
	if x != nil {
		return x.Base
	}
	return nil
}

type FRPAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Ok     bool    `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *FRPAuthResponse) Reset() {
	*x = FRPAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FRPAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FRPAuthResponse) ProtoMessage() {}

func (x *FRPAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FRPAuthResponse.ProtoReflect.Descriptor instead.
func (*FRPAuthResponse) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{9}
}

func (x *FRPAuthResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *FRPAuthResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_rpc_master_proto protoreflect.FileDescriptor

var file_rpc_master_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x4e, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x9c, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e,
	0x0a, 0x13, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0xff, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x66,
	0x0a, 0x14, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26,
	0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x13, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x66, 0x0a, 0x14, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x63,
	0x0a, 0x0e, 0x46, 0x52, 0x50, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x22, 0x49, 0x0a, 0x0f, 0x46, 0x52, 0x50, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x2a, 0xd1,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x52, 0x50, 0x43, 0x10, 0x05, 0x12, 0x15, 0x0a,
	0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x46, 0x52,
	0x50, 0x43, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x52, 0x50, 0x53, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x46, 0x52, 0x50, 0x53,
	0x10, 0x08, 0x32, 0xa3, 0x02, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x0a,
	0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4d, 0x0a,
	0x10, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1c,
	0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4d, 0x0a, 0x10,
	0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x08, 0x46,
	0x52, 0x50, 0x43, 0x41, 0x75, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x46, 0x52, 0x50, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x52, 0x50, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_master_proto_rawDescOnce sync.Once
	file_rpc_master_proto_rawDescData = file_rpc_master_proto_rawDesc
)

func file_rpc_master_proto_rawDescGZIP() []byte {
	file_rpc_master_proto_rawDescOnce.Do(func() {
		file_rpc_master_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_master_proto_rawDescData)
	})
	return file_rpc_master_proto_rawDescData
}

var file_rpc_master_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_master_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rpc_master_proto_goTypes = []interface{}{
	(Event)(0),                   // 0: master.Event
	(*ServerBase)(nil),           // 1: master.ServerBase
	(*ClientBase)(nil),           // 2: master.ClientBase
	(*ServerMessage)(nil),        // 3: master.ServerMessage
	(*ClientMessage)(nil),        // 4: master.ClientMessage
	(*PullClientConfigReq)(nil),  // 5: master.PullClientConfigReq
	(*PullClientConfigResp)(nil), // 6: master.PullClientConfigResp
	(*PullServerConfigReq)(nil),  // 7: master.PullServerConfigReq
	(*PullServerConfigResp)(nil), // 8: master.PullServerConfigResp
	(*FRPAuthRequest)(nil),       // 9: master.FRPAuthRequest
	(*FRPAuthResponse)(nil),      // 10: master.FRPAuthResponse
	(*Status)(nil),               // 11: common.Status
	(*Client)(nil),               // 12: common.Client
	(*Server)(nil),               // 13: common.Server
}
var file_rpc_master_proto_depIdxs = []int32{
	0,  // 0: master.ServerMessage.event:type_name -> master.Event
	0,  // 1: master.ClientMessage.event:type_name -> master.Event
	2,  // 2: master.PullClientConfigReq.base:type_name -> master.ClientBase
	11, // 3: master.PullClientConfigResp.status:type_name -> common.Status
	12, // 4: master.PullClientConfigResp.client:type_name -> common.Client
	1,  // 5: master.PullServerConfigReq.base:type_name -> master.ServerBase
	11, // 6: master.PullServerConfigResp.status:type_name -> common.Status
	13, // 7: master.PullServerConfigResp.server:type_name -> common.Server
	1,  // 8: master.FRPAuthRequest.base:type_name -> master.ServerBase
	11, // 9: master.FRPAuthResponse.status:type_name -> common.Status
	4,  // 10: master.Master.ServerSend:input_type -> master.ClientMessage
	5,  // 11: master.Master.PullClientConfig:input_type -> master.PullClientConfigReq
	7,  // 12: master.Master.PullServerConfig:input_type -> master.PullServerConfigReq
	9,  // 13: master.Master.FRPCAuth:input_type -> master.FRPAuthRequest
	3,  // 14: master.Master.ServerSend:output_type -> master.ServerMessage
	6,  // 15: master.Master.PullClientConfig:output_type -> master.PullClientConfigResp
	8,  // 16: master.Master.PullServerConfig:output_type -> master.PullServerConfigResp
	10, // 17: master.Master.FRPCAuth:output_type -> master.FRPAuthResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_rpc_master_proto_init() }
func file_rpc_master_proto_init() {
	if File_rpc_master_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_master_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerBase); i {
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
		file_rpc_master_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientBase); i {
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
		file_rpc_master_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMessage); i {
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
		file_rpc_master_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessage); i {
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
		file_rpc_master_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullClientConfigReq); i {
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
		file_rpc_master_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullClientConfigResp); i {
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
		file_rpc_master_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullServerConfigReq); i {
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
		file_rpc_master_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullServerConfigResp); i {
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
		file_rpc_master_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FRPAuthRequest); i {
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
		file_rpc_master_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FRPAuthResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_master_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_master_proto_goTypes,
		DependencyIndexes: file_rpc_master_proto_depIdxs,
		EnumInfos:         file_rpc_master_proto_enumTypes,
		MessageInfos:      file_rpc_master_proto_msgTypes,
	}.Build()
	File_rpc_master_proto = out.File
	file_rpc_master_proto_rawDesc = nil
	file_rpc_master_proto_goTypes = nil
	file_rpc_master_proto_depIdxs = nil
}
