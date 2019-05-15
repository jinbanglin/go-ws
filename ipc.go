package ws

import (
  "github.com/jinbanglin/go-micro/client"
  "context"
  "github.com/jinbanglin/micro/opts"
  "github.com/jinbanglin/go-micro"
  "github.com/jinbanglin/go-ws/proto"
)

const _WS_SERVER_NAME = "go.micro.srv.ws"

var GWsRpc *WsRpc

type WsRpc struct {
  Client client.Client
}

func (w *WsRpc) Request(ctx context.Context, req *wsp.RpcReq, rsp *wsp.RpcRsp) error {
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
