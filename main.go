package main

import (
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
	e.Use(middleware.Recover())

	e.GET("/hello", handler.Hello())
	e.POST("/login", handler.Login())
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("", handler.Restricted())

	e.Start(":3000")
}
