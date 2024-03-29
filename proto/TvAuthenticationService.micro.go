// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: TvAuthenticationService.proto

package TvAuthenticationService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for TvAuthenticationService service

type TvAuthenticationService interface {
	CloudwalkerAuthWall(ctx context.Context, in *TargetTv, opts ...client.CallOption) (*IsAuthorised, error)
}

type tvAuthenticationService struct {
	c    client.Client
	name string
}

func NewTvAuthenticationService(name string, c client.Client) TvAuthenticationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "TvAuthenticationService"
	}
	return &tvAuthenticationService{
		c:    c,
		name: name,
	}
}

func (c *tvAuthenticationService) CloudwalkerAuthWall(ctx context.Context, in *TargetTv, opts ...client.CallOption) (*IsAuthorised, error) {
	req := c.c.NewRequest(c.name, "TvAuthenticationService.CloudwalkerAuthWall", in)
	out := new(IsAuthorised)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TvAuthenticationService service

type TvAuthenticationServiceHandler interface {
	CloudwalkerAuthWall(context.Context, *TargetTv, *IsAuthorised) error
}

func RegisterTvAuthenticationServiceHandler(s server.Server, hdlr TvAuthenticationServiceHandler, opts ...server.HandlerOption) error {
	type tvAuthenticationService interface {
		CloudwalkerAuthWall(ctx context.Context, in *TargetTv, out *IsAuthorised) error
	}
	type TvAuthenticationService struct {
		tvAuthenticationService
	}
	h := &tvAuthenticationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TvAuthenticationService{h}, opts...))
}

type tvAuthenticationServiceHandler struct {
	TvAuthenticationServiceHandler
}

func (h *tvAuthenticationServiceHandler) CloudwalkerAuthWall(ctx context.Context, in *TargetTv, out *IsAuthorised) error {
	return h.TvAuthenticationServiceHandler.CloudwalkerAuthWall(ctx, in, out)
}
