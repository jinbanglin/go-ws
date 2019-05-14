package go_ws

import (
  "net/http"
  "sync"
  "github.com/alex023/clock"
  "github.com/jinbanglin/log"
  "github.com/gorilla/websocket"
  "context"
  "github.com/jinbanglin/go-ws/ws_proto"
  "strings"
  "github.com/jinbanglin/go-micro"
  "github.com/jinbanglin/helper"
)

type WS struct {
  lock             *sync.WaitGroup
  clients          *sync.Map
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
  GWS = &WS{
    lock:       new(sync.WaitGroup),
    clients:    new(sync.Map),
    register:   make(chan *Client),
    unregister: make(chan *Client),
    broadcast:  make(chan *BroadcastData, 10240),
    clock:      clock.NewClock(),
  }

  GWS.lock.Add(1)

  var socketService micro.Service
  socketService = GWS.setupWSRpcServer(func() micro.Option {
    return micro.AfterStart(func() error {

      GWS.lock.Done()

      GWS.ServerName = socketService.Server().Options().Name
      GWS.ServerID = socketService.Server().Options().Id
      GWS.ServerAddress = helper.GetLocalIP() + ":" + getPort(socketService.Server().Options().Address)
      GWS.ServerNameLen = len(GWS.ServerName)
      GWS.ServerIDLen = len(GWS.ServerID)
      GWS.ServerAddressLen = len(GWS.ServerAddress)
      RegisterEndpoint(HeartbeatMsgID, &ws_proto.PingReq{}, Heartbeat)

      go GWS.Run()
      return nil
    })
  })

  go func() {
    if err := socketService.Run(); err != nil {
      panic(err)
    }
  }()

  GWS.lock.Wait()

  socketService.Init()
  gWsRpc = &WsRpc{Client: socketService.Client()}
  ws_proto.RegisterWsRpcHandler(socketService.Server(), gWsRpc)

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

  for {
    select {
    case client := <-w.register:

      w.clients.Store(client.UserID, client)

      client.userOnline()

    case client := <-w.unregister:

      log.Debugf("OFFLINE |user_id=%s", client.UserID)

      client.userOffline()
      if _, ok := w.clients.Load(client.UserID); ok {
        close(client.send)
        w.clients.Delete(client.UserID)
      }

    case packet := <-w.broadcast:

      //if !strings.EqualFold(packet.RoomID, "") {
      //  w.clients.Range(func(key, value interface{}) bool {
      //    client := value.(*Client)
      //    if client.RoomID == packet.RoomID &&client.getState()==_IS_ONLINE_STATE{
      //      if cnn, ok := w.clients.Load(client.UserID); ok {
      //        cnn.(*Client).send <- packet.data
      //      }
      //    }
      //    return true
      //  })
      //} else {
      if cnn, ok := w.clients.Load(packet.userID);
        ok && cnn.(*Client).getState() == _IS_ONLINE_STATE {
        cnn.(*Client).send <- packet.data
      }
      //}
    }
  }
}
