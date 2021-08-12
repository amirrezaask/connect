# Connect
A chat application with (hopefully)
- Text chat
- support various file types
- voice chat
- video chat

# Terminology
- *Hub* => A hub is like a Slack workspace or Discord server. A user can have different nickname, profile avatar and customization settings for each hub.
- *Channel* => A hub can have multiple channels, channels have various types ( voice/video, text ).

# Architecture
All WebSocket events that are received in the server are emitted into the Bus ( can be Go channels, NATS, ... ) and then from there the registered handler for the given event type will handle the event.
