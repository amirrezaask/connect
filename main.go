package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/amirrezaask/connect/bus"
	"github.com/amirrezaask/connect/domain"
	"github.com/amirrezaask/connect/handlers"
	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

func setupAPIServer(db *sql.DB) http.Handler {
	e := echo.New()

	hubHandler := handlers.HubHandler{DB: db}
	channelHandler := handlers.ChannelHandler{DB: db}

	e.POST("/hub", hubHandler.CreateHub)
	e.POST("/hub_users", hubHandler.AddUserToHub)
	e.POST("/channel", channelHandler.CreateChannel)

	return e.Server.Handler
}

func setupWSServer(h *handlers.WSHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.WSHandler)
	return mux
}

func regiterServers() {
	l, _ := zap.NewDevelopment()
	logger := l.Sugar()

	b := bus.NewChannelBus()
	uc := handlers.UserConnections{}

	// FIX
	db := &sql.DB{}
	WSHandler := &handlers.WSHandler{
		Users:  uc,
		Logger: logger,
		Bus:    b,
		DB:     db,
	}

	WSServer := setupWSServer(WSHandler)
	apiServer := setupAPIServer(db)

	b.Register(domain.EventType_NewMessage, WSHandler.NewMessageEventHandler())
	http.Handle("/ws", WSServer)
	http.Handle("/api", apiServer)
}

func main() {
	regiterServers()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
