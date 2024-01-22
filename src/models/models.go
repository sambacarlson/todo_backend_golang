package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Heading    string    `json:"Heading"`
	Body       string    `json:"body"`
	CategoryID string    `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Category struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Reminder struct {
	ID            uuid.UUID     `json:"id"`
	TodoID        uuid.UUID     `json:"todo_id"`
	Snooze        bool          `json:"snooze"`
	SnoozeTimeout time.Duration `json:"snooze_timeout"`
	RemindTime    time.Time     `json:"remind_time"`
	Description   string        `json:"description"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
