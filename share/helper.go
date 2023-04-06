package share

const absoluteZeroCelsius float64 = 273.15

func KelvinToFahrenheit(k float64) float64 {
	return (k-absoluteZeroCelsius)*9/5 + 32
}

func KelvinToCelsius(k float64) float64 {
	return k - absoluteZeroCelsius
}

func MpsToMph(mps float64) float64 {
	return mps * 2.23694
}