package grpc

import (
	"github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	protocolClient "im/protocol/client"
	"log"
)

type Rpc struct {
	ClientConn *grpc.ClientConn
}

func (s *Rpc) Rpc(c context.Context, request *protocolClient.RpcRequest) (*protocolClient.RpcResponse, error) {

	rpcResponse := &protocolClient.RpcResponse{
		Rid: request.GetRid(),
	}
	if request.MessageType == grpcPb.MessageTypeCreateSessionRequest {

		//远程调用im逻辑服务器
		sessionClient := grpcPb.NewSessionClient(s.ClientConn)
		protoMessage := grpcPb.Factory((grpcPb.MessageType)(request.MessageType))

		err := proto.Unmarshal(request.ProtoBuf, protoMessage)

		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		log.Println(protoMessage.String())

		reply, err := sessionClient.CreateSession(context.Background(), protoMessage.(*grpcPb.CreateSessionRequest))
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		log.Println(reply.String())
		protoBuf, err := proto.Marshal(reply)
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		rpcResponse = &protocolClient.RpcResponse{
			Rid:         request.GetRid(),
			MessageType: grpcPb.MessageTypeCreateSessionReply,
			ProtoBuf:    protoBuf,
		}
		return rpcResponse, nil
	}

	return rpcResponse, nil
}
