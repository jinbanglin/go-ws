package go_ws

import (
  "net/http"
  "sync"
  "github.com/alex023/clock"
  "github.com/jinbanglin/log"
  "github.com/gorilla/websocket"
  "context"
  "strings"
)

type WS struct {
  Clients *sync.Map

  lock       *sync.Mutex
  register   chan *Client
  unregister chan *Client
  broadcast  chan *broadcastData

  clock *clock.Clock
}

type broadcastData struct {
  roomId string
  userId string
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
  log.Debugf("HANDSHAKE |user_id=%", userId)
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
      if !strings.EqualFold(packet.roomId, "") {
        w.Clients.Range(func(key, value interface{}) bool {
          client := value.(*Client)
          if client.roomID == packet.roomId {
            if cnn, ok := w.Clients.Load(client.userID); ok {
              cnn.(*Client).send <- packet.data
            }
          }
          return true
        })
      } else {
        if cnn, ok := w.Clients.Load(packet.userId); ok {
          cnn.(*Client).send <- packet.data
        }
      }
    }
  }
  w.lock.Unlock()
}
