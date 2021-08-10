package auth

import (
	"context"
	"database/sql"
	"testing"

	"github.com/amirrezaask/connect/models"
	"github.com/stretchr/testify/assert"
)

func TestHasRoleInHub(t *testing.T) {
	clean := func(db *sql.DB) error {
		return nil
	}
	db, err := testgetDB()
	assert.NoError(t, err)
	defer clean(db)
	// Creating basic things
	assert.NoError(t, (&models.Hub{ID: "hubid"}).AddUsers(context.Background(), db, false, &models.User{ID: "userid"}))
}

func TestHasRoleInChannel(t *testing.T) {}
