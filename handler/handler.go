package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"../graphql"
	"../models"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Hello() echo.HandlerFunc {
	log.Debug("Created user")
	println("bar")
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		fmt.Printf("Login user %s\n", username)
		fmt.Printf("Login password %s\n", password)

		user := []models.User{}
		db.Find(&user, "email=? and password=?", username, password)
		if len(user) > 0 && username == user[0].Email {
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = username
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
			fmt.Printf("%+v\n", claims)
			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}

		return echo.ErrUnauthorized
	}
}

func Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"token": "dfsfd",
		})
	}
}

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		_ = user.Claims.(jwt.MapClaims)
		bufBody := new(bytes.Buffer)
		bufBody.ReadFrom(c.Request().Body)
		query := bufBody.String()
		log.Printf(query)
		result := graphql.ExecuteQuery(query)
		return c.JSON(http.StatusOK, result)
	}
}
