package models

import "time"

type Todo struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Heading    string    `json:"Heading"`
	Body       string    `json:"body"`
	CategoryID string    `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Category struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Reminder struct {
	ID          int       `json:"id"`
	TodoID      int       `json:"todo_id"`
	RemindTime  time.Time `json:"remind_time"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
