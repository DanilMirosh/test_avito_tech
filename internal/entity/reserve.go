package entity

import "github.com/google/uuid"

type Reserve struct {
	UserID  uuid.UUID `json:"user_uuid"`
	Reserve int64     `json:"reserve"`
}
