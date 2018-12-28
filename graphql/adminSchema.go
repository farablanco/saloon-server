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

var queryAdminType = graphql.NewObject(
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
					db.First(&user)
					fmt.Printf(">>> %v", user)
					return &user, nil
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
					db := db.DBManager()

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
					db := db.DBManager()

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
					db := db.DBManager()

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
					db := db.DBManager()

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
					db := db.DBManager()

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
					db := db.DBManager()

					outlet := []models.Outlet{}
					err := db.Find(&outlet).Error
					if err == nil {
						return &outlet, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"CompanyAdmin": &graphql.Field{
				Type: graphql.NewList(outletType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.DBManager()

					company := []models.Company{}
					err := db.Find(&company).Error
					if err == nil {
						return &company, nil
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
					db := db.DBManager()
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
					db := db.DBManager()

					membershipProduct := []models.MembershipProduct{}
					err = db.Where("membership_id = ?", membershipId).Find(&membershipProduct).Error
					if err == nil {
						return &membershipProduct, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"BookingAdmin": &graphql.Field{
				Type: graphql.NewList(bookingType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.DBManager()

					booking := []models.Booking{}
					err := db.Find(&booking).Error
					if err == nil {
						return &booking, nil
					}
					log.Fatal(err)
					return nil, nil
				},
			},

			"HairDresserAdmin": &graphql.Field{
				Type: graphql.NewList(hairDresserType),
				Args: graphql.FieldConfigArgument{
					"dummy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.DBManager()

					hairDresser := []models.Hairdresser{}
					err := db.Find(&hairDresser).Error
					if err == nil {
						return &hairDresser, nil
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
					db := db.DBManager()
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

					db := db.DBManager()

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

					db := db.DBManager()
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

					db := db.DBManager()

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

					db := db.DBManager()

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

			"createCompany": &graphql.Field{
				Type:        outletType,
				Description: "Create Outlet",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"registrationNo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"contactNo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"companyOwner": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					registrationNo, _ := params.Args["registrationNo"].(string)
					contactNo, _ := params.Args["contactNo"].(string)
					email, _ := params.Args["email"].(string)
					companyOwner, _ := params.Args["companyOwner"].(string)

					db := db.DBManager()

					createdAt := time.Now()
					company := models.Company{}
					company.Name = name
					company.RegistrationCode = registrationNo
					company.CreatedAt = createdAt
					company.ContactNo = contactNo
					company.Email = email
					company.CompanyOwner = companyOwner

					if errs := validator.Validate(company); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&company)
					return &company, nil
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

					db := db.DBManager()

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
					db := db.DBManager()

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

			// hairdresser
			"createHairdresser": &graphql.Field{
				Type:        membershipProductType,
				Description: "Create hairdresser",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"yearsOfExp": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"joinedDate": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"contactNo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					yearsOfExp, _ := params.Args["yearsOfExp"].(string)
					joinedDate, _ := params.Args["joinedDate"].(string)
					contactNo, _ := params.Args["contactNo"].(string)
					db := db.DBManager()

					createdAt := time.Now()
					hairdresser := models.Hairdresser{}
					n3, _ := strconv.ParseInt(yearsOfExp, 10, 64)
					hairdresser.YearsOfExp = n3
					t, err := time.Parse("2006-01-02 3:04PM", joinedDate)
					if err != nil {
						return &err, nil
					}
					hairdresser.JoinedDate = &t
					hairdresser.Name = name
					hairdresser.ContactNo = contactNo

					hairdresser.CreatedAt = createdAt
					if errs := validator.Validate(hairdresser); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&hairdresser)
					return &hairdresser, nil
				},
			},

			// booking
			"createBooking": &graphql.Field{
				Type:        bookingType,
				Description: "Create booking",
				Args: graphql.FieldConfigArgument{
					"hairDresserId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"bookingDateTime": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"outlet": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"customer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"remarks": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					hairDresserId, _ := params.Args["hairDresserId"].(string)
					bookingDatetime, _ := params.Args["bookingDateTime"].(string)
					outletId, _ := params.Args["outlet"].(string)
					userId, _ := params.Args["customer"].(string)
					remarks, _ := params.Args["remarks"].(string)

					db := db.DBManager()

					createdAt := time.Now()
					booking := models.Booking{}
					n3, _ := strconv.ParseInt(hairDresserId, 10, 64)
					booking.Hairdresser = n3
					fmt.Printf("bookingDatetime > %s\n", bookingDatetime)
					t, err := time.Parse("2006-01-02 3:04PM", bookingDatetime)
					if err != nil {
						return &err, nil
					}
					booking.BookingDatetime = &t
					fmt.Printf("%s", remarks)
					booking.Remarks = remarks
					n2, _ := strconv.ParseInt(outletId, 10, 64)
					booking.Outlet = n2
					n1, _ := strconv.ParseInt(userId, 10, 64)
					booking.Customer = n1

					booking.CreatedAt = createdAt
					if errs := validator.Validate(booking); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Save(&booking)
					return &booking, nil
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

					db := db.DBManager()

					n, _ := strconv.ParseInt(userId, 10, 64)
					n2, _ := strconv.ParseInt(productId, 10, 64)
					n3, _ := strconv.ParseInt(allocatedPts, 10, 64)
					if errs := validator.Validate(userPts); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&userPts).Where("id = " + id).UpdateColumns(models.UserPts{UserID: n, AllocatedPts: n3, ProductID: n2, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					userPts.Id = n4
					return &userPts, nil
				},
			},

			// update product
			"updateProduct": &graphql.Field{
				Type:        productType,
				Description: "Update Product",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
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
					id, _ := params.Args["id"].(string)
					name, _ := params.Args["name"].(string)
					pts, _ := params.Args["pts"].(string)
					price, _ := params.Args["price"].(string)
					remarks, _ := params.Args["remarks"].(string)
					gender, _ := params.Args["gender"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					product := models.Product{}
					n2, _ := strconv.ParseInt(pts, 10, 64)
					n, _ := strconv.ParseFloat(price, 32)
					if errs := validator.Validate(product); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&product).Where("id = " + id).UpdateColumns(models.Product{Name: name, Price: float32(n), Pts: n2, Remarks: remarks, Gender: gender, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					product.Id = n4
					return &product, nil
				},
			},

			"updatePayment": &graphql.Field{
				Type:        paymentType,
				Description: "update Payment",
				Args: graphql.FieldConfigArgument{
					"paymentRefNo": &graphql.ArgumentConfig{
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					paymentRefNo, _ := params.Args["paymentRefNo"].(string)
					outlet, _ := params.Args["outlet"].(string)
					discount, _ := params.Args["discount"].(string)
					userId, _ := params.Args["userId"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					payment := models.Payment{}
					n2, _ := strconv.ParseInt(outlet, 10, 64)
					n, _ := strconv.ParseFloat(discount, 32)
					payment.Discount = float32(n)
					n6, _ := strconv.ParseInt(userId, 10, 64)

					if errs := validator.Validate(payment); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&payment).Where("PaymentRefNo = " + paymentRefNo).UpdateColumns(models.Payment{Outlet: n2, Discount: float32(n), User: n6, UpdatedAt: updatedAt})
					payment.PaymentRefNo = paymentRefNo
					return &payment, nil
				},
			},

			// update payment item
			"updatePaymentItem": &graphql.Field{
				Type:        paymentItemsType,
				Description: "Update Payment Item",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
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
					id, _ := params.Args["id"].(string)
					subTotal, _ := params.Args["subTotal"].(string)
					productId, _ := params.Args["productId"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					paymentItems := models.PaymentItems{}
					n3, _ := strconv.ParseFloat(subTotal, 32)
					n5, _ := strconv.ParseInt(productId, 10, 64)
					if errs := validator.Validate(paymentItems); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&paymentItems).Where("id = " + id).UpdateColumns(models.PaymentItems{SubTotal: float32(n3), Product: n5, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					paymentItems.Id = n4
					return &paymentItems, nil
				},
			},

			// update outlet
			"updateOutlet": &graphql.Field{
				Type:        outletType,
				Description: "Update Outlet",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
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
					id, _ := params.Args["id"].(string)
					name, _ := params.Args["name"].(string)
					postalCode, _ := params.Args["postalCode"].(string)
					contactNo, _ := params.Args["contactNo"].(string)
					email, _ := params.Args["email"].(string)
					outletSupervisor, _ := params.Args["outletSupervisor"].(string)
					longtitude, _ := params.Args["longtitude"].(string)
					latitude, _ := params.Args["latitude"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					outlet := models.Outlet{}
					n2, _ := strconv.ParseInt(postalCode, 10, 64)
					if errs := validator.Validate(outlet); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&outlet).Where("id = " + id).UpdateColumns(models.Outlet{Name: name, PostalCode: int(n2), ContactNo: contactNo, Email: email, OutletSupervisor: outletSupervisor, Longtitude: longtitude, Latitude: latitude, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					outlet.Id = n4
					return &outlet, nil
				},
			},

			// update company
			"updateCompany": &graphql.Field{
				Type:        companyType,
				Description: "Update Company",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"registrationCode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"contactNo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"companyOwner": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)
					name, _ := params.Args["name"].(string)
					registrationCode, _ := params.Args["registrationCode"].(string)
					contactNo, _ := params.Args["contactNo"].(string)
					email, _ := params.Args["email"].(string)
					companyOwner, _ := params.Args["companyOwner"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					company := models.Company{}
					if errs := validator.Validate(company); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&company).Where("id = " + id).UpdateColumns(models.Company{Name: name, RegistrationCode: registrationCode, ContactNo: contactNo, Email: email, CompanyOwner: companyOwner, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					company.Id = n4
					return &company, nil
				},
			},

			// update Membership
			"updateMembership": &graphql.Field{
				Type:        membershipType,
				Description: "Update membership",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"ptsGiven": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)
					name, _ := params.Args["name"].(string)
					ptsGiven, _ := params.Args["ptsGiven"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					membership := models.Membership{}
					n2, _ := strconv.ParseInt(ptsGiven, 10, 64)
					if errs := validator.Validate(membership); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&membership).Where("id = " + id).UpdateColumns(models.Membership{Name: name, PtsGiven: n2, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					membership.Id = n4
					return &membership, nil
				},
			},

			"updateMembershipProduct": &graphql.Field{
				Type:        membershipProductType,
				Description: "Update membership product",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
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
					id, _ := params.Args["id"].(string)
					productId, _ := params.Args["productId"].(string)
					membershipId, _ := params.Args["membershipId"].(string)
					count, _ := params.Args["count"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					membershipProduct := models.MembershipProduct{}
					n3, _ := strconv.ParseInt(productId, 10, 64)
					n2, _ := strconv.ParseInt(membershipId, 10, 64)
					n, _ := strconv.ParseInt(count, 10, 64)

					if errs := validator.Validate(membershipProduct); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&membershipProduct).Where("id = " + id).UpdateColumns(models.MembershipProduct{Product: n3, Membership: n2, Count: n, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					membershipProduct.Id = n4
					return &membershipProduct, nil
				},
			},

			// update hairdresser
			"updateHairdresser": &graphql.Field{
				Type:        membershipProductType,
				Description: "update hairdresser",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"yearsOfExp": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"joinedDate": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"contactNo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)
					name, _ := params.Args["name"].(string)
					yearsOfExp, _ := params.Args["yearsOfExp"].(string)
					joinedDate, _ := params.Args["joinedDate"].(string)
					contactNo, _ := params.Args["contactNo"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					hairdresser := models.Hairdresser{}
					n3, _ := strconv.ParseInt(yearsOfExp, 10, 64)
					t, err := time.Parse("2006-01-02 3:04PM", joinedDate)
					if err != nil {
						return &err, nil
					}

					if errs := validator.Validate(hairdresser); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&hairdresser).Where("id = " + id).UpdateColumns(models.Hairdresser{YearsOfExp: n3, JoinedDate: &t, Name: name, ContactNo: contactNo, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					hairdresser.Id = n4
					return &hairdresser, nil
				},
			},

			// update booking
			"updateBooking": &graphql.Field{
				Type:        bookingType,
				Description: "update booking",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"hairdresserId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"bookingDatetime": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"outletId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"remarks": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)
					hairdresserId, _ := params.Args["hairdresserId"].(string)
					bookingDatetime, _ := params.Args["bookingDatetime"].(string)
					outletId, _ := params.Args["outletId"].(string)
					userId, _ := params.Args["userId"].(string)
					remarks, _ := params.Args["remarks"].(string)

					db := db.DBManager()

					updatedAt := time.Now()
					booking := models.Booking{}
					n3, _ := strconv.ParseInt(hairdresserId, 10, 64)
					t, err := time.Parse("2006-01-02 3:04PM", bookingDatetime)
					if err != nil {
						return &err, nil
					}
					n2, _ := strconv.ParseInt(outletId, 10, 64)
					n1, _ := strconv.ParseInt(userId, 10, 64)

					if errs := validator.Validate(booking); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&booking).Where("id = " + id).UpdateColumns(models.Booking{Hairdresser: n3, BookingDatetime: &t, Remarks: remarks, Outlet: n2, Customer: n1, UpdatedAt: updatedAt})
					n4, _ := strconv.ParseInt(id, 10, 64)
					booking.Id = n4
					return &booking, nil
				},
			},

			// delete user pts
			"deleteUserPts": &graphql.Field{
				Type:        userPtsType,
				Description: "delete User Pts",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					fmt.Printf("email == %s", emailAddress)
					id, _ := params.Args["id"].(string)

					userPts := models.UserPts{}
					db := db.DBManager()

					if errs := validator.Validate(userPts); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}

					db.Model(&userPts).Where("id = " + id).Delete(&userPts)
					n4, _ := strconv.ParseInt(id, 10, 64)
					userPts.Id = n4
					return &userPts, nil
				},
			},

			// delete product
			"deleteProduct": &graphql.Field{
				Type:        productType,
				Description: "delete Product",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)

					db := db.DBManager()

					product := models.Product{}
					if errs := validator.Validate(product); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&product).Where("id = " + id).Delete(&product)
					n4, _ := strconv.ParseInt(id, 10, 64)
					product.Id = n4
					return &product, nil
				},
			},

			"deleteOutlet": &graphql.Field{
				Type:        outletType,
				Description: "delete Outlet",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)

					db := db.DBManager()

					outlet := models.Outlet{}
					if errs := validator.Validate(outlet); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&outlet).Where("id = " + id).Delete(&outlet)
					n4, _ := strconv.ParseInt(id, 10, 64)
					outlet.Id = n4
					return &outlet, nil
				},
			},

			// delete Membership
			"deleteMembership": &graphql.Field{
				Type:        membershipType,
				Description: "delete membership",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)

					db := db.DBManager()

					membership := models.Membership{}
					if errs := validator.Validate(membership); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&membership).Where("id = " + id).Delete(&membership)
					n4, _ := strconv.ParseInt(id, 10, 64)
					membership.Id = n4
					return &membership, nil
				},
			},

			"deleteMembershipProduct": &graphql.Field{
				Type:        membershipProductType,
				Description: "delete membership product",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)

					db := db.DBManager()

					membershipProduct := models.MembershipProduct{}

					if errs := validator.Validate(membershipProduct); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&membershipProduct).Where("id = " + id).Delete(&membershipProduct)
					n4, _ := strconv.ParseInt(id, 10, 64)
					membershipProduct.Id = n4
					return &membershipProduct, nil
				},
			},

			// delete hairdresser
			"deleteHairdresser": &graphql.Field{
				Type:        membershipProductType,
				Description: "delete hairdresser",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)

					db := db.DBManager()

					hairdresser := models.Hairdresser{}

					if errs := validator.Validate(hairdresser); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&hairdresser).Where("id = " + id).Delete(&hairdresser)
					n4, _ := strconv.ParseInt(id, 10, 64)
					hairdresser.Id = n4
					return &hairdresser, nil
				},
			},

			// delete booking
			"deleteBooking": &graphql.Field{
				Type:        bookingType,
				Description: "delete booking",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)

					db := db.DBManager()

					booking := models.Booking{}

					if errs := validator.Validate(booking); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&booking).Where("id = " + id).Delete(&booking)
					n4, _ := strconv.ParseInt(id, 10, 64)
					booking.Id = n4
					return &booking, nil
				},
			},

			// delete company
			"deleteCompany": &graphql.Field{
				Type:        companyType,
				Description: "delete company",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)

					db := db.DBManager()

					company := models.Company{}

					if errs := validator.Validate(company); errs != nil {
						// values not valid, deal with errors here
						return &errs, nil
					}
					db.Model(&company).Where("id = " + id).Delete(&company)
					n4, _ := strconv.ParseInt(id, 10, 64)
					company.Id = n4
					return &company, nil
				},
			},
		},
	},
)

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
