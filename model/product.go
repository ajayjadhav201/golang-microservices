package model

import "time"

type Product struct {
	ID                string
	ProductID         string
	Name              string
	Description       string
	Price             float64
	Quantity          int32
	Category          string
	Brand             string
	Images            string //json marshaled list of strings/ image urls
	Reviews           string
	SellerInformation string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
