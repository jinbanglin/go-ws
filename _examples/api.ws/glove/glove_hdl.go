package glove

import (
  "context"
  "github.com/gogo/protobuf/proto"
  "github.com/jinbanglin/micro/message"
  "github.com/jinbanglin/go-ws"
  "github.com/jinbanglin/go-ws/proto"
)

func SendMsg(ctx context.Context, client *ws.Client, req proto.Message) (rsp proto.Message, err error) {
  userLoad, sendPacket := req.(*wsp.SendMsgTestReq), &wsp.SendMsgTestRsp{Message: &msg.Message{
    Code: 200,
    Msg:  "hello go websocket!",
  }}
  ws.BroadcastToOne(ctx, 10001, userLoad.UserId, "", client, sendPacket)
  return sendPacket, nil
}
