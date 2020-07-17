package model

import "time"

type Todo struct {
	ID          int32
	Title       string
	Description string
	CreatedAt   time.Time
}
