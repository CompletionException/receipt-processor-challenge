package business

import (
	"log"
	"math"
	"receipt-processor-challenge/internal/model"
	"strconv"
	"strings"
	"time"
)

// Define the expected date format
const dateFormat = "2006-01-02" // Change if your format is different

// CalculatePoints calculates the points based on a receipt
func CalculatePoints(receipt model.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += countAlphanumericCharacters(receipt.Retailer)

	// Rule 2/3: Parse total string
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		//Receipt total not parsable, no points awarded.
		log.Print(err)
	} else {
		// Rule 2: 50 points if the total is a round dollar Total with no cents
		if total == float64(int(total)) {
			points += 50
		}
		// Rule 3: 25 points if the total is a multiple of 0.25
		if math.Mod(total, 0.25) == 0 {
			points += 25
		}
	}

	// Rule 4: 5 points for every two items on the receipt
	points += 5 * (len(receipt.Items) / 2)

	// Rule 5: Points for each item based on the description length
	for _, item := range receipt.Items {
		descriptionLength := len(strings.TrimSpace(item.ShortDescription))
		if descriptionLength%3 == 0 {
			//Parse price
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				//Receipt total not parsable, no points awarded.
				log.Print(err)
			} else {
				itemPoints := int(math.Ceil(price * 0.2))
				points += itemPoints
			}
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse(dateFormat, receipt.Date)
	if err != nil {
		//Receipt date not parsable, no points awarded.
		log.Print(err)
	} else {
		if purchaseDate.Day()%2 != 0 {
			points += 6
		}
	}

	// Rule 7: 10 points if the time of purchase is between 2:00pm and 4:00pm
	parsedTime := parseTime(receipt.Time)
	if parsedTime.After(parseTime("14:00")) && parsedTime.Before(parseTime("16:00")) {
		points += 10
	}
	return points
}

// countAlphanumericCharacters counts alphanumeric characters in a string
func countAlphanumericCharacters(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			count++
		}
	}
	return count
}

// parseTime parses a time string into a time.Time object. Ignore any internal thoughts about time zones
func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("15:04", timeStr)
	return t
}
