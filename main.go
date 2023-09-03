package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int64     `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int64     `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMB       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int64     `json:"humidity"`
	Cloud            int64     `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	VisKM            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	Uv               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int64  `json:"code"`
}

type Forecast struct {
	Forecastday []Forecastday `json:"forecastday"`
}

type Forecastday struct {
	Date      string `json:"date"`
	DateEpoch int64  `json:"date_epoch"`
	Day       Day    `json:"day"`
	Astro     Astro  `json:"astro"`
	Hour      []Hour `json:"hour"`
}

type Astro struct {
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	MoonPhase        string `json:"moon_phase"`
	MoonIllumination string `json:"moon_illumination"`
	IsMoonUp         int64  `json:"is_moon_up"`
	IsSunUp          int64  `json:"is_sun_up"`
}

type Day struct {
	MaxtempC          float64   `json:"maxtemp_c"`
	MaxtempF          float64   `json:"maxtemp_f"`
	MintempC          float64   `json:"mintemp_c"`
	MintempF          float64   `json:"mintemp_f"`
	AvgtempC          float64   `json:"avgtemp_c"`
	AvgtempF          float64   `json:"avgtemp_f"`
	MaxwindMph        float64   `json:"maxwind_mph"`
	MaxwindKph        float64   `json:"maxwind_kph"`
	TotalprecipMm     float64   `json:"totalprecip_mm"`
	TotalprecipIn     float64   `json:"totalprecip_in"`
	TotalsnowCM       float64   `json:"totalsnow_cm"`
	AvgvisKM          float64   `json:"avgvis_km"`
	AvgvisMiles       float64   `json:"avgvis_miles"`
	Avghumidity       float64   `json:"avghumidity"`
	DailyWillItRain   int64     `json:"daily_will_it_rain"`
	DailyChanceOfRain int64     `json:"daily_chance_of_rain"`
	DailyWillItSnow   int64     `json:"daily_will_it_snow"`
	DailyChanceOfSnow int64     `json:"daily_chance_of_snow"`
	Condition         Condition `json:"condition"`
	Uv                float64   `json:"uv"`
}

type Hour struct {
	TimeEpoch    int64     `json:"time_epoch"`
	Time         string    `json:"time"`
	TempC        float64   `json:"temp_c"`
	TempF        float64   `json:"temp_f"`
	IsDay        int64     `json:"is_day"`
	Condition    Condition `json:"condition"`
	WindMph      float64   `json:"wind_mph"`
	WindKph      float64   `json:"wind_kph"`
	WindDegree   int64     `json:"wind_degree"`
	WindDir      string    `json:"wind_dir"`
	PressureMB   float64   `json:"pressure_mb"`
	PressureIn   float64   `json:"pressure_in"`
	PrecipMm     float64   `json:"precip_mm"`
	PrecipIn     float64   `json:"precip_in"`
	Humidity     int64     `json:"humidity"`
	Cloud        int64     `json:"cloud"`
	FeelslikeC   float64   `json:"feelslike_c"`
	FeelslikeF   float64   `json:"feelslike_f"`
	WindchillC   float64   `json:"windchill_c"`
	WindchillF   float64   `json:"windchill_f"`
	HeatindexC   float64   `json:"heatindex_c"`
	HeatindexF   float64   `json:"heatindex_f"`
	DewpointC    float64   `json:"dewpoint_c"`
	DewpointF    float64   `json:"dewpoint_f"`
	WillItRain   int64     `json:"will_it_rain"`
	ChanceOfRain int64     `json:"chance_of_rain"`
	WillItSnow   int64     `json:"will_it_snow"`
	ChanceOfSnow int64     `json:"chance_of_snow"`
	VisKM        float64   `json:"vis_km"`
	VisMiles     float64   `json:"vis_miles"`
	GustMph      float64   `json:"gust_mph"`
	GustKph      float64   `json:"gust_kph"`
	Uv           float64   `json:"uv"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

func main() {
	var API_KEY string
	if val, exists := os.LookupEnv("WEATHER_API_KEY"); exists {
		API_KEY = val
	} else {
		fmt.Println("key not found")
	}

	q := "Manila"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?q=%s&key=%s", q, API_KEY)
	res, err := http.Get(url)
	if err != nil {
		panic("there is an error, please try again")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("there is an error, please try again")
	}

	var weather Weather
	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		panic("there is an error, please try again")
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf("%s, %s: %.0fC, %s\n",
		location.Name, location.Country, current.TempC, current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf("%s - %0.fC, %d%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}
