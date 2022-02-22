package trade_tools

func CalculatePivotPoint(high , low, close float64) float64 {
	return (high + low + close) / 3
}

func CalculateFirstSupport(pivotPoint, high float64) float64 {
	return (pivotPoint * 2) - high
}

func CalculateSecondSupport(pivotPoint, high, low float64) float64 {
	return pivotPoint - (high - low)
}

func CalculateFirstResistance(pivotPoint, low float64) float64 {
	return (pivotPoint * 2) - low
}

func CalculateSecondResistance(pivotPoint, high , low float64) float64 {
	return pivotPoint + (high - low)
}
