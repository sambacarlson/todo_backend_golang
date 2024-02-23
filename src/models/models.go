package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID         uuid.UUID `json:"id" db:"id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	Heading    string    `json:"heading" db:"heading"`
	Body       string    `json:"body" db:"body"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type Category struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Reminder struct {
	ID            uuid.UUID     `json:"id" db:"id"`
	TodoID        uuid.UUID     `json:"todo_id" db:"todo_id"`
	Snooze        bool          `json:"snooze" db:"snooze"`
	SnoozeTimeout time.Duration `json:"snooze_timeout" db:"snooze_timeout"`
	RemindTime    time.Time     `json:"remind_time" db:"remind_time"`
	Description   string        `json:"description" db:"description"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`
}

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
