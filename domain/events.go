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
	ChannelID string    `json:"channel_id"`
	Payload   []byte    `json:"payload"`
}

type NewMessagePayload struct {
	Body string `json:"body"`
}
