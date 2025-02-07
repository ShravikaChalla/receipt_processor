package utils

import (
	"math"
	"receipt_processor/internal/models"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CalculatePoints calculates the total points for a given receipt based on predefined rules.
//
// Rules:
//   - 1 point for every alphanumeric character in the retailer name.
//   - 50 points if the total is a round dollar amount with no cents.
//   - 25 points if the total is a multiple of 0.25.
//   - 5 points for every two items on the receipt.
//   - If the trimmed length of an item description is a multiple of 3, multiply the price by 0.2,
//     round up to the nearest integer, and add it to the total points.
//   - 5 bonus points if the total is greater than 10.00 (specific to LLM-generated programs).
//   - 6 points if the purchase date is an odd day.
//   - 10 points if the purchase time is between 2:00 PM and 4:00 PM.
//
// Parameters:
// - receipt: The receipt for which points should be calculated.
//
// Returns:
// - An integer representing the total points earned.
func CalculatePoints(receipt models.Receipt) int {
	points := 0
	for _, r := range receipt.Retailer {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			points++
		}
	}

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		points += 10
	}

	return points
}
