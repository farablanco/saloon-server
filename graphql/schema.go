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
	"strconv"
	"time"

	"../db"
	"../models"
	validator "../utils"
	"github.com/graphql-go/graphql"
)

var emailAddress string = ""

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"firstname": &graphql.Field{
				Type: graphql.String,
			},
			"lastname": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"isAdmin": &graphql.Field{
				Type: graphql.String,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"expiry_date": &graphql.Field{
				Type: graphql.String,
			},
			"ptsBalance": &graphql.Field{
				Type: graphql.Int,
			},
			"membership": &graphql.Field{
				Type: graphql.String,
			},
			"outlet": &graphql.Field{
				Type: graphql.String,
			},
			"cutCnt": &graphql.Field{
				Type: graphql.Int,
			},
			"treatment_cnt": &graphql.Field{
				Type: graphql.Int,
			},
			"hairloss_treatment_cnt": &graphql.Field{
				Type: graphql.Int,
			},
			"contactNo": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var userPtsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserPts",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"productId": &graphql.Field{
				Type: graphql.Int,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"allocatedPts": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"pts": &graphql.Field{
				Type: graphql.Int,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"remarks": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var paymentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Payment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"paymentRefNo": &graphql.Field{
				Type: graphql.String,
			},
			"total": &graphql.Field{
				Type: graphql.Float,
			},
			"outletId": &graphql.Field{
				Type: graphql.Int,
			},
			"discount": &graphql.Field{
				Type: graphql.Float,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"paymentMode": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"gst": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var paymentItemsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PaymentItems",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"paymentId": &graphql.Field{
				Type: graphql.Int,
			},
			"subTotal": &graphql.Field{
				Type: graphql.Float,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"productId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var outletType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Outlet",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"postalCode": &graphql.Field{
				Type: graphql.Int,
			},
			"contactNo": &graphql.Field{
				Type: graphql.Int,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"outletSupervisor": &graphql.Field{
				Type: graphql.String,
			},
			"longtitude": &graphql.Field{
				Type: graphql.String,
			},
			"latitude": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var membershipType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Membership",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"ptsGiven": &graphql.Field{
				Type: graphql.Int,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var membershipProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MembershipProduct",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"productId": &graphql.Field{
				Type: graphql.Int,
			},
			"membershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"count": &graphql.Field{
				Type: graphql.Int,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var bookingType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Booking",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"hairDresserId": &graphql.Field{
				Type: graphql.Int,
			},
			"bookingDateTime": &graphql.Field{
				Type: graphql.String,
			},
			"outlet": &graphql.Field{
				Type: graphql.Int,
			},
			"customer": &graphql.Field{
				Type: graphql.Int,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var hairDresserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Hairdresser",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"yearsOfExp": &graphql.Field{
				Type: graphql.Int,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"User": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, err := strconv.ParseInt(p.Args["id"].(string), 10, 64)
					db := db.ConnectGORM()
					db.SingularTable(true)
					if err == nil {
						user := models.User{}
						user.Id = idQuery
						log.Print(idQuery)
						db.First(&user)
						fmt.Printf(">>> %v", user)
						return &user, nil
					}

					return nil, nil
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
					db := db.ConnectGORM()
					db.SingularTable(true)

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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					email := emailAddress
					contactNo, _ := params.Args["contactNo"].(string)
					firstName, _ := params.Args["firstname"].(string)
					lastName, _ := params.Args["lastname"].(string)

					updatedAt := time.Now()
					user := models.User{}
					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)
					db.Model(&user).Where("email = ?", email).UpdateColumns(models.User{ContactNo: contactNo, Firstname: firstName, Lastname: lastName, UpdatedAt: updatedAt})
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
					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)
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
					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)
					err := user.HashPassword(password)
					if err != nil {
						return err, nil
					}
					db.Model(&user).Where("reset_password_token = ? AND email = ?", resetPasswordToken, email).UpdateColumns(models.User{UpdatedAt: updatedAt, ResetPasswordDate: &updatedAt})
					return &user, nil
				},
			},
		},
	},
)

var queryAdminType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"User": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, err := strconv.ParseInt(p.Args["id"].(string), 10, 64)
					db := db.ConnectGORM()
					db.SingularTable(true)
					if err == nil {
						user := models.User{}
						user.Id = idQuery
						log.Print(idQuery)
						db.First(&user)
						fmt.Printf(">>> %v", user)
						return &user, nil
					}

					return nil, nil
				},
			},

			"UserPts": &graphql.Field{
				Type: graphql.NewList(userPtsType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.ConnectGORM()
					db.SingularTable(true)

					userPts := []models.UserPts{}
					err := db.Find(&userPts).Error
					if err == nil {
						return &userPts, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"UsersAdmin": &graphql.Field{
				Type: graphql.NewList(userType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.ConnectGORM()
					db.SingularTable(true)

					users := []models.User{}
					err := db.Find(&users).Error
					if err == nil {
						return &users, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"ProductAdmin": &graphql.Field{
				Type: graphql.NewList(productType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.ConnectGORM()
					db.SingularTable(true)

					products := []models.Product{}
					err := db.Find(&products).Error
					if err == nil {
						return &products, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"PaymentAdmin": &graphql.Field{
				Type: graphql.NewList(paymentType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.ConnectGORM()
					db.SingularTable(true)

					payments := []models.Payment{}
					err := db.Find(&payments).Error
					if err == nil {
						return &payments, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"PaymentItemAdmin": &graphql.Field{
				Type: graphql.NewList(paymentItemsType),
				Args: graphql.FieldConfigArgument{
					"paymentId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					paymentId, err := strconv.ParseInt(p.Args["paymentId"].(string), 10, 64)
					if err != nil {
						return nil, nil
					}
					db := db.ConnectGORM()
					db.SingularTable(true)

					paymentItems := []models.PaymentItems{}
					err = db.Where("payment_id = ?", paymentId).Find(&paymentItems).Error
					if err == nil {
						return &paymentItems, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"OutletAdmin": &graphql.Field{
				Type: graphql.NewList(outletType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.ConnectGORM()
					db.SingularTable(true)

					outlet := []models.Outlet{}
					err := db.Find(&outlet).Error
					if err == nil {
						return &outlet, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"MembershipAdmin": &graphql.Field{
				Type: graphql.NewList(membershipType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.ConnectGORM()
					db.SingularTable(true)

					membership := []models.Membership{}
					err := db.Find(&membership).Error
					if err == nil {
						return &membership, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"MembershipProductAdmin": &graphql.Field{
				Type: graphql.NewList(membershipProductType),
				Args: graphql.FieldConfigArgument{
					"membershipId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					membershipId, err := strconv.ParseInt(p.Args["membershipId"].(string), 10, 64)
					if err != nil {
						return nil, nil
					}
					db := db.ConnectGORM()
					db.SingularTable(true)

					membershipProduct := []models.MembershipProduct{}
					err = db.Where("membership_id = ?", membershipId).Find(&membershipProduct).Error
					if err == nil {
						return &membershipProduct, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},
		},
	},
)

var rootAdminMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// User Pts
			"createUserPts": &graphql.Field{
				Type:        userPtsType,
				Description: "Create user points",
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"allocatedPts": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"productId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					userId, _ := params.Args["userId"].(string)
					fmt.Printf("<<<< userId: %s", userId)
					allocatedPts, _ := params.Args["allocatedPts"].(string)
					productId, _ := params.Args["productId"].(string)
					db := db.ConnectGORM()
					db.SingularTable(true)

					createdAt := time.Now()
					userPts := models.UserPts{}
					n, _ := strconv.ParseInt(userId, 10, 64)
					userPts.UserID = n
					n2, _ := strconv.ParseInt(productId, 10, 64)
					userPts.ProductID = n2
					userPts.CreatedAt = createdAt
					n3, _ := strconv.ParseInt(allocatedPts, 10, 64)
					userPts.AllocatedPts = n3
					if errs := validator.Validate(userPts); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&userPts)
					return &userPts, nil
				},
			},

			// Product
			"createProduct": &graphql.Field{
				Type:        productType,
				Description: "Create Product",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"pts": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"remarks": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"gender": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					pts, _ := params.Args["pts"].(string)
					price, _ := params.Args["price"].(string)
					remarks, _ := params.Args["remarks"].(string)
					gender, _ := params.Args["gender"].(string)

					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)

					createdAt := time.Now()
					product := models.Product{}
					product.Name = name
					n2, _ := strconv.ParseInt(pts, 10, 64)
					product.Pts = n2
					n, _ := strconv.ParseFloat(price, 32)
					product.Price = float32(n)
					product.CreatedAt = createdAt
					product.Remarks = remarks
					product.Gender = gender
					if errs := validator.Validate(product); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&product)
					return &product, nil
				},
			},

			// Payment
			"createPayment": &graphql.Field{
				Type:        paymentType,
				Description: "Create Payment",
				Args: graphql.FieldConfigArgument{
					"total": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"outlet": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"discount": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"paymentMode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"gst": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					total, _ := params.Args["total"].(string)
					outlet, _ := params.Args["outlet"].(string)
					discount, _ := params.Args["discount"].(string)
					userId, _ := params.Args["userId"].(string)
					paymentMode, _ := params.Args["paymentMode"].(string)
					gst, _ := params.Args["gst"].(string)

					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)

					createdAt := time.Now()
					payment := models.Payment{}
					n3, _ := strconv.ParseFloat(total, 32)
					n4 := createdAt.Format("20060102150405")
					payment.PaymentRefNo = "REF" + n4
					payment.Total = float32(n3)
					n2, _ := strconv.ParseInt(outlet, 10, 64)
					payment.Outlet = n2
					n, _ := strconv.ParseFloat(discount, 32)
					payment.Discount = float32(n)
					n5, _ := strconv.ParseInt(paymentMode, 10, 64)
					payment.CreatedAt = createdAt
					payment.PaymentMode = int(n5)
					n6, _ := strconv.ParseInt(userId, 10, 64)
					payment.User = n6
					n7, _ := strconv.ParseFloat(gst, 32)
					payment.Gst = float32(n7)
					if errs := validator.Validate(payment); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&payment)
					return &payment, nil
				},
			},

			// Payment
			"createPaymentItem": &graphql.Field{
				Type:        paymentItemsType,
				Description: "Create Payment Item",
				Args: graphql.FieldConfigArgument{
					"paymentId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"subTotal": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"productId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					paymentId, _ := params.Args["paymentId"].(string)
					subTotal, _ := params.Args["subTotal"].(string)
					productId, _ := params.Args["productId"].(string)

					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)

					createdAt := time.Now()
					paymentItems := models.PaymentItems{}
					n3, _ := strconv.ParseFloat(subTotal, 32)
					paymentItems.SubTotal = float32(n3)
					n2, _ := strconv.ParseInt(paymentId, 10, 64)
					paymentItems.Payment = n2
					n5, _ := strconv.ParseInt(productId, 10, 64)
					paymentItems.CreatedAt = createdAt
					paymentItems.Product = n5
					if errs := validator.Validate(paymentItems); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&paymentItems)
					return &paymentItems, nil
				},
			},

			// Outlet
			"createOutlet": &graphql.Field{
				Type:        outletType,
				Description: "Create Outlet",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"postalCode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"contactNo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"outletSupervisor": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"longtitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					postalCode, _ := params.Args["postalCode"].(string)
					contactNo, _ := params.Args["contactNo"].(string)
					email, _ := params.Args["email"].(string)
					outletSupervisor, _ := params.Args["outletSupervisor"].(string)
					longtitude, _ := params.Args["longtitude"].(string)
					latitude, _ := params.Args["latitude"].(string)

					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)

					createdAt := time.Now()
					outlet := models.Outlet{}
					outlet.Name = name
					n2, _ := strconv.ParseInt(postalCode, 10, 64)
					outlet.PostalCode = int(n2)
					outlet.CreatedAt = createdAt
					outlet.ContactNo = contactNo
					outlet.Email = email
					outlet.OutletSupervisor = outletSupervisor
					outlet.Longtitude = longtitude
					outlet.Latitude = latitude
					if errs := validator.Validate(outlet); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&outlet)
					return &outlet, nil
				},
			},

			// Membership
			"createMembership": &graphql.Field{
				Type:        membershipType,
				Description: "Create membership",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"ptsGiven": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					ptsGiven, _ := params.Args["ptsGiven"].(string)

					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)

					createdAt := time.Now()
					membership := models.Membership{}
					membership.Name = name
					n2, _ := strconv.ParseInt(ptsGiven, 10, 64)
					membership.PtsGiven = n2
					membership.CreatedAt = createdAt
					if errs := validator.Validate(membership); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&membership)
					return &membership, nil
				},
			},

			// Membership
			"createMembershipProduct": &graphql.Field{
				Type:        membershipProductType,
				Description: "Create membership product",
				Args: graphql.FieldConfigArgument{
					"productId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"membershipId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"count": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					productId, _ := params.Args["productId"].(string)
					membershipId, _ := params.Args["membershipId"].(string)
					count, _ := params.Args["count"].(string)
					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)

					createdAt := time.Now()
					membershipProduct := models.MembershipProduct{}
					n3, _ := strconv.ParseInt(productId, 10, 64)
					membershipProduct.Product = n3
					n2, _ := strconv.ParseInt(membershipId, 10, 64)
					membershipProduct.Membership = n2
					n, _ := strconv.ParseInt(count, 10, 64)
					membershipProduct.Count = n

					membershipProduct.CreatedAt = createdAt
					if errs := validator.Validate(membershipProduct); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&membershipProduct)
					return &membershipProduct, nil
				},
			},
			// update user points
			"updateUserPts": &graphql.Field{
				Type:        userPtsType,
				Description: "Update User Pts",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"allocatedPts": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"productId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					fmt.Printf("email == %s", emailAddress)
					id, _ := params.Args["id"].(string)
					userId, _ := params.Args["userId"].(string)
					allocatedPts, _ := params.Args["allocatedPts"].(string)
					productId, _ := params.Args["productId"].(string)
					updatedAt := time.Now()
					userPts := models.UserPts{}
					db := db.ConnectGORM()
					// Disable table name's pluralization globally
					db.SingularTable(true)
					n, _ := strconv.ParseInt(userId, 10, 64)
					n2, _ := strconv.ParseInt(productId, 10, 64)
					n3, _ := strconv.ParseInt(allocatedPts, 10, 64)
					db.Model(&userPts).Where("id = " + id).UpdateColumns(models.UserPts{UserID: n, AllocatedPts: n3, ProductID: n2, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					userPts.Id = n4
					return &userPts, nil
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

func ExecuteAdminQuery(query string, usrEmail string) *graphql.Result {
	emailAddress = usrEmail
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryAdminType,
			Mutation: rootAdminMutation,
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
