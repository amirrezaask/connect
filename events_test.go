package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/amirrezaask/connect/auth"
	"github.com/amirrezaask/connect/domain"
	"github.com/amirrezaask/connect/testutils"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
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
	db, err := testutils.GetDB()
	assert.NoError(t, err)
	assert.NoError(t, testutils.AddUser(db, "user1"))
	assert.NoError(t, testutils.AddUser(db, "user2"))

	assert.NoError(t, testutils.AddHub(db, "hubid"))
	assert.NoError(t, testutils.AddChannel(db, "hubid", "channelid"))
	assert.NoError(t, testutils.AddUserToHub(db, "user1", "hubid"))
	assert.NoError(t, testutils.AddUserToHub(db, "user2", "hubid"))
	assert.NoError(t, testutils.AddUserToChannel(db, "user1", "channelid"))
	assert.NoError(t, testutils.AddUserToChannel(db, "user2", "channelid"))

	assert.NoError(t, testutils.AddRoleForUserInChannel(db, "user1", "channelid", auth.ROLE_CHANNEL_WRITE))
	assert.NoError(t, testutils.AddRoleForUserInChannel(db, "user2", "channelid", auth.ROLE_CHANNEL_WRITE))

	clean := func(db *sql.DB) {

		count1, err := testutils.RemoveRoleForUserInChannel(db, "user1", "channelid", auth.ROLE_CHANNEL_WRITE)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), count1)

		count2, err := testutils.RemoveRoleForUserInChannel(db, "user2", "channelid", auth.ROLE_CHANNEL_WRITE)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), count2)

		count1 = 0
		count2 = 0

		count1, err = testutils.RemoveUserFromChannel(db, "user1", "channelid")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), count1)

		count2, err = testutils.RemoveUserFromChannel(db, "user2", "channelid")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), count2)

		count1 = 0
		count2 = 0

		count1, err = testutils.RemoveUserFromHub(db, "user1", "hubid")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), count1)

		count2, err = testutils.RemoveUserFromHub(db, "user2", "hubid")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), count2)

		countMessage, err := testutils.RemoveMessage(db, "user1", "channelid")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), countMessage)

		assert.NoError(t, testutils.RemoveChannel(db, "channelid"))
		assert.NoError(t, testutils.RemoveHub(db, "hubid"))

		assert.NoError(t, testutils.RemoveUser(db, "user1"))
		assert.NoError(t, testutils.RemoveUser(db, "user2"))

	}
	defer clean(db)
	regiterServers()
	// Create test server with the echo handler.
	s := httptest.NewServer(http.DefaultServeMux)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// first client
	ws1, _, err := websocket.DefaultDialer.Dial(u+"/ws?id=user1&token=something", nil)
	assert.NoError(t, err)

	defer ws1.Close()
	assert.True(t, isConnected(ws1))

	// second client
	ws2, _, err := websocket.DefaultDialer.Dial(u+"/ws?id=user2&token=something", nil)
	assert.NoError(t, err)

	defer ws2.Close()
	assert.True(t, isConnected(ws2))

	bs, _ := json.Marshal(domain.NewMessagePayload{
		Body:      "salam",
		ChannelID: "channelid",
	})
	err = ws1.WriteJSON(domain.Event{
		EventType: domain.EventType_NewMessage,
		Payload:   bs,
	})
	assert.NoError(t, err)

	e := &domain.Event{}
	err = ws2.ReadJSON(e)
	assert.NoError(t, err)

	assert.NotEmpty(t, e.Payload)

	p := &domain.NewMessagePayload{}
	err = json.Unmarshal(e.Payload, p)
	assert.NoError(t, err)
	assert.Equal(t, "salam", p.Body)
}
