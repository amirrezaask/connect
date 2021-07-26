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

type Connect struct {
	Users  UserConnections
	Logger *zap.SugaredLogger
}
type EventType uint8

const (
	_ = iota
	EventType_NewMessage
)

type Event struct {
	Creator   string
	EventType EventType `json:"event_type"`
	Payload   []byte `json:"payload"`
}

func (c *Connect) HandleEvent(e *Event) {
	if e.EventType == EventType_NewMessage {
		for _, conn := range c.Users {
			conn.WriteJSON(c)
		}
	}
}

type UserConnections map[string]*websocket.Conn

func (uc UserConnections) Add(nickname string, conn *websocket.Conn) {
	uc[nickname] = conn
}

func (s *Connect) WSHandler(w http.ResponseWriter, r *http.Request) {
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
	go func(c *websocket.Conn) {
		for {
			e := new(Event)
			// err = c.ReadJSON(e)
			t, b, err := c.ReadMessage()
			if t != websocket.TextMessage {
				s.Logger.Errorf("unsupported message type: %v", t)
				continue
			}
			err = json.Unmarshal(b, e)
			if err != nil {
				s.Logger.Errorf("error in reading event from client: %v", err)
				continue
			}
			// e.Creator = nickName
			// s.HandleEvent(e)
		}
	}(conn)
	conn.WriteMessage(websocket.TextMessage, []byte("Connected"))
}

func main() {
	l, _ := zap.NewDevelopment()
	logger := l.Sugar()

	c := &Connect{
		Users:  make(UserConnections),
		Logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/ws", c.WSHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
