package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
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
				s.Logger.Errorf("error in reading event from client: %v", err)
				continue
			}
			e.Creator = nickName
			s.Bus.Emit(e)
		}
	}(conn, nickName)
	conn.WriteMessage(websocket.TextMessage, []byte("Connected"))
}


func main() {
	l, _ := zap.NewDevelopment()
	logger := l.Sugar()

	c := &ConnectServer{
		Users:  make(UserConnections),
		Logger: logger,
	}

	b := NewChannelBus()
	b.Register(EventType_NewMessage, newMessageHandler(c))
	c.Bus = b

	mux := http.NewServeMux()
	mux.HandleFunc("/api/ws", c.WSHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func newMessageHandler(c *ConnectServer) func(e *Event) error {
	return func(e *Event) error {
		if e.EventType == EventType_NewMessage {
			p := &NewMessagePayload{}
			err := json.Unmarshal(e.Payload, p)
			if err != nil {
				c.Logger.Errorf("error in unmarshaling new message payload: %v", err)
				return nil
			}
			c.Users.Get(p.Receiver).WriteJSON(p)
		}
		return nil
	}
}
