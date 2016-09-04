package main

import (
	"errors"
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

var (
	apiKey  *string
	weather *cityWeather
)
var client = gentleman.New()
var cities = map[string]string{
	"Nantes":    "2990969",
	"Palo Alto": "5380748",
	"Prague":    "3067696",
}

func main() {
	parseFlags()
	for _, cityID := range cities {
		err := fetchWeather(cityID)
		if err == nil {
			fmt.Printf("%v: %v (%v), %v°C\n",
				weather.Name,
				weather.Weather[0].Main,
				weather.Weather[0].Description,
				kelvinToCelsius(weather.Main.Temp))
		} else {
			fmt.Print(err)
		}
	}
}

func parseFlags() {
	apiKey = flag.String("key", "", "OpenWeatherMap API key")
	flag.Parse()
}

func fetchWeather(cityID string) error {
	req := client.Request().URL("api.openweathermap.org")
	req.Path("/data/2.5/weather")
	req.SetQuery("id", cityID)
	req.SetQuery("appid", *apiKey)

	res, err := req.Send()
	if err != nil {
		return errors.New("Request error: " + err.Error())
	}
	if !res.Ok {
		return errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
	}

	weather = &cityWeather{}
	err = res.JSON(weather)
	if err != nil {
		return errors.New("JSON parse error: " + err.Error())
	}

	return nil
}

func kelvinToCelsius(temp float64) int {
	return int(math.Floor(temp - 273.15))
}
