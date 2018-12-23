package graphql

import "github.com/graphql-go/graphql"

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
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"bio": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
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
			"remarks": &graphql.Field{
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
