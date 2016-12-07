package grpc

import (
	"im/grpc/session"
	"log"
)

type Sesion struct{}

func (s *Sesion) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionReply, error) {

	log.Println("creteSession")

	return nil, nil
}
