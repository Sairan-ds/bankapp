package deposit


func Calculate(amount int, currency string) (min, max int) {
	minPercent, maxPercent := PercentFor(currency)

	min = amount * minPercent / 100 / 12
	max = amount * maxPercent / 100 / 12

	return min, max
}

func PercentFor(currency string) (int,int) {
	switch currency {
	case "TJS":
		return 4,6
	case "USD":
		return 1,2
	default:
		return 0,0
	}
}