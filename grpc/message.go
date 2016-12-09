package grpc

import (
	context "golang.org/x/net/context"
	//"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	"log"
)

type Message struct {
}

func (s *Message) CreateMessage(context.Context, *grpcPb.CreateMessageRequest) (*grpcPb.CreateMessageReply, error) {

	log.Println("CreateMessage")

	return nil, nil
}
