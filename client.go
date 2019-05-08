package go_ws

import (
  "github.com/gorilla/websocket"
  "sync/atomic"
  "context"
  "github.com/alex023/clock"
  "github.com/jinbanglin/log"
  "time"
  "github.com/spf13/viper"
  "github.com/google/uuid"
  "github.com/jinbanglin/go-micro/metadata"
  "github.com/jinbanglin/go-ws/ws_proto"
  "github.com/jinbanglin/helper"
  "encoding/json"
  "strings"
  "github.com/jinbanglin/go-micro/client"
)

type state = int32

const (
  _                 state = iota
  _IS_OFFLINE_STATE
  _IS_ONLINE_STATE
)

var (
  // Time allowed to write a message to the peer.
  WriteWait = 10 * time.Second

  // Time allowed to read the next pong message from the peer.
  PongWait = 60 * time.Second

  // send pings to peer with this period. Must be less than PongWait.
  PingPeriod = (PongWait * 9) / 10

  // Maximum message size allowed from peer.
  MaxMessageSize int64 = 1024

  // read channel cache
  MaxReadCache = 10240

  // write channel cache
  MaxWriteCache = 10240
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
  if v := viper.GetInt("ws.read_cache"); v > 0 {
    MaxReadCache = v
  }
  if v := viper.GetInt("ws.write_cache"); v > 0 {
    MaxWriteCache = v
  }
}

type Client struct {
  ctx    context.Context
  conn   *websocket.Conn
  send   chan []byte
  appID  string
  userID string
  roomID string

  State         int32
  ServerName    string
  ServerID      string
  ServerAddress string
}

func (c *Client) setLogTraceID() {
  c.ctx = metadata.NewContext(c.ctx, metadata.Metadata{log.GContextKey: uuid.New().String()})
}

func (c *Client) getLogTraceID() string {
  md, ok := metadata.FromContext(c.ctx)
  if !ok {
    md = metadata.Metadata{}
    return ""
  }

  return md[log.GContextKey]
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
    c.setLogTraceID()

    b, err := gDispatch.Invoking(c.ctx, c, packet)
    if err != nil {
      log.Error2(c.ctx, err)
      continue
    }
    broadcastLocalServer(&BroadcastData{
      roomID: c.roomID,
      userID: c.userID,
      data:   b,
    })
  }
  c.conn.Close()
}

func (c *Client) writeLoop() {
  defer func() {
    c.conn.Close()
  }()
  var clockQuit = make(chan struct{})
  var job clock.Job
  job, _ = GWS.clock.AddJobRepeat(PingPeriod, 0, func() {
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

func BroadcastMulti(userID []string, msg *BroadcastData) {
  for _, v := range userID {
    if c := getUserState(v); c != nil {

      msg.userID = v
      if strings.EqualFold(c.ServerAddress, GWS.serverAddress) {
        broadcastLocalServer(msg)
      } else {
        broadcastOtherServer(msg)
      }
    }
  }
}

func BroadcastSingle(msg *BroadcastData) {
  if c := getUserState(msg.userID); c != nil {
    if strings.EqualFold(c.ServerAddress, GWS.serverAddress) {

      broadcastLocalServer(msg)
    } else {

      broadcastOtherServer(msg)
    }
  }
}

func BroadcastByAddress(address string, msg *BroadcastData) {
  if strings.EqualFold(address, GWS.serverAddress) {
    broadcastLocalServer(msg)
  } else {
    broadcastOtherServer(msg)
  }
}

func broadcastLocalServer(msg *BroadcastData) {
  GWS.broadcast <- msg
}

func broadcastOtherServer(msg *BroadcastData) {

  GWSService.Request(context.Background(), &ws_proto.RpcReq{
    UserId: msg.userID,
    RoomId: msg.roomID,
    Packet: msg.data,
    Seq:    msg.seq,
  }, client.WithAddress(msg.address))
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
  return atomic.LoadInt32(&c.State)
}

func (c *Client) setState(new state) {
  atomic.SwapInt32(&c.State, new)
}

func (c *Client) LocalAddr() string {
  return c.conn.LocalAddr().String()
}

const _REDIS_KEY_USER_STATE = "ws:user:state:"

func (c *Client) userOnline() {
  c.setState(_IS_ONLINE_STATE)
  helper.GRedisRing.Set(_REDIS_KEY_USER_STATE+c.userID, helper.Marshal2Bytes(c), time.Hour*24)
}

func (c *Client) userOffline() {
  c.setState(_IS_OFFLINE_STATE)
  helper.GRedisRing.Del(_REDIS_KEY_USER_STATE + c.userID)
}

func getUserState(userID string) *Client {
  c := &Client{}
  b, err := helper.GRedisRing.Get(_REDIS_KEY_USER_STATE + userID).Bytes()
  if err != nil {
    return nil
  }
  if json.Unmarshal(b, c) != nil {
    return nil
  }
  return c
}

// PipelineUserState:get all user's state,
// eg.find from a room
func PipelineUserState(userID []string) []*Client {
  pipe := helper.GRedisRing.Pipeline()
  results := make([]*Client, len(userID))
  for _, v := range userID {
    c := &Client{}
    if b, err := pipe.Get(_REDIS_KEY_USER_STATE + v).Bytes(); err == nil {
      if json.Unmarshal(b, c) == nil {
        results = append(results, c)
      }
    }
  }
  return results
}
