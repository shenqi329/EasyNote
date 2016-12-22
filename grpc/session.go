package grpc

import (
	"easynote/util/key"
	"github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcPb "im/logicserver/grpc/pb"
	"log"
)

// type Sesion struct {
// }

// func (s *Sesion) CreateSession(context.Context, *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {

// 	log.Println("creteSession")
// 	return nil, nil
// }

// func (s *Sesion) DeleteUsers(context.Context, *grpcPb.DeleteSessionUsersRequest) (*grpcPb.Response, error) {

// 	log.Println("DeleteUsers")
// 	return nil, nil
// }

// func (s *Sesion) AddUsers(context.Context, *grpcPb.AddSessionUsersRequest) (*grpcPb.Response, error) {

// 	log.Println("AddUsers")
// 	return nil, nil
// }

func CreateSession(ctx context.Context, message proto.Message) (proto.Message, error) {

	log.Println("creteSession")
	rpcInfo := ctx.Value(key.RpcInfo).(*grpcPb.RpcInfo)
	request, ok := message.(*grpcPb.CreateSessionRequest)
	if !ok {
		return nil, nil
	}

	logicServerConn := ctx.Value(key.LogicServerConn).(*grpc.ClientConn)
	sessionClient := grpcPb.NewSessionClient(logicServerConn)

	request.RpcInfo = rpcInfo
	response, err := sessionClient.CreateSession(context.Background(), request)

	return response, err
}

func DeleteSessionUsers(ctx context.Context, message proto.Message) (proto.Message, error) {
	log.Println("DeleteUsers")
	return nil, nil
}

func AddSessionUsers(ctx context.Context, message proto.Message) (proto.Message, error) {

	log.Println("AddUsers")
	rpcInfo := ctx.Value(key.RpcInfo).(*grpcPb.RpcInfo)
	request, ok := message.(*grpcPb.AddSessionUsersRequest)
	if !ok {
		return nil, nil
	}

	logicServerConn := ctx.Value(key.LogicServerConn).(*grpc.ClientConn)
	sessionClient := grpcPb.NewSessionClient(logicServerConn)

	request.RpcInfo = rpcInfo
	response, err := sessionClient.AddSessionUsers(context.Background(), request)

	return response, err
}
