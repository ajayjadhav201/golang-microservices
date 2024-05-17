package model

import "time"

type Review struct {
	ID          string
	SubmittedBy string
	Rating      int
	Message     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
