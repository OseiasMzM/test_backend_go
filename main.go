package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"test_backend_go/handler"
)

func main() {
	r := gin.Default()

	/*
		GET /user: Receberá um parâmetro id (ex: ?id=1) e consultará uma API externa
		de usuários (ex: JSONPlaceholder) para trazer informações de um usuário.

		cURL: curl -X GET http://localhost:8080/user/1
	*/
	r.GET("/user/:id", handler.GetUser)

	/*
		GET /weather: Receberá um parâmetro de cidade (ex: ?city=London) e
		consultará uma API externa (por exemplo, OpenWeatherMap) para trazer o clima
		atual daquela cidade.

		cURL: curl -X GET http://localhost:8080/weather/London
	*/

	r.GET("/weather/:city", handler.GetWeather)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("%v", err)
	}
}
