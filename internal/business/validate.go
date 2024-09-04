package business

import (
	"receipt-processor-challenge/internal/model"
	"regexp"
)

// Exported map of precompiled regexes
var RegexMap = map[string]*regexp.Regexp{
	"retailer":         regexp.MustCompile(`^[&\w\s\-]+$`),
	"total":            regexp.MustCompile(`^\d+\.\d{2}$`),
	"shortDescription": regexp.MustCompile(`^[\w\s\-]+$`),
}

func Validate(receipt model.Receipt) bool {
	if !RegexMap["retailer"].MatchString(receipt.Retailer) ||
		!RegexMap["total"].MatchString(receipt.Total) {
		return false
	}
	for _, item := range receipt.Items {
		if !RegexMap["shortDescription"].MatchString(item.ShortDescription) ||
			!RegexMap["total"].MatchString(item.Price) {
			return false
		}
	}
	return true
}
