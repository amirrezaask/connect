package handlers

import (
	"context"
	"database/sql"

	"github.com/amirrezaask/connect/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ChannelHandler struct {
	DB *sql.DB
}

func (c *ChannelHandler) CreateChannel(ctx echo.Context) error {
	h := &models.Channel{}
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
