package handlers_test

import (
	"receipt_processor/internal/models"
	"receipt_processor/internal/services"
	"testing"
)

func TestProcessReceiptHandler(t *testing.T) {
	receiptService := services.NewReceiptService()

	tests := []struct {
		name           string
		receipt        models.Receipt
		expectedPoints int
	}{
		{
			name: "Test Case 1",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expectedPoints: 28,
		},
		{
			name: "Test Case 2",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expectedPoints: 109,
		},
		{
			name: "Test Case 3",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "13:13",
				Items: []models.Item{
					{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
				},
				Total: "1.25",
			},
			expectedPoints: 31,
		},
		{
			name: "Test Case 4",
			receipt: models.Receipt{
				Retailer:     "Walgreens",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "08:13",
				Items: []models.Item{
					{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
					{ShortDescription: "Dasani", Price: "1.40"},
				},
				Total: "2.65",
			},
			expectedPoints: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiptID := receiptService.ProcessReceipt(tt.receipt)
			points, exists := receiptService.GetPoints(receiptID)

			if !exists {
				t.Errorf("No receipt found for ID: %s", receiptID)
			} else {
				if points != tt.expectedPoints {
					t.Errorf("Expected points %d, but got %d", tt.expectedPoints, points)
				}
			}
		})
	}
}
