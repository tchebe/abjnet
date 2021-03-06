// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/souscription/souscription.proto

package souscription

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

type Souscription struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nom                  string   `protobuf:"bytes,2,opt,name=nom,proto3" json:"nom,omitempty"`
	Prenom               string   `protobuf:"bytes,3,opt,name=prenom,proto3" json:"prenom,omitempty"`
	Dateofbirth          string   `protobuf:"bytes,4,opt,name=dateofbirth,proto3" json:"dateofbirth,omitempty"`
	Telephone            string   `protobuf:"bytes,5,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Abjcardno            uint32   `protobuf:"fixed32,6,opt,name=abjcardno,proto3" json:"abjcardno,omitempty"`
	Montant              uint64   `protobuf:"varint,7,opt,name=montant,proto3" json:"montant,omitempty"`
	Codeproduit          int32    `protobuf:"varint,8,opt,name=codeproduit,proto3" json:"codeproduit,omitempty"`
	Datepayment          string   `protobuf:"bytes,9,opt,name=datepayment,proto3" json:"datepayment,omitempty"`
	Echeance             string   `protobuf:"bytes,10,opt,name=echeance,proto3" json:"echeance,omitempty"`
	Beneficiaire         string   `protobuf:"bytes,11,opt,name=beneficiaire,proto3" json:"beneficiaire,omitempty"`
	Email                string   `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	Etattraitement       string   `protobuf:"bytes,13,opt,name=etattraitement,proto3" json:"etattraitement,omitempty"`
	CreatedAt            string   `protobuf:"bytes,14,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Firstassociate       string   `protobuf:"bytes,15,opt,name=firstassociate,proto3" json:"firstassociate,omitempty"`
	DateofbirthOne       string   `protobuf:"bytes,16,opt,name=dateofbirth_one,json=dateofbirthOne,proto3" json:"dateofbirth_one,omitempty"`
	DateofbirthTwo       string   `protobuf:"bytes,17,opt,name=dateofbirth_two,json=dateofbirthTwo,proto3" json:"dateofbirth_two,omitempty"`
	DateofbirthThree     string   `protobuf:"bytes,18,opt,name=dateofbirth_three,json=dateofbirthThree,proto3" json:"dateofbirth_three,omitempty"`
	LienparenteOne       string   `protobuf:"bytes,19,opt,name=lienparente_one,json=lienparenteOne,proto3" json:"lienparente_one,omitempty"`
	Secondtassociate     string   `protobuf:"bytes,20,opt,name=secondtassociate,proto3" json:"secondtassociate,omitempty"`
	LienparenteTwo       string   `protobuf:"bytes,21,opt,name=lienparente_two,json=lienparenteTwo,proto3" json:"lienparente_two,omitempty"`
	AbjCardNum           string   `protobuf:"bytes,23,opt,name=abj_card_num,json=abjCardNum,proto3" json:"abj_card_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Souscription) Reset()         { *m = Souscription{} }
func (m *Souscription) String() string { return proto.CompactTextString(m) }
func (*Souscription) ProtoMessage()    {}
func (*Souscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_7284ccaf0eb6016c, []int{0}
}

func (m *Souscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Souscription.Unmarshal(m, b)
}
func (m *Souscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Souscription.Marshal(b, m, deterministic)
}
func (m *Souscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Souscription.Merge(m, src)
}
func (m *Souscription) XXX_Size() int {
	return xxx_messageInfo_Souscription.Size(m)
}
func (m *Souscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Souscription.DiscardUnknown(m)
}

var xxx_messageInfo_Souscription proto.InternalMessageInfo

func (m *Souscription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Souscription) GetNom() string {
	if m != nil {
		return m.Nom
	}
	return ""
}

func (m *Souscription) GetPrenom() string {
	if m != nil {
		return m.Prenom
	}
	return ""
}

func (m *Souscription) GetDateofbirth() string {
	if m != nil {
		return m.Dateofbirth
	}
	return ""
}

func (m *Souscription) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *Souscription) GetAbjcardno() uint32 {
	if m != nil {
		return m.Abjcardno
	}
	return 0
}

func (m *Souscription) GetMontant() uint64 {
	if m != nil {
		return m.Montant
	}
	return 0
}

func (m *Souscription) GetCodeproduit() int32 {
	if m != nil {
		return m.Codeproduit
	}
	return 0
}

func (m *Souscription) GetDatepayment() string {
	if m != nil {
		return m.Datepayment
	}
	return ""
}

func (m *Souscription) GetEcheance() string {
	if m != nil {
		return m.Echeance
	}
	return ""
}

func (m *Souscription) GetBeneficiaire() string {
	if m != nil {
		return m.Beneficiaire
	}
	return ""
}

func (m *Souscription) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Souscription) GetEtattraitement() string {
	if m != nil {
		return m.Etattraitement
	}
	return ""
}

func (m *Souscription) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Souscription) GetFirstassociate() string {
	if m != nil {
		return m.Firstassociate
	}
	return ""
}

func (m *Souscription) GetDateofbirthOne() string {
	if m != nil {
		return m.DateofbirthOne
	}
	return ""
}

func (m *Souscription) GetDateofbirthTwo() string {
	if m != nil {
		return m.DateofbirthTwo
	}
	return ""
}

func (m *Souscription) GetDateofbirthThree() string {
	if m != nil {
		return m.DateofbirthThree
	}
	return ""
}

func (m *Souscription) GetLienparenteOne() string {
	if m != nil {
		return m.LienparenteOne
	}
	return ""
}

func (m *Souscription) GetSecondtassociate() string {
	if m != nil {
		return m.Secondtassociate
	}
	return ""
}

func (m *Souscription) GetLienparenteTwo() string {
	if m != nil {
		return m.LienparenteTwo
	}
	return ""
}

func (m *Souscription) GetAbjCardNum() string {
	if m != nil {
		return m.AbjCardNum
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7284ccaf0eb6016c, []int{1}
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

type Response struct {
	Done                 bool            `protobuf:"varint,1,opt,name=done,proto3" json:"done,omitempty"`
	Description          string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Souscription         *Souscription   `protobuf:"bytes,3,opt,name=souscription,proto3" json:"souscription,omitempty"`
	Souscriptions        []*Souscription `protobuf:"bytes,4,rep,name=souscriptions,proto3" json:"souscriptions,omitempty"`
	Errors               []*Error        `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7284ccaf0eb6016c, []int{2}
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

func (m *Response) GetSouscription() *Souscription {
	if m != nil {
		return m.Souscription
	}
	return nil
}

func (m *Response) GetSouscriptions() []*Souscription {
	if m != nil {
		return m.Souscriptions
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
	return fileDescriptor_7284ccaf0eb6016c, []int{3}
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
	proto.RegisterType((*Souscription)(nil), "souscription.Souscription")
	proto.RegisterType((*Request)(nil), "souscription.Request")
	proto.RegisterType((*Response)(nil), "souscription.Response")
	proto.RegisterType((*Error)(nil), "souscription.Error")
}

func init() {
	proto.RegisterFile("proto/souscription/souscription.proto", fileDescriptor_7284ccaf0eb6016c)
}

var fileDescriptor_7284ccaf0eb6016c = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x6b, 0xdb, 0x3c,
	0x14, 0xc7, 0x1f, 0xb7, 0x71, 0x12, 0x9f, 0xa6, 0x6d, 0xaa, 0xbe, 0x89, 0xf0, 0x5c, 0x18, 0xc3,
	0xb6, 0xb0, 0x42, 0x07, 0xdd, 0xd5, 0x60, 0x1b, 0x0b, 0xdb, 0x18, 0xec, 0x62, 0x03, 0x67, 0xf7,
	0x41, 0xb6, 0x4f, 0x88, 0x82, 0x2d, 0x79, 0xb2, 0xbc, 0xb2, 0x8f, 0xb9, 0x0f, 0x32, 0xf6, 0x15,
	0x86, 0xe4, 0xbc, 0xc8, 0x09, 0xb4, 0xb0, 0x3b, 0x9f, 0xdf, 0xff, 0x7f, 0x8e, 0xce, 0x51, 0x74,
	0x02, 0x4f, 0x4a, 0x25, 0xb5, 0x7c, 0x51, 0xc9, 0xba, 0x4a, 0x15, 0x2f, 0x35, 0x97, 0xa2, 0x15,
	0xdc, 0x5a, 0x9d, 0x0c, 0x5c, 0x16, 0xfd, 0xf6, 0x61, 0x30, 0x75, 0x00, 0x39, 0x81, 0x03, 0x9e,
	0x51, 0x2f, 0xf4, 0xc6, 0x41, 0x7c, 0xc0, 0x33, 0x32, 0x84, 0x43, 0x21, 0x0b, 0x7a, 0x60, 0x81,
	0xf9, 0x24, 0x57, 0xd0, 0x2d, 0x15, 0x1a, 0x78, 0x68, 0xe1, 0x2a, 0x22, 0x21, 0x1c, 0x65, 0x4c,
	0xa3, 0x9c, 0x27, 0x5c, 0xe9, 0x05, 0xed, 0x58, 0xd1, 0x45, 0xe4, 0x7f, 0x08, 0x34, 0xe6, 0x58,
	0x2e, 0xa4, 0x40, 0xea, 0x5b, 0x7d, 0x0b, 0x8c, 0xca, 0x92, 0x65, 0xca, 0x54, 0x26, 0x24, 0xed,
	0x86, 0xde, 0xb8, 0x17, 0x6f, 0x01, 0xa1, 0xd0, 0x2b, 0xa4, 0xd0, 0x4c, 0x68, 0xda, 0x0b, 0xbd,
	0x71, 0x27, 0x5e, 0x87, 0xe6, 0xdc, 0x54, 0x66, 0x58, 0x2a, 0x99, 0xd5, 0x5c, 0xd3, 0x7e, 0xe8,
	0x8d, 0xfd, 0xd8, 0x45, 0xeb, 0xce, 0x4a, 0xf6, 0xb3, 0x40, 0xa1, 0x69, 0xb0, 0xed, 0x6c, 0x85,
	0xc8, 0x08, 0xfa, 0x98, 0x2e, 0x90, 0x89, 0x14, 0x29, 0x58, 0x79, 0x13, 0x93, 0x08, 0x06, 0x09,
	0x0a, 0x9c, 0xf3, 0x94, 0x33, 0xae, 0x90, 0x1e, 0x59, 0xbd, 0xc5, 0xc8, 0x05, 0xf8, 0x58, 0x30,
	0x9e, 0xd3, 0x81, 0x15, 0x9b, 0x80, 0x3c, 0x85, 0x13, 0xd4, 0x4c, 0x6b, 0xc5, 0xb8, 0x46, 0x7b,
	0xf4, 0xb1, 0x95, 0x77, 0xa8, 0x99, 0x3c, 0x55, 0xc8, 0x34, 0x66, 0x13, 0x4d, 0x4f, 0x9a, 0x7b,
	0xd9, 0x00, 0x53, 0x65, 0xce, 0x55, 0xa5, 0x59, 0x55, 0xc9, 0x94, 0x33, 0x8d, 0xf4, 0xb4, 0xa9,
	0xd2, 0xa6, 0xe4, 0x19, 0x9c, 0x3a, 0x97, 0x3d, 0x33, 0x77, 0x3c, 0x6c, 0x8c, 0x0e, 0xfe, 0x2a,
	0xf6, 0x8c, 0xfa, 0x5e, 0xd2, 0xb3, 0x3d, 0xe3, 0xb7, 0x7b, 0x49, 0x6e, 0xe0, 0xac, 0x65, 0x5c,
	0x28, 0x44, 0x4a, 0xac, 0x75, 0xe8, 0x5a, 0x0d, 0x37, 0x55, 0x73, 0x8e, 0xa2, 0x64, 0x0a, 0x85,
	0x46, 0x7b, 0xfc, 0x79, 0x53, 0xd5, 0xc1, 0xe6, 0xf8, 0xe7, 0x30, 0xac, 0x30, 0x95, 0x22, 0x73,
	0x26, 0xba, 0x68, 0x8a, 0xee, 0xf2, 0xdd, 0xa2, 0xa6, 0xd5, 0xcb, 0xbd, 0xa2, 0xa6, 0xd5, 0x10,
	0x06, 0x2c, 0x59, 0xce, 0xcc, 0x63, 0x99, 0x89, 0xba, 0xa0, 0xd7, 0xd6, 0x05, 0x2c, 0x59, 0xbe,
	0x67, 0x2a, 0xfb, 0x52, 0x17, 0x9f, 0x3b, 0xfd, 0xab, 0xe1, 0x75, 0x14, 0x40, 0x2f, 0xc6, 0xef,
	0x35, 0x56, 0x3a, 0xfa, 0xe3, 0x41, 0x3f, 0xc6, 0xaa, 0x94, 0xa2, 0x42, 0x42, 0xa0, 0x93, 0x99,
	0x96, 0xcd, 0xc3, 0xef, 0xc7, 0xf6, 0xdb, 0x3e, 0x1b, 0xdc, 0x6c, 0xc6, 0x6a, 0x05, 0x5c, 0x44,
	0xde, 0x42, 0x6b, 0x9b, 0xec, 0x42, 0x1c, 0xdd, 0x8d, 0x6e, 0x5b, 0x6b, 0xe7, 0xae, 0x57, 0xdc,
	0xf2, 0x93, 0x77, 0x70, 0xec, 0xc6, 0x15, 0xed, 0x84, 0x87, 0x8f, 0x14, 0x68, 0x27, 0x90, 0x1b,
	0xe8, 0xa2, 0x52, 0x52, 0x55, 0xd4, 0xb7, 0xa9, 0xe7, 0xed, 0xd4, 0x8f, 0x46, 0x8b, 0x57, 0x96,
	0xe8, 0x0d, 0xf8, 0x16, 0x98, 0x69, 0xcd, 0x7e, 0xd8, 0x69, 0xfd, 0xd8, 0x7e, 0x3f, 0x3e, 0xed,
	0xdd, 0x2f, 0x0f, 0xce, 0xdd, 0x5e, 0xa6, 0xa8, 0x7e, 0xf0, 0x14, 0xc9, 0x04, 0x82, 0x69, 0x9d,
	0x18, 0x9c, 0x20, 0x79, 0xa0, 0xf7, 0xd1, 0x55, 0x5b, 0x5b, 0x5f, 0x7e, 0xf4, 0x1f, 0x79, 0x05,
	0xdd, 0x4f, 0xa8, 0x27, 0x79, 0x4e, 0x2e, 0x77, 0x3d, 0xf6, 0xc7, 0x7a, 0x20, 0xf5, 0x35, 0x04,
	0x1f, 0x30, 0x47, 0x8d, 0xff, 0x92, 0x9d, 0x74, 0xed, 0x9f, 0xe2, 0xcb, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x42, 0xac, 0xfd, 0x01, 0x3d, 0x05, 0x00, 0x00,
}
