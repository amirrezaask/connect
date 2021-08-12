package main

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"testing"

	"github.com/amirrezaask/connect/handlers"
	"github.com/amirrezaask/connect/models"
	"github.com/amirrezaask/connect/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupChannelHandler(t *testing.T) *handlers.ChannelHandler {
	db, err := sql.Open("postgres", "user=connect password=connect dbname=connect sslmode=disable")
	assert.NoError(t, err)
	err = db.Ping()
	assert.NoError(t, err)
	channelHandler := &handlers.ChannelHandler{DB: db}
	return channelHandler
}

func TestCreateChannel(t *testing.T) {
	channelHandler := setupChannelHandler(t)
	assert.NoError(t, testutils.AddHub(channelHandler.DB, "hubid"))
	_, rec, ctx := testutils.MakeRequest(http.MethodPost, strings.NewReader(`{"id": "myid", "name": "channelTestText", "type": "text", "hub_id": "hubid"}`))
	err := channelHandler.CreateChannel(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// CheckDB
	db := channelHandler.DB
	row := db.QueryRow(`SELECT COUNT(id) FROM channels WHERE id='myid' AND name='channelTestText' AND "type"='text' AND hub_id='hubid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 1, count)

	// clean up
	_, err = models.Channels(models.ChannelWhere.ID.EQ("myid")).DeleteAll(context.TODO(), channelHandler.DB)
	assert.NoError(t, testutils.RemoveHub(db, "hubid"))
	assert.NoError(t, err)

}
func TestRemoveChannel(t *testing.T) {
	channelHandler := setupChannelHandler(t)
	// Adding a channel
	db := channelHandler.DB
	channel := &models.Channel{
		ID:    "myid",
		HubID: "hubid",
	}
	assert.NoError(t, testutils.AddHub(db, "hubid"))

	assert.NoError(t, channel.Insert(context.TODO(), db, boil.Infer()))

	// Removing using API
	req, rec, ctx := testutils.MakeRequest(http.MethodPost, strings.NewReader(`{"id": "myid"}`))
	_ = req
	err := channelHandler.RemoveChannel(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// check DB if it's removed
	row := db.QueryRow(`SELECT COUNT(id) FROM channels WHERE id='myid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 0, count)

	// clean
	assert.NoError(t, testutils.RemoveHub(db, "hubid"))
}

func TestAddUserToChannel(t *testing.T) {
	clean := func(db *sql.DB, t *testing.T) {
		_, err := db.Exec(`DELETE FROM channel_users WHERE user_id='userid' AND channel_id='channelid'`)
		assert.NoError(t, err)
		assert.NoError(t, testutils.RemoveChannel(db, "channelid"))
		assert.NoError(t, testutils.RemoveHub(db, "hubid"))
		assert.NoError(t, testutils.RemoveUser(db, "userid"))
	}
	channelHandler := setupChannelHandler(t)
	db := channelHandler.DB
	defer clean(channelHandler.DB, t)
	assert.NoError(t, testutils.AddHub(channelHandler.DB, "hubid"))
	assert.NoError(t, testutils.AddUser(channelHandler.DB, "userid"))
	assert.NoError(t, testutils.AddChannel(channelHandler.DB, "hubid", "channelid"))

	// Do the thing
	req, rec, ctx := testutils.MakeRequest(http.MethodPost, strings.NewReader(`{"user_id": "userid", "channel_id": "channelid"}`))
	_ = req
	err := channelHandler.AddUserToChannel(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// 	// CheckDB
	row := db.QueryRow(`SELECT COUNT(user_id) FROM channel_users WHERE user_id='userid' AND channel_id='channelid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 1, count)
}
func TestRemoveUserFromChannel(t *testing.T) {
	channelHandler := setupChannelHandler(t)
	db := channelHandler.DB
	clean := func(db *sql.DB, t *testing.T) {
		_, err := db.Exec(`DELETE FROM channel_users WHERE user_id='userid' AND channel_id='channelid'`)
		assert.NoError(t, err)
		assert.NoError(t, testutils.RemoveChannel(db, "channelid"))
		assert.NoError(t, testutils.RemoveHub(db, "hubid"))
		assert.NoError(t, testutils.RemoveUser(db, "userid"))
	}
	defer clean(db, t)

	assert.NoError(t, testutils.AddUser(db, "userid"))
	assert.NoError(t, testutils.AddHub(db, "hubid"))
	assert.NoError(t, testutils.AddChannel(db, "hubid", "channelid"))
	assert.NoError(t, testutils.AddUserToChannel(db, "userid", "channelid"))

	req, rec, ctx := testutils.MakeRequest(http.MethodPost, strings.NewReader(`{"user_id": "userid", "channel_id": "channelid"}`))
	_ = req
	err := channelHandler.RemoveUserFromChannel(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	row := db.QueryRow(`SELECT COUNT(user_id) FROM channel_users WHERE user_id='userid' AND channel_id='channelid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 0, count)

}
