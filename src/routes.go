package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RestGetEmployees(c echo.Context) error {
	var response = EmployeeResponse{
		Employees,
		len(Employees),
	}

	return c.JSON(http.StatusOK, response)
}

func RestGetProducts(c echo.Context) error {
	var response = ProductResponse{
		Products,
		len(Products),
	}

	return c.JSON(http.StatusOK, response)
}

func SoapGetEmployees(c echo.Context) error {
	var response = EmployeeResponse{
		Employees,
		len(Employees),
	}

	return c.XML(http.StatusOK, response)
}

func SoapGetProducts(c echo.Context) error {
	var response = ProductResponse{
		Products,
		len(Products),
	}

	return c.XML(http.StatusOK, response)
}
