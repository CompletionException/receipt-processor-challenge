package api

import (
	"receipt-processor-challenge/internal/business"
	"receipt-processor-challenge/internal/model"
	"testing"

	"github.com/google/uuid"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  model.Receipt
		expected int
	}{
		{
			name: "README.md example 1",
			receipt: model.Receipt{
				ID:       uuid.New(),
				Retailer: "Target",
				Date:     "2022-01-01",
				Time:     "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			name: "README.md example 2",
			receipt: model.Receipt{
				ID:       uuid.New(),
				Retailer: "M&M Corner Market",
				Date:     "2022-03-20",
				Time:     "14:33",
				Items: []model.Item{
					{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					}, {
						ShortDescription: "Gatorade",
						Price:            "2.25",
					}, {
						ShortDescription: "Gatorade",
						Price:            "2.25",
					}, {
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
				},
				Total: "9.00",
			},
			expected: 109,
		},
		{
			name: "Challenge morning receipt",
			receipt: model.Receipt{
				ID:       uuid.New(),
				Retailer: "Walgreens",
				Date:     "2022-01-02",
				Time:     "08:13",
				Total:    "2.65",
				Items: []model.Item{
					{
						ShortDescription: "Pepsi - 12-oz",
						Price:            "1.25",
					}, {
						ShortDescription: "Dasani",
						Price:            "1.40",
					},
				},
			},
			expected: 15,
		},
		{
			name: "Challenge simple receipt",
			receipt: model.Receipt{
				ID:       uuid.New(),
				Retailer: "Target",
				Date:     "2022-01-02",
				Time:     "13:13",
				Total:    "1.25",
				Items: []model.Item{
					{
						ShortDescription: "Pepsi - 12-oz",
						Price:            "1.25",
					},
				},
			},
			expected: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := business.CalculatePoints(tt.receipt, 5)
			if got != tt.expected {
				t.Errorf("CalculatePoints() = %v, want %v", got, tt.expected)
			}
		})
	}
}
