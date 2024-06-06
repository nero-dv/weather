package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: %s <city>", args[0])
	}
	city := strings.Join(args[1:], " ")
	var inCoords []float64
	inCoords, err := getCoords(city)
	if err != nil {
		return err
	}
	coords := fmt.Sprintf("%f,%f", inCoords[1], inCoords[0])
	weather, err := GetWeather(coords)
	if err != nil {
		return err
	}
	fmt.Printf("Weather in %s: %f\n", city, weather)

	return nil
}

func GetWeather(city string) (float64, error) {
	API_KEY := os.Getenv("WEATHER_API_KEY")
	var coords string
	if city != "" {
		coords = city
	} else {
		coords = "29.97255655272305, -95.69789623941507"

	}
	url := "https://api.pirateweather.net/forecast/" + API_KEY + "/" + coords
	fmt.Printf("URL: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0.0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}
	var weather WeatherData
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return 0.0, err
	}
	return float64(weather.Currently.Temperature), nil
}

func getCoords(city string) ([]float64, error) {

	paramAddress := city

	encodedAddress := url.QueryEscape(paramAddress)
	GEO_API_KEY := os.Getenv("GEO_API_KEY")
	geocodeURL := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=%s", encodedAddress, GEO_API_KEY)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, geocodeURL, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var geocode Geocode
	err = json.Unmarshal(body, &geocode)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return geocode.Features[0].Geometry.Coordinates, nil
}
