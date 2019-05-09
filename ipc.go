package go_ws

import (
  "github.com/jinbanglin/go-micro/client"
  "github.com/jinbanglin/go-ws/ws_proto"
  "context"
  "github.com/jinbanglin/micro/opts"
  "github.com/jinbanglin/go-micro"
)

const _WS_SERVER_NAME = "go.micro.srv.ws"

var gWsRpc *WsRpc

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

func (w *WS) setupWSRpcServer(f func() micro.Option) micro.Service {
  return micro.NewService(append(opts.SServerOptions(_WS_SERVER_NAME), f())...)
}
