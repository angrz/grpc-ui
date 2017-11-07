// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package http_server is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ReflectionResponse
	Reflection
	Error
*/
package http_server

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

type ReflectionResponse struct {
	// Types that are valid to be assigned to Response:
	//	*ReflectionResponse_Reflection
	//	*ReflectionResponse_Error
	Response isReflectionResponse_Response `protobuf_oneof:"response"`
}

func (m *ReflectionResponse) Reset()                    { *m = ReflectionResponse{} }
func (m *ReflectionResponse) String() string            { return proto.CompactTextString(m) }
func (*ReflectionResponse) ProtoMessage()               {}
func (*ReflectionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isReflectionResponse_Response interface {
	isReflectionResponse_Response()
}

type ReflectionResponse_Reflection struct {
	Reflection *Reflection `protobuf:"bytes,1,opt,name=reflection,oneof"`
}
type ReflectionResponse_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,oneof"`
}

func (*ReflectionResponse_Reflection) isReflectionResponse_Response() {}
func (*ReflectionResponse_Error) isReflectionResponse_Response()      {}

func (m *ReflectionResponse) GetResponse() isReflectionResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ReflectionResponse) GetReflection() *Reflection {
	if x, ok := m.GetResponse().(*ReflectionResponse_Reflection); ok {
		return x.Reflection
	}
	return nil
}

func (m *ReflectionResponse) GetError() *Error {
	if x, ok := m.GetResponse().(*ReflectionResponse_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReflectionResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReflectionResponse_OneofMarshaler, _ReflectionResponse_OneofUnmarshaler, _ReflectionResponse_OneofSizer, []interface{}{
		(*ReflectionResponse_Reflection)(nil),
		(*ReflectionResponse_Error)(nil),
	}
}

func _ReflectionResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReflectionResponse)
	// response
	switch x := m.Response.(type) {
	case *ReflectionResponse_Reflection:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Reflection); err != nil {
			return err
		}
	case *ReflectionResponse_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReflectionResponse.Response has unexpected type %T", x)
	}
	return nil
}

func _ReflectionResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReflectionResponse)
	switch tag {
	case 1: // response.reflection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Reflection)
		err := b.DecodeMessage(msg)
		m.Response = &ReflectionResponse_Reflection{msg}
		return true, err
	case 2: // response.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Response = &ReflectionResponse_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReflectionResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReflectionResponse)
	// response
	switch x := m.Response.(type) {
	case *ReflectionResponse_Reflection:
		s := proto.Size(x.Reflection)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReflectionResponse_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Reflection struct {
	Service        []string `protobuf:"bytes,1,rep,name=service" json:"service,omitempty"`
	FileDescriptor [][]byte `protobuf:"bytes,2,rep,name=file_descriptor,json=fileDescriptor,proto3" json:"file_descriptor,omitempty"`
}

func (m *Reflection) Reset()                    { *m = Reflection{} }
func (m *Reflection) String() string            { return proto.CompactTextString(m) }
func (*Reflection) ProtoMessage()               {}
func (*Reflection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Error struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*ReflectionResponse)(nil), "ReflectionResponse")
	proto.RegisterType((*Reflection)(nil), "Reflection")
	proto.RegisterType((*Error)(nil), "Error")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x3f, 0xaf, 0x82, 0x30,
	0x14, 0x47, 0x1f, 0x8f, 0xc0, 0x7b, 0x5c, 0xde, 0x9f, 0xa4, 0x13, 0x93, 0x51, 0x16, 0x59, 0x64,
	0xd0, 0x6f, 0x40, 0x34, 0x71, 0x33, 0xe9, 0xe8, 0x42, 0x10, 0x2f, 0xda, 0x04, 0x69, 0x73, 0xdb,
	0xf8, 0xf9, 0xa5, 0x15, 0xc1, 0xf1, 0x77, 0x72, 0x72, 0x6e, 0x0b, 0x51, 0xa5, 0x44, 0xae, 0x48,
	0x1a, 0x99, 0x4a, 0x60, 0x1c, 0x9b, 0x16, 0x6b, 0x23, 0x64, 0xc7, 0x51, 0x2b, 0xd9, 0x69, 0x64,
	0x2b, 0x00, 0x1a, 0x69, 0xe2, 0xcd, 0xbd, 0x2c, 0x5e, 0xc7, 0xf9, 0x24, 0xee, 0x3f, 0xf8, 0x9b,
	0xc0, 0x66, 0x10, 0x20, 0x91, 0xa4, 0xe4, 0xd3, 0x99, 0x61, 0xbe, 0xb3, 0xab, 0x97, 0x9e, 0xb8,
	0x00, 0xf8, 0xa6, 0x21, 0x9d, 0x1e, 0x00, 0xa6, 0x0e, 0x4b, 0xe0, 0x4b, 0x23, 0xdd, 0x45, 0x8d,
	0xfd, 0x15, 0x3f, 0x8b, 0xf8, 0x6b, 0xb2, 0x25, 0xfc, 0x37, 0xa2, 0xc5, 0xf2, 0x8c, 0xba, 0x26,
	0xa1, 0x8c, 0xab, 0xfb, 0xd9, 0x0f, 0xff, 0xb3, 0x78, 0x3b, 0xd2, 0x74, 0x01, 0x81, 0x3b, 0x67,
	0x5b, 0x37, 0xd4, 0xba, 0xba, 0xa0, 0x7b, 0x71, 0xdf, 0x1a, 0x66, 0xf1, 0x7b, 0x8c, 0xaf, 0xc6,
	0xa8, 0xd2, 0xb6, 0x91, 0x4e, 0xa1, 0xfb, 0xfa, 0xe6, 0x11, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xba,
	0x71, 0x25, 0x07, 0x01, 0x00, 0x00,
}