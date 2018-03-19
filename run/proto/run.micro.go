// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/micro/run/proto/run.proto

/*
Package go_micro_run is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/micro/run/proto/run.proto

It has these top-level messages:
	Source
	Binary
	Process
	FetchRequest
	FetchResponse
	BuildRequest
	BuildResponse
	ExecRequest
	ExecResponse
	KillRequest
	KillResponse
	WaitRequest
	WaitResponse
	RunRequest
	RunResponse
	StopRequest
	StopResponse
	StatusRequest
	StatusResponse
*/
package go_micro_run

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Runtime service

type RuntimeClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...client.CallOption) (*FetchResponse, error)
	Build(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error)
	Kill(ctx context.Context, in *KillRequest, opts ...client.CallOption) (*KillResponse, error)
	Wait(ctx context.Context, in *WaitRequest, opts ...client.CallOption) (Runtime_WaitClient, error)
}

type runtimeClient struct {
	c           client.Client
	serviceName string
}

func NewRuntimeClient(serviceName string, c client.Client) RuntimeClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.run"
	}
	return &runtimeClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *runtimeClient) Fetch(ctx context.Context, in *FetchRequest, opts ...client.CallOption) (*FetchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Fetch", in)
	out := new(FetchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Build(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Build", in)
	out := new(BuildResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Exec", in)
	out := new(ExecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Kill(ctx context.Context, in *KillRequest, opts ...client.CallOption) (*KillResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Kill", in)
	out := new(KillResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Wait(ctx context.Context, in *WaitRequest, opts ...client.CallOption) (Runtime_WaitClient, error) {
	req := c.c.NewRequest(c.serviceName, "Runtime.Wait", &WaitRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &runtimeWaitClient{stream}, nil
}

type Runtime_WaitClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*WaitResponse, error)
}

type runtimeWaitClient struct {
	stream client.Streamer
}

func (x *runtimeWaitClient) Close() error {
	return x.stream.Close()
}

func (x *runtimeWaitClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeWaitClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeWaitClient) Recv() (*WaitResponse, error) {
	m := new(WaitResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Runtime service

type RuntimeHandler interface {
	Fetch(context.Context, *FetchRequest, *FetchResponse) error
	Build(context.Context, *BuildRequest, *BuildResponse) error
	Exec(context.Context, *ExecRequest, *ExecResponse) error
	Kill(context.Context, *KillRequest, *KillResponse) error
	Wait(context.Context, *WaitRequest, Runtime_WaitStream) error
}

func RegisterRuntimeHandler(s server.Server, hdlr RuntimeHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Runtime{hdlr}, opts...))
}

type Runtime struct {
	RuntimeHandler
}

func (h *Runtime) Fetch(ctx context.Context, in *FetchRequest, out *FetchResponse) error {
	return h.RuntimeHandler.Fetch(ctx, in, out)
}

func (h *Runtime) Build(ctx context.Context, in *BuildRequest, out *BuildResponse) error {
	return h.RuntimeHandler.Build(ctx, in, out)
}

func (h *Runtime) Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error {
	return h.RuntimeHandler.Exec(ctx, in, out)
}

func (h *Runtime) Kill(ctx context.Context, in *KillRequest, out *KillResponse) error {
	return h.RuntimeHandler.Kill(ctx, in, out)
}

func (h *Runtime) Wait(ctx context.Context, stream server.Streamer) error {
	m := new(WaitRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RuntimeHandler.Wait(ctx, m, &runtimeWaitStream{stream})
}

type Runtime_WaitStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*WaitResponse) error
}

type runtimeWaitStream struct {
	stream server.Streamer
}

func (x *runtimeWaitStream) Close() error {
	return x.stream.Close()
}

func (x *runtimeWaitStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeWaitStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeWaitStream) Send(m *WaitResponse) error {
	return x.stream.Send(m)
}

// Client API for Service service

type ServiceClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...client.CallOption) (*StopResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error)
}

type serviceClient struct {
	c           client.Client
	serviceName string
}

func NewServiceClient(serviceName string, c client.Client) ServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.run"
	}
	return &serviceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *serviceClient) Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Service.Run", in)
	out := new(RunResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Stop(ctx context.Context, in *StopRequest, opts ...client.CallOption) (*StopResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Service.Stop", in)
	out := new(StopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Service.Status", in)
	out := new(StatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	Run(context.Context, *RunRequest, *RunResponse) error
	Stop(context.Context, *StopRequest, *StopResponse) error
	Status(context.Context, *StatusRequest, *StatusResponse) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Service{hdlr}, opts...))
}

type Service struct {
	ServiceHandler
}

func (h *Service) Run(ctx context.Context, in *RunRequest, out *RunResponse) error {
	return h.ServiceHandler.Run(ctx, in, out)
}

func (h *Service) Stop(ctx context.Context, in *StopRequest, out *StopResponse) error {
	return h.ServiceHandler.Stop(ctx, in, out)
}

func (h *Service) Status(ctx context.Context, in *StatusRequest, out *StatusResponse) error {
	return h.ServiceHandler.Status(ctx, in, out)
}