package validation

import "time"

// ValidateAndParseDate validates and parses a date string.
//
// Expected date format: "2006-01-02"
func ValidateAndParseDate(dateStr string) (time.Time, error) {
	// Define the expected date format
	dateFormat := "2006-01-02"

	// Parse the date string into a time.Time object
	date, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
