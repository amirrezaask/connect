package main

import (
	"context"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/amirrezaask/connect/handlers"
	"github.com/amirrezaask/connect/models"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupReq(body io.Reader) (*http.Request, *httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	return req, rec, echo.New().NewContext(req, rec)
}

func setupHubHandler(t *testing.T) *handlers.HubHandler {
	db, err := sql.Open("postgres", "user=connect password=connect dbname=connect sslmode=disable")
	assert.NoError(t, err)
	err = db.Ping()
	assert.NoError(t, err)
	hubHandler := &handlers.HubHandler{DB: db}
	return hubHandler
}

func TestCreateHub(t *testing.T) {
	// Do the thing
	hubHandler := setupHubHandler(t)
	req, rec, ctx := setupReq(strings.NewReader(`{"id": "myid", "name": "handlerTesthub"}`))
	_ = req
	err := hubHandler.CreateHub(ctx)
	assert.NoError(t, err)
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
	req, rec, ctx := setupReq(strings.NewReader(`{"id": "myid"}`))
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
