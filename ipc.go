package go_ws

import (
  "github.com/jinbanglin/go-micro/client"
  "github.com/jinbanglin/go-ws/ws_proto"
  "context"
  "github.com/jinbanglin/micro/opts"
  "github.com/jinbanglin/go-micro"
)

const _WS_SERVER_NAME = "go.micro.srv.ws"

type WsRpc struct {
  Client client.Client
}

func (w *WsRpc) Request(ctx context.Context, req *ws_proto.RpcReq, rsp *ws_proto.RpcRsp) error {

  broadcastLocalServer(&BroadcastData{
    roomID: req.RoomId,
    userID: req.UserId,
    data:   req.Packet,
  })
  return nil
}

var GWSService ws_proto.WsRpcService

func initRpcClient() {

  GWSService = ws_proto.NewWsRpcService(_WS_SERVER_NAME, client.NewClient(opts.WClientOptions()...))
}

func (w *WS) setupWSRpcServer() micro.Service {
  return micro.NewService(opts.SServerOptions(_WS_SERVER_NAME)...)
}
