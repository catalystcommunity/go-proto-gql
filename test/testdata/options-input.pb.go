// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: options-input.proto

package pb

import (
	_ "github.com/catalystsquad/go-proto-gql/pkg/graphqlpb"
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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringX string    `protobuf:"bytes,1,opt,name=string_x,json=stringX,proto3" json:"string_x,omitempty"` // must be required
	Foo     *Foo2     `protobuf:"bytes,2,opt,name=foo,proto3" json:"foo,omitempty"`                        // must be required
	Double  []float64 `protobuf:"fixed64,3,rep,packed,name=double,proto3" json:"double,omitempty"`         // must be required because its greater than 0
	String2 string    `protobuf:"bytes,4,opt,name=string2,proto3" json:"string2,omitempty"`                // simple
	Foo2    *Foo2     `protobuf:"bytes,5,opt,name=foo2,proto3" json:"foo2,omitempty"`                      // simple
	Double2 []float64 `protobuf:"fixed64,6,rep,packed,name=double2,proto3" json:"double2,omitempty"`       // simple
	Bar     string    `protobuf:"bytes,7,opt,name=bar,proto3" json:"bar,omitempty"`
	String_ string    `protobuf:"bytes,8,opt,name=string,proto3" json:"string,omitempty"`
	Ignore  string    `protobuf:"bytes,9,opt,name=ignore,proto3" json:"ignore,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_options_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_options_input_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetStringX() string {
	if x != nil {
		return x.StringX
	}
	return ""
}

func (x *Data) GetFoo() *Foo2 {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *Data) GetDouble() []float64 {
	if x != nil {
		return x.Double
	}
	return nil
}

func (x *Data) GetString2() string {
	if x != nil {
		return x.String2
	}
	return ""
}

func (x *Data) GetFoo2() *Foo2 {
	if x != nil {
		return x.Foo2
	}
	return nil
}

func (x *Data) GetDouble2() []float64 {
	if x != nil {
		return x.Double2
	}
	return nil
}

func (x *Data) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

func (x *Data) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Data) GetIgnore() string {
	if x != nil {
		return x.Ignore
	}
	return ""
}

type Foo2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param1 string `protobuf:"bytes,1,opt,name=param1,proto3" json:"param1,omitempty"`
}

func (x *Foo2) Reset() {
	*x = Foo2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo2) ProtoMessage() {}

func (x *Foo2) ProtoReflect() protoreflect.Message {
	mi := &file_options_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo2.ProtoReflect.Descriptor instead.
func (*Foo2) Descriptor() ([]byte, []int) {
	return file_options_input_proto_rawDescGZIP(), []int{1}
}

func (x *Foo2) GetParam1() string {
	if x != nil {
		return x.Param1
	}
	return ""
}

var File_options_input_proto protoreflect.FileDescriptor

var file_options_input_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x27,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x21, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x01, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x58, 0x12, 0x27, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x6f, 0x6f, 0x32, 0x42,
	0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x01, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x1e, 0x0a, 0x06,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01, 0x42, 0x06, 0xb2, 0xe0,
	0x1f, 0x02, 0x08, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x21, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x32, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46,
	0x6f, 0x6f, 0x32, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x32, 0x18, 0x06, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x32, 0x12, 0x1c, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xb2, 0xe0, 0x1f, 0x06, 0x2a, 0x04, 0x62, 0x61, 0x72, 0x73, 0x52, 0x03, 0x62, 0x61,
	0x72, 0x12, 0x21, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xb2, 0xe0, 0x1f, 0x05, 0x2a, 0x03, 0x73, 0x74, 0x72, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x06, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x22, 0x1e, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x32, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x31, 0x32, 0xe3, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x07, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x31, 0x12, 0x0d, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x07, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x32, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x2e, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x72, 0x79, 0x31, 0x12, 0x0d, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02,
	0x08, 0x02, 0x12, 0x29, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x0d, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x12, 0x2b, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x12, 0x2b, 0x0a, 0x07, 0x50, 0x75,
	0x62, 0x53, 0x75, 0x62, 0x31, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x31, 0x12, 0x0d, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02,
	0x08, 0x02, 0x28, 0x01, 0x12, 0x3b, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x32, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x01, 0x30,
	0x01, 0x12, 0x3d, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x33, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x02, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x33, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x32, 0x12, 0x0d, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x06, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12,
	0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2,
	0xe0, 0x1f, 0x02, 0x10, 0x01, 0x12, 0x33, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0d, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x0d, 0xb2, 0xe0, 0x1f,
	0x09, 0x1a, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x5d, 0x0a, 0x04, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x1a, 0x00, 0x32, 0xbd, 0x01, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x72, 0x79, 0x31, 0x12, 0x0d, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x32, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x07, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x31, 0x12, 0x0d,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0,
	0x1f, 0x02, 0x08, 0x01, 0x12, 0x2b, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x0d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x30,
	0x01, 0x1a, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x02, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c,
	0x61, 0x64, 0x63, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x71,
	0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_input_proto_rawDescOnce sync.Once
	file_options_input_proto_rawDescData = file_options_input_proto_rawDesc
)

func file_options_input_proto_rawDescGZIP() []byte {
	file_options_input_proto_rawDescOnce.Do(func() {
		file_options_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_input_proto_rawDescData)
	})
	return file_options_input_proto_rawDescData
}

var file_options_input_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_options_input_proto_goTypes = []interface{}{
	(*Data)(nil), // 0: options.Data
	(*Foo2)(nil), // 1: options.Foo2
}
var file_options_input_proto_depIdxs = []int32{
	1,  // 0: options.Data.foo:type_name -> options.Foo2
	1,  // 1: options.Data.foo2:type_name -> options.Foo2
	0,  // 2: options.Service.Mutate1:input_type -> options.Data
	0,  // 3: options.Service.Mutate2:input_type -> options.Data
	0,  // 4: options.Service.Query1:input_type -> options.Data
	0,  // 5: options.Service.Publish:input_type -> options.Data
	0,  // 6: options.Service.Subscribe:input_type -> options.Data
	0,  // 7: options.Service.PubSub1:input_type -> options.Data
	0,  // 8: options.Service.InvalidSubscribe1:input_type -> options.Data
	0,  // 9: options.Service.InvalidSubscribe2:input_type -> options.Data
	0,  // 10: options.Service.InvalidSubscribe3:input_type -> options.Data
	0,  // 11: options.Service.PubSub2:input_type -> options.Data
	0,  // 12: options.Service.Ignore:input_type -> options.Data
	0,  // 13: options.Service.Name:input_type -> options.Data
	0,  // 14: options.Test.Name:input_type -> options.Data
	0,  // 15: options.Test.NewName:input_type -> options.Data
	0,  // 16: options.Query.Query1:input_type -> options.Data
	0,  // 17: options.Query.Query2:input_type -> options.Data
	0,  // 18: options.Query.Mutate1:input_type -> options.Data
	0,  // 19: options.Query.Subscribe:input_type -> options.Data
	0,  // 20: options.Service.Mutate1:output_type -> options.Data
	0,  // 21: options.Service.Mutate2:output_type -> options.Data
	0,  // 22: options.Service.Query1:output_type -> options.Data
	0,  // 23: options.Service.Publish:output_type -> options.Data
	0,  // 24: options.Service.Subscribe:output_type -> options.Data
	0,  // 25: options.Service.PubSub1:output_type -> options.Data
	0,  // 26: options.Service.InvalidSubscribe1:output_type -> options.Data
	0,  // 27: options.Service.InvalidSubscribe2:output_type -> options.Data
	0,  // 28: options.Service.InvalidSubscribe3:output_type -> options.Data
	0,  // 29: options.Service.PubSub2:output_type -> options.Data
	0,  // 30: options.Service.Ignore:output_type -> options.Data
	0,  // 31: options.Service.Name:output_type -> options.Data
	0,  // 32: options.Test.Name:output_type -> options.Data
	0,  // 33: options.Test.NewName:output_type -> options.Data
	0,  // 34: options.Query.Query1:output_type -> options.Data
	0,  // 35: options.Query.Query2:output_type -> options.Data
	0,  // 36: options.Query.Mutate1:output_type -> options.Data
	0,  // 37: options.Query.Subscribe:output_type -> options.Data
	20, // [20:38] is the sub-list for method output_type
	2,  // [2:20] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_options_input_proto_init() }
func file_options_input_proto_init() {
	if File_options_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_options_input_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo2); i {
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
			RawDescriptor: file_options_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_options_input_proto_goTypes,
		DependencyIndexes: file_options_input_proto_depIdxs,
		MessageInfos:      file_options_input_proto_msgTypes,
	}.Build()
	File_options_input_proto = out.File
	file_options_input_proto_rawDesc = nil
	file_options_input_proto_goTypes = nil
	file_options_input_proto_depIdxs = nil
}
