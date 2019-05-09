package go_ws

import (
  "net/http"
  "sync"
  "github.com/alex023/clock"
  "github.com/jinbanglin/log"
  "github.com/gorilla/websocket"
  "context"
  "github.com/jinbanglin/go-ws/ws_proto"
  "time"
  "github.com/jinbanglin/helper"
  "strings"
)

type WS struct {
  Clients          *sync.Map
  lock             *sync.Mutex
  register         chan *Client
  unregister       chan *Client
  broadcast        chan *BroadcastData
  clock            *clock.Clock
  ServerName       string
  ServerID         string
  ServerAddress    string
  ServerNameLen    int
  ServerIDLen      int
  ServerAddressLen int
}

var GWS *WS

func SetupWS() {

  socketService := GWS.setupWSRpcServer()
  socketService.Init()
  gWsRpc = &WsRpc{Client: socketService.Client()}
  ws_proto.RegisterWsRpcHandler(socketService.Server(), gWsRpc)
  go func() {
    if err := socketService.Run(); err != nil {
      log.Error(err)
    }
  }()
  time.Sleep(time.Second * 1)
  {
    GWS = &WS{
      lock:          new(sync.Mutex),
      Clients:       new(sync.Map),
      register:      make(chan *Client),
      unregister:    make(chan *Client),
      broadcast:     make(chan *BroadcastData, 10240),
      clock:         clock.NewClock(),
      ServerName:    socketService.Server().Options().Name,
      ServerID:      socketService.Server().Options().Id,
      ServerAddress: helper.GetLocalIP() + ":" + getPort(socketService.Server().Options().Address),
    }

    GWS.ServerNameLen = len(GWS.ServerName)
    GWS.ServerIDLen = len(GWS.ServerID)
    GWS.ServerAddressLen = len(GWS.ServerAddress)
    RegisterEndpoint(HeartbeatMsgID, &ws_proto.PingReq{}, Heartbeat)

    go GWS.Run()
  }
  log.Debugf("socket server start at: name=%s id=%s address=%s",
    GWS.ServerName,
    GWS.ServerID,
    GWS.ServerAddress)
}

func getPort(address string) string {
  return strings.Split(address, ":")[3]
}

type BroadcastData struct {
  roomID string
  userID string
  data   []byte
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
    WS:     GWS,
    UserID: userId,
    conn:   conn,
    send:   make(chan []byte),
    ctx:    context.Background(),
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

      w.Clients.Store(client.UserID, client)

      client.userOnline()

    case client := <-w.unregister:

      log.Debugf("OFFLINE |user_id=%s", client.UserID)

      client.userOffline()
      if _, ok := w.Clients.Load(client.UserID); ok {
        close(client.send)
        w.Clients.Delete(client.UserID)
      }

    case packet := <-w.broadcast:

      //if !strings.EqualFold(packet.RoomID, "") {
      //  w.Clients.Range(func(key, value interface{}) bool {
      //    client := value.(*Client)
      //    if client.RoomID == packet.RoomID &&client.getState()==_IS_ONLINE_STATE{
      //      if cnn, ok := w.Clients.Load(client.UserID); ok {
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
