package go_ws

import (
  "github.com/gorilla/websocket"
  "github.com/jinbanglin/go-ws/bufferpool"
)

type Client struct {
  userID string

  appID string

  roomID string

  state int32

  send chan *bufferpool.ByteBuffer

  conn *websocket.Conn
}

func (c *Client) remoteIP() string {
  return c.conn.RemoteAddr().String()
}
