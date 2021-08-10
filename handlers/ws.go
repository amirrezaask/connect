package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/amirrezaask/connect/bus"
	"github.com/amirrezaask/connect/domain"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WSHandler struct {
	Users  UserConnections
	Logger *zap.SugaredLogger
	Bus    bus.Bus
	DB     *sql.DB
}

type UserConnections map[string]*websocket.Conn

func (uc UserConnections) Add(nickname string, conn *websocket.Conn) {
	uc[nickname] = conn
}

func (uc UserConnections) Get(nickName string) *websocket.Conn {
	return uc[nickName]
}

func (s *WSHandler) WSHandler(w http.ResponseWriter, r *http.Request) {
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
			e := &domain.Event{}
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

func (c *WSHandler) NewMessageEventHandler() func(e *domain.Event) error {
	return func(e *domain.Event) error {
		if e.EventType == domain.EventType_NewMessage {
			p := &domain.NewMessagePayload{}
			err := json.Unmarshal(e.Payload, p)
			if err != nil {
				c.Logger.Errorf("error in unmarshaling new message payload: %v", err)
				return nil
			}
			// save message into database
			go func(db *sql.DB) {
			}(c.DB)
			c.Users.Get(p.Receiver).WriteJSON(e)
		}
		return nil
	}
}
