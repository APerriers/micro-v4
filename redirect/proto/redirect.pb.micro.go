// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/redirect.proto

package redirect

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

// Api Endpoints for Redirect service

func NewRedirectEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Redirect service

type RedirectService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Redirect_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Redirect_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Redirect_BidiStreamService, error)
}

type redirectService struct {
	c    client.Client
	name string
}

func NewRedirectService(name string, c client.Client) RedirectService {
	return &redirectService{
		c:    c,
		name: name,
	}
}

func (c *redirectService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Redirect.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redirectService) ClientStream(ctx context.Context, opts ...client.CallOption) (Redirect_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Redirect.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &redirectServiceClientStream{stream}, nil
}

type Redirect_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ClientStreamRequest) error
}

type redirectServiceClientStream struct {
	stream client.Stream
}

func (x *redirectServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *redirectServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *redirectServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *redirectServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *redirectServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *redirectService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Redirect_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Redirect.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &redirectServiceServerStream{stream}, nil
}

type Redirect_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type redirectServiceServerStream struct {
	stream client.Stream
}

func (x *redirectServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *redirectServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *redirectServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *redirectServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *redirectServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *redirectService) BidiStream(ctx context.Context, opts ...client.CallOption) (Redirect_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Redirect.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &redirectServiceBidiStream{stream}, nil
}

type Redirect_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type redirectServiceBidiStream struct {
	stream client.Stream
}

func (x *redirectServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *redirectServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *redirectServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *redirectServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *redirectServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *redirectServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Redirect service

type RedirectHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Redirect_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Redirect_ServerStreamStream) error
	BidiStream(context.Context, Redirect_BidiStreamStream) error
}

func RegisterRedirectHandler(s server.Server, hdlr RedirectHandler, opts ...server.HandlerOption) error {
	type redirect interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type Redirect struct {
		redirect
	}
	h := &redirectHandler{hdlr}
	return s.Handle(s.NewHandler(&Redirect{h}, opts...))
}

type redirectHandler struct {
	RedirectHandler
}

func (h *redirectHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.RedirectHandler.Call(ctx, in, out)
}

func (h *redirectHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.RedirectHandler.ClientStream(ctx, &redirectClientStreamStream{stream})
}

type Redirect_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type redirectClientStreamStream struct {
	stream server.Stream
}

func (x *redirectClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *redirectClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *redirectClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *redirectClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *redirectClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *redirectHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RedirectHandler.ServerStream(ctx, m, &redirectServerStreamStream{stream})
}

type Redirect_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type redirectServerStreamStream struct {
	stream server.Stream
}

func (x *redirectServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *redirectServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *redirectServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *redirectServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *redirectServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *redirectHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.RedirectHandler.BidiStream(ctx, &redirectBidiStreamStream{stream})
}

type Redirect_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type redirectBidiStreamStream struct {
	stream server.Stream
}

func (x *redirectBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *redirectBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *redirectBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *redirectBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *redirectBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *redirectBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
