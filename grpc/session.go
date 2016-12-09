package grpc

import (
	context "golang.org/x/net/context"
	//"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	"log"
)

type Sesion struct {
}

func (s *Sesion) CreateSession(context.Context, *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionReply, error) {

	log.Println("creteSession")

	return nil, nil
}
