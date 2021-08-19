package handlers

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/amirrezaask/connect/domain"
	"github.com/amirrezaask/connect/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type HubHandler struct {
	DB     *sql.DB
	Users  UserConnections
	Logger *zap.SugaredLogger
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
	creatorCon := c.Users.Get(h.Creator.String)
	if creatorCon != nil {
		payload, _ := json.Marshal(domain.HubCreatedPayload{HubID: h.ID})
		err = creatorCon.WriteJSON(&domain.Event{
			Creator:   h.Creator.String,
			EventType: domain.EventType_HubCreated,
			Payload:   []byte(payload),
		})
		if err != nil {
			c.Logger.Errorf("error in sending event to user: %v", err)
		}
	}
	return OK(ctx, h)
}

func (c *HubHandler) AddUserToHub(ctx echo.Context) error {
	type HubUser struct {
		UserID string `json:"user_id"`
		HubID  string `json:"hub_id"`
	}
	hu := &HubUser{}
	err := ctx.Bind(hu)
	if err != nil {
		return ClientErr(ctx, err)
	}
	h := &models.Hub{ID: hu.HubID}
	err = h.AddUsers(context.TODO(), c.DB, false, &models.User{ID: hu.UserID})
	if err != nil {
		return ServerErr(ctx, err)
	}
	h, err = models.Hubs(models.HubWhere.ID.EQ(h.ID), qm.Load(models.HubRels.Users)).One(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}
	for _, u := range h.R.Users {
		userCon := c.Users.Get(u.ID)
		if userCon != nil {
			err = userCon.WriteJSON(&domain.Event{
				Creator:   hu.UserID,
				EventType: domain.EventType_HubUserAdded,
				Payload:   domain.MakePayload(&domain.HubUserAddedPayload{UserID: hu.UserID, HubID: hu.HubID}),
			})
			if err != nil {
				c.Logger.Errorf("error in publishing event to user: %v", err)
			}
		}
	}
	return OK(ctx, hu)
}

func (c *HubHandler) RemoveUserFromHub(ctx echo.Context) error {
	type HubUser struct {
		UserID string `json:"user_id"`
		HubID  string `json:"hub_id"`
	}
	hu := &HubUser{}
	err := ctx.Bind(hu)
	if err != nil {
		return ClientErr(ctx, err)
	}
	h := &models.Hub{ID: hu.HubID}
	err = h.RemoveUsers(context.TODO(), c.DB, &models.User{ID: hu.UserID})
	if err != nil {
		return ServerErr(ctx, err)
	}
	h, err = models.Hubs(models.HubWhere.ID.EQ(h.ID), qm.Load(models.HubRels.Users)).One(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}
	for _, u := range h.R.Users {
		userCon := c.Users.Get(u.ID)
		if userCon != nil {
			err = userCon.WriteJSON(&domain.Event{
				Creator:   hu.UserID,
				EventType: domain.EventType_HubUserDeleted,
				Payload:   domain.MakePayload(&domain.HubUserDeletedPayload{UserID: hu.UserID, HubID: hu.HubID}),
			})
			if err != nil {
				c.Logger.Errorf("error in publishing event to user: %v", err)
			}
		}
	}
	return OK(ctx, hu)
}

func (c *HubHandler) RemoveHub(ctx echo.Context) error {
	h := &models.Hub{}
	err := ctx.Bind(h)
	if err != nil {
		return ClientErr(ctx, err)
	}
	h, err = models.Hubs(models.HubWhere.ID.EQ(h.ID), qm.Load(models.HubRels.Users)).One(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}
	_, err = models.Hubs(models.HubWhere.ID.EQ(h.ID)).DeleteAll(ctx.Request().Context(), c.DB)
	if err != nil {
		return ServerErr(ctx, err)
	}

	for _, u := range h.R.Users {
		userCon := c.Users.Get(u.ID)
		if userCon != nil {
			err = userCon.WriteJSON(&domain.Event{
				Creator:   h.Creator.String,
				EventType: domain.EventType_HubDeleted,
				Payload:   domain.MakePayload(&domain.HubDeletedPayload{HubID: h.ID}),
			})
			if err != nil {
				c.Logger.Errorf("error in publishing event to user: %v", err)
			}
		}
	}
	return OK(ctx)
}
