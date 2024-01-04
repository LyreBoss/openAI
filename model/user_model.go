package model

import "time"

type User struct {
	ID          int       `json:"id"`
	WxID        string    `json:"wxId"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsTop       bool      `json:"is_top"`
}
