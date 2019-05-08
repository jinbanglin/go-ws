package go_ws

import (
  "net/http"
  "sync"
  "github.com/alex023/clock"
  "github.com/jinbanglin/log"
  "github.com/gorilla/websocket"
  "context"
  "github.com/jinbanglin/go-ws/ws_proto"
)

type WS struct {
  Clients       *sync.Map
  lock          *sync.Mutex
  register      chan *Client
  unregister    chan *Client
  broadcast     chan *BroadcastData
  clock         *clock.Clock
  serverName    string
  serverID      string
  serverAddress string
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

func (w *WS) getServerAddressLen() int {
  return len(w.serverAddress)
}

func (w *WS) getServerAddress() string {
  return w.serverAddress
}

var GWS *WS

func SetupWS() {

  initRpcClient()

  socketService := GWS.setupWSRpcServer()
  socketService.Init()
  go func() {
    if err := socketService.Run(); err != nil {
      log.Error(err)
    }
  }()

  {
    GWS = &WS{
      lock:          new(sync.Mutex),
      Clients:       new(sync.Map),
      register:      make(chan *Client),
      unregister:    make(chan *Client),
      broadcast:     make(chan *BroadcastData, 10240),
      clock:         clock.NewClock(),
      serverName:    socketService.Server().Options().Name,
      serverID:      socketService.Server().Options().Id,
      serverAddress: socketService.Server().Options().Address,
    }
    RegisterEndpoint(HeartbeatMsgID, &ws_proto.PingReq{}, Heartbeat)

    go GWS.Run()
  }
  log.Debugf("socket server start at: name=%s id=%s address=%s", GWS.serverName, GWS.serverID, GWS.serverAddress)
}

type BroadcastData struct {
  roomID  string
  userID  string
  address string
  seq     string
  data    []byte
}

func SetBroadcastData(roomID, userID string, b []byte) *BroadcastData {
  return &BroadcastData{
    roomID: roomID,
    userID: userID,
    data:   b,
  }
}

var gUpGrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}

func Handshake(userId string, w http.ResponseWriter, r *http.Request) {
  log.Debugf("HANDSHAKE |user_id=%s", userId)
  if GWS == nil {
    panic("no setup websocket!")
  }

  conn, err := gUpGrader.Upgrade(w, r, nil)
  if err != nil {
    log.Error(err)
    return
  }

  client := &Client{
    userID:        userId,
    conn:          conn,
    send:          make(chan []byte),
    ctx:           context.Background(),
    ServerID:      GWS.serverID,
    ServerName:    GWS.serverName,
    ServerAddress: GWS.serverAddress,
  }
  GWS.register <- client
  go client.readLoop()
  go client.writeLoop()
}

func (w *WS) Run() {

  w.lock.Lock()

  for {
    select {
    case client := <-w.register:

      w.Clients.Store(client.userID, client)

      client.userOnline()

    case client := <-w.unregister:

      log.Debugf("OFFLINE |user_id=%s", client.userID)

      client.userOffline()
      if _, ok := w.Clients.Load(client.userID); ok {
        close(client.send)
        w.Clients.Delete(client.userID)
      }

    case packet := <-w.broadcast:

      //if !strings.EqualFold(packet.roomID, "") {
      //  w.Clients.Range(func(key, value interface{}) bool {
      //    client := value.(*Client)
      //    if client.roomID == packet.roomID &&client.getState()==_IS_ONLINE_STATE{
      //      if cnn, ok := w.Clients.Load(client.userID); ok {
      //        cnn.(*Client).send <- packet.data
      //      }
      //    }
      //    return true
      //  })
      //} else {
      if cnn, ok := w.Clients.Load(packet.userID);
        ok && cnn.(*Client).getState() == _IS_ONLINE_STATE {

        cnn.(*Client).send <- packet.data
      }
      //}
    }
  }

  w.lock.Unlock()
}
