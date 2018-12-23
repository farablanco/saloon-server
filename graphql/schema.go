// Package graphql schema
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
package graphql

import (
	"fmt"
	"log"
	"time"

	"../db"
	"../models"
	validator "../utils"
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"User": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					email := emailAddress
					db := db.DBManager()
					user := models.User{}
					user.Email = email
					log.Print(email)
					db.First(&user)
					fmt.Printf(">>> %v", user)
					user.Password = ""
					return &user, nil
				},
			},

			"UserPtsHistory": &graphql.Field{
				Type: graphql.NewList(userPtsType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					email := emailAddress
					db := db.DBManager()

					userPts := []models.UserPts{}
					err := db.Joins("JOIN user on user.id=user_pts.user_id").
						Where("user.email=?", email).
						Find(&userPts).Error
					if err == nil {
						fmt.Printf(">>> %v", userPts)
						return &userPts, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},
		},
	},
)

var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{

			// update user profile
			"updateUserDetails": &graphql.Field{
				Type:        userType,
				Description: "Update User Details",
				Args: graphql.FieldConfigArgument{
					"contactNo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"firstname": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastname": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"bio": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					email := emailAddress
					contactNo, _ := params.Args["contactNo"].(string)
					firstName, _ := params.Args["firstname"].(string)
					lastName, _ := params.Args["lastname"].(string)
					image, _ := params.Args["image"].(string)
					bio, _ := params.Args["bio"].(string)

					updatedAt := time.Now()
					user := models.User{}
					db := db.DBManager()
					db.Model(&user).Where("email = ?", email).UpdateColumns(models.User{ContactNo: contactNo, Firstname: firstName, Lastname: lastName, Image: image, Bio: bio, UpdatedAt: updatedAt})
					return &user, nil
				},
			},

			"changePassword": &graphql.Field{
				Type:        userType,
				Description: "Change password",
				Args: graphql.FieldConfigArgument{
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					email := emailAddress
					password, _ := params.Args["password"].(string)

					updatedAt := time.Now()
					user := models.User{}
					db := db.DBManager()
					err := user.HashPassword(password)
					if err != nil {
						return err, nil
					}
					db.Model(&user).Where("email = ?", email).UpdateColumns(models.User{UpdatedAt: updatedAt})
					return &user, nil
				},
			},

			"changePasswordWithChallengePhrase": &graphql.Field{
				Type:        userType,
				Description: "Change password with Challenge phrase",
				Args: graphql.FieldConfigArgument{
					"resetPasswordToken": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					email := emailAddress
					password, _ := params.Args["password"].(string)
					resetPasswordToken, _ := params.Args["resetPasswordToken"].(string)
					updatedAt := time.Now()
					user := models.User{}
					db := db.DBManager()
					err := user.HashPassword(password)
					if err != nil {
						return err, nil
					}
					if errs := validator.Validate(user); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&user).Where("reset_password_token = ? AND email = ?", resetPasswordToken, email).UpdateColumns(models.User{UpdatedAt: updatedAt, ResetPasswordDate: &updatedAt})
					return &user, nil
				},
			},
		},
	},
)

func ExecuteQuery(query string, usrEmail string) *graphql.Result {
	emailAddress = usrEmail
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: rootMutation,
		},
	)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}
