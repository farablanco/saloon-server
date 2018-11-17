package graphql

import (
	"fmt"
	"log"
	"strconv"

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
		},
	},
)

func ExecuteQuery(query string) *graphql.Result {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
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
