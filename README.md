## Install dependecies

```bash
go get github.com/labstack/echo
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
go get github.com/graphql-go/graphql
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
```

- Using Echo middleware framework

https://github.com/labstack/echox/tree/master/cookbook/jwt

- Using JWT go lang libary
  https://github.com/dgrijalva/jwt-go

```bash
curl -X GET  localhost:3000/hello
```

## Registration end point

```bash
curl -X POST -d "username=bunnyppl@gmail.com&password=password1234&contactNo=91450518" localhost:3000/api/register -H "Content-Type: application/x-www-form-urlencoded"

curl -X POST -d "username=bunnyppl@hotmail.com&password=password1234&contactNo=91450516" localhost:3000/api/register -H "Content-Type: application/x-www-form-urlencoded"

curl -X POST -d "username=bunnyppl@yahoo.com&password=password1234&contactNo=91450519" localhost:3000/api/register -H "Content-Type: application/x-www-form-urlencoded"
```

## Login end point

```bash
curl -X POST -d "email=bunnyppl@gmail.com&password=password1234" localhost:3000/api/login -H "Content-Type: application/x-www-form-urlencoded"

curl -X POST -d "email=bunnyppl@yahoo.com&password=password1234" localhost:3000/api/login -H "Content-Type: application/x-www-form-urlencoded"
```

## Common end point to access the database record

```bash
curl -X POST localhost:3000/api/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQyNTI4MTMyLCJuYW1lIjoiYnVubnlwcGxAZ21haWwuY29tIn0.K0NOmo2uEd10iKkhEy16gbPfZVfkT9KPLGVyXf7bkm4"
```

```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQyNTI4MTMyLCJuYW1lIjoiYnVubnlwcGxAZ21haWwuY29tIn0.K0NOmo2uEd10iKkhEy16gbPfZVfkT9KPLGVyXf7bkm4" -X POST -d '{ query: User(id: "1") { id, firstname, lastname }}' http://localhost:3000/api/restricted
```

An identifier may be exported to permit access to it from another package. An identifier is exported if both:

## Graphql

### Get user by id

```
{
  query: User(id: "1") {id, firstname, lastname, email, isAdmin, createdAt,
  	updatedAt, contactNo, password, hairloss_treatment_cnt, treatment_cnt, cutCnt,
  expiry_date, ptsBalance}
}

{
  query: User(id: "3") {id, firstname, lastname, email, isAdmin, createdAt,
  	updatedAt, contactNo, password, hairloss_treatment_cnt, treatment_cnt, cutCnt,
  expiry_date, ptsBalance}
}
```

### Create user's point

```
mutation {
  createUserPts(userId: "1", allocatedPts: "11", productId: "3") {
    userId
    allocatedPts
    productId
  }
}
```

### Create product

```
mutation {
  createProduct(userId: 1, allocatedPts: 11, productId: 3) {
    userId
    allocatedPts
    productId
  }
}
```

### Get all user' points

```
{
  query: UserPtsHistory(dummy: "") {id, userId, productId, allocatedPts}
}
```

### Update product

```
mutation {
  updateUserPts(id: "12", userId: "3", allocatedPts: "11", productId: "232") {
    id
    userId
    allocatedPts
    productId
  }
}
```
