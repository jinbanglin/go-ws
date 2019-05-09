package glove

import (
  "moss/general/message"
  "context"
  "github.com/jinbanglin/go-ws"
  "github.com/gogo/protobuf/proto"
  "github.com/jinbanglin/go-ws/ws_proto"
)

func SendMsg(ctx context.Context, client *go_ws.Client, req proto.Message) (rsp proto.Message, err error) {
  userLoad, sendPacket := req.(*ws_proto.SendMsgTestReq), &ws_proto.SendMsgTestRsp{Message: &message.MsgServerSuccess}
  go_ws.BroadcastToOne(ctx, 10001, userLoad.UserId, "", client, sendPacket)
  return sendPacket, nil
}
