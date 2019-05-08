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
  endpoint    Endpoint
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
    endpoint:    endpoint,
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
  packet []byte) (b []byte, err error) {

  header, payload, err := ParseRemotePacket(packet)
  defer func() {
    if payload!=nil{
      payload.Release()
    }
    packetHeaderRelease(header)
  }()
  if err != nil {
    return nil, err
  }

  if !verifyPacket(len(packet), int(header.PacketLen)) {
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
  log.Debugf2(
    ctx,
    "FROM"+
      " |user_id=%s"+
      " |data=%v"+
      " |packet_header=%v"+
      " |next_seq=%s",
    client.userID,
    helper.Marshal2String(req),
    helper.Marshal2String(header),
    client.getLogTraceID(),
  )

  rsp, err := endpoint.endpoint(ctx, client, req)
  if err != nil {
    log.Error(err)
    return nil, err
  }

  body, err := proto.Marshal(rsp)
  if err != nil {
    log.Error(err)
    return nil, err
  }

  b = PackLocalPacket(
    header,
    body,
    GWS.getServerNameLen(),
    GWS.getServerIDLen(),
    GWS.getServerAddressLen(),
    GWS.serverName,
    GWS.GetServerID(),
    GWS.getServerAddress(),
    client.getLogTraceID(),
  )

  log.Debugf2(
    ctx,
    "TO"+
      " |user_id=%s"+
      " |data=%v"+
      " |packet_header=%v"+
      " |next_seq=%s",
    client.userID,
    helper.Marshal2String(rsp),
    helper.Marshal2String(header),
    header.Seq,
  )
  return
}

const HeartbeatMsgID MessageIDType = 10000

func Heartbeat(ctx context.Context, client *Client, req proto.Message) (rsp proto.Message, err error) {
  rsp = &ws_proto.PongRsp{
    Appid:         client.appID,
    UserId:        client.userID,
    RoomId:        client.roomID,
    ServerName:    client.ServerName,
    ServerId:      client.ServerID,
    ServerAddress: client.ServerAddress,
    Message:       &msg.Message{Code: 200, Msg: ""}}
  return
}
