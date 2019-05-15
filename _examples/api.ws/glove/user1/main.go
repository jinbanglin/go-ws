package main

import (
  "github.com/gorilla/websocket"
  "fmt"
  "github.com/gogo/protobuf/proto"
  "github.com/jinbanglin/helper"
  "github.com/jinbanglin/go-ws/bufferpool"
  "time"
  "github.com/jinbanglin/go-ws"
  "github.com/jinbanglin/go-ws/proto"
)

func main() {
  var dialer *websocket.Dialer

  conn, _, err := dialer.Dial("wss://xxxx/ws/handshake/test/11111", nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  var header *ws.PacketHeader

  go func() {
    ping(conn)
    time.Sleep(time.Second * 2)
    sendMsgTest(conn, header)
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
      return
    }
    fmt.Println(helper.Marshal2String(header))
    b := &wsp.PongRsp{}

    proto.Unmarshal(payload.Bytes(), b)
    fmt.Println(helper.Marshal2String(b))
  }
}

var userid2 = "22222"

func sendMsgTest(conn *websocket.Conn, header *ws.PacketHeader) {
  b, _ := proto.Marshal(&wsp.SendMsgTestReq{
    UserId: userid2,
  })
  c := &ws.Client{
    WS: &ws.WS{
      ServerName:       header.ServerName,
      ServerID:         header.ServerID,
      ServerAddress:    header.ServerAddress,
      ServerNameLen:    len(header.ServerName),
      ServerIDLen:      len(header.ServerID),
      ServerAddressLen: len(header.ServerAddress),
    },
  }
  c.SetLogTraceID()
  packet := ws.PackLocalPacket(&ws.PacketHeader{MsgID: 60002},
    b, c)
  conn.WriteMessage(websocket.TextMessage, packet)
}

func ping(conn *websocket.Conn) {
  b, _ := proto.Marshal(&wsp.PingReq{Ping: "ping"})

  c := &ws.Client{
    WS: &ws.WS{
      ServerName:       "1",
      ServerID:         "1",
      ServerAddress:    "1",
      ServerNameLen:    1,
      ServerIDLen:      1,
      ServerAddressLen: 1,
    },
  }
  c.SetLogTraceID()
  packet := ws.PackLocalPacket(&ws.PacketHeader{MsgID: ws.HeartbeatMsgID},
    b, c)

  conn.WriteMessage(websocket.TextMessage, packet)
}
