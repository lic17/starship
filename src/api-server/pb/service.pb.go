// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: src/api-server/pb/service.proto

package servicepb

import (
	context "context"
	module "github.com/tricorder/src/pb/module"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DeploymentState int32

const (
	DeploymentState_CREATED                  DeploymentState = 0
	DeploymentState_TO_BE_DEPLOYED           DeploymentState = 1
	DeploymentState_DEPLOYMENT_IN_PROGRESS   DeploymentState = 2
	DeploymentState_DEPLOYMENT_FAILED        DeploymentState = 3
	DeploymentState_DEPLOYMENT_SUCCEEDED     DeploymentState = 4
	DeploymentState_TO_BE_UNDEPLOYED         DeploymentState = 5
	DeploymentState_UNDEPLOYMENT_IN_PROGRESS DeploymentState = 6
	DeploymentState_UNDEPLOYMENT_FAILED      DeploymentState = 7
	DeploymentState_UNDEPLOYMENT_SUCCEEDED   DeploymentState = 8
)

// Enum value maps for DeploymentState.
var (
	DeploymentState_name = map[int32]string{
		0: "CREATED",
		1: "TO_BE_DEPLOYED",
		2: "DEPLOYMENT_IN_PROGRESS",
		3: "DEPLOYMENT_FAILED",
		4: "DEPLOYMENT_SUCCEEDED",
		5: "TO_BE_UNDEPLOYED",
		6: "UNDEPLOYMENT_IN_PROGRESS",
		7: "UNDEPLOYMENT_FAILED",
		8: "UNDEPLOYMENT_SUCCEEDED",
	}
	DeploymentState_value = map[string]int32{
		"CREATED":                  0,
		"TO_BE_DEPLOYED":           1,
		"DEPLOYMENT_IN_PROGRESS":   2,
		"DEPLOYMENT_FAILED":        3,
		"DEPLOYMENT_SUCCEEDED":     4,
		"TO_BE_UNDEPLOYED":         5,
		"UNDEPLOYMENT_IN_PROGRESS": 6,
		"UNDEPLOYMENT_FAILED":      7,
		"UNDEPLOYMENT_SUCCEEDED":   8,
	}
)

func (x DeploymentState) Enum() *DeploymentState {
	p := new(DeploymentState)
	*p = x
	return p
}

func (x DeploymentState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentState) Descriptor() protoreflect.EnumDescriptor {
	return file_src_api_server_pb_service_proto_enumTypes[0].Descriptor()
}

func (DeploymentState) Type() protoreflect.EnumType {
	return &file_src_api_server_pb_service_proto_enumTypes[0]
}

func (x DeploymentState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentState.Descriptor instead.
func (DeploymentState) EnumDescriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{0}
}

type DeployModuleReq_DEPLOY_STATUS int32

const (
	DeployModuleReq_UNDEPLOY DeployModuleReq_DEPLOY_STATUS = 0
	DeployModuleReq_DEPLOY   DeployModuleReq_DEPLOY_STATUS = 1
)

// Enum value maps for DeployModuleReq_DEPLOY_STATUS.
var (
	DeployModuleReq_DEPLOY_STATUS_name = map[int32]string{
		0: "UNDEPLOY",
		1: "DEPLOY",
	}
	DeployModuleReq_DEPLOY_STATUS_value = map[string]int32{
		"UNDEPLOY": 0,
		"DEPLOY":   1,
	}
)

func (x DeployModuleReq_DEPLOY_STATUS) Enum() *DeployModuleReq_DEPLOY_STATUS {
	p := new(DeployModuleReq_DEPLOY_STATUS)
	*p = x
	return p
}

func (x DeployModuleReq_DEPLOY_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeployModuleReq_DEPLOY_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_src_api_server_pb_service_proto_enumTypes[1].Descriptor()
}

func (DeployModuleReq_DEPLOY_STATUS) Type() protoreflect.EnumType {
	return &file_src_api_server_pb_service_proto_enumTypes[1]
}

func (x DeployModuleReq_DEPLOY_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeployModuleReq_DEPLOY_STATUS.Descriptor instead.
func (DeployModuleReq_DEPLOY_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{0, 0}
}

type DeployModuleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleId        string                        `protobuf:"bytes,1,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	LifeTimeSeconds int32                         `protobuf:"varint,2,opt,name=life_time_seconds,json=lifeTimeSeconds,proto3" json:"life_time_seconds,omitempty"`
	Module          *module.Module                `protobuf:"bytes,7,opt,name=module,proto3" json:"module,omitempty"`
	Deploy          DeployModuleReq_DEPLOY_STATUS `protobuf:"varint,6,opt,name=deploy,proto3,enum=tricorder.deployer.servicepb.DeployModuleReq_DEPLOY_STATUS" json:"deploy,omitempty"`
}

func (x *DeployModuleReq) Reset() {
	*x = DeployModuleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_server_pb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployModuleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployModuleReq) ProtoMessage() {}

func (x *DeployModuleReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_server_pb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployModuleReq.ProtoReflect.Descriptor instead.
func (*DeployModuleReq) Descriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{0}
}

func (x *DeployModuleReq) GetModuleId() string {
	if x != nil {
		return x.ModuleId
	}
	return ""
}

func (x *DeployModuleReq) GetLifeTimeSeconds() int32 {
	if x != nil {
		return x.LifeTimeSeconds
	}
	return 0
}

func (x *DeployModuleReq) GetModule() *module.Module {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *DeployModuleReq) GetDeploy() DeployModuleReq_DEPLOY_STATUS {
	if x != nil {
		return x.Deploy
	}
	return DeployModuleReq_UNDEPLOY
}

type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PodId    string `protobuf:"bytes,2,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_server_pb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_server_pb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{1}
}

func (x *Agent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Agent) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

func (x *Agent) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type DeployModuleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleId string          `protobuf:"bytes,1,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	Agent    *Agent          `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	State    DeploymentState `protobuf:"varint,3,opt,name=state,proto3,enum=tricorder.deployer.servicepb.DeploymentState" json:"state,omitempty"`
	Desc     string          `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *DeployModuleResp) Reset() {
	*x = DeployModuleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_server_pb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployModuleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployModuleResp) ProtoMessage() {}

func (x *DeployModuleResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_server_pb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployModuleResp.ProtoReflect.Descriptor instead.
func (*DeployModuleResp) Descriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeployModuleResp) GetModuleId() string {
	if x != nil {
		return x.ModuleId
	}
	return ""
}

func (x *DeployModuleResp) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

func (x *DeployModuleResp) GetState() DeploymentState {
	if x != nil {
		return x.State
	}
	return DeploymentState_CREATED
}

func (x *DeployModuleResp) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type ProcessWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*ProcessWrapper_NodeName
	//	*ProcessWrapper_Process
	Msg isProcessWrapper_Msg `protobuf_oneof:"msg"`
}

func (x *ProcessWrapper) Reset() {
	*x = ProcessWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_server_pb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessWrapper) ProtoMessage() {}

func (x *ProcessWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_server_pb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessWrapper.ProtoReflect.Descriptor instead.
func (*ProcessWrapper) Descriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{3}
}

func (m *ProcessWrapper) GetMsg() isProcessWrapper_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ProcessWrapper) GetNodeName() string {
	if x, ok := x.GetMsg().(*ProcessWrapper_NodeName); ok {
		return x.NodeName
	}
	return ""
}

func (x *ProcessWrapper) GetProcess() *ProcessInfo {
	if x, ok := x.GetMsg().(*ProcessWrapper_Process); ok {
		return x.Process
	}
	return nil
}

type isProcessWrapper_Msg interface {
	isProcessWrapper_Msg()
}

type ProcessWrapper_NodeName struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3,oneof"`
}

type ProcessWrapper_Process struct {
	Process *ProcessInfo `protobuf:"bytes,2,opt,name=process,proto3,oneof"`
}

func (*ProcessWrapper_NodeName) isProcessWrapper_Msg() {}

func (*ProcessWrapper_Process) isProcessWrapper_Msg() {}

type ProcessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcList  []*Process     `protobuf:"bytes,1,rep,name=proc_list,json=procList,proto3" json:"proc_list,omitempty"`
	Container *ContainerInfo `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
}

func (x *ProcessInfo) Reset() {
	*x = ProcessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_server_pb_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessInfo) ProtoMessage() {}

func (x *ProcessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_server_pb_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessInfo.ProtoReflect.Descriptor instead.
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{4}
}

func (x *ProcessInfo) GetProcList() []*Process {
	if x != nil {
		return x.ProcList
	}
	return nil
}

func (x *ProcessInfo) GetContainer() *ContainerInfo {
	if x != nil {
		return x.Container
	}
	return nil
}

type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime int64 `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_server_pb_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_server_pb_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{5}
}

func (x *Process) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Process) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type ContainerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PodUid   string `protobuf:"bytes,3,opt,name=pod_uid,json=podUid,proto3" json:"pod_uid,omitempty"`
	PodName  string `protobuf:"bytes,4,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	QosClass string `protobuf:"bytes,5,opt,name=qos_class,json=qosClass,proto3" json:"qos_class,omitempty"`
}

func (x *ContainerInfo) Reset() {
	*x = ContainerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_server_pb_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerInfo) ProtoMessage() {}

func (x *ContainerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_server_pb_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerInfo.ProtoReflect.Descriptor instead.
func (*ContainerInfo) Descriptor() ([]byte, []int) {
	return file_src_api_server_pb_service_proto_rawDescGZIP(), []int{6}
}

func (x *ContainerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContainerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerInfo) GetPodUid() string {
	if x != nil {
		return x.PodUid
	}
	return ""
}

func (x *ContainerInfo) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *ContainerInfo) GetQosClass() string {
	if x != nil {
		return x.QosClass
	}
	return ""
}

var File_src_api_server_pb_service_proto protoreflect.FileDescriptor

var file_src_api_server_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x1a,
	0x1a, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x0f,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11,
	0x6c, 0x69, 0x66, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6c, 0x69, 0x66, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x53, 0x0a,
	0x06, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e,
	0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x44, 0x45, 0x50,
	0x4c, 0x4f, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x22, 0x29, 0x0a, 0x0d, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x10, 0x01, 0x22, 0x4b, 0x0a,
	0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x10, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x05,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x72,
	0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x22, 0x7d, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x57, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x45, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x9c, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x42, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x3a,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6f, 0x64, 0x55, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x71, 0x6f, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x6f, 0x73, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x2a, 0xe8, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x4f, 0x5f, 0x42, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c,
	0x4f, 0x59, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x50,
	0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x4f, 0x5f, 0x42, 0x45, 0x5f, 0x55, 0x4e, 0x44,
	0x45, 0x50, 0x4c, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x4e, 0x44,
	0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f,
	0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x44, 0x45, 0x50,
	0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07,
	0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x08, 0x32, 0x85, 0x01, 0x0a,
	0x0e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x12,
	0x73, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x2e, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x1a,
	0x2d, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x32, 0x84, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x70, 0x0a, 0x0d, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x2e, 0x74, 0x72, 0x69,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x1a, 0x2b, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_api_server_pb_service_proto_rawDescOnce sync.Once
	file_src_api_server_pb_service_proto_rawDescData = file_src_api_server_pb_service_proto_rawDesc
)

func file_src_api_server_pb_service_proto_rawDescGZIP() []byte {
	file_src_api_server_pb_service_proto_rawDescOnce.Do(func() {
		file_src_api_server_pb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_api_server_pb_service_proto_rawDescData)
	})
	return file_src_api_server_pb_service_proto_rawDescData
}

var file_src_api_server_pb_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_src_api_server_pb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_src_api_server_pb_service_proto_goTypes = []interface{}{
	(DeploymentState)(0),               // 0: tricorder.deployer.servicepb.DeploymentState
	(DeployModuleReq_DEPLOY_STATUS)(0), // 1: tricorder.deployer.servicepb.DeployModuleReq.DEPLOY_STATUS
	(*DeployModuleReq)(nil),            // 2: tricorder.deployer.servicepb.DeployModuleReq
	(*Agent)(nil),                      // 3: tricorder.deployer.servicepb.Agent
	(*DeployModuleResp)(nil),           // 4: tricorder.deployer.servicepb.DeployModuleResp
	(*ProcessWrapper)(nil),             // 5: tricorder.deployer.servicepb.ProcessWrapper
	(*ProcessInfo)(nil),                // 6: tricorder.deployer.servicepb.ProcessInfo
	(*Process)(nil),                    // 7: tricorder.deployer.servicepb.Process
	(*ContainerInfo)(nil),              // 8: tricorder.deployer.servicepb.ContainerInfo
	(*module.Module)(nil),              // 9: tricorder.pb.module.Module
}
var file_src_api_server_pb_service_proto_depIdxs = []int32{
	9, // 0: tricorder.deployer.servicepb.DeployModuleReq.module:type_name -> tricorder.pb.module.Module
	1, // 1: tricorder.deployer.servicepb.DeployModuleReq.deploy:type_name -> tricorder.deployer.servicepb.DeployModuleReq.DEPLOY_STATUS
	3, // 2: tricorder.deployer.servicepb.DeployModuleResp.agent:type_name -> tricorder.deployer.servicepb.Agent
	0, // 3: tricorder.deployer.servicepb.DeployModuleResp.state:type_name -> tricorder.deployer.servicepb.DeploymentState
	6, // 4: tricorder.deployer.servicepb.ProcessWrapper.process:type_name -> tricorder.deployer.servicepb.ProcessInfo
	7, // 5: tricorder.deployer.servicepb.ProcessInfo.proc_list:type_name -> tricorder.deployer.servicepb.Process
	8, // 6: tricorder.deployer.servicepb.ProcessInfo.container:type_name -> tricorder.deployer.servicepb.ContainerInfo
	4, // 7: tricorder.deployer.servicepb.ModuleDeployer.DeployModule:input_type -> tricorder.deployer.servicepb.DeployModuleResp
	5, // 8: tricorder.deployer.servicepb.ProcessCollector.ReportProcess:input_type -> tricorder.deployer.servicepb.ProcessWrapper
	2, // 9: tricorder.deployer.servicepb.ModuleDeployer.DeployModule:output_type -> tricorder.deployer.servicepb.DeployModuleReq
	8, // 10: tricorder.deployer.servicepb.ProcessCollector.ReportProcess:output_type -> tricorder.deployer.servicepb.ContainerInfo
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_src_api_server_pb_service_proto_init() }
func file_src_api_server_pb_service_proto_init() {
	if File_src_api_server_pb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_api_server_pb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployModuleReq); i {
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
		file_src_api_server_pb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
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
		file_src_api_server_pb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployModuleResp); i {
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
		file_src_api_server_pb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessWrapper); i {
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
		file_src_api_server_pb_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessInfo); i {
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
		file_src_api_server_pb_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
		file_src_api_server_pb_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerInfo); i {
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
	file_src_api_server_pb_service_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ProcessWrapper_NodeName)(nil),
		(*ProcessWrapper_Process)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_api_server_pb_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_src_api_server_pb_service_proto_goTypes,
		DependencyIndexes: file_src_api_server_pb_service_proto_depIdxs,
		EnumInfos:         file_src_api_server_pb_service_proto_enumTypes,
		MessageInfos:      file_src_api_server_pb_service_proto_msgTypes,
	}.Build()
	File_src_api_server_pb_service_proto = out.File
	file_src_api_server_pb_service_proto_rawDesc = nil
	file_src_api_server_pb_service_proto_goTypes = nil
	file_src_api_server_pb_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ModuleDeployerClient is the client API for ModuleDeployer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModuleDeployerClient interface {
	DeployModule(ctx context.Context, opts ...grpc.CallOption) (ModuleDeployer_DeployModuleClient, error)
}

type moduleDeployerClient struct {
	cc grpc.ClientConnInterface
}

func NewModuleDeployerClient(cc grpc.ClientConnInterface) ModuleDeployerClient {
	return &moduleDeployerClient{cc}
}

func (c *moduleDeployerClient) DeployModule(ctx context.Context, opts ...grpc.CallOption) (ModuleDeployer_DeployModuleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ModuleDeployer_serviceDesc.Streams[0], "/tricorder.deployer.servicepb.ModuleDeployer/DeployModule", opts...)
	if err != nil {
		return nil, err
	}
	x := &moduleDeployerDeployModuleClient{stream}
	return x, nil
}

type ModuleDeployer_DeployModuleClient interface {
	Send(*DeployModuleResp) error
	Recv() (*DeployModuleReq, error)
	grpc.ClientStream
}

type moduleDeployerDeployModuleClient struct {
	grpc.ClientStream
}

func (x *moduleDeployerDeployModuleClient) Send(m *DeployModuleResp) error {
	return x.ClientStream.SendMsg(m)
}

func (x *moduleDeployerDeployModuleClient) Recv() (*DeployModuleReq, error) {
	m := new(DeployModuleReq)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ModuleDeployerServer is the server API for ModuleDeployer service.
type ModuleDeployerServer interface {
	DeployModule(ModuleDeployer_DeployModuleServer) error
}

// UnimplementedModuleDeployerServer can be embedded to have forward compatible implementations.
type UnimplementedModuleDeployerServer struct {
}

func (*UnimplementedModuleDeployerServer) DeployModule(ModuleDeployer_DeployModuleServer) error {
	return status.Errorf(codes.Unimplemented, "method DeployModule not implemented")
}

func RegisterModuleDeployerServer(s *grpc.Server, srv ModuleDeployerServer) {
	s.RegisterService(&_ModuleDeployer_serviceDesc, srv)
}

func _ModuleDeployer_DeployModule_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModuleDeployerServer).DeployModule(&moduleDeployerDeployModuleServer{stream})
}

type ModuleDeployer_DeployModuleServer interface {
	Send(*DeployModuleReq) error
	Recv() (*DeployModuleResp, error)
	grpc.ServerStream
}

type moduleDeployerDeployModuleServer struct {
	grpc.ServerStream
}

func (x *moduleDeployerDeployModuleServer) Send(m *DeployModuleReq) error {
	return x.ServerStream.SendMsg(m)
}

func (x *moduleDeployerDeployModuleServer) Recv() (*DeployModuleResp, error) {
	m := new(DeployModuleResp)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ModuleDeployer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tricorder.deployer.servicepb.ModuleDeployer",
	HandlerType: (*ModuleDeployerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeployModule",
			Handler:       _ModuleDeployer_DeployModule_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/api-server/pb/service.proto",
}

// ProcessCollectorClient is the client API for ProcessCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessCollectorClient interface {
	ReportProcess(ctx context.Context, opts ...grpc.CallOption) (ProcessCollector_ReportProcessClient, error)
}

type processCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessCollectorClient(cc grpc.ClientConnInterface) ProcessCollectorClient {
	return &processCollectorClient{cc}
}

func (c *processCollectorClient) ReportProcess(ctx context.Context, opts ...grpc.CallOption) (ProcessCollector_ReportProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessCollector_serviceDesc.Streams[0], "/tricorder.deployer.servicepb.ProcessCollector/ReportProcess", opts...)
	if err != nil {
		return nil, err
	}
	x := &processCollectorReportProcessClient{stream}
	return x, nil
}

type ProcessCollector_ReportProcessClient interface {
	Send(*ProcessWrapper) error
	Recv() (*ContainerInfo, error)
	grpc.ClientStream
}

type processCollectorReportProcessClient struct {
	grpc.ClientStream
}

func (x *processCollectorReportProcessClient) Send(m *ProcessWrapper) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processCollectorReportProcessClient) Recv() (*ContainerInfo, error) {
	m := new(ContainerInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessCollectorServer is the server API for ProcessCollector service.
type ProcessCollectorServer interface {
	ReportProcess(ProcessCollector_ReportProcessServer) error
}

// UnimplementedProcessCollectorServer can be embedded to have forward compatible implementations.
type UnimplementedProcessCollectorServer struct {
}

func (*UnimplementedProcessCollectorServer) ReportProcess(ProcessCollector_ReportProcessServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportProcess not implemented")
}

func RegisterProcessCollectorServer(s *grpc.Server, srv ProcessCollectorServer) {
	s.RegisterService(&_ProcessCollector_serviceDesc, srv)
}

func _ProcessCollector_ReportProcess_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessCollectorServer).ReportProcess(&processCollectorReportProcessServer{stream})
}

type ProcessCollector_ReportProcessServer interface {
	Send(*ContainerInfo) error
	Recv() (*ProcessWrapper, error)
	grpc.ServerStream
}

type processCollectorReportProcessServer struct {
	grpc.ServerStream
}

func (x *processCollectorReportProcessServer) Send(m *ContainerInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processCollectorReportProcessServer) Recv() (*ProcessWrapper, error) {
	m := new(ProcessWrapper)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ProcessCollector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tricorder.deployer.servicepb.ProcessCollector",
	HandlerType: (*ProcessCollectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReportProcess",
			Handler:       _ProcessCollector_ReportProcess_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/api-server/pb/service.proto",
}
