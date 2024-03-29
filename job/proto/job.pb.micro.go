// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/job.proto

package job

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

// Api Endpoints for Job service

func NewJobEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Job service

type JobService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type jobService struct {
	c    client.Client
	name string
}

func NewJobService(name string, c client.Client) JobService {
	return &jobService{
		c:    c,
		name: name,
	}
}

func (c *jobService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Job.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Job service

type JobHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
}

func RegisterJobHandler(s server.Server, hdlr JobHandler, opts ...server.HandlerOption) error {
	type job interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
	}
	type Job struct {
		job
	}
	h := &jobHandler{hdlr}
	return s.Handle(s.NewHandler(&Job{h}, opts...))
}

type jobHandler struct {
	JobHandler
}

func (h *jobHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.JobHandler.Call(ctx, in, out)
}
