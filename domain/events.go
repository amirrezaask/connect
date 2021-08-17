package domain

type EventType string

const (
	EventType_NewMessage         = "new_message"
	EventType_FileChunk          = "file_chunk"
	EventType_Handshake_Request  = "handshake_req"
	EventType_Handshake_Response = "handshake_res"
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
