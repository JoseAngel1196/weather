package pkg

import (
	"fmt"
	"time"

	"github.com/JoseAngel1196/weather/api"
	"github.com/JoseAngel1196/weather/share"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Current(cmd *cobra.Command, args []string) {
	// TODO: Change this.
	userAddress := share.UserAddress{
		Street:     "38 West, 129 street",
		City:       "New York",
		State:      "New York",
		Country:    "USA",
		PostalCode: "10027",
	}

	location, err := api.GetGeolocation(userAddress)
	if err != nil {
		exitOnError("Sorry, looks like there was an error getting your location 😭", err)
	}

	weatherPointsData, err := api.GetWeatherPoints(location.Latitude, location.Longitude)
	if err != nil {
		exitOnError("Sorry, looks like there was an error getting the weather data 😭", err)
	}

	forecastData, err := api.GetForecastHourly(weatherPointsData.Properties.ForecastHourly)
	if err != nil {
		exitOnError("Sorry, looks like there was an error getting the forecast data 😭", err)
	}

	currentPeriod := getCurrentPeriod(forecastData)

	printCurrent(currentPeriod)
}

func getCurrentPeriod(forecastData share.ForecastHourly) share.ForecastPeriod {
	var currentPeriod share.ForecastPeriod
	now := time.Now()

	for _, period := range forecastData.Properties.Periods {
		start, _ := time.Parse(time.RFC3339, period.StartTime)
		end, _ := time.Parse(time.RFC3339, period.EndTime)

		if start.Before(now) && end.After(now) {
			currentPeriod = period
			break
		}
	}
	return currentPeriod
}

func printCurrent(currentWeather share.ForecastPeriod) {
	blue := color.New(color.FgBlue)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)
	white := color.New(color.FgWhite)

	dt := time.Now()
	month := dt.Month().String()
	day := dt.Day()
	year := dt.Year()
	currentTime := dt.Format(time.Kitchen)

	fmt.Println()
	blue.Printf("🌤️  Current Weather Conditions for %v %d, %d: 🌤️\n", month, day, year)
	fmt.Println("----------------------------------")
	white.Printf("📍 Location: %s\n", "New York")
	green.Printf("🌡️ Temperature: %s\n", currentWeather.Temperature)
	yellow.Printf("💧 Humidity: %s\n", currentWeather.RelativeHumidity.Value)
	green.Printf("💨 Wind: %s\n\n", currentWeather.WindSpeed)

	blue.Printf("🌡️ Additional Information 🌡️\n")
	fmt.Println("------------------------------")
	yellow.Printf("💦 Dew Point: %.2f\n", currentWeather.Dewpoint.Value)

	switch currentWeather.ShortForecast {
	case "Sunny":
		green.Printf("🌅 Sunset: %v\n", currentTime)
	default:
		green.Printf("🌥️ Partly Cloudy: %v\n", currentTime)
	}
}
