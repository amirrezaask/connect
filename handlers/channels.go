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

func (c *ChannelHandler) RemoveChannel(ctx echo.Context) error {
	h := &models.Channel{}
	err := ctx.Bind(h)
	if err != nil {
		return ClientErr(ctx, err)
	}
	_, err = models.Channels(models.ChannelWhere.ID.EQ(h.ID)).DeleteAll(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}
	return OK(ctx)

}

func (c *ChannelHandler) AddUserToChannel(ctx echo.Context) error {
	type ChannelUser struct {
		UserID    string `json:"user_id"`
		ChannelID string `json:"channel_id"`
	}
	cu := &ChannelUser{}
	err := ctx.Bind(cu)
	if err != nil {
		return ClientErr(ctx, err)
	}
	channel := &models.Channel{ID: cu.ChannelID}
	err = channel.AddUsers(ctx.Request().Context(), c.DB, false, &models.User{ID: cu.UserID})
	if err != nil {
		return ServerErr(ctx, err)
	}
	return OK(ctx)
}

func (c *ChannelHandler) RemoveUserFromChannel(ctx echo.Context) error {
	type ChannelUser struct {
		UserID    string `json:"user_id"`
		ChannelID string `json:"channel_id"`
	}
	cu := &ChannelUser{}
	err := ctx.Bind(cu)
	if err != nil {
		return ClientErr(ctx, err)
	}
	channel := &models.Channel{ID: cu.ChannelID}
	err = channel.RemoveUsers(ctx.Request().Context(), c.DB, &models.User{ID: cu.UserID})
	if err != nil {
		return ServerErr(ctx, err)
	}
	return OK(ctx)
}
