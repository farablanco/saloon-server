## Install dependecies

```bash
go get github.com/labstack/echo
 go get github.com/dgrijalva/jwt-go
```

Using echo middleware framework

https://github.com/labstack/echox/tree/master/cookbook/jwt

Using jwt go lang libary
https://github.com/dgrijalva/jwt-go

```bash
curl -X GET  localhost:3000/hello
```

```bash
curl -X POST -d "username=bunnyppl@gmail.com&password=password@123" localhost:3000/login -H "Content-Type: application/x-www-form-urlencoded"
```

```bash
curl -X POST localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQyNTI4MTMyLCJuYW1lIjoiYnVubnlwcGxAZ21haWwuY29tIn0.K0NOmo2uEd10iKkhEy16gbPfZVfkT9KPLGVyXf7bkm4"
```

```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQyNTI4MTMyLCJuYW1lIjoiYnVubnlwcGxAZ21haWwuY29tIn0.K0NOmo2uEd10iKkhEy16gbPfZVfkT9KPLGVyXf7bkm4" -X POST -d '{ query: User(id: "1") { id, firstname, lastname }}' http://localhost:3000/restricted
```

An identifier may be exported to permit access to it from another package. An identifier is exported if both: