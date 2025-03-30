package model

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	location    string
	DateTime    time.Time
	UserID      int
}
