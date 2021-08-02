package handlers

import (
	"context"
	"database/sql"

	"github.com/amirrezaask/connect/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type HubHandler struct {
	DB *sql.DB
}

func (c *HubHandler) CreateHub(ctx echo.Context) error {
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

func (c *HubHandler) AddUserToHub(ctx echo.Context) error {
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


