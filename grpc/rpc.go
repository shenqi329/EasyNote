package grpc

import (
	easynoteError "easynote/error"
	"github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	protocolClient "im/protocol/client"
	"log"
)

type Rpc struct {
}

func (r *Rpc) Rpc(ctx context.Context, request *protocolClient.RpcRequest) (*protocolClient.RpcResponse, error) {

	accessServerConn := ctx.Value("AccessServerConn").(*grpc.ClientConn)
	imServerConn := ctx.Value("ImServerConn").(*grpc.ClientConn)

	accessServerConn = accessServerConn

	rpcResponse := &protocolClient.RpcResponse{
		Rid:    request.Rid,
		Code:   easynoteError.CommonInternalServerError,
		Desc:   easynoteError.ErrorCodeToText(easynoteError.CommonInternalServerError),
		ConnId: request.ConnId,
	}
	if request.MessageType == grpcPb.MessageTypeCreateSessionRequest {

		//远程调用im逻辑服务器
		sessionClient := grpcPb.NewSessionClient(imServerConn)
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
		log.Println(request.ConnId)
		protoBuf, err := proto.Marshal(reply)
		log.Println(protoBuf)
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		rpcResponse = &protocolClient.RpcResponse{
			Rid:         request.GetRid(),
			MessageType: grpcPb.MessageTypeCreateSessionReply,
			ProtoBuf:    protoBuf,
			ConnId:      request.ConnId,
		}
		return rpcResponse, nil
	}

	return rpcResponse, nil
}
