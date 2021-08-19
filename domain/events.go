package domain

import "encoding/json"

type EventType string

const (
	EventType_NewMessage         = "new_message"
	EventType_FileChunk          = "file_chunk"
	EventType_HubCreated         = "hub_created"
	EventType_HubUserAdded       = "hub_user_added"
	EventType_HubUserDeleted     = "hub_user_deleted"
	EventType_HubDeleted         = "hub_deleted"
	EventType_ChannelCreated     = "channel_created"
	EventType_ChannelDeleted     = "channel_deleted"
	EventType_ChanenlUserAdded   = "channel_user_added"
	EventType_ChannelUserDeleted = "channel_user_deleted"
)

type Event struct {
	Creator   string
	EventType EventType `json:"event_type"`
	Payload   []byte    `json:"payload"`
}

type NewMessagePayload struct {
	ChannelID string `json:"channel_id"`
	Body      string `json:"body"`
}

type HubCreatedPayload struct {
	HubID string `json:"hub_id"`
}

type HubDeletedPayload struct {
	HubID string `json:"hub_id"`
}
type HubUserAddedPayload struct {
	HubID  string `json:"hub_id"`
	UserID string `json:"user_id"`
}
type HubUserDeletedPayload struct {
	HubID  string `json:"hub_id"`
	UserID string `json:"user_id"`
}
type ChannelCreatedPayload struct {
	HubID     string `json:"hub_id"`
	ChannelID string `json:"channel_id"`
}
type ChannelDeletedPayload struct {
	HubID     string `json:"hub_id"`
	ChannelID string `json:"channel_id"`
}
type ChannelUserAddedPayload struct {
	ChannelID string `json:"channel_id"`
	UserID    string `json:"user_id"`
}
type ChannelUserDeletedPayload struct {
	ChannelID string `json:"channel_id"`
	UserID    string `json:"user_id"`
}

func MakePayload(v interface{}) []byte {
	bs, _ := json.Marshal(v)
	return bs
}
