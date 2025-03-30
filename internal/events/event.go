package events

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}
