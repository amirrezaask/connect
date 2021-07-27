package main

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Connect struct {
	Users  UserConnections
	Logger *zap.SugaredLogger
	Bus    Bus
}



type UserConnections map[string]*websocket.Conn

func (uc UserConnections) Add(nickname string, conn *websocket.Conn) {
	uc[nickname] = conn
}

func (uc UserConnections) Get(nickName string) *websocket.Conn {
	return uc[nickName]
}
