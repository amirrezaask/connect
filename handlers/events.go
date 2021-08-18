package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/amirrezaask/connect/auth"
	"github.com/amirrezaask/connect/bus"
	"github.com/amirrezaask/connect/domain"
	"github.com/amirrezaask/connect/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type EventsHandler struct {
	Users  UserConnections
	Logger *zap.SugaredLogger
	Bus    bus.Bus
	DB     *sql.DB
}

type UserConnections map[string]*websocket.Conn

func (uc UserConnections) Add(id string, conn *websocket.Conn) {
	uc[id] = conn
}

func (uc UserConnections) Get(id string) *websocket.Conn {
	return uc[id]
}

func (s *EventsHandler) WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Logger.Errorf("error in upgrading user connection to ws protocol: %v", err)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		s.Logger.Error("user has not send a nickname")
		return
	}
	token := r.URL.Query().Get("token")
	if token == "" {
		s.Logger.Error("user has not send a token")
		return
	}
	exists, err := models.Users(models.UserWhere.ID.EQ(id)).Exists(r.Context(), s.DB)
	if err != nil {
		s.Logger.Errorf("error in looking up user: %v", err)
		return
	}
	if !exists {
		s.Logger.Error("No user with given ID <%s> found", id)
		return
	}
	//TODO(amirreza): Check the token here
	s.Logger.Debugf("%s connected", id)
	s.Users.Add(id, conn)
	go func(c *websocket.Conn, nickName string) {
		for {
			e := &domain.Event{}
			err = c.ReadJSON(e)
			if err != nil {
				//TODO(amirreza): please fix this:))
				if !strings.Contains(err.Error(), "EOF") {
					s.Logger.Errorf("error in reading event from client: %v", err)
				}
				break
			}
			e.Creator = nickName
			s.Logger.Debugf("received from %s: %+v", nickName, e)
			s.Bus.Emit(e)
		}
	}(conn, id)
	conn.WriteMessage(websocket.TextMessage, []byte("Connected"))
}

func (c *EventsHandler) NewMessageEventHandler() func(e *domain.Event) error {
	return func(e *domain.Event) error {
		if e.EventType == domain.EventType_NewMessage {
			p := &domain.NewMessagePayload{}
			err := json.Unmarshal(e.Payload, p)
			if err != nil {
				c.Logger.Errorf("error in unmarshaling new message payload: %v", err)
				return err
			}
			rm := auth.RoleManager{DB: c.DB}
			has, err := rm.HasRoleInChannel(e.Creator, p.ChannelID, auth.ROLE_CHANNEL_WRITE)
			if err != nil {
				c.Logger.Error(err)
				return err
			}
			if !has {
				c.Logger.Debugf("user %s has no write permission in %s channel", e.Creator, p.ChannelID)
				return auth.Unauthorized(e.Creator, p.ChannelID)
			}
			// save message into database
			go func(db *sql.DB, logger *zap.SugaredLogger) {
				id := uuid.New().String()
				if id == "" {
					logger.Error("cannot create UUID")
					return
				}
				(&models.Message{
					ID:        id,
					UserID:    e.Creator,
					ChannelID: p.ChannelID,
					Payload:   string(e.Payload),
				}).Insert(context.Background(), db, boil.Infer())
			}(c.DB, c.Logger)

			channel, err := models.Channels(models.ChannelWhere.ID.EQ(p.ChannelID),
				qm.Load(models.ChannelRels.Users),
			).One(context.TODO(), c.DB)
			if err != nil {
				return err
			}
			for _, u := range channel.R.Users {
				userC := c.Users.Get(u.ID)
				if userC == nil {
					c.Logger.Errorf("error, user has no active connection: %s", u.ID)
					continue
				}
				if err := userC.WriteJSON(e); err != nil {
					c.Logger.Errorf("error in pushing message to user: %v", err)
					continue
				}
			}
		}
		return nil
	}
}
