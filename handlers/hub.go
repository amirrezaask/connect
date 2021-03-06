package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/amirrezaask/connect/domain"
	"github.com/amirrezaask/connect/models"
	"github.com/golobby/sql/builder"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type HubHandler struct {
	DB     *sql.DB
	Users  UserConnections
	Logger *zap.SugaredLogger
}
type Hub struct {
	ID      string  `bind:"id" json:"id" toml:"id" yaml:"id"`
	Name    *string `bind:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Creator *string `bind:"creator" json:"creator,omitempty" toml:"creator" yaml:"creator,omitempty"`
}

func valueOrEmptyString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

var hubsMeta = builder.ObjectMetadataFrom(&Hub{})

func (c *HubHandler) CreateHub(ctx echo.Context) error {
	h := &Hub{}
	err := ctx.Bind(h)
	if err != nil {
		return ClientErr(ctx, err)
	}

	_, err = builder.
		NewInsert(hubsMeta.Table).
		Into(hubsMeta.Columns...).
		Values(h.ID, h.Name, h.Creator).
		ExecContext(ctx.Request().Context(), c.DB)

	if err != nil {
		log.Printf("error in creating hub: %v", err)
		return ServerErr(ctx, err)
	}

	creatorCon := c.Users.Get(valueOrEmptyString(h.Creator))
	if creatorCon != nil {
		payload, _ := json.Marshal(domain.HubCreatedPayload{HubID: h.ID})
		err = creatorCon.WriteJSON(&domain.Event{
			Creator:   valueOrEmptyString(h.Creator),
			EventType: domain.EventType_HubCreated,
			Payload:   []byte(payload),
		})
		if err != nil {
			c.Logger.Errorf("error in sending event to user: %v", err)
		}
	}
	return OK(ctx, h)
}

func (c *HubHandler) GetHub(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ClientErr(ctx, fmt.Errorf("no id parameter"))
	}
	hub := &Hub{}
	err := builder.NewQuery().Table(hubsMeta.Table).Where(builder.WhereHelpers.EqualID("$1")).Query().BindContext(ctx.Request().Context(), c.DB, hub, id)
	if err != nil {
		return ServerErr(ctx, err)
	}

	return OK(ctx, hub)
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
