// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/srv.proto

package srv

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Srv service

func NewSrvEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Srv service

type SrvService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Srv_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Srv_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Srv_BidiStreamService, error)
}

type srvService struct {
	c    client.Client
	name string
}

func NewSrvService(name string, c client.Client) SrvService {
	return &srvService{
		c:    c,
		name: name,
	}
}

func (c *srvService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Srv.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvService) ClientStream(ctx context.Context, opts ...client.CallOption) (Srv_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Srv.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &srvServiceClientStream{stream}, nil
}

type Srv_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ClientStreamRequest) error
}

type srvServiceClientStream struct {
	stream client.Stream
}

func (x *srvServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *srvServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *srvServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *srvServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *srvServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *srvService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Srv_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Srv.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &srvServiceServerStream{stream}, nil
}

type Srv_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type srvServiceServerStream struct {
	stream client.Stream
}

func (x *srvServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *srvServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *srvServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *srvServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *srvServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *srvService) BidiStream(ctx context.Context, opts ...client.CallOption) (Srv_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Srv.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &srvServiceBidiStream{stream}, nil
}

type Srv_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type srvServiceBidiStream struct {
	stream client.Stream
}

func (x *srvServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *srvServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *srvServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *srvServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *srvServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *srvServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Srv service

type SrvHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Srv_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Srv_ServerStreamStream) error
	BidiStream(context.Context, Srv_BidiStreamStream) error
}

func RegisterSrvHandler(s server.Server, hdlr SrvHandler, opts ...server.HandlerOption) error {
	type srv interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type Srv struct {
		srv
	}
	h := &srvHandler{hdlr}
	return s.Handle(s.NewHandler(&Srv{h}, opts...))
}

type srvHandler struct {
	SrvHandler
}

func (h *srvHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.SrvHandler.Call(ctx, in, out)
}

func (h *srvHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.SrvHandler.ClientStream(ctx, &srvClientStreamStream{stream})
}

type Srv_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type srvClientStreamStream struct {
	stream server.Stream
}

func (x *srvClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *srvClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *srvClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *srvClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *srvClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *srvHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SrvHandler.ServerStream(ctx, m, &srvServerStreamStream{stream})
}

type Srv_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type srvServerStreamStream struct {
	stream server.Stream
}

func (x *srvServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *srvServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *srvServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *srvServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *srvServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *srvHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.SrvHandler.BidiStream(ctx, &srvBidiStreamStream{stream})
}

type Srv_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type srvBidiStreamStream struct {
	stream server.Stream
}

func (x *srvBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *srvBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *srvBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *srvBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *srvBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *srvBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
