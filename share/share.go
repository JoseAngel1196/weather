package share

import (
	"time"
)

type WeatherPoints struct {
	Context  []interface{} `json:"@context"`
	ID       string        `json:"id"`
	Type     string        `json:"type"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		ID                  string `json:"@id"`
		Type                string `json:"@type"`
		Cwa                 string `json:"cwa"`
		ForecastOffice      string `json:"forecastOffice"`
		GridID              string `json:"gridId"`
		GridX               int    `json:"gridX"`
		GridY               int    `json:"gridY"`
		Forecast            string `json:"forecast"`
		ForecastHourly      string `json:"forecastHourly"`
		ForecastGridData    string `json:"forecastGridData"`
		ObservationStations string `json:"observationStations"`
		RelativeLocation    struct {
			Type     string `json:"type"`
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Properties struct {
				City     string `json:"city"`
				State    string `json:"state"`
				Distance struct {
					UnitCode string  `json:"unitCode"`
					Value    float64 `json:"value"`
				} `json:"distance"`
				Bearing struct {
					UnitCode string `json:"unitCode"`
					Value    int    `json:"value"`
				} `json:"bearing"`
			} `json:"properties"`
		} `json:"relativeLocation"`
		ForecastZone    string `json:"forecastZone"`
		County          string `json:"county"`
		FireWeatherZone string `json:"fireWeatherZone"`
		TimeZone        string `json:"timeZone"`
		RadarStation    string `json:"radarStation"`
	} `json:"properties"`
}

type ForecastHourly struct {
	Context  []interface{} `json:"@context"`
	Type     string        `json:"type"`
	Geometry struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		Updated           time.Time `json:"updated"`
		Units             string    `json:"units"`
		ForecastGenerator string    `json:"forecastGenerator"`
		GeneratedAt       time.Time `json:"generatedAt"`
		UpdateTime        time.Time `json:"updateTime"`
		ValidTimes        time.Time `json:"validTimes"`
		Elevation         struct {
			UnitCode string  `json:"unitCode"`
			Value    float64 `json:"value"`
		} `json:"elevation"`
		Periods []struct {
			Number                     int         `json:"number"`
			Name                       string      `json:"name"`
			StartTime                  string      `json:"startTime"`
			EndTime                    string      `json:"endTime"`
			IsDaytime                  bool        `json:"isDaytime"`
			Temperature                int         `json:"temperature"`
			TemperatureUnit            string      `json:"temperatureUnit"`
			TemperatureTrend           interface{} `json:"temperatureTrend"`
			ProbabilityOfPrecipitation struct {
				UnitCode string `json:"unitCode"`
				Value    int    `json:"value"`
			} `json:"probabilityOfPrecipitation"`
			Dewpoint struct {
				UnitCode string  `json:"unitCode"`
				Value    float64 `json:"value"`
			} `json:"dewpoint"`
			RelativeHumidity struct {
				UnitCode string `json:"unitCode"`
				Value    int    `json:"value"`
			} `json:"relativeHumidity"`
			WindSpeed        string `json:"windSpeed"`
			WindDirection    string `json:"windDirection"`
			Icon             string `json:"icon"`
			ShortForecast    string `json:"shortForecast"`
			DetailedForecast string `json:"detailedForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

type WeatherInfo struct {
	Number                     int
	Name                       string
	StartTime                  string
	EndTime                    string
	IsDaytime                  bool
	Temperature                int
	TemperatureUnit            string
	TemperatureTrend           interface{}
	ProbabilityOfPrecipitation struct {
		UnitCode string
		Value    int
	}
	Dewpoint struct {
		UnitCode string
		Value    float64
	}
	RelativeHumidity struct {
		UnitCode string
		Value    int
	}
	WindSpeed        string
	WindDirection    string
	Icon             string
	ShortForecast    string
	DetailedForecast string
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type UserAddress struct {
	Street string
	// Number     int
	City       string
	State      string
	Country    string
	PostalCode string
}

func (address *UserAddress) FormatAddress() string {

	// Creats a slice with all content from the Address struct
	var content []string
	// if address.Number > 0 {
	// 	content = append(content, strconv.Itoa(address.Number))
	// }
	content = append(content, address.Street)
	content = append(content, address.PostalCode)
	content = append(content, address.City)
	content = append(content, address.State)
	content = append(content, address.Country)

	var formattedAddress string

	// For each value in the content slice check if it is valid
	// and add to the formattedAddress string
	for _, value := range content {
		if value != "" {
			if formattedAddress != "" {
				formattedAddress += ", "
			}
			formattedAddress += value
		}
	}

	return formattedAddress
}

type GeoResults struct {
	Results []GeoResult `json:"results"`
	Status  string      `json:"status"`
}

type GeoResult struct {
	AddressComponents []Address `json:"address_components"`
	FormattedAddress  string    `json:"formatted_address"`
	Geometry          Geometry  `json:"geometry"`
	PlaceId           string    `json:"place_id"`
	Types             []string  `json:"types"`
}

type Address struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Geometry struct {
	Bounds       Bounds `json:"bounds"`
	Location     LatLng `json:"location"`
	LocationType string `json:"location_type"`
	Viewport     Bounds `json:"viewport"`
}

type Bounds struct {
	Northeast LatLng `json:"northeast"`
	Southwest LatLng `json:"southwest"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
