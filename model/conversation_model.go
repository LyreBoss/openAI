package model

import "time"

type Conversation struct {
	ID               int       `json:"id"`
	ConversationName string    `json:"conversation_name"`
	Avatar           string    `json:"avatar"`
	WxID             string    `json:"wx_id"`
	ConversationID   string    `json:"conversation_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	IsPinned         bool      `json:"is_pinned"`
	Description      string    `json:"description"`
}
