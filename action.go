package bot

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Emote string
type Facing string

const (
	EmoteAngry    Emote = "emoji-angry"
	EmoteThumbsUp Emote = "emoji-thumbsup"
	EmoteHello    Emote = "emoji-hello"
	EmoteTired    Emote = "emoji-tired"
	EmoteMacarena Emote = "emoji-macarena"
)

const (
	FacingFrontRight Facing = "FrontRight"
	FacingFrontLeft  Facing = "FrontLeft"
	FacingBackRight  Facing = "BackRight"
	FacingBackLeft   Facing = "BackLeft"
)

func (b *Bot) SendPublicMessage(msg string) error {
	return b.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"_type": "ChatRequest", "message": "%s"}`, msg)))
}

func (b *Bot) SendPrivateMessage(msg string, userID string) error {
	return b.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"_type": "ChatRequest", "message": "%s", "whisper_target_id": "%s"}`, msg, userID)))
}

func (b *Bot) SendEmote(emotion Emote) error {
	return b.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"_type": "EmoteRequest", "emote_id": "%s"}`, emotion)))
}

func (b *Bot) TeleportUser(userID string, x, y, z float32) error {
	return b.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"_type": "TeleportRequest", "user_id": "%s", "destination": {"x": %.1f, "y": %.1f, "z": %.1f, "facing": "FrontRight"}}`, userID, x, y, z)))
}

func (b *Bot) FloorHit(x, y, z float32, facing Facing) error {
	return b.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"_type": "FloorHitRequest", "destination": {"x": %.1f, "y": %.1f, "z": %.1f, "facing": "%s"}}`, x, y, z, facing)))
}
