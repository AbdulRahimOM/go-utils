package myMath

import (
	"math"
	"strconv"
)

func RoundFloat32(value float32, decimalPlaces float64) float32 {
	divisor := math.Pow(10, decimalPlaces)
	return float32(math.Round(float64(value)*divisor) / divisor)
}

// StringToInt32 converts a string to an int32.
func StringToInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}
