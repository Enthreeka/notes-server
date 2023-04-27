package entity

import "time"

type Notes struct {
	ID        int       `json:"id"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}
