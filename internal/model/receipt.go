package model

import "github.com/google/uuid"

// Receipt represents the structure of receipt data
type Receipt struct {
	ID       uuid.UUID `json:"id"`
	Retailer string    `json:"retailer"`
	Date     string    `json:"purchaseDate"`
	Time     string    `json:"purchaseTime"`
	Items    []Item    `json:"items"`
	Total    string    `json:"total"`
	Points   int       `json:"points"`
}

// Item represents an individual item in a receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
