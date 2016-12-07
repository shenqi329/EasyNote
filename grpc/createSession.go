package grpc

import (
	context "golang.org/x/net/context"
	"im/grpc/session"
	"log"
)

type Sesion struct{}

func (s *Sesion) CreateSession(context.Context, *session.CreateSessionRequest) (*session.CreateSessionReply, error) {

	log.Println("creteSession")

	return nil, nil
}
