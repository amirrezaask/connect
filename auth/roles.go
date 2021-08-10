package auth

import (
	"context"
	"database/sql"

	"github.com/amirrezaask/connect/models"
)

type RoleManager struct {
	DB *sql.DB
}

func (r *RoleManager) HasRoleInHub(userID string, hubID string, role string) (bool, error) {
	return models.
		HubPermissions(models.HubPermissionWhere.UserID.EQ(userID), models.HubPermissionWhere.HubID.EQ(hubID), models.HubPermissionWhere.RoleName.EQ(role)).
		Exists(context.TODO(), r.DB)
}

func (r *RoleManager) HasRoleInChannel(userID string, channelID string, role string) (bool, error) {
	return models.
		ChannelPermissions(models.ChannelPermissionWhere.UserID.EQ(userID), models.ChannelPermissionWhere.ChannelID.EQ(channelID), models.ChannelPermissionWhere.RoleName.EQ(role)).
		Exists(context.TODO(), r.DB)
}
