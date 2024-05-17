package model

import "time"

type Order struct {
	ID        string
	OrderId   string
	OrderedBy string //this wii be users Id
	CreatedAt time.Time
	UpdatedAt time.Time
}
