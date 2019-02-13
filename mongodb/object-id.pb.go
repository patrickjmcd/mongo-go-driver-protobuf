// Code generated by protoc-gen-go. DO NOT EDIT.
// source: object-id.proto

package mongodb // import "github.com/amsokol/mongo-go-driver-protobuf/mongodb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ObjectId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectId) Reset()         { *m = ObjectId{} }
func (m *ObjectId) String() string { return proto.CompactTextString(m) }
func (*ObjectId) ProtoMessage()    {}
func (*ObjectId) Descriptor() ([]byte, []int) {
	return fileDescriptor_object_id_9932e51e332ec59d, []int{0}
}
func (m *ObjectId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectId.Unmarshal(m, b)
}
func (m *ObjectId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectId.Marshal(b, m, deterministic)
}
func (dst *ObjectId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectId.Merge(dst, src)
}
func (m *ObjectId) XXX_Size() int {
	return xxx_messageInfo_ObjectId.Size(m)
}
func (m *ObjectId) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectId.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectId proto.InternalMessageInfo

func (m *ObjectId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*ObjectId)(nil), "mongodb.ObjectId")
}

func init() { proto.RegisterFile("object-id.proto", fileDescriptor_object_id_9932e51e332ec59d) }

var fileDescriptor_object_id_9932e51e332ec59d = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4f, 0xca, 0x4a,
	0x4d, 0x2e, 0xd1, 0xcd, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0xcd, 0xcf,
	0x4b, 0xcf, 0x4f, 0x49, 0x52, 0x92, 0xe2, 0xe2, 0xf0, 0x07, 0xcb, 0x79, 0xa6, 0x08, 0xf1, 0x71,
	0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x38, 0x99, 0x46,
	0x19, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe6, 0x16, 0xe7,
	0x67, 0xe7, 0xe7, 0xe8, 0x83, 0x75, 0xea, 0xa6, 0xe7, 0xeb, 0xa6, 0x14, 0x65, 0x96, 0xa5, 0x16,
	0xe9, 0x82, 0x0d, 0x4c, 0x2a, 0x4d, 0xd3, 0x87, 0x1a, 0x99, 0xc4, 0x06, 0x16, 0x31, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0xf4, 0x60, 0x7c, 0x75, 0x00, 0x00, 0x00,
}