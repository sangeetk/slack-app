package models

import "time"

type User struct {
	ID        string    `json:"id"`
	SlackID   string    `json:"slack_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
