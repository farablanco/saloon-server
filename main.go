package main

import (
	"os"

	"./db"
	"./handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	db := db.ConnectGORM()
	db.SingularTable(true)
	e.GET("/hello", handler.Hello())
	e.POST("/login", handler.Login(db))
	e.POST("/register", handler.Register())

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	r.POST("", handler.Restricted())

	e.Start(":3000")
}
