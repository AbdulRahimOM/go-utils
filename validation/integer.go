package validation

import "strconv"

func ValidateAndParseInt(param string) (int, error) {
	return strconv.Atoi(param)
}

func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}