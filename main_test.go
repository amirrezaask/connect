package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/amirrezaask/connect/domain"
	"github.com/gorilla/websocket"
)

func isConnected(ws *websocket.Conn) bool {
	_, data, err := ws.ReadMessage()
	if err != nil {
		return false
	}

	if string(data) != "Connected" {
		return false
	}
	return true
}

func TestPerson2PersonMessage(t *testing.T) {
	regiterServers()
	// Create test server with the echo handler.
	s := httptest.NewServer(http.DefaultServeMux)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// first client
	ws1, _, err := websocket.DefaultDialer.Dial(u+"/ws?nickname=user1", nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws1.Close()
	if !isConnected(ws1) {
		t.Fatalf("connection is not ok for ws1 since we can't read connected message from stream")
	}

	// second client
	ws2, _, err := websocket.DefaultDialer.Dial(u+"/ws?nickname=user2", nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws2.Close()
	if !isConnected(ws2) {
		t.Fatalf("connection is not ok for ws2 since we can't read connected message from stream")
	}
	bs, _ := json.Marshal(domain.NewMessagePayload{
		Sender:   "user1",
		Receiver: "user2",
		Body:     "salam",
	})
	err = ws1.WriteJSON(domain.Event{
		EventType: domain.EventType_NewMessage,
		Payload:   bs,
	})
	if err != nil {
		t.Fatalf("cannot write salam message: %v", err)
	}
	e := &domain.Event{}
	err = ws2.ReadJSON(e)
	if err != nil {
		t.Fatalf("cannot read salam message: %v", err)
	}
	t.Logf("%+v", e)
	if len(e.Payload) == 0 {
		t.Fatal("cannot read salam message payload")
	}
	p := &domain.NewMessagePayload{}
	err = json.Unmarshal(e.Payload, p)
	if err != nil {
		t.Fatalf("cannot read salam message payload: %v", err)
	}
	if p.Body != "salam" {
		t.Fatalf("message is not salam")
	}
}
