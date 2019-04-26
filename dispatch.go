package go_ws

import (
  "reflect"
  "context"
  "github.com/gogo/protobuf/proto"
  "sync"
  "errors"
  "fmt"
  "github.com/jinbanglin/log"
  "github.com/jinbanglin/helper"
  "github.com/jinbanglin/go-ws/ws_proto"
  "github.com/jinbanglin/micro/message"
)

type MessageIDType = uint16
type Endpoint func(ctx context.Context, client *Client, req proto.Message) (rsp proto.Message, err error)

type Dispatch struct {
  dispatch map[uint16]*SchedulerEndpoint
  lock     *sync.RWMutex
}

type SchedulerEndpoint struct {
  RequestType reflect.Type
  handler     Endpoint
}

var gDispatch *Dispatch

func RegisterEndpoint(msgID uint16, req proto.Message, endpoint Endpoint) {

  if gDispatch == nil {
    gDispatch = &Dispatch{lock: new(sync.RWMutex), dispatch: make(map[uint16]*SchedulerEndpoint)}
  }

  gDispatch.lock.Lock()
  defer gDispatch.lock.Unlock()

  if _, ok := gDispatch.dispatch[msgID]; ok {
    panic("is already register")
  }
  gDispatch.dispatch[msgID] = &SchedulerEndpoint{
    RequestType: reflect.TypeOf(req).Elem(),
    handler:     endpoint,
  }
}

func getEndpoint(msgID uint16) (endpoint *SchedulerEndpoint, err error) {

  var ok = false
  if endpoint, ok = gDispatch.dispatch[msgID]; !ok {
    return nil, errors.New(fmt.Sprintf("no endpoint: %v", msgID))
  }
  return
}

func (d *Dispatch) Invoking(
  ctx context.Context,
  client *Client,
  packet []byte, seq uint64) (b []byte, err error) {

  header, payload, packetLength := ParseRemotePacket(packet)

  if !verifyPacket(len(packet), packetLength) {
    log.Error("verifyPacket |err=invalid packet")
    return nil, errors.New("invalid packet")
  }

  endpoint, err := getEndpoint(header.MsgID)
  if err != nil {
    log.Errorf2(ctx, "getEndpoint |err=%v", err)
    return nil, err
  }

  req := reflect.New(endpoint.RequestType).Interface().(proto.Message)
  err = proto.Unmarshal(payload.Bytes(), req)
  if err != nil {
    log.Error(err)
    return nil, err
  }
  log.Debugf("FROM |user_id=%s", client.userID, helper.Marshal2String(req))

  rsp, err := endpoint.handler(ctx, client, req)
  if err != nil {
    log.Error(err)
    return nil, err
  }
  log.Debugf("TO |user_id=%s", client.userID, helper.Marshal2String(rsp))

  body, err := proto.Marshal(rsp)
  if err != nil {
    log.Error(err)
    return nil, err
  }

  b = PackLocalPacket(header, body, seq)

  return
}

const HeartbeatMsgID MessageIDType = 10000

func Heartbeat(ctx context.Context, client *Client, req proto.Message) (rsp proto.Message, err error) {
  rsp = &ws_proto.PongRsp{
    Appid:   client.appID,
    UserId:  client.userID,
    RoomId:  client.roomID,
    Message: &msg.Message{Code: 200, Msg: ""}}
  return
}
