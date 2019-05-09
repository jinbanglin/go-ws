package main

import (
  "github.com/gorilla/websocket"
  "fmt"
  "github.com/jinbanglin/go-ws/ws_proto"
  "github.com/gogo/protobuf/proto"
  "github.com/jinbanglin/go-ws"
  "github.com/jinbanglin/helper"
  "github.com/jinbanglin/go-ws/bufferpool"
)

func main() {
  var dialer *websocket.Dialer

  conn, _, err := dialer.Dial("wss://xxxxxxx/ws/handshake/test/22222", nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  var header *go_ws.PacketHeader

  go func() {
    ping(conn)
  }()

  for {
    _, message, err := conn.ReadMessage()
    if err != nil {
      fmt.Println("read:", err)
      return
    }
    var payload *bufferpool.ByteBuffer

    header, payload, err = go_ws.ParseRemotePacket(message)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(helper.Marshal2String(header))

    switch header.MsgID {
    case 10000:
      b := &ws_proto.PongRsp{}

      proto.Unmarshal(payload.Bytes(), b)
      fmt.Println(helper.Marshal2String(b))
    case 10001:
      b := &ws_proto.SendMsgTestRsp{}

      proto.Unmarshal(payload.Bytes(), b)
      fmt.Println(helper.Marshal2String(b))
    }

  }
}

func ping(conn *websocket.Conn) {
  b, _ := proto.Marshal(&ws_proto.PingReq{Ping: "ping"})
  c := &go_ws.Client{WS: &go_ws.WS{
    ServerName:       "1",
    ServerID:         "1",
    ServerAddress:    "1",
    ServerNameLen:    1,
    ServerIDLen:      1,
    ServerAddressLen: 1,
  }}
  c.SetLogTraceID()
  packet := go_ws.PackLocalPacket(&go_ws.PacketHeader{MsgID: go_ws.HeartbeatMsgID},
    b, c)

  conn.WriteMessage(websocket.TextMessage, packet)
}
