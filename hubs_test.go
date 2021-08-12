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
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupHubHandler(t *testing.T) *handlers.HubHandler {
	db, err := testutils.GetDB()
	assert.NoError(t, err)
	err = db.Ping()
	assert.NoError(t, err)
	hubHandler := &handlers.HubHandler{DB: db}
	return hubHandler
}

func TestCreateHub(t *testing.T) {
	// Do the thing
	hubHandler := setupHubHandler(t)
	req, rec, ctx := testutils.MakeRequest(http.MethodPost, strings.NewReader(`{"id": "myid", "name": "handlerTesthub"}`))
	_ = req
	err := hubHandler.CreateHub(ctx)
	assert.NoError(t, err)
	// assert.Equal(t, "", string(rec.Body.Bytes()))
	assert.Equal(t, http.StatusOK, rec.Code)

	// CheckDB
	db := hubHandler.DB
	row := db.QueryRow(`SELECT COUNT(id) FROM hubs WHERE id='myid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 1, count)

	// clean up
	_, err = models.Hubs(models.HubWhere.ID.EQ("myid")).DeleteAll(context.TODO(), hubHandler.DB)
	assert.NoError(t, err)
}

func TestRemoveHub(t *testing.T) {
	hubHandler := setupHubHandler(t)
	// Adding a hub
	db := hubHandler.DB
	hub := &models.Hub{
		ID: "myid",
	}

	assert.NoError(t, hub.Insert(context.TODO(), db, boil.Infer()))

	// Removing using API
	req, rec, ctx := testutils.MakeRequest(http.MethodPost, strings.NewReader(`{"id": "myid"}`))
	_ = req
	err := hubHandler.RemoveHub(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// check DB if it's removed
	row := db.QueryRow(`SELECT COUNT(id) FROM hubs WHERE id='myid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 0, count)

}

func TestAddUserToHub(t *testing.T) {
	clean := func(db *sql.DB, t *testing.T) {
		_, err := db.Exec(`DELETE FROM hub_users WHERE user_id='userid' AND hub_id='hubid'`)
		assert.NoError(t, err)
		assert.NoError(t, testutils.RemoveHub(db, "hubid"))
		assert.NoError(t, testutils.RemoveUser(db, "userid"))
	}
	hubHandler := setupHubHandler(t)
	db := hubHandler.DB
	defer clean(hubHandler.DB, t)
	clean(hubHandler.DB, t)
	assert.NoError(t, testutils.AddUser(hubHandler.DB, "userid"))
	assert.NoError(t, testutils.AddHub(hubHandler.DB, "hubid"))

	// Do the thing
	req, rec, ctx := testutils.MakeRequest(http.MethodPost, strings.NewReader(`{"user_id": "userid", "hub_id": "hubid"}`))
	_ = req
	err := hubHandler.AddUserToHub(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// 	// CheckDB
	row := db.QueryRow(`SELECT COUNT(user_id) FROM hub_users WHERE user_id='userid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 1, count)
}

func TestRemoveUserFromHub(t *testing.T) {
	hubHandler := setupHubHandler(t)
	db := hubHandler.DB
	clean := func(db *sql.DB, t *testing.T) {
	}
	defer clean(db, t)

	assert.NoError(t, testutils.AddUser(db, "userid"))
	assert.NoError(t, testutils.AddHub(db, "hubid"))
	assert.NoError(t, testutils.AddUserToHub(db, "userid", "hubid"))

	req, rec, ctx := testutils.MakeRequest(http.MethodPost, strings.NewReader(`{"user_id": "userid", "hub_id": "hubid"}`))
	_ = req
	err := hubHandler.RemoveUserFromHub(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	row := db.QueryRow(`SELECT COUNT(user_id) FROM hub_users WHERE user_id='userid'`)
	assert.NoError(t, row.Err())
	var count int
	assert.NoError(t, row.Scan(&count))
	assert.Equal(t, 0, count)
}
