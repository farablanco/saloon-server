package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func Login() echo.HandlerFunc {
	log.Debug("Created user")

	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		log.Debug("Created user", username)
		log.Debug("Created password", password)
		if username == "test" && password == "test" {
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = "test"
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
			fmt.Printf("%+v\n", claims)
			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secret"))
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

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.String(http.StatusOK, "Welcome "+name+"!")
	}
}
