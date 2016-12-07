package grpc

import (
	context "golang.org/x/net/context"
	"im/grpc/message"
	"log"
)

type Message struct{}

func (s *Message) CreateMessage(context.Context, *message.CreateMessageRequest) (*message.CreateMessageReply, error) {

	log.Println("CreateMessage")

	return nil, nil
}
