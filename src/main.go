package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

var (
	Employees = []string{
		"Jhon Doe",
		"Theodorus Martell",
		"Radovan Alvarado",
		"Neven Schultes",
		"Lysandra Gorecki",
		"Erhard Douglas",
	}

	Products = []string{
		"Mechanical Keyboard",
		"Mouse",
		"Monitor",
		"Headset",
		"Mousepad",
		"Webcam",
	}
)

func main() {
	web := echo.New()

	rest := web.Group("/rest")
	rest.GET("/employees", RestGetEmployees)
	rest.GET("/products", RestGetProducts)

	soap := web.Group("/soap")
	soap.GET("/employees", SoapGetEmployees)
	soap.GET("/products", SoapGetProducts)

	log.Fatal(web.Start(":8080"))
}
