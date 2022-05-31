package main

import (
	"api-redeem-point/api"
	"api-redeem-point/config"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title API Jasa Pengiriman
// @version 1.0
// @description Berikut API Jasa Pengiriman
// @host localhost:8080
// @BasePath /
func main() {
	config := config.GetConfig()

	e := echo.New()
	handleSwagger := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwagger)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})
	api.RegistrationPath(e, api.Controller{})

	go func() {
		address := fmt.Sprintf(":%d", config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
