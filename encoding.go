package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"os"
)

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64  `json:"lat"`
}

type WeatherObject struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}
type Weather struct {
	Weatherdata [] WeatherObject
}
type MainWeather struct {
	Temp float64 `json:"temp"`
	Feels_like float64 `json:"feels like"`
	Temp_min float64 `json:"temp min"`
	Temp_max float64 `json:"temp max"`
	Pressure int `json:"pressure`
	Humidity int `json:"humidity"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg int `json:"deg"`
}
type Clouds struct {
	All int `json:"all"`
}
type Sys struct {
	Type int `json:"type"`
	Id int `json:"id"`
	Country string `json:"country"`
	Sunrise int `json:"sunrise"`
	Sunset int `json:"sunset"`
}
type Response struct {
	Coord Coord `json:"coord"`
	Weather []struct {
		Id int `json:"id"`
		Main string `json:"main"`
		Description string `json:"description"`
		Icon string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main MainWeather `json:"main"`
	Visibility int `json:"visibility"`
	Wind Wind `json:"wind"`
	Clouds Clouds `json:"clouds"`
	Dt int `json:"dt"`
	Sys Sys `json:"sys"`
	Cod int `json:"cod"`
	Timezone int `json:"timezone"`
	ID int `json:"id"`
	Name string `json:"name"`
}
func main() {
	API_KEY := "95d398820e5d4d98dd19b5c59c59ea18"
	city := "Nairobi"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, API_KEY)
	// io.WriteString(os.Stdout, url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err :=ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	resJson := responseData
	var resConverted Response
	err = json.Unmarshal([]byte (resJson), &resConverted)
	if err !=nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Printf("Response Struct: %#v\n", resConverted)
}
