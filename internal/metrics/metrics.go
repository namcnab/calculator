package metrics

func CalculateResponseTime(startTime, endTime int64) float64 {
	return float64(endTime-startTime) / 1000000 // Convert nanoseconds to milliseconds
}
