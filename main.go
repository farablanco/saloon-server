package main

import (
	"os"

	"./db"
	"./handler"
	"./models"
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

	user := models.User{}
	userPts := models.UserPts{}
	product := models.Product{}
	points := models.Points{}
	payment := models.Payment{}
	paymentItems := models.PaymentItems{}
	outlet := models.Outlet{}
	membership := models.Membership{}
	membershipProduct := models.MembershipProduct{}

	db.AutoMigrate(&user, &userPts,
		&product, &points, &payment, &paymentItems, &outlet, &membership, &membershipProduct)
	e.GET("/hello", handler.Hello())
	e.POST("/login", handler.Login(db))
	e.POST("/register", handler.Register(db))

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	r.POST("", handler.Restricted())

	e.Start(":3000")
}
