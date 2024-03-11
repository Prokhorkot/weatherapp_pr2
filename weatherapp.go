package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("API ключ не найден. Убедитесь, что вы задали переменную окружения WEATHER_API_KEY.")
		os.Exit(1)
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?id=524901&appid=%s&units=metric", apiKey)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при запросе к API погоды:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		var weatherResponse WeatherResponse
		json.Unmarshal(body, &weatherResponse)

		temperature := weatherResponse.Main.Temp
		weatherDescription := weatherResponse.Weather[0].Description
		report := fmt.Sprintf("Текущая температура в Москве: %.2f°C, погода: %s", temperature, weatherDescription)

		ioutil.WriteFile("weather_report.txt", []byte(report), 0644)
	} else {
		fmt.Println("Ошибка при получении данных о погоде")
	}
}

func main() {
	getWeather()
}
