package models

import "time"

type Calendar struct {
	Id        int64
	EventDate time.Time
	Effective time.Time
	Status    string
	Notes     string
}
