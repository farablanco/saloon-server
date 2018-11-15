package graphql

import (
	"fmt"
	"log"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/kenken64/saloon-server/db"
	"github.com/kenken64/saloon-server/models"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
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
				Type: graphql.Int,
			},
			"deleted_at": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
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
					if err == nil {
						db := db.ConnectGORM()
						db.SingularTable(true)
						user := models.User{}
						user.Id = idQuery
						db.First(&user)
						log.Print(idQuery)
						log.Print(&user)
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
