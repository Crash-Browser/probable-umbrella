// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: eth/v1/events.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	github_com_prysmaticlabs_prysm_v4_consensus_types_primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	_ "github.com/prysmaticlabs/prysm/v4/proto/engine/v1"
	_ "github.com/prysmaticlabs/prysm/v4/eth/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot                      github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot"`
	Block                     []byte                                                            `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty" ssz-size:"32"`
	State                     []byte                                                            `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty" ssz-size:"32"`
	EpochTransition           bool                                                              `protobuf:"varint,4,opt,name=epoch_transition,json=epochTransition,proto3" json:"epoch_transition,omitempty"`
	PreviousDutyDependentRoot []byte                                                            `protobuf:"bytes,5,opt,name=previous_duty_dependent_root,json=previousDutyDependentRoot,proto3" json:"previous_duty_dependent_root,omitempty" ssz-size:"32"`
	CurrentDutyDependentRoot  []byte                                                            `protobuf:"bytes,6,opt,name=current_duty_dependent_root,json=currentDutyDependentRoot,proto3" json:"current_duty_dependent_root,omitempty" ssz-size:"32"`
	ExecutionOptimistic       bool                                                              `protobuf:"varint,7,opt,name=execution_optimistic,json=executionOptimistic,proto3" json:"execution_optimistic,omitempty"`
}

func (x *EventHead) Reset() {
	*x = EventHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHead) ProtoMessage() {}

func (x *EventHead) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHead.ProtoReflect.Descriptor instead.
func (*EventHead) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *EventHead) GetSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(0)
}

func (x *EventHead) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *EventHead) GetState() []byte {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *EventHead) GetEpochTransition() bool {
	if x != nil {
		return x.EpochTransition
	}
	return false
}

func (x *EventHead) GetPreviousDutyDependentRoot() []byte {
	if x != nil {
		return x.PreviousDutyDependentRoot
	}
	return nil
}

func (x *EventHead) GetCurrentDutyDependentRoot() []byte {
	if x != nil {
		return x.CurrentDutyDependentRoot
	}
	return nil
}

func (x *EventHead) GetExecutionOptimistic() bool {
	if x != nil {
		return x.ExecutionOptimistic
	}
	return false
}

type EventBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot                github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot"`
	Block               []byte                                                            `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty" ssz-size:"32"`
	ExecutionOptimistic bool                                                              `protobuf:"varint,3,opt,name=execution_optimistic,json=executionOptimistic,proto3" json:"execution_optimistic,omitempty"`
}

func (x *EventBlock) Reset() {
	*x = EventBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBlock) ProtoMessage() {}

func (x *EventBlock) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventBlock.ProtoReflect.Descriptor instead.
func (*EventBlock) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_proto_rawDescGZIP(), []int{1}
}

func (x *EventBlock) GetSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(0)
}

func (x *EventBlock) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *EventBlock) GetExecutionOptimistic() bool {
	if x != nil {
		return x.ExecutionOptimistic
	}
	return false
}

type EventChainReorg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot                github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot  `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot"`
	Depth               uint64                                                             `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"`
	OldHeadBlock        []byte                                                             `protobuf:"bytes,3,opt,name=old_head_block,json=oldHeadBlock,proto3" json:"old_head_block,omitempty" ssz-size:"32"`
	NewHeadBlock        []byte                                                             `protobuf:"bytes,4,opt,name=new_head_block,json=newHeadBlock,proto3" json:"new_head_block,omitempty" ssz-size:"32"`
	OldHeadState        []byte                                                             `protobuf:"bytes,5,opt,name=old_head_state,json=oldHeadState,proto3" json:"old_head_state,omitempty" ssz-size:"32"`
	NewHeadState        []byte                                                             `protobuf:"bytes,6,opt,name=new_head_state,json=newHeadState,proto3" json:"new_head_state,omitempty" ssz-size:"32"`
	Epoch               github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch `protobuf:"varint,7,opt,name=epoch,proto3" json:"epoch,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Epoch"`
	ExecutionOptimistic bool                                                               `protobuf:"varint,8,opt,name=execution_optimistic,json=executionOptimistic,proto3" json:"execution_optimistic,omitempty"`
}

func (x *EventChainReorg) Reset() {
	*x = EventChainReorg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventChainReorg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventChainReorg) ProtoMessage() {}

func (x *EventChainReorg) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventChainReorg.ProtoReflect.Descriptor instead.
func (*EventChainReorg) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_proto_rawDescGZIP(), []int{2}
}

func (x *EventChainReorg) GetSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(0)
}

func (x *EventChainReorg) GetDepth() uint64 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *EventChainReorg) GetOldHeadBlock() []byte {
	if x != nil {
		return x.OldHeadBlock
	}
	return nil
}

func (x *EventChainReorg) GetNewHeadBlock() []byte {
	if x != nil {
		return x.NewHeadBlock
	}
	return nil
}

func (x *EventChainReorg) GetOldHeadState() []byte {
	if x != nil {
		return x.OldHeadState
	}
	return nil
}

func (x *EventChainReorg) GetNewHeadState() []byte {
	if x != nil {
		return x.NewHeadState
	}
	return nil
}

func (x *EventChainReorg) GetEpoch() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch {
	if x != nil {
		return x.Epoch
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch(0)
}

func (x *EventChainReorg) GetExecutionOptimistic() bool {
	if x != nil {
		return x.ExecutionOptimistic
	}
	return false
}

type EventFinalizedCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block               []byte                                                             `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty" ssz-size:"32"`
	State               []byte                                                             `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty" ssz-size:"32"`
	Epoch               github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Epoch"`
	ExecutionOptimistic bool                                                               `protobuf:"varint,4,opt,name=execution_optimistic,json=executionOptimistic,proto3" json:"execution_optimistic,omitempty"`
}

func (x *EventFinalizedCheckpoint) Reset() {
	*x = EventFinalizedCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventFinalizedCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventFinalizedCheckpoint) ProtoMessage() {}

func (x *EventFinalizedCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventFinalizedCheckpoint.ProtoReflect.Descriptor instead.
func (*EventFinalizedCheckpoint) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_proto_rawDescGZIP(), []int{3}
}

func (x *EventFinalizedCheckpoint) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *EventFinalizedCheckpoint) GetState() []byte {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *EventFinalizedCheckpoint) GetEpoch() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch {
	if x != nil {
		return x.Epoch
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch(0)
}

func (x *EventFinalizedCheckpoint) GetExecutionOptimistic() bool {
	if x != nil {
		return x.ExecutionOptimistic
	}
	return false
}

var File_eth_v1_events_proto protoreflect.FileDescriptor

var file_eth_v1_events_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x12, 0x59, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x45, 0x82, 0xb5, 0x18, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18,
	0x02, 0x33, 0x32, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33,
	0x32, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x1c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f,
	0x64, 0x75, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33,
	0x32, 0x52, 0x19, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x44, 0x75, 0x74, 0x79, 0x44,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x45, 0x0a, 0x1b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x18, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x44, 0x75, 0x74, 0x79, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x59, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x45, 0x82, 0xb5, 0x18, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x12, 0x1c, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x31,
	0x0a, 0x14, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x22, 0xcb, 0x03, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x6f, 0x72, 0x67, 0x12, 0x59, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x45, 0x82, 0xb5, 0x18, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x0e, 0x6f, 0x6c, 0x64, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06,
	0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0c, 0x6f, 0x6c, 0x64, 0x48, 0x65, 0x61, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5,
	0x18, 0x02, 0x33, 0x32, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x48, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x0e, 0x6f, 0x6c, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02,
	0x33, 0x32, 0x52, 0x0c, 0x6f, 0x6c, 0x64, 0x48, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x2c, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32,
	0x52, 0x0c, 0x6e, 0x65, 0x77, 0x48, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x5c,
	0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x46, 0x82,
	0xb5, 0x18, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73,
	0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x31, 0x0a, 0x14,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22,
	0xe7, 0x01, 0x0a, 0x18, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18,
	0x02, 0x33, 0x32, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33,
	0x32, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x5c, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x46, 0x82, 0xb5, 0x18, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69,
	0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52,
	0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x42, 0x7e, 0x0a, 0x13, 0x6f, 0x72, 0x67,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x42, 0x11, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x45, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eth_v1_events_proto_rawDescOnce sync.Once
	file_eth_v1_events_proto_rawDescData = file_eth_v1_events_proto_rawDesc
)

func file_eth_v1_events_proto_rawDescGZIP() []byte {
	file_eth_v1_events_proto_rawDescOnce.Do(func() {
		file_eth_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1_events_proto_rawDescData)
	})
	return file_eth_v1_events_proto_rawDescData
}

var file_eth_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eth_v1_events_proto_goTypes = []interface{}{
	(*EventHead)(nil),                // 0: ethereum.eth.v1.EventHead
	(*EventBlock)(nil),               // 1: ethereum.eth.v1.EventBlock
	(*EventChainReorg)(nil),          // 2: ethereum.eth.v1.EventChainReorg
	(*EventFinalizedCheckpoint)(nil), // 3: ethereum.eth.v1.EventFinalizedCheckpoint
}
var file_eth_v1_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eth_v1_events_proto_init() }
func file_eth_v1_events_proto_init() {
	if File_eth_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eth_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHead); i {
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
		file_eth_v1_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventBlock); i {
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
		file_eth_v1_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventChainReorg); i {
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
		file_eth_v1_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventFinalizedCheckpoint); i {
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
			RawDescriptor: file_eth_v1_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1_events_proto_goTypes,
		DependencyIndexes: file_eth_v1_events_proto_depIdxs,
		MessageInfos:      file_eth_v1_events_proto_msgTypes,
	}.Build()
	File_eth_v1_events_proto = out.File
	file_eth_v1_events_proto_rawDesc = nil
	file_eth_v1_events_proto_goTypes = nil
	file_eth_v1_events_proto_depIdxs = nil
}
