package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/amirrezaask/connect/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"go.uber.org/zap"
)

func OK(ctx echo.Context, data interface{}) error {
    return ctx.JSON(http.StatusOK, data)
}

func (c *ConnectServer) CreateHub(ctx echo.Context) error {
	h := &models.Hub{}
	err := ctx.Bind(h)
	if err != nil {
		return ctx.String(400, err.Error())
	}
	err = h.Insert(context.TODO(), c.DB, boil.Infer())
	if err != nil {
		return ctx.String(400, err.Error())
	}
    return OK(ctx, h)
}

func (c *ConnectServer) CreateChannel(ctx echo.Context) error {
	h := &models.Channel{}
	err := ctx.Bind(h)
	if err != nil {
		return ctx.String(400, err.Error())
	}
	err = h.Insert(context.TODO(), c.DB, boil.Infer())
	if err != nil {
		return ctx.String(400, err.Error())
	}
    return OK(ctx, h)
}

func (c *ConnectServer) AddUserToHub(ctx echo.Context) error {
    type HubUser struct {
        UserID string `json:"user_id"`
        HubID string `json:"hub_id"`
    }
    hu := &HubUser{}
    err := ctx.Bind(hu)
    if err != nil {
		return ctx.String(400, err.Error())
	}
    h := models.Hub{ID: hu.HubID}
    err = h.AddUsers(context.TODO(), c.DB, false)
    if err != nil {
		return ctx.String(400, err.Error())
	}
    return OK(ctx, hu)
}

func setupAPIServer(c *ConnectServer) http.Handler {
	e := echo.New()
	e.POST("/hub", c.CreateHub)
    e.POST("/hub_users", c.AddUserToHub)
    e.POST("/channel", c.CreateChannel)
	return e.Server.Handler
}

func setupWSServer(c *ConnectServer) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", c.WSHandler)
	return mux
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

	WSServer := setupWSServer(c)
	apiServer := setupAPIServer(c)

	http.Handle("/ws", WSServer)
	http.Handle("/api", apiServer)
	err := http.ListenAndServe(":8080", nil)
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
			c.Users.Get(p.Receiver).WriteJSON(e)
		}
		return nil
	}
}
