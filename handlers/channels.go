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
        return ClientErr(ctx, err)
	}
	err = h.Insert(context.TODO(), c.DB, boil.Infer())
	if err != nil {
        return ServerErr(ctx, err)
	}
	return OK(ctx, h)
}
