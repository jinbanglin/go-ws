/*
 * Copyright (c) 2018 All Rights Reserved.
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
 *
 * Please contact me:
 * Author:jinbanglin
 * File:main.go
 * EMAIL:570751295@qq.com
 * LastModified:2018/08/01 11:45:01
 */

package main

import (
  "github.com/gin-gonic/gin"
  "github.com/jinbanglin/log"
  "github.com/jinbanglin/go-web"
  "github.com/jinbanglin/helper"
  "github.com/jinbanglin/micro/opts"
  "github.com/jinbanglin/go-ws"
  "github.com/jinbanglin/go-ws/ws_proto"
  "github.com/jinbanglin/go-ws/_examples/api.ws/glove"
)

var _WEB_WS_API_NAME = "go.micro.web.ws"

func main() {
  defer opts.Recover()
  helper.Chaos("api.ws.toml", log.SetupMossLog, helper.MgoChaos, helper.RedisChaos, go_ws.WsChaos)
  service := web.NewService(opts.WServerWithOptions(_WEB_WS_API_NAME, nil)...)
  if err := service.Init(); err != nil {
    log.Fatal(err)
  }
  app := gin.Default()

  go_ws.SetupWS()

  RegisterEndpoint()

  app.GET("/handshake/test/:userid", func(context *gin.Context) {
    go_ws.Handshake(context.Param("userid"), context.Writer, context.Request)
  })

  service.Handle("/", app)
  if err := service.Run(); err != nil {
    log.Fatal(err)
  }
}

func RegisterEndpoint() {
  go_ws.RegisterEndpoint(60002, &ws_proto.SendMsgTestReq{}, glove.SendMsg)
}
