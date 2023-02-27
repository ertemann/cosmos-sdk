// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: testpb/test_schema.proto

package testpb

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum int32

const (
	Enum_ENUM_UNSPECIFIED Enum = 0
	Enum_ENUM_ONE         Enum = 1
	Enum_ENUM_TWO         Enum = 2
	Enum_ENUM_FIVE        Enum = 5
	Enum_ENUM_NEG_THREE   Enum = -3
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0:  "ENUM_UNSPECIFIED",
		1:  "ENUM_ONE",
		2:  "ENUM_TWO",
		5:  "ENUM_FIVE",
		-3: "ENUM_NEG_THREE",
	}
	Enum_value = map[string]int32{
		"ENUM_UNSPECIFIED": 0,
		"ENUM_ONE":         1,
		"ENUM_TWO":         2,
		"ENUM_FIVE":        5,
		"ENUM_NEG_THREE":   -3,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_testpb_test_schema_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_testpb_test_schema_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{0}
}

type ExampleTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Valid key fields:
	U32  uint32                 `protobuf:"varint,1,opt,name=u32,proto3" json:"u32,omitempty"`
	U64  uint64                 `protobuf:"varint,2,opt,name=u64,proto3" json:"u64,omitempty"`
	Str  string                 `protobuf:"bytes,3,opt,name=str,proto3" json:"str,omitempty"`
	Bz   []byte                 `protobuf:"bytes,4,opt,name=bz,proto3" json:"bz,omitempty"`
	Ts   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ts,proto3" json:"ts,omitempty"`
	Dur  *durationpb.Duration   `protobuf:"bytes,6,opt,name=dur,proto3" json:"dur,omitempty"`
	I32  int32                  `protobuf:"varint,7,opt,name=i32,proto3" json:"i32,omitempty"`
	S32  int32                  `protobuf:"zigzag32,8,opt,name=s32,proto3" json:"s32,omitempty"`
	Sf32 int32                  `protobuf:"fixed32,9,opt,name=sf32,proto3" json:"sf32,omitempty"`
	I64  int64                  `protobuf:"varint,10,opt,name=i64,proto3" json:"i64,omitempty"`
	S64  int64                  `protobuf:"zigzag64,11,opt,name=s64,proto3" json:"s64,omitempty"`
	Sf64 int64                  `protobuf:"fixed64,12,opt,name=sf64,proto3" json:"sf64,omitempty"`
	F32  uint32                 `protobuf:"fixed32,13,opt,name=f32,proto3" json:"f32,omitempty"`
	F64  uint64                 `protobuf:"fixed64,14,opt,name=f64,proto3" json:"f64,omitempty"`
	B    bool                   `protobuf:"varint,15,opt,name=b,proto3" json:"b,omitempty"`
	E    Enum                   `protobuf:"varint,16,opt,name=e,proto3,enum=testpb.Enum" json:"e,omitempty"`
	// Invalid key fields:
	Repeated []uint32                     `protobuf:"varint,17,rep,packed,name=repeated,proto3" json:"repeated,omitempty"`
	Map      map[string]uint32            `protobuf:"bytes,18,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Msg      *ExampleTable_ExampleMessage `protobuf:"bytes,19,opt,name=msg,proto3" json:"msg,omitempty"`
	// Types that are assignable to Sum:
	//	*ExampleTable_Oneof
	Sum isExampleTable_Sum `protobuf_oneof:"sum"`
}

func (x *ExampleTable) Reset() {
	*x = ExampleTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleTable) ProtoMessage() {}

func (x *ExampleTable) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleTable.ProtoReflect.Descriptor instead.
func (*ExampleTable) Descriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleTable) GetU32() uint32 {
	if x != nil {
		return x.U32
	}
	return 0
}

func (x *ExampleTable) GetU64() uint64 {
	if x != nil {
		return x.U64
	}
	return 0
}

func (x *ExampleTable) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *ExampleTable) GetBz() []byte {
	if x != nil {
		return x.Bz
	}
	return nil
}

func (x *ExampleTable) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ExampleTable) GetDur() *durationpb.Duration {
	if x != nil {
		return x.Dur
	}
	return nil
}

func (x *ExampleTable) GetI32() int32 {
	if x != nil {
		return x.I32
	}
	return 0
}

func (x *ExampleTable) GetS32() int32 {
	if x != nil {
		return x.S32
	}
	return 0
}

func (x *ExampleTable) GetSf32() int32 {
	if x != nil {
		return x.Sf32
	}
	return 0
}

func (x *ExampleTable) GetI64() int64 {
	if x != nil {
		return x.I64
	}
	return 0
}

func (x *ExampleTable) GetS64() int64 {
	if x != nil {
		return x.S64
	}
	return 0
}

func (x *ExampleTable) GetSf64() int64 {
	if x != nil {
		return x.Sf64
	}
	return 0
}

func (x *ExampleTable) GetF32() uint32 {
	if x != nil {
		return x.F32
	}
	return 0
}

func (x *ExampleTable) GetF64() uint64 {
	if x != nil {
		return x.F64
	}
	return 0
}

func (x *ExampleTable) GetB() bool {
	if x != nil {
		return x.B
	}
	return false
}

func (x *ExampleTable) GetE() Enum {
	if x != nil {
		return x.E
	}
	return Enum_ENUM_UNSPECIFIED
}

func (x *ExampleTable) GetRepeated() []uint32 {
	if x != nil {
		return x.Repeated
	}
	return nil
}

func (x *ExampleTable) GetMap() map[string]uint32 {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *ExampleTable) GetMsg() *ExampleTable_ExampleMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (m *ExampleTable) GetSum() isExampleTable_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (x *ExampleTable) GetOneof() uint32 {
	if x, ok := x.GetSum().(*ExampleTable_Oneof); ok {
		return x.Oneof
	}
	return 0
}

type isExampleTable_Sum interface {
	isExampleTable_Sum()
}

type ExampleTable_Oneof struct {
	Oneof uint32 `protobuf:"varint,20,opt,name=oneof,proto3,oneof"`
}

func (*ExampleTable_Oneof) isExampleTable_Sum() {}

type ExampleAutoIncrementTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	X  string `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y  int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *ExampleAutoIncrementTable) Reset() {
	*x = ExampleAutoIncrementTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleAutoIncrementTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleAutoIncrementTable) ProtoMessage() {}

func (x *ExampleAutoIncrementTable) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleAutoIncrementTable.ProtoReflect.Descriptor instead.
func (*ExampleAutoIncrementTable) Descriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleAutoIncrementTable) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExampleAutoIncrementTable) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *ExampleAutoIncrementTable) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type ExampleSingleton struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar int32  `protobuf:"varint,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *ExampleSingleton) Reset() {
	*x = ExampleSingleton{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleSingleton) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleSingleton) ProtoMessage() {}

func (x *ExampleSingleton) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleSingleton.ProtoReflect.Descriptor instead.
func (*ExampleSingleton) Descriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{2}
}

func (x *ExampleSingleton) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

func (x *ExampleSingleton) GetBar() int32 {
	if x != nil {
		return x.Bar
	}
	return 0
}

type ExampleTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ts   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *ExampleTimestamp) Reset() {
	*x = ExampleTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleTimestamp) ProtoMessage() {}

func (x *ExampleTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleTimestamp.ProtoReflect.Descriptor instead.
func (*ExampleTimestamp) Descriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{3}
}

func (x *ExampleTimestamp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExampleTimestamp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExampleTimestamp) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type ExampleDuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dur  *durationpb.Duration `protobuf:"bytes,3,opt,name=dur,proto3" json:"dur,omitempty"`
}

func (x *ExampleDuration) Reset() {
	*x = ExampleDuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleDuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleDuration) ProtoMessage() {}

func (x *ExampleDuration) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleDuration.ProtoReflect.Descriptor instead.
func (*ExampleDuration) Descriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{4}
}

func (x *ExampleDuration) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExampleDuration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExampleDuration) GetDur() *durationpb.Duration {
	if x != nil {
		return x.Dur
	}
	return nil
}

type SimpleExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Unique    string `protobuf:"bytes,2,opt,name=unique,proto3" json:"unique,omitempty"`
	NotUnique string `protobuf:"bytes,3,opt,name=not_unique,json=notUnique,proto3" json:"not_unique,omitempty"`
}

func (x *SimpleExample) Reset() {
	*x = SimpleExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleExample) ProtoMessage() {}

func (x *SimpleExample) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleExample.ProtoReflect.Descriptor instead.
func (*SimpleExample) Descriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{5}
}

func (x *SimpleExample) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleExample) GetUnique() string {
	if x != nil {
		return x.Unique
	}
	return ""
}

func (x *SimpleExample) GetNotUnique() string {
	if x != nil {
		return x.NotUnique
	}
	return ""
}

// ExampleAutoIncFieldName is a table for testing InsertReturning<FieldName>.
type ExampleAutoIncFieldName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo uint64 `protobuf:"varint,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar uint64 `protobuf:"varint,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *ExampleAutoIncFieldName) Reset() {
	*x = ExampleAutoIncFieldName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleAutoIncFieldName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleAutoIncFieldName) ProtoMessage() {}

func (x *ExampleAutoIncFieldName) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleAutoIncFieldName.ProtoReflect.Descriptor instead.
func (*ExampleAutoIncFieldName) Descriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{6}
}

func (x *ExampleAutoIncFieldName) GetFoo() uint64 {
	if x != nil {
		return x.Foo
	}
	return 0
}

func (x *ExampleAutoIncFieldName) GetBar() uint64 {
	if x != nil {
		return x.Bar
	}
	return 0
}

type ExampleTable_ExampleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar int32  `protobuf:"varint,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *ExampleTable_ExampleMessage) Reset() {
	*x = ExampleTable_ExampleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_schema_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleTable_ExampleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleTable_ExampleMessage) ProtoMessage() {}

func (x *ExampleTable_ExampleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_schema_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleTable_ExampleMessage.ProtoReflect.Descriptor instead.
func (*ExampleTable_ExampleMessage) Descriptor() ([]byte, []int) {
	return file_testpb_test_schema_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ExampleTable_ExampleMessage) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

func (x *ExampleTable_ExampleMessage) GetBar() int32 {
	if x != nil {
		return x.Bar
	}
	return 0
}

var File_testpb_test_schema_proto protoreflect.FileDescriptor

var file_testpb_test_schema_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x05, 0x0a,
	0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x33, 0x32, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x36,
	0x34, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x74, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x62, 0x7a, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12,
	0x2b, 0x0a, 0x03, 0x64, 0x75, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x64, 0x75, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x33, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x11, 0x52, 0x03, 0x73, 0x33, 0x32,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x66, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x04,
	0x73, 0x66, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x69, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x36, 0x34, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x03, 0x73, 0x36, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x66, 0x36, 0x34,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x04, 0x73, 0x66, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x33, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x07, 0x52, 0x03, 0x66, 0x33, 0x32, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x36, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x06, 0x52, 0x03, 0x66, 0x36, 0x34,
	0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x01, 0x62, 0x12, 0x1a,
	0x0a, 0x01, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x62, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x01, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x35, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16,
	0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x34,
	0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66,
	0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x62, 0x61, 0x72, 0x3a, 0x3f, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x39, 0x0a, 0x0d, 0x0a, 0x0b,
	0x75, 0x33, 0x32, 0x2c, 0x69, 0x36, 0x34, 0x2c, 0x73, 0x74, 0x72, 0x12, 0x0d, 0x0a, 0x07, 0x75,
	0x36, 0x34, 0x2c, 0x73, 0x74, 0x72, 0x10, 0x01, 0x18, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x73, 0x74,
	0x72, 0x2c, 0x75, 0x33, 0x32, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x62, 0x7a, 0x2c, 0x73, 0x74,
	0x72, 0x10, 0x03, 0x18, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x62, 0x0a, 0x19,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x79, 0x3a, 0x19, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x13, 0x0a, 0x06, 0x0a,
	0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x01, 0x78, 0x10, 0x01, 0x18, 0x01, 0x18, 0x03,
	0x22, 0x40, 0x0a, 0x10, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x74, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x61, 0x72, 0x3a, 0x08, 0xfa, 0x9e, 0xd3, 0x8e, 0x03, 0x02,
	0x08, 0x02, 0x22, 0x7c, 0x0a, 0x10, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x3a, 0x18, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x12, 0x0a, 0x06,
	0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x74, 0x73, 0x10, 0x01, 0x18, 0x04,
	0x22, 0x7d, 0x0a, 0x0f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x64, 0x75, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x03, 0x64, 0x75, 0x72, 0x3a, 0x19, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x13, 0x0a, 0x06, 0x0a, 0x02,
	0x69, 0x64, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x75, 0x72, 0x10, 0x01, 0x18, 0x04, 0x22,
	0x7a, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x6f, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x6f, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x3a, 0x1e, 0xf2, 0x9e, 0xd3,
	0x8e, 0x03, 0x18, 0x0a, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x06, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x10, 0x01, 0x18, 0x01, 0x18, 0x05, 0x22, 0x50, 0x0a, 0x17, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x49, 0x6e, 0x63, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x62, 0x61, 0x72, 0x3a, 0x11, 0xf2, 0x9e, 0xd3, 0x8e,
	0x03, 0x0b, 0x0a, 0x07, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x10, 0x01, 0x18, 0x06, 0x2a, 0x64, 0x0a,
	0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45,
	0x4e, 0x55, 0x4d, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4e, 0x55, 0x4d, 0x5f,
	0x46, 0x49, 0x56, 0x45, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x0e, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4e,
	0x45, 0x47, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x10, 0xfd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x42, 0x87, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x62, 0x42, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x06,
	0x54, 0x65, 0x73, 0x74, 0x70, 0x62, 0xca, 0x02, 0x06, 0x54, 0x65, 0x73, 0x74, 0x70, 0x62, 0xe2,
	0x02, 0x12, 0x54, 0x65, 0x73, 0x74, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x54, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testpb_test_schema_proto_rawDescOnce sync.Once
	file_testpb_test_schema_proto_rawDescData = file_testpb_test_schema_proto_rawDesc
)

func file_testpb_test_schema_proto_rawDescGZIP() []byte {
	file_testpb_test_schema_proto_rawDescOnce.Do(func() {
		file_testpb_test_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_testpb_test_schema_proto_rawDescData)
	})
	return file_testpb_test_schema_proto_rawDescData
}

var file_testpb_test_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testpb_test_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_testpb_test_schema_proto_goTypes = []interface{}{
	(Enum)(0),                           // 0: testpb.Enum
	(*ExampleTable)(nil),                // 1: testpb.ExampleTable
	(*ExampleAutoIncrementTable)(nil),   // 2: testpb.ExampleAutoIncrementTable
	(*ExampleSingleton)(nil),            // 3: testpb.ExampleSingleton
	(*ExampleTimestamp)(nil),            // 4: testpb.ExampleTimestamp
	(*ExampleDuration)(nil),             // 5: testpb.ExampleDuration
	(*SimpleExample)(nil),               // 6: testpb.SimpleExample
	(*ExampleAutoIncFieldName)(nil),     // 7: testpb.ExampleAutoIncFieldName
	nil,                                 // 8: testpb.ExampleTable.MapEntry
	(*ExampleTable_ExampleMessage)(nil), // 9: testpb.ExampleTable.ExampleMessage
	(*timestamppb.Timestamp)(nil),       // 10: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),         // 11: google.protobuf.Duration
}
var file_testpb_test_schema_proto_depIdxs = []int32{
	10, // 0: testpb.ExampleTable.ts:type_name -> google.protobuf.Timestamp
	11, // 1: testpb.ExampleTable.dur:type_name -> google.protobuf.Duration
	0,  // 2: testpb.ExampleTable.e:type_name -> testpb.Enum
	8,  // 3: testpb.ExampleTable.map:type_name -> testpb.ExampleTable.MapEntry
	9,  // 4: testpb.ExampleTable.msg:type_name -> testpb.ExampleTable.ExampleMessage
	10, // 5: testpb.ExampleTimestamp.ts:type_name -> google.protobuf.Timestamp
	11, // 6: testpb.ExampleDuration.dur:type_name -> google.protobuf.Duration
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_testpb_test_schema_proto_init() }
func file_testpb_test_schema_proto_init() {
	if File_testpb_test_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testpb_test_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleTable); i {
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
		file_testpb_test_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleAutoIncrementTable); i {
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
		file_testpb_test_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleSingleton); i {
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
		file_testpb_test_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleTimestamp); i {
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
		file_testpb_test_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleDuration); i {
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
		file_testpb_test_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleExample); i {
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
		file_testpb_test_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleAutoIncFieldName); i {
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
		file_testpb_test_schema_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleTable_ExampleMessage); i {
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
	file_testpb_test_schema_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ExampleTable_Oneof)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testpb_test_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testpb_test_schema_proto_goTypes,
		DependencyIndexes: file_testpb_test_schema_proto_depIdxs,
		EnumInfos:         file_testpb_test_schema_proto_enumTypes,
		MessageInfos:      file_testpb_test_schema_proto_msgTypes,
	}.Build()
	File_testpb_test_schema_proto = out.File
	file_testpb_test_schema_proto_rawDesc = nil
	file_testpb_test_schema_proto_goTypes = nil
	file_testpb_test_schema_proto_depIdxs = nil
}
