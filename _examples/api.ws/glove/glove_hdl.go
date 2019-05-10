package glove

import (
  "context"
  "github.com/jinbanglin/go-ws"
  "github.com/gogo/protobuf/proto"
  "github.com/jinbanglin/go-ws/ws_proto"
  "github.com/jinbanglin/micro/message"
)

func SendMsg(ctx context.Context, client *go_ws.Client, req proto.Message) (rsp proto.Message, err error) {
  userLoad, sendPacket := req.(*ws_proto.SendMsgTestReq), &ws_proto.SendMsgTestRsp{Message: &msg.Message{
    Code: 200,
    Msg:  "hello go websocket!",
  }}
  go_ws.BroadcastToOne(ctx, 10001, userLoad.UserId, "", client, sendPacket)
  return sendPacket, nil
}
