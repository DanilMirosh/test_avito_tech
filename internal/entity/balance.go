package entity

import "github.com/google/uuid"

type Balance struct {
	UserID  uuid.UUID `json:"user_uuid"`
	Balance int64     `json:"balance"`
}
