package model

import "time"

type Agent struct {
	ID          int       `json:"id"`
	Avatar      string    `json:"avatar"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Firepower   int       `json:"firepower"`
	Author      string    `json:"author"`
	WxId        string    `json:"wx_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsFollowed  bool      `json:"is_followed"`
}
