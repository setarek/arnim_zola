package utils

import (
	"fmt"
	"strconv"
)

func ParseFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case string:
		val, _ := strconv.ParseFloat(v, 64)
		return val
	case float64:
		return v
	case int:
		return float64(v)
	default:
		val, _ := strconv.ParseFloat(fmt.Sprintf("%+v", v), 64)
		return val
	}
}
