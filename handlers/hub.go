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
        return ClientErr(ctx, err)
	}
	err = h.Insert(context.TODO(), c.DB, boil.Infer())
	if err != nil {
        return ServerErr(ctx, err)
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
        return ClientErr(ctx, err)
	}
    h := models.Hub{ID: hu.HubID}
    err = h.AddUsers(context.TODO(), c.DB, false, &models.User{ID: hu.UserID})
    if err != nil {
        return ServerErr(ctx, err)
	}
    return OK(ctx, hu)
}

func (c *HubHandler) RemoveUserFromHub(ctx echo.Context) error {
    type HubUser struct {
        UserID string `json:"user_id"`
        HubID string `json:"hub_id"`
    }
    hu := &HubUser{}
    err := ctx.Bind(hu)
    if err != nil {
        return ClientErr(ctx, err)
	}
    h := models.Hub{ID: hu.HubID}
    err = h.RemoveUsers(context.TODO(), c.DB, &models.User{ID: hu.UserID})
    if err != nil {
        return ServerErr(ctx, err)
	}
    return OK(ctx, hu)
}

func (c *HubHandler) RemoveHub(ctx echo.Context) error {
    h := &models.Hub{}
    err := ctx.Bind(h)
    if err != nil {
        return ClientErr(ctx, err)
    }
    _, err = models.Hubs(models.HubWhere.ID.EQ(h.ID)).DeleteAll(ctx.Request().Context(), c.DB)
    if err != nil {
        return ServerErr(ctx, err)
    }
    return OK(ctx)
}

