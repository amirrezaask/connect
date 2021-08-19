package handlers

import (
	"context"
	"database/sql"

	"github.com/amirrezaask/connect/domain"
	"github.com/amirrezaask/connect/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type ChannelHandler struct {
	Users  UserConnections
	DB     *sql.DB
	Logger *zap.SugaredLogger
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
	channel, err := models.Channels(models.ChannelWhere.ID.EQ(h.ID), qm.Load(qm.Rels(models.ChannelRels.Hub, models.HubRels.Users))).One(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}
	for _, u := range channel.R.Users {
		userCon := c.Users.Get(u.ID)
		if userCon != nil {
			err = userCon.WriteJSON(&domain.Event{
				EventType: domain.EventType_ChannelCreated,
				Payload:   domain.MakePayload(&domain.ChannelCreatedPayload{HubID: h.HubID, ChannelID: channel.ID}),
			})
			if err != nil {
				c.Logger.Errorf("error in publishing event to user: %v", err)
			}
		}
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
	channel, err := models.Channels(models.ChannelWhere.ID.EQ(h.ID), qm.Load(qm.Rels(models.ChannelRels.Hub, models.HubRels.Users))).One(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}
	for _, u := range channel.R.Users {
		userCon := c.Users.Get(u.ID)
		if userCon != nil {
			err = userCon.WriteJSON(&domain.Event{
				EventType: domain.EventType_ChannelDeleted,
				Payload:   domain.MakePayload(&domain.ChannelDeletedPayload{HubID: h.HubID, ChannelID: channel.ID}),
			})
			if err != nil {
				c.Logger.Errorf("error in publishing event to user: %v", err)
			}
		}
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
	channel, err = models.Channels(models.ChannelWhere.ID.EQ(cu.ChannelID), qm.Load(qm.Rels(models.ChannelRels.Hub, models.HubRels.Users))).One(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}
	for _, u := range channel.R.Users {
		userCon := c.Users.Get(u.ID)
		if userCon != nil {
			err = userCon.WriteJSON(&domain.Event{
				EventType: domain.EventType_ChanenlUserAdded,
				Payload:   domain.MakePayload(&domain.ChannelUserAddedPayload{ChannelID: channel.ID, UserID: cu.UserID}),
			})
			if err != nil {
				c.Logger.Errorf("error in publishing event to user: %v", err)
			}
		}
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
	channel, err = models.Channels(models.ChannelWhere.ID.EQ(cu.ChannelID), qm.Load(qm.Rels(models.ChannelRels.Hub, models.HubRels.Users))).One(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}
	for _, u := range channel.R.Users {
		userCon := c.Users.Get(u.ID)
		if userCon != nil {
			err = userCon.WriteJSON(&domain.Event{
				EventType: domain.EventType_ChannelUserDeleted,
				Payload:   domain.MakePayload(&domain.ChannelUserDeletedPayload{ChannelID: channel.ID, UserID: cu.UserID}),
			})
			if err != nil {
				c.Logger.Errorf("error in publishing event to user: %v", err)
			}
		}
	}

	return OK(ctx)
}
