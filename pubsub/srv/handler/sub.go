package handler

import (
	"context"
	proto "go-micro.dev/v4/api/proto"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/util/log"
)

// All methods of Sub will be executed when
// a message is received
type Sub struct{}

// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}
