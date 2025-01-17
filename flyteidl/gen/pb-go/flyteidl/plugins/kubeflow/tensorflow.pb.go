// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/kubeflow/tensorflow.proto

package plugins

import (
	fmt "fmt"
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
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

// Proto for plugin that enables distributed training using https://github.com/kubeflow/tf-operator
type DistributedTensorflowTrainingTask struct {
	// Worker replicas spec
	WorkerReplicas *DistributedTensorflowTrainingReplicaSpec `protobuf:"bytes,1,opt,name=worker_replicas,json=workerReplicas,proto3" json:"worker_replicas,omitempty"`
	// Parameter server replicas spec
	PsReplicas *DistributedTensorflowTrainingReplicaSpec `protobuf:"bytes,2,opt,name=ps_replicas,json=psReplicas,proto3" json:"ps_replicas,omitempty"`
	// Chief replicas spec
	ChiefReplicas *DistributedTensorflowTrainingReplicaSpec `protobuf:"bytes,3,opt,name=chief_replicas,json=chiefReplicas,proto3" json:"chief_replicas,omitempty"`
	// RunPolicy encapsulates various runtime policies of the distributed training
	// job, for example how to clean up resources and how long the job can stay
	// active.
	RunPolicy            *RunPolicy `protobuf:"bytes,4,opt,name=run_policy,json=runPolicy,proto3" json:"run_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DistributedTensorflowTrainingTask) Reset()         { *m = DistributedTensorflowTrainingTask{} }
func (m *DistributedTensorflowTrainingTask) String() string { return proto.CompactTextString(m) }
func (*DistributedTensorflowTrainingTask) ProtoMessage()    {}
func (*DistributedTensorflowTrainingTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_93de2bd764ddf01a, []int{0}
}

func (m *DistributedTensorflowTrainingTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Unmarshal(m, b)
}
func (m *DistributedTensorflowTrainingTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Marshal(b, m, deterministic)
}
func (m *DistributedTensorflowTrainingTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedTensorflowTrainingTask.Merge(m, src)
}
func (m *DistributedTensorflowTrainingTask) XXX_Size() int {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Size(m)
}
func (m *DistributedTensorflowTrainingTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedTensorflowTrainingTask.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedTensorflowTrainingTask proto.InternalMessageInfo

func (m *DistributedTensorflowTrainingTask) GetWorkerReplicas() *DistributedTensorflowTrainingReplicaSpec {
	if m != nil {
		return m.WorkerReplicas
	}
	return nil
}

func (m *DistributedTensorflowTrainingTask) GetPsReplicas() *DistributedTensorflowTrainingReplicaSpec {
	if m != nil {
		return m.PsReplicas
	}
	return nil
}

func (m *DistributedTensorflowTrainingTask) GetChiefReplicas() *DistributedTensorflowTrainingReplicaSpec {
	if m != nil {
		return m.ChiefReplicas
	}
	return nil
}

func (m *DistributedTensorflowTrainingTask) GetRunPolicy() *RunPolicy {
	if m != nil {
		return m.RunPolicy
	}
	return nil
}

type DistributedTensorflowTrainingReplicaSpec struct {
	// Number of replicas
	Replicas int32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// Image used for the replica group
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// Resources required for the replica group
	Resources *core.Resources `protobuf:"bytes,3,opt,name=resources,proto3" json:"resources,omitempty"`
	// RestartPolicy Determines whether pods will be restarted when they exit
	RestartPolicy        RestartPolicy `protobuf:"varint,4,opt,name=restart_policy,json=restartPolicy,proto3,enum=flyteidl.plugins.kubeflow.RestartPolicy" json:"restart_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DistributedTensorflowTrainingReplicaSpec) Reset() {
	*m = DistributedTensorflowTrainingReplicaSpec{}
}
func (m *DistributedTensorflowTrainingReplicaSpec) String() string { return proto.CompactTextString(m) }
func (*DistributedTensorflowTrainingReplicaSpec) ProtoMessage()    {}
func (*DistributedTensorflowTrainingReplicaSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_93de2bd764ddf01a, []int{1}
}

func (m *DistributedTensorflowTrainingReplicaSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.Unmarshal(m, b)
}
func (m *DistributedTensorflowTrainingReplicaSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.Marshal(b, m, deterministic)
}
func (m *DistributedTensorflowTrainingReplicaSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.Merge(m, src)
}
func (m *DistributedTensorflowTrainingReplicaSpec) XXX_Size() int {
	return xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.Size(m)
}
func (m *DistributedTensorflowTrainingReplicaSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec proto.InternalMessageInfo

func (m *DistributedTensorflowTrainingReplicaSpec) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *DistributedTensorflowTrainingReplicaSpec) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DistributedTensorflowTrainingReplicaSpec) GetResources() *core.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DistributedTensorflowTrainingReplicaSpec) GetRestartPolicy() RestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return RestartPolicy_RESTART_POLICY_NEVER
}

func init() {
	proto.RegisterType((*DistributedTensorflowTrainingTask)(nil), "flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask")
	proto.RegisterType((*DistributedTensorflowTrainingReplicaSpec)(nil), "flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/kubeflow/tensorflow.proto", fileDescriptor_93de2bd764ddf01a)
}

var fileDescriptor_93de2bd764ddf01a = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x99, 0x73, 0xe2, 0x32, 0x56, 0xa1, 0x78, 0x98, 0x3b, 0xe9, 0x10, 0x19, 0x82, 0x0d,
	0x4c, 0x50, 0xbc, 0x3a, 0xef, 0x4a, 0xdc, 0xc9, 0xcb, 0x68, 0xb3, 0xb7, 0x2e, 0xb6, 0x4d, 0xc2,
	0x4b, 0xc2, 0xd8, 0x37, 0xf2, 0x8b, 0xf9, 0x3d, 0x64, 0xed, 0xda, 0x4e, 0x61, 0xc3, 0xc3, 0x6e,
	0x79, 0xcd, 0x3f, 0xff, 0xdf, 0x7b, 0xaf, 0x7f, 0x72, 0x3b, 0x4f, 0x57, 0x16, 0xc4, 0x2c, 0xa5,
	0x3a, 0x75, 0xb1, 0x90, 0x86, 0x26, 0x2e, 0x82, 0x79, 0xaa, 0x96, 0xd4, 0x82, 0x34, 0x0a, 0xd7,
	0xc7, 0x40, 0xa3, 0xb2, 0xca, 0xbf, 0x28, 0xb5, 0xc1, 0x46, 0x1b, 0x94, 0xda, 0x7e, 0x75, 0x45,
	0xb9, 0x42, 0xa0, 0x36, 0x34, 0x89, 0x29, 0x5e, 0xf5, 0x6f, 0x76, 0x13, 0xb8, 0xca, 0x32, 0x25,
	0x0b, 0xdd, 0xe0, 0xab, 0x49, 0xae, 0x5e, 0x84, 0xb1, 0x28, 0x22, 0x67, 0x61, 0x36, 0xa9, 0xe8,
	0x13, 0x0c, 0x85, 0x14, 0x32, 0x9e, 0x84, 0x26, 0xf1, 0x53, 0x72, 0xb6, 0x54, 0x98, 0x00, 0x4e,
	0x11, 0x74, 0x2a, 0x78, 0x68, 0x7a, 0x8d, 0xcb, 0xc6, 0xb0, 0x33, 0x1a, 0x07, 0x3b, 0xbb, 0x0b,
	0xf6, 0xda, 0xb2, 0xc2, 0xe7, 0x5d, 0x03, 0x67, 0x5e, 0xe1, 0xbd, 0xf9, 0x64, 0xfc, 0x19, 0xe9,
	0x68, 0x53, 0x93, 0x8e, 0x0e, 0x47, 0x22, 0xda, 0x54, 0x94, 0x4f, 0xe2, 0xf1, 0x85, 0x80, 0x79,
	0x0d, 0x6a, 0x1e, 0x0e, 0xd4, 0xcd, 0xad, 0x2b, 0xd6, 0x98, 0x10, 0x74, 0x72, 0xaa, 0x55, 0x2a,
	0xf8, 0xaa, 0x77, 0x9c, 0x73, 0xae, 0xf7, 0x70, 0x98, 0x93, 0x6f, 0xb9, 0x96, 0xb5, 0xb1, 0x3c,
	0x0e, 0xbe, 0x1b, 0x64, 0xf8, 0xdf, 0x06, 0xfc, 0x3e, 0x39, 0xfd, 0xf5, 0xab, 0x5a, 0xac, 0xaa,
	0xfd, 0x73, 0xd2, 0x12, 0x59, 0x18, 0x43, 0xbe, 0xd9, 0x36, 0x2b, 0x0a, 0xff, 0x81, 0xb4, 0x11,
	0x8c, 0x72, 0xc8, 0xa1, 0x5c, 0x45, 0xaf, 0x6e, 0x71, 0x1d, 0xb0, 0x80, 0x95, 0xf7, 0xac, 0x96,
	0xfa, 0xaf, 0xc4, 0x43, 0x30, 0x36, 0x44, 0xbb, 0x3d, 0x9f, 0x37, 0x1a, 0xee, 0x9b, 0xaf, 0x78,
	0xb0, 0x99, 0xb1, 0x8b, 0xdb, 0xe5, 0xf3, 0xd3, 0xc7, 0x63, 0x2c, 0xec, 0xc2, 0x45, 0x01, 0x57,
	0x19, 0xcd, 0x4d, 0x14, 0xc6, 0xb4, 0x0a, 0x74, 0x0c, 0x92, 0xea, 0xe8, 0x2e, 0x56, 0xf4, 0x6f,
	0xc6, 0xa3, 0x93, 0x3c, 0xd4, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x01, 0x7d, 0x16,
	0x60, 0x03, 0x00, 0x00,
}
