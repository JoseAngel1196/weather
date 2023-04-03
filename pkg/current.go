package pkg

import (
	"fmt"

	"github.com/JoseAngel1196/weather/api"
	"github.com/JoseAngel1196/weather/share"

	// "github.com/fatih/color"
	"github.com/spf13/cobra"
)

type weather struct {
	Location   string
	Temp       string
	Humidity   string
	Wind       string
	FeelsLike  string
	Sky        string
	UVIndex    string
	Visibility string
	Pressure   string
	DewPoint   string
	Sunrise    string
	Sunset     string
}

func Current(cmd *cobra.Command, args []string) {
	// Get user location (e.g New York)
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
	fmt.Println(err)

	// Fetch the weather data
	weatherPointsData, err := api.GetWeatherPoints(location.Latitude, location.Longitude)
	if err != nil {
		exitOnError("Sorry, looks like there was an error getting the weather data 😭", err)
	}
	fmt.Println(weatherPointsData)

	// Print the current weather
	// printCurrent(weatherData)
}

// func printCurrent(weatherData internal.WeatherDataResponse) {
// 	// Create Weather struct instance with sample data
// 	fahrenheit := internal.KelvinToFahrenheit(weatherData.Main.Temp)
// 	celsius := internal.KelvinToCelsius(weatherData.Main.Temp)
// 	feelsLikeFahrenheit := internal.KelvinToFahrenheit(weatherData.Main.FeelsLike)
// 	feelsLikeCelsius := internal.KelvinToCelsius(weatherData.Main.FeelsLike)
// 	windSpeed := internal.MpsToMph(weatherData.Wind.Speed)
// 	weather := weather{
// 		Location:   weatherData.Sys.Country,
// 		Temp:       fmt.Sprintf("%.2f°F (%2.f°C)", fahrenheit, celsius),
// 		Humidity:   fmt.Sprintf("%d%v", weatherData.Main.Humidity, "%"),
// 		Wind:       fmt.Sprintf("%.2f mph W", windSpeed),
// 		FeelsLike:  fmt.Sprintf("%.2f°F (%2.f°C)", feelsLikeFahrenheit, feelsLikeCelsius),
// 		Sky:        weatherData.Weather[0].Main,
// 		UVIndex:    "7",
// 		Visibility: "10 mi",
// 		Pressure:   "1012 hPa",
// 		DewPoint:   "55°F (13°C)",
// 		Sunrise:    "6:30 AM",
// 		Sunset:     "7:15 PM",
// 	}

// 	// Create color instances for styling text output
// 	blue := color.New(color.FgBlue)
// 	yellow := color.New(color.FgYellow)
// 	green := color.New(color.FgGreen)
// 	white := color.New(color.FgWhite)

// 	// Print formatted weather data with styled text
// 	fmt.Println()
// 	blue.Printf("🌤️  Current Weather Conditions 🌤️\n")
// 	fmt.Println("----------------------------------")
// 	white.Printf("📍 Location: %s\n", weather.Location)
// 	green.Printf("🌡️ Temperature: %s\n", weather.Temp)
// 	yellow.Printf("💧 Humidity: %s\n", weather.Humidity)
// 	green.Printf("💨 Wind: %s\n", weather.Wind)
// 	green.Printf("🌡️ Feels Like: %s\n\n", weather.FeelsLike)

// 	blue.Printf("🌥️ Sky Conditions 🌥️\n")
// 	fmt.Println("---------------------")
// 	green.Printf("☀️ Sky: %s\n", weather.Sky)
// 	yellow.Printf("☀️ UV Index: %s\n", weather.UVIndex)
// 	green.Printf("👀 Visibility: %s\n\n", weather.Visibility)

// 	blue.Printf("🌡️ Additional Information 🌡️\n")
// 	fmt.Println("------------------------------")
// 	yellow.Printf("🗜️ Pressure: %s\n", weather.Pressure)
// 	yellow.Printf("💦 Dew Point: %s\n", weather.DewPoint)
// 	green.Printf("🌅 Sunrise: %s\n", weather.Sunrise)
// 	green.Printf("🌇 Sunset: %s\n", weather.Sunset)
// 	fmt.Println()
// }
