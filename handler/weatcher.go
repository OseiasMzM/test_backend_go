package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test_backend_go/service"
)

func GetWeather(c *gin.Context) {
	city := c.Param("city")

	weather, err := service.FetchWeatherByCity(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Desculpe, não conseguimos obter as informações do clima."})
		return
	}

	state := weather.Sys.State
	if state == "" {
		state = "unavailable"
	}

	response := gin.H{
		"city":        weather.Name,
		"state":       state,
		"country":     weather.Sys.Country,
		"temperature": fmt.Sprintf("%.1f°C", weather.Main.Temp),
		"humidity":    fmt.Sprintf("%.0f%%", weather.Main.Humidity),
		"pressure":    fmt.Sprintf("%.0f hPa", weather.Main.Pressure),
		"description": weather.Weather[0].Description,
	}

	c.JSON(http.StatusOK, response)
}
