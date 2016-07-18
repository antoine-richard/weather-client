package main

import (
	"flag"
	"fmt"
	"math"

	"gopkg.in/h2non/gentleman.v1"
)

type cityWeather struct {
	Name    string
	Weather []struct {
		Main        string
		Description string
	}
	Wind map[string]float64
	Main struct {
		Temp     float64
		Humidity float64
	}
}

var apiKey *string
var weather *cityWeather
var cities = map[string]string{
	"Nantes":    "2990969",
	"Palo Alto": "5380748",
	"Prague":    "3067696",
}
var http = gentleman.New()

func main() {
	parseFlags()
	for _, cityID := range cities {
		getWeather(cityID)
		fmt.Printf("%v: %v (%v), %v°C\n",
			weather.Name,
			weather.Weather[0].Main,
			weather.Weather[0].Description,
			kelvinToCelsius(weather.Main.Temp))
	}
}

func parseFlags() {
	apiKey = flag.String("key", "", "Open Weather Map API key")
	flag.Parse()
}

func getWeather(cityID string) {
	req := http.Request().URL("api.openweathermap.org")
	req.Path("/data/2.5/weather")
	req.SetQuery("id", cityID)
	req.SetQuery("appid", *apiKey)

	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	weather = &cityWeather{}
	err = res.JSON(weather)
	if err != nil {
		fmt.Printf("JSON parse error: %s\n", err)
		return
	}
}

func kelvinToCelsius(temp float64) int {
	return int(math.Floor(temp - 273.15))
}
