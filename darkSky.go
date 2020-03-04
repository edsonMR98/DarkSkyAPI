package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/shawntoffel/darksky"
)

func main() {
	client := darksky.New("KEY")

	request := darksky.ForecastRequest{
		Latitude:  19.4327224,
		Longitude: -99.1333089,
		Options: darksky.ForecastRequestOptions{
			Exclude: "hourly,minutely",
			Units:   "auto", // SI = Celsius
		},
	}

	forecast, err := client.Forecast(request)
	if err != nil {
		panic(err)
	}

	e := reflect.ValueOf(&forecast).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}

	fmt.Printf("Latitud: %v - Longitud: %v\n", forecast.Latitude, forecast.Longitude)
	fmt.Printf("Temperature: %v°\n", forecast.Currently.Temperature)
	fmt.Printf("Current summary: %v\n", forecast.Currently.Summary)
	fmt.Printf("Timezone: %v\n", forecast.Timezone)

	fmt.Printf("Forecasted week summary: %v\n", forecast.Daily.Summary)
	for _, day := range forecast.Daily.Data {
		tm := time.Unix(int64(day.Time), 0) // Converts timestamp to int64. Gets DateTime
		fmt.Println(tm.Date())
		fmt.Println(day.Summary)
		avg := (day.TemperatureMax + day.TemperatureMin) / 2
		fmt.Printf("%.1f°\n", avg) // Rounds to a floating-point value
	}
}
