// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/prestation/prestation.proto

package prestation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PrestationService service

type PrestationService interface {
	Rachat(ctx context.Context, in *Prestation, opts ...client.CallOption) (*Response, error)
	ValeurRachat(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type prestationService struct {
	c    client.Client
	name string
}

func NewPrestationService(name string, c client.Client) PrestationService {
	return &prestationService{
		c:    c,
		name: name,
	}
}

func (c *prestationService) Rachat(ctx context.Context, in *Prestation, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PrestationService.Rachat", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prestationService) ValeurRachat(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PrestationService.ValeurRachat", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PrestationService service

type PrestationServiceHandler interface {
	Rachat(context.Context, *Prestation, *Response) error
	ValeurRachat(context.Context, *Request, *Response) error
}

func RegisterPrestationServiceHandler(s server.Server, hdlr PrestationServiceHandler, opts ...server.HandlerOption) error {
	type prestationService interface {
		Rachat(ctx context.Context, in *Prestation, out *Response) error
		ValeurRachat(ctx context.Context, in *Request, out *Response) error
	}
	type PrestationService struct {
		prestationService
	}
	h := &prestationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PrestationService{h}, opts...))
}

type prestationServiceHandler struct {
	PrestationServiceHandler
}

func (h *prestationServiceHandler) Rachat(ctx context.Context, in *Prestation, out *Response) error {
	return h.PrestationServiceHandler.Rachat(ctx, in, out)
}

func (h *prestationServiceHandler) ValeurRachat(ctx context.Context, in *Request, out *Response) error {
	return h.PrestationServiceHandler.ValeurRachat(ctx, in, out)
}
