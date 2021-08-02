package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (s *ConnectServer) WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Logger.Errorf("error in upgrading user connection to ws protocol: %v", err)
		return
	}
	nickName := r.URL.Query().Get("nickname")
	if nickName == "" {
		s.Logger.Error("user has not send a nickname")
		return
	}
	s.Logger.Debugf("%s connected", nickName)
	s.Users.Add(nickName, conn)
	go func(c *websocket.Conn, nickName string) {
		for {
			e := &Event{}
			err = c.ReadJSON(e)
			if err != nil {
				//TODO(amirreza): please fix this:))
				if !strings.Contains(err.Error(), "EOF") {
					s.Logger.Errorf("error in reading event from client: %v", err)
				}
				continue
			}
			e.Creator = nickName
			s.Logger.Debugf("received from %s: %+v", nickName, e)
			s.Bus.Emit(e)
		}
	}(conn, nickName)
	conn.WriteMessage(websocket.TextMessage, []byte("Connected"))
}
