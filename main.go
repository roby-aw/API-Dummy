package main

import (
	"api-redeem-point/api"
	"api-redeem-point/api/dummy"
	_ "api-redeem-point/docs"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title API-Loyalty Point Agent
// @version 1.0
// @description Berikut API-Loyalty Point Agent
// @host api-dummy.herokuapp.com
// @BasePath /
func main() {
	dummy.InitiateDB()
	port := os.Getenv("PORT")
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
		if port == "" {
			port = "8080"
		}
		address := fmt.Sprintf(":%s", port)
		fmt.Println(address)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
