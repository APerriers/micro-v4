package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "job/proto"
)

type Job struct{}

func (e *Job) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Job.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
