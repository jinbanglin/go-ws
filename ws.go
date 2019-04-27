package go_ws

import (
  "net/http"
  "sync"
  "github.com/alex023/clock"
  "github.com/jinbanglin/log"
  "github.com/gorilla/websocket"
  "context"
  "strings"
  "github.com/jinbanglin/go-ws/ws_proto"
)

type WS struct {
  Clients *sync.Map

  lock       *sync.Mutex
  register   chan *Client
  unregister chan *Client
  broadcast  chan *broadcastData

  clock *clock.Clock

  serverName string
  serverID   string
}

func (w *WS) GetServerName() string {
  return w.serverName
}

func (w *WS) GetServerID() string {
  return w.serverID
}

func (w *WS) getServerNameLen() int {
  return len(w.serverName)
}

func (w *WS) getServerIDLen() int {
  return len(w.serverID)
}

var GWS *WS

func SetupWS(ServerName, ServerID string) {
  GWS = &WS{
    lock:       new(sync.Mutex),
    Clients:    new(sync.Map),
    register:   make(chan *Client),
    unregister: make(chan *Client),
    broadcast:  make(chan *broadcastData),
    clock:      clock.NewClock(),
    serverName: ServerName,
    serverID:   ServerID,
  }
  RegisterEndpoint(HeartbeatMsgID, &ws_proto.PingReq{}, Heartbeat)

  log.Debugf("server start: ServerName=%s ServerID=%s", ServerName, ServerID)
  go GWS.Run()
}

type broadcastData struct {
  roomID string
  userID string
  data   []byte
}

var gUpGrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}

func Handshake(ws *WS, userId string, w http.ResponseWriter, r *http.Request) {
  log.Debugf("HANDSHAKE |user_id=%s", userId)
  conn, err := gUpGrader.Upgrade(w, r, nil)
  if err != nil {
    log.Error(err)
    return
  }

  //if client was connected,kicked it
  client, ok := Search(ws, userId)
  if ok {
    client.conn = conn
    client.setState(_IS_CONNECTED_STATE)
  } else {
    client = &Client{
      userID: userId,
      ws:     ws,
      conn:   conn,
      send:   make(chan []byte),
      ctx:    context.Background(),
    }
    client.setState(_NORMAL_STATE)
    client.ws.register <- client
    go client.readLoop()
    go client.writeLoop()
  }
}

func Search(ws *WS, userId string) (*Client, bool) {
  v, ok := ws.Clients.Load(userId)
  if ok {
    return v.(*Client), ok
  }
  return nil, false
}

func (w *WS) Run() {
  w.lock.Lock()
  for {
    select {
    case client := <-w.register:
      w.Clients.Store(client.userID, client)
    case client := <-w.unregister:
      log.Debugf("OFFLINE |user_id=%s", client.userID)
      if _, ok := w.Clients.Load(client.userID); ok {
        close(client.send)
        w.Clients.Delete(client.userID)
      }
    case packet := <-w.broadcast:
      if !strings.EqualFold(packet.roomID, "") {
        w.Clients.Range(func(key, value interface{}) bool {
          client := value.(*Client)
          if client.roomID == packet.roomID {
            if cnn, ok := w.Clients.Load(client.userID); ok {
              cnn.(*Client).send <- packet.data
            }
          }
          return true
        })
      } else {
        if cnn, ok := w.Clients.Load(packet.userID); ok {
          cnn.(*Client).send <- packet.data
        }
      }
    }
  }
  w.lock.Unlock()
}
