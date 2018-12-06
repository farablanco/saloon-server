// author : Kenneth Phang
package graphql

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"../db"
	"../models"
	"github.com/graphql-go/graphql"
)

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
			"pts": &graphql.Field{
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
			"deletedAt": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
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

var pointsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Points",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"productId": &graphql.Field{
				Type: graphql.Int,
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
			"price": &graphql.Field{
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
			"UserPts": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, err := strconv.ParseInt(p.Args["userId"].(string), 10, 64)
					db := db.ConnectGORM()
					db.SingularTable(true)
					if err == nil {
						userPts := models.UserPts{}
						userPts.UserID = idQuery
						log.Print(idQuery)
						db.Find(&userPts)
						fmt.Printf(">>> %v", userPts)
						return &userPts, nil
					}

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
			// User Pts
			"createUserPts": &graphql.Field{
				Type:        userPtsType,
				Description: "Create user points",
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"pts": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"productId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					userId, _ := params.Args["userId"].(string)
					fmt.Printf("<<<< userId: %s", userId)
					pts, _ := params.Args["pts"].(string)
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
					n3, _ := strconv.ParseInt(pts, 10, 64)
					userPts.AllocatedPts = n3
					db.Save(&userPts)
					return &userPts, nil
				},
			},

			// Product
			"createProduct": &graphql.Field{
				Type:        userPtsType,
				Description: "Create user points",
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"pts": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"productId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					userId, _ := params.Args["userId"].(int64)
					pts, _ := params.Args["pts"].(int64)
					productId, _ := params.Args["productId"].(int64)
					db := db.ConnectGORM()
					db.SingularTable(true)

					createdAt := time.Now()
					userPts := models.UserPts{}
					userPts.UserID = userId
					userPts.ProductID = productId
					userPts.CreatedAt = createdAt
					userPts.AllocatedPts = pts
					db.Save(&userPts)
					return &userPts, nil
				},
			},
		},
	},
)

func ExecuteQuery(query string) *graphql.Result {
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
