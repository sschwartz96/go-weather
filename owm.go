package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

const owmURL = "https://api.openweathermap.org/data/2.5/weather"

func getCurrentTempurature(cfg *config) (string, error) {
	cW, err := getCurrentWeather(cfg)
	if err != nil {
		return "", fmt.Errorf("getCurrentTempurature(), cannot get currentWeather: \n\t%v", err)
	}
	return fmt.Sprintf("%.0f", math.Round(cW.Main.Temp)), nil
}

func getCurrentWeather(cfg *config) (*currentWeather, error) {
	req, err := http.NewRequest("GET", owmURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("getCurrentWeather(), cannot build request: ", err)
	}
	query := req.URL.Query()
	query.Add("appid", cfg.ApiKey)
	query.Add("units", cfg.Units)
	query.Add("id", cfg.CityID)
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getCurrentWeather(), cannot retrieve weather: ", err)
	}

	cW := &currentWeather{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(cW)
	if err != nil {
		return nil, fmt.Errorf("getCurrentWeather(), cannot decode currentWeather: ", err)
	}

	return cW, nil
}

// GENERATED CODE
type currentWeather struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
