package tax

func CalculateTax(amout float64) float64 {
	if amout >= 1000 {
		return 10.0
	}
	return 5.0
}
