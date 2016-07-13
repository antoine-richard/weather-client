package main

import (
	"fmt"
	"math"

	"gopkg.in/h2non/gentleman.v1"
)

type CityWeather struct {
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

func kelvinToCelsius(temp float64) int {
	return int(math.Floor(temp - 273.15))
}

func main() {
	cli := gentleman.New()

	req := cli.Request().URL("api.openweathermap.org")
	req.Path("/data/2.5/weather")
	req.SetQuery("id", "2990969")
	req.SetQuery("appid", "07b1c6874d437a56457a1d6d175e67ff")

	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	weather := &CityWeather{}
	err = res.JSON(weather)
	if err != nil {
		fmt.Printf("JSON parse error: %s\n", err)
		return
	}

	fmt.Printf("%#v\n", weather)
	fmt.Print("###\n\n")

	fmt.Printf("%v: %v, %v°C\n", weather.Name, weather.Weather[0].Description, kelvinToCelsius(weather.Main.Temp))

}
