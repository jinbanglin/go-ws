package go_ws

import (
  "github.com/gorilla/websocket"
  "sync/atomic"
  "context"
  "github.com/alex023/clock"
  "github.com/jinbanglin/log"
  "time"
  "github.com/spf13/viper"
  "strconv"
)

type state = int32

const _NORMAL_STATE state = 0
const _IS_CONNECTED_STATE state = 1

var (
  // Time allowed to write a message to the peer.
  WriteWait = 10 * time.Second

  // Time allowed to read the next pong message from the peer.
  PongWait = 60 * time.Second

  // send pings to peer with this period. Must be less than PongWait.
  PingPeriod = (PongWait * 9) / 10

  // Maximum message size allowed from peer.
  MaxMessageSize int64 = 1024
)

func WsChaos() {
  if v := viper.GetInt("ws.write_wait"); v > 0 {
    WriteWait = time.Duration(v) * time.Second
  }
  if v := viper.GetInt("ws.pong_wait"); v > 0 {
    PongWait = time.Duration(v) * time.Second
  }
  if v := viper.GetInt64("ws.max_message_size"); v > 0 {
    MaxMessageSize = v
  }
}

type Client struct {
  ws     *WS
  ctx    context.Context
  conn   *websocket.Conn
  send   chan []byte
  state  int32
  appID  string
  userID string
  roomID string
}

func (c *Client) readLoop() {

  c.conn.SetReadLimit(MaxMessageSize)

  for {
    _, packet, err := c.conn.ReadMessage()
    if err != nil {
      if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
        log.Error(err)
      }
      break
    }

    c.conn.SetReadDeadline(time.Now().Add(PongWait))

    seq := MakeSeq()

    c.ctx = context.WithValue(c.ctx, log.GContextKey, strconv.FormatUint(seq, 10))

    b, err := gDispatch.Invoking(c.ctx, c, packet, seq)
    if err != nil {
      log.Error2(c.ctx, err)
      continue
    }
    Broadcast(&broadcastData{
      roomId: c.roomID,
      userId: c.userID,
      data:   b,
    }, c.ws)
  }
  c.conn.Close()
}

func (c *Client) writeLoop() {
  defer func() {
    c.conn.Close()
  }()
  var clockQuit = make(chan struct{})
  var job clock.Job
  job, _ = c.ws.clock.AddJobRepeat(PingPeriod, 0, func() {
    c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
    if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
      clockQuit <- struct{}{}
    }
  })
  for {
    select {
    case packet, ok := <-c.send:
      c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
      if !ok {
        c.conn.WriteMessage(websocket.CloseMessage, []byte{})
        job.Cancel()
        return
      }
      w, err := c.conn.NextWriter(websocket.TextMessage)
      if err != nil {
        log.Error2(c.ctx, err)
        job.Cancel()
        return
      }
      w.Write(packet)
      if err := w.Close(); err != nil {
        log.Error2(c.ctx, err)
        job.Cancel()
        return
      }
    case <-clockQuit:
      job.Cancel()
      return
    }
  }
}

func Broadcast(msg *broadcastData, ws *WS) {
  ws.broadcast <- msg
}

func (c *Client) RemoteIP() string {
  return c.conn.RemoteAddr().String()
}

func (c *Client) GetAppID() string {
  return c.appID
}

func (c *Client) GetUserID() string {
  return c.userID
}

func (c *Client) GetRoomID() string {
  return c.roomID
}

func (c *Client) getConn() *websocket.Conn {
  return c.conn
}

func (c *Client) getState() int32 {
  return atomic.LoadInt32(&c.state)
}

func (c *Client) setState(new state) {
  atomic.SwapInt32(&c.state, new)
}

func (c *Client) LocalAddr() string {
  return c.conn.LocalAddr().String()
}
