// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/apiuser.proto

package apiuser

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

// Api Endpoints for Apiuser service

func NewApiuserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Apiuser service

type ApiuserService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Apiuser_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Apiuser_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Apiuser_BidiStreamService, error)
}

type apiuserService struct {
	c    client.Client
	name string
}

func NewApiuserService(name string, c client.Client) ApiuserService {
	return &apiuserService{
		c:    c,
		name: name,
	}
}

func (c *apiuserService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Apiuser.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiuserService) ClientStream(ctx context.Context, opts ...client.CallOption) (Apiuser_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Apiuser.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &apiuserServiceClientStream{stream}, nil
}

type Apiuser_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ClientStreamRequest) error
}

type apiuserServiceClientStream struct {
	stream client.Stream
}

func (x *apiuserServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *apiuserServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *apiuserServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *apiuserServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *apiuserServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *apiuserService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Apiuser_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Apiuser.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &apiuserServiceServerStream{stream}, nil
}

type Apiuser_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type apiuserServiceServerStream struct {
	stream client.Stream
}

func (x *apiuserServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *apiuserServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *apiuserServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *apiuserServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *apiuserServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apiuserService) BidiStream(ctx context.Context, opts ...client.CallOption) (Apiuser_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Apiuser.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &apiuserServiceBidiStream{stream}, nil
}

type Apiuser_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type apiuserServiceBidiStream struct {
	stream client.Stream
}

func (x *apiuserServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *apiuserServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *apiuserServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *apiuserServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *apiuserServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *apiuserServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Apiuser service

type ApiuserHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Apiuser_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Apiuser_ServerStreamStream) error
	BidiStream(context.Context, Apiuser_BidiStreamStream) error
}

func RegisterApiuserHandler(s server.Server, hdlr ApiuserHandler, opts ...server.HandlerOption) error {
	type apiuser interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type Apiuser struct {
		apiuser
	}
	h := &apiuserHandler{hdlr}
	return s.Handle(s.NewHandler(&Apiuser{h}, opts...))
}

type apiuserHandler struct {
	ApiuserHandler
}

func (h *apiuserHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.ApiuserHandler.Call(ctx, in, out)
}

func (h *apiuserHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.ApiuserHandler.ClientStream(ctx, &apiuserClientStreamStream{stream})
}

type Apiuser_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type apiuserClientStreamStream struct {
	stream server.Stream
}

func (x *apiuserClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *apiuserClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *apiuserClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *apiuserClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *apiuserClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *apiuserHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ApiuserHandler.ServerStream(ctx, m, &apiuserServerStreamStream{stream})
}

type Apiuser_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type apiuserServerStreamStream struct {
	stream server.Stream
}

func (x *apiuserServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *apiuserServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *apiuserServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *apiuserServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *apiuserServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *apiuserHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.ApiuserHandler.BidiStream(ctx, &apiuserBidiStreamStream{stream})
}

type Apiuser_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type apiuserBidiStreamStream struct {
	stream server.Stream
}

func (x *apiuserBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *apiuserBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *apiuserBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *apiuserBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *apiuserBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *apiuserBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
