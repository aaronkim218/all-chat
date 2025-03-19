package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID `json:"id" db:"id"`
	RoomId    uuid.UUID `json:"room_id" db:"room_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Author    uuid.UUID `json:"author" db:"author"`
	Content   string    `json:"content" db:"content"`
}
