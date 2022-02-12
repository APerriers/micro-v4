package handler

import (
	"context"
	api "go-micro.dev/v4/api/proto"
	"io"
	"time"

	log "go-micro.dev/v4/logger"

	pb "redirect/proto"
)

type Redirect struct{}

func (r *Redirect) Url(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = int32(301)
	rsp.Header = map[string]*api.Pair{
		"Location": &api.Pair{
			Key:    "Location",
			Values: []string{"https://google.com"},
		},
	}
	return nil
}

func (e *Redirect) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Redirect.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	time.Sleep(10000)
	return nil
}

func (e *Redirect) ClientStream(ctx context.Context, stream pb.Redirect_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Redirect) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Redirect_ServerStreamStream) error {
	log.Infof("Received Redirect.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Redirect) BidiStream(ctx context.Context, stream pb.Redirect_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
