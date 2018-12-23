// author : Kenneth Phang
package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"../graphql"
	"../models"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

/**
 *
 */
func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("Context %v\n", c)
		fmt.Printf("Context %v\n", c.Request().Body)
		bufBody := new(bytes.Buffer)

		bufBody.ReadFrom(c.Request().Body)
		var apolloQuery map[string]interface{}
		readBuf, _ := ioutil.ReadAll(bufBody)
		if err := json.Unmarshal(readBuf, &apolloQuery); err != nil { // unmarshall body contents as a type query
			fmt.Println(err)
			fmt.Println("Error on Unmarshalling!!!")
			return echo.ErrUnauthorized
		}
		username := apolloQuery["email"].(string)
		password := apolloQuery["password"].(string)
		fmt.Printf("Login user %s\n", username)
		fmt.Printf("Login password %s\n", password)

		user := []models.User{}
		db.Find(&user, "email=?", username)
		if len(user) > 0 && username == user[0].Email {
			// Create token
			isPasswordCorrect := user[0].CheckPassword(password)
			fmt.Printf("isPasswordCorrect %t\n", isPasswordCorrect)
			if !isPasswordCorrect {
				return echo.ErrUnauthorized
			}
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = username
			claims["admin"] = user[0].IsAdmin
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

func Register(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		contactNo := c.FormValue("contactNo")

		fmt.Printf("Register user %s\n", username)
		fmt.Printf("Register password %s\n", password)
		fmt.Printf("Register contactNo %s\n", contactNo)
		user := []models.User{}
		db.Find(&user, "email=? ", username)
		fmt.Printf("Register password %d\n", len(user))
		if len(user) <= 0 {
			insertUser := models.User{Email: username, Password: password, ContactNo: contactNo}

			fmt.Printf("Register password %v\n", &insertUser)
			err2 := insertUser.HashPassword(password)
			if err2 != nil {
				return err2
			}
			err := db.Save(&insertUser).Error
			if err != nil {
				return echo.ErrUnauthorized
			}
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)
			fmt.Printf("isAdmin %d\n", &insertUser.IsAdmin)
			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = username
			claims["admin"] = false
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

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		isAdminstr := fmt.Sprint(user.Claims.(jwt.MapClaims)["admin"])
		userEmail := fmt.Sprint(user.Claims.(jwt.MapClaims)["name"])
		isAdmin, _ := strconv.ParseBool(isAdminstr)
		bufBody := new(bytes.Buffer)

		bufBody.ReadFrom(c.Request().Body)
		var apolloQuery map[string]interface{}
		fmt.Println("Received request body")
		readBuf, _ := ioutil.ReadAll(bufBody)
		//bodyContent := bufBody.String()
		if err := json.Unmarshal(readBuf, &apolloQuery); err != nil { // unmarshall body contents as a type query
			fmt.Println(err)
			fmt.Println("Error on Unmarshalling!!!")
			return echo.ErrUnauthorized
		}
		query := apolloQuery["query"]
		variables := apolloQuery["variables"]

		log.Printf("%v", query)
		log.Printf("%v", variables)
		if isAdmin {
			fmt.Printf(" ADMIN !")
			result := graphql.ExecuteAdminQuery(query.(string), userEmail)
			return c.JSON(http.StatusOK, result)
		} else {
			fmt.Printf(" NON ADMIN !")
			result := graphql.ExecuteQuery(query.(string), userEmail)
			return c.JSON(http.StatusOK, result)
		}
	}
}
