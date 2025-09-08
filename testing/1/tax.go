package tax

import "time"

func CalculateTax(amout float64) float64 {
	if amout >= 1000 {
		return 10.0
	}
	return 5.0
}

func CalculateTax2(amout float64) float64 {
	time.Sleep(time.Millisecond)
	if amout >= 1000 {
		return 10.0
	}
	return 5.0
}
