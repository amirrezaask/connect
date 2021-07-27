package main

type EventType uint8

const (
	_ = iota
	EventType_NewMessage
	EventType_FileChunk
	EventType_Handshake_Request
	EventType_Handshake_Response
)

type Event struct {
	Creator   string
	EventType EventType `json:"event_type"`
	Payload   []byte    `json:"payload"`
}

type NewMessagePayload struct {
	Sender   string
	Receiver string
	Body     string
}
