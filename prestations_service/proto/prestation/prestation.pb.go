// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/prestation/prestation.proto

package prestation

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

type Prestation struct {
	Transacno            int64    `protobuf:"varint,1,opt,name=transacno,proto3" json:"transacno,omitempty"`
	Nomclient            string   `protobuf:"bytes,2,opt,name=nomclient,proto3" json:"nomclient,omitempty"`
	Prenomclient         string   `protobuf:"bytes,3,opt,name=prenomclient,proto3" json:"prenomclient,omitempty"`
	Telephone            string   `protobuf:"bytes,4,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Datedamandeuser      string   `protobuf:"bytes,5,opt,name=datedamandeuser,proto3" json:"datedamandeuser,omitempty"`
	Conventionno         int32    `protobuf:"varint,6,opt,name=conventionno,proto3" json:"conventionno,omitempty"`
	Policeno             int32    `protobuf:"varint,7,opt,name=policeno,proto3" json:"policeno,omitempty"`
	Montantdemande       int64    `protobuf:"varint,8,opt,name=montantdemande,proto3" json:"montantdemande,omitempty"`
	Montantrestant       int64    `protobuf:"varint,9,opt,name=montantrestant,proto3" json:"montantrestant,omitempty"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prestation) Reset()         { *m = Prestation{} }
func (m *Prestation) String() string { return proto.CompactTextString(m) }
func (*Prestation) ProtoMessage()    {}
func (*Prestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4002c97688aaf2, []int{0}
}

func (m *Prestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prestation.Unmarshal(m, b)
}
func (m *Prestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prestation.Marshal(b, m, deterministic)
}
func (m *Prestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prestation.Merge(m, src)
}
func (m *Prestation) XXX_Size() int {
	return xxx_messageInfo_Prestation.Size(m)
}
func (m *Prestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Prestation.DiscardUnknown(m)
}

var xxx_messageInfo_Prestation proto.InternalMessageInfo

func (m *Prestation) GetTransacno() int64 {
	if m != nil {
		return m.Transacno
	}
	return 0
}

func (m *Prestation) GetNomclient() string {
	if m != nil {
		return m.Nomclient
	}
	return ""
}

func (m *Prestation) GetPrenomclient() string {
	if m != nil {
		return m.Prenomclient
	}
	return ""
}

func (m *Prestation) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *Prestation) GetDatedamandeuser() string {
	if m != nil {
		return m.Datedamandeuser
	}
	return ""
}

func (m *Prestation) GetConventionno() int32 {
	if m != nil {
		return m.Conventionno
	}
	return 0
}

func (m *Prestation) GetPoliceno() int32 {
	if m != nil {
		return m.Policeno
	}
	return 0
}

func (m *Prestation) GetMontantdemande() int64 {
	if m != nil {
		return m.Montantdemande
	}
	return 0
}

func (m *Prestation) GetMontantrestant() int64 {
	if m != nil {
		return m.Montantrestant
	}
	return 0
}

func (m *Prestation) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type Request struct {
	Police               int32    `protobuf:"varint,1,opt,name=police,proto3" json:"police,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4002c97688aaf2, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPolice() int32 {
	if m != nil {
		return m.Police
	}
	return 0
}

type Response struct {
	Done                 bool        `protobuf:"varint,1,opt,name=done,proto3" json:"done,omitempty"`
	Description          string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Payment              *Prestation `protobuf:"bytes,3,opt,name=payment,proto3" json:"payment,omitempty"`
	Errors               []*Error    `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4002c97688aaf2, []int{2}
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

func (m *Response) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Response) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Response) GetPayment() *Prestation {
	if m != nil {
		return m.Payment
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4002c97688aaf2, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Prestation)(nil), "prestation.Prestation")
	proto.RegisterType((*Request)(nil), "prestation.Request")
	proto.RegisterType((*Response)(nil), "prestation.Response")
	proto.RegisterType((*Error)(nil), "prestation.Error")
}

func init() {
	proto.RegisterFile("proto/prestation/prestation.proto", fileDescriptor_9d4002c97688aaf2)
}

var fileDescriptor_9d4002c97688aaf2 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcd, 0xea, 0x13, 0x31,
	0x14, 0xc5, 0x9d, 0xb6, 0x33, 0x9d, 0xde, 0x16, 0xa5, 0x51, 0x4a, 0x28, 0x2e, 0xa6, 0xb3, 0x90,
	0x71, 0x53, 0xa5, 0x6e, 0x04, 0x71, 0xe1, 0xc2, 0xbd, 0x44, 0x70, 0x1f, 0x33, 0x17, 0x3a, 0x30,
	0x93, 0xc4, 0x24, 0x2d, 0xf8, 0x0a, 0xbe, 0x83, 0x2f, 0xe7, 0x93, 0x48, 0xd2, 0xf9, 0x2c, 0x94,
	0xff, 0x2e, 0xf9, 0xdd, 0x73, 0xef, 0x9c, 0x7b, 0xc2, 0xc0, 0x41, 0x1b, 0xe5, 0xd4, 0x3b, 0x6d,
	0xd0, 0x3a, 0xee, 0x2a, 0x25, 0x47, 0xc7, 0x63, 0xa8, 0x11, 0x18, 0x48, 0xfe, 0x6f, 0x06, 0xf0,
	0xad, 0xbf, 0x92, 0xd7, 0xb0, 0x72, 0x86, 0x4b, 0xcb, 0x85, 0x54, 0x34, 0xca, 0xa2, 0x62, 0xce,
	0x06, 0xe0, 0xab, 0x52, 0x35, 0xa2, 0xae, 0x50, 0x3a, 0x3a, 0xcb, 0xa2, 0x62, 0xc5, 0x06, 0x40,
	0x72, 0xd8, 0x68, 0x83, 0x83, 0x60, 0x1e, 0x04, 0x13, 0x16, 0xe6, 0x63, 0x8d, 0xfa, 0xac, 0x24,
	0xd2, 0xc5, 0x6d, 0x42, 0x0f, 0x48, 0x01, 0x2f, 0x4a, 0xee, 0xb0, 0xe4, 0x0d, 0x97, 0x25, 0x5e,
	0x2c, 0x1a, 0x1a, 0x07, 0xcd, 0x3d, 0xf6, 0xdf, 0x12, 0x4a, 0x5e, 0x51, 0x7a, 0xd7, 0x52, 0xd1,
	0x24, 0x8b, 0x8a, 0x98, 0x4d, 0x18, 0xd9, 0x43, 0xaa, 0x55, 0x5d, 0x09, 0x94, 0x8a, 0x2e, 0x43,
	0xbd, 0xbf, 0x93, 0x37, 0xf0, 0xbc, 0x51, 0xd2, 0x71, 0xe9, 0x4a, 0x0c, 0x53, 0x69, 0x1a, 0x96,
	0xbd, 0xa3, 0x23, 0x5d, 0xc8, 0x48, 0x3a, 0xba, 0x9a, 0xe8, 0x5a, 0xea, 0xf7, 0x12, 0x06, 0xbd,
	0xc9, 0x2f, 0x8e, 0xc2, 0x6d, 0xaf, 0x1e, 0xe4, 0x07, 0x58, 0x32, 0xfc, 0x75, 0x41, 0xeb, 0xc8,
	0x0e, 0x92, 0x9b, 0x89, 0x90, 0x6e, 0xcc, 0xda, 0x5b, 0xfe, 0x37, 0x82, 0x94, 0xa1, 0xd5, 0x4a,
	0x5a, 0x24, 0x04, 0x16, 0xa5, 0x0f, 0xc8, 0x4b, 0x52, 0x16, 0xce, 0x24, 0x83, 0x75, 0x89, 0x56,
	0x98, 0x4a, 0xfb, 0xf5, 0xda, 0xf4, 0xc7, 0x88, 0xbc, 0x87, 0xa5, 0xe6, 0xbf, 0x9b, 0x2e, 0xfa,
	0xf5, 0x69, 0x77, 0x1c, 0x3d, 0xfd, 0xf0, 0xc8, 0xac, 0x93, 0x91, 0xb7, 0x90, 0xa0, 0x31, 0xca,
	0x58, 0xba, 0xc8, 0xe6, 0xc5, 0xfa, 0xb4, 0x1d, 0x37, 0x7c, 0xf5, 0x15, 0xd6, 0x0a, 0xf2, 0xcf,
	0x10, 0x07, 0xe0, 0xbd, 0x09, 0x55, 0x76, 0xf6, 0xc3, 0xf9, 0x69, 0x6f, 0xa7, 0x3f, 0x11, 0x6c,
	0x07, 0x07, 0xdf, 0xd1, 0x5c, 0x2b, 0x81, 0xe4, 0x23, 0x24, 0x8c, 0x8b, 0x33, 0x77, 0xe4, 0x81,
	0xd5, 0xfd, 0xab, 0x31, 0xef, 0xf2, 0xc9, 0x9f, 0x91, 0x4f, 0xb0, 0xf9, 0xc1, 0x6b, 0xbc, 0x98,
	0xb6, 0xff, 0xe5, 0x54, 0x17, 0xb2, 0x7e, 0xd4, 0xfc, 0x33, 0x09, 0xbf, 0xc1, 0x87, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0x19, 0x5e, 0xd1, 0x2b, 0x03, 0x00, 0x00,
}
