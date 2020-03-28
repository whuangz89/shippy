// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Created              bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Vessel               *Vessel   `protobuf:"bytes,2,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,3,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "Vessel")
	proto.RegisterType((*Specification)(nil), "Specification")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() {
	proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716)
}

var fileDescriptor_04ef66862bb50716 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0xb4, 0xcd, 0xc7, 0x68, 0x7a, 0xd8, 0xd3, 0xb6, 0x28, 0xc6, 0x9c, 0x82, 0x87,
	0x08, 0xf5, 0x17, 0x88, 0x20, 0xe8, 0x71, 0x0b, 0x7a, 0x0c, 0xdb, 0xcd, 0x68, 0x17, 0xf2, 0x45,
	0x12, 0xd2, 0xfa, 0x6f, 0xfc, 0xa9, 0xc2, 0xe4, 0xc3, 0xd6, 0x83, 0xa7, 0xcc, 0x3c, 0x99, 0x7d,
	0x79, 0xdf, 0x19, 0x58, 0x55, 0x75, 0xd9, 0x96, 0xf7, 0x1d, 0x36, 0x0d, 0x66, 0xc3, 0x27, 0x26,
	0x16, 0x7e, 0x33, 0xb0, 0xdf, 0x08, 0xf0, 0x25, 0x58, 0x26, 0x15, 0x2c, 0x60, 0x91, 0x27, 0x2d,
	0x93, 0xf2, 0x35, 0xb8, 0x5a, 0x55, 0x4a, 0x9b, 0xf6, 0x4b, 0x58, 0x01, 0x8b, 0x16, 0x72, 0xea,
	0xf9, 0x35, 0x40, 0xae, 0x8e, 0xc9, 0x01, 0xcd, 0xe7, 0xbe, 0x15, 0x33, 0xfa, 0xeb, 0xe5, 0xea,
	0xf8, 0x4e, 0x80, 0x73, 0x98, 0x17, 0x2a, 0x47, 0x31, 0x27, 0x31, 0xaa, 0xf9, 0x15, 0x78, 0xaa,
	0x53, 0x26, 0x53, 0xbb, 0x0c, 0xc5, 0x22, 0x60, 0x91, 0x2b, 0x7f, 0x01, 0x5f, 0x81, 0x5b, 0x1e,
	0x0a, 0xac, 0x13, 0x93, 0x0a, 0x9b, 0x5e, 0x39, 0xd4, 0xbf, 0xa4, 0xe1, 0x2b, 0xf8, 0xdb, 0x0a,
	0xb5, 0xf9, 0x30, 0x5a, 0xb5, 0xa6, 0x2c, 0xce, 0x8c, 0xb1, 0x7f, 0x8d, 0x59, 0x7f, 0x8c, 0x85,
	0x7b, 0x70, 0x25, 0x36, 0x55, 0x59, 0x34, 0xc8, 0x05, 0x38, 0xba, 0x46, 0xd5, 0x62, 0x1f, 0xda,
	0x95, 0x63, 0xcb, 0x6f, 0xc0, 0xee, 0x97, 0x44, 0x02, 0x17, 0x1b, 0x27, 0xee, 0x57, 0x24, 0x07,
	0xcc, 0x6f, 0xc1, 0xe9, 0xab, 0x46, 0xcc, 0x82, 0xd9, 0xe9, 0xc4, 0xc8, 0x37, 0x09, 0xf8, 0x3d,
	0xda, 0x62, 0xdd, 0x19, 0x8d, 0xfc, 0x0e, 0xfc, 0x67, 0x53, 0xa4, 0x8f, 0x53, 0xe4, 0x65, 0x7c,
	0x16, 0x6b, 0xed, 0xc5, 0x93, 0xb5, 0x10, 0x2e, 0x9f, 0xc8, 0xcb, 0x70, 0x9a, 0x51, 0xfe, 0x64,
	0x66, 0x67, 0xd3, 0x01, 0x1f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xdb, 0x01, 0xba, 0xdd,
	0x01, 0x00, 0x00,
}
