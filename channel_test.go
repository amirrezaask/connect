package main

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"testing"

	"github.com/amirrezaask/connect/handlers"
	"github.com/amirrezaask/connect/models"
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
	assert.NoError(t, addHub(channelHandler.DB, "hubid"))
	_, rec, ctx := setupReq(strings.NewReader(`{"id": "myid", "name": "channelTestText", "type": "text", "hub_id": "hubid"}`))
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
	assert.NoError(t, removeHub(db, "hubid"))
	assert.NoError(t, err)

}
func TestRemoveChannel(t *testing.T) {
	channelHandler := setupChannelHandler(t)
	// Adding a channel
	db := channelHandler.DB
	assert.NoError(t, addHub(db, "hubid"))
	channel := &models.Channel{
		ID:    "myid",
		HubID: "hubid",
	}

	assert.NoError(t, channel.Insert(context.TODO(), db, boil.Infer()))

	// Removing using API
	req, rec, ctx := setupReq(strings.NewReader(`{"id": "myid"}`))
	_ = req
	err := channelHandler.RemoveChannel(ctx)
	assert.NoError(t, removeHub(db, "hubid"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// check DB if it's removed
	row := db.QueryRow(`SELECT COUNT(id) FROM channels WHERE id='myid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 0, count)

}
func removeChannel(db *sql.DB, channelID string) error {
	_, err := models.Channels(models.ChannelWhere.ID.EQ(channelID)).DeleteAll(context.Background(), db)
	return err
}

func addChannel(db *sql.DB, hubID, channelID string) error {
	return (&models.Channel{ID: channelID, HubID: hubID}).Insert(context.TODO(), db, boil.Infer())
}

func TestAddUserToChannel(t *testing.T) {
	clean := func(db *sql.DB, t *testing.T) {
		_, err := db.Exec(`DELETE FROM channel_users WHERE user_id='userid' AND channel_id='channelid'`)
		assert.NoError(t, err)
		assert.NoError(t, removeChannel(db, "channelid"))
		assert.NoError(t, removeHub(db, "hubid"))
		assert.NoError(t, removeUser(db, "userid"))
	}
	channelHandler := setupChannelHandler(t)
	db := channelHandler.DB
	defer clean(channelHandler.DB, t)
	assert.NoError(t, addHub(channelHandler.DB, "hubid"))
	assert.NoError(t, addUser(channelHandler.DB, "userid"))
	assert.NoError(t, addChannel(channelHandler.DB, "hubid", "channelid"))

	// Do the thing
	req, rec, ctx := setupReq(strings.NewReader(`{"user_id": "userid", "channel_id": "channelid"}`))
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

func addUserToChannel(db *sql.DB, userID string, channelID string) error {
	_, err := db.Exec(`INSERT INTO channel_users (user_id, channel_id) VALUES ($1, $2)`, userID, channelID)
	return err
}
func TestRemoveUserFromChannel(t *testing.T) {
	channelHandler := setupChannelHandler(t)
	db := channelHandler.DB
	clean := func(db *sql.DB, t *testing.T) {
		_, err := db.Exec(`DELETE FROM channel_users WHERE user_id='userid' AND channel_id='channelid'`)
		assert.NoError(t, err)
		assert.NoError(t, removeChannel(db, "channelid"))
		assert.NoError(t, removeHub(db, "hubid"))
		assert.NoError(t, removeUser(db, "userid"))
	}
	defer clean(db, t)

	assert.NoError(t, addUser(db, "userid"))
	assert.NoError(t, addHub(db, "hubid"))
	assert.NoError(t, addChannel(db, "hubid", "channelid"))
	assert.NoError(t, addUserToChannel(db, "userid", "channelid"))

	req, rec, ctx := setupReq(strings.NewReader(`{"user_id": "userid", "channel_id": "channelid"}`))
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
