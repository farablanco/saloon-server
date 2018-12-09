// Package web
//
// Copyright 2018 Kenneth Phang <bunnyppl@yahoo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"os"

	"../../db"
	"../../handler"
	"../../models"
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
	defer db.Close()

	user := models.User{}
	userPts := models.UserPts{}
	product := models.Product{}
	payment := models.Payment{}
	paymentItems := models.PaymentItems{}
	outlet := models.Outlet{}
	membership := models.Membership{}
	membershipProduct := models.MembershipProduct{}

	db.AutoMigrate(&user, &userPts,
		&product, &payment, &paymentItems, &outlet, &membership, &membershipProduct)

	e.GET("/hello", handler.Hello())
	e.POST("/login", handler.Login(db))
	e.POST("/register", handler.Register(db))

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	r.POST("", handler.Restricted())

	e.Start(":3000")
}
