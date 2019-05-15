package main

import (
  "github.com/gorilla/websocket"
  "fmt"
  "github.com/gogo/protobuf/proto"
  "github.com/jinbanglin/helper"
  "github.com/jinbanglin/go-ws/bufferpool"
  "github.com/jinbanglin/go-ws"
  "github.com/jinbanglin/go-ws/proto"
)

func main() {
  var dialer *websocket.Dialer

  conn, _, err := dialer.Dial("wss://xxxxxxx/ws/handshake/test/22222", nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  var header *ws.PacketHeader

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

    header, payload, err = ws.ParseRemotePacket(message)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(helper.Marshal2String(header))

    switch header.MsgID {
    case 10000:
      b := &wsp.PongRsp{}

      proto.Unmarshal(payload.Bytes(), b)
      fmt.Println(helper.Marshal2String(b))
    case 10001:
      b := &wsp.SendMsgTestRsp{}

      proto.Unmarshal(payload.Bytes(), b)
      fmt.Println(helper.Marshal2String(b))
    }

  }
}

func ping(conn *websocket.Conn) {
  b, _ := proto.Marshal(&wsp.PingReq{Ping: "ping"})
  c := &ws.Client{WS: &ws.WS{
    ServerName:       "1",
    ServerID:         "1",
    ServerAddress:    "1",
    ServerNameLen:    1,
    ServerIDLen:      1,
    ServerAddressLen: 1,
  }}
  c.SetLogTraceID()
  packet := ws.PackLocalPacket(&ws.PacketHeader{MsgID: ws.HeartbeatMsgID},
    b, c)

  conn.WriteMessage(websocket.TextMessage, packet)
}
