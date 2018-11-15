## Install dependecies

```
go get github.com/labstack/echo
 go get github.com/dgrijalva/jwt-go
```

https://github.com/labstack/echox/tree/master/cookbook/jwt
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

```
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQyNTI4MTMyLCJuYW1lIjoiYnVubnlwcGxAZ21haWwuY29tIn0.K0NOmo2uEd10iKkhEy16gbPfZVfkT9KPLGVyXf7bkm4" -X POST -d '{ query: User(id: "1") { id, firstname, lastname }}' http://localhost:3000/restricted
```

CREATE TABLE IF NOT EXISTS `crm`.`user` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `firstname` VARCHAR(100) NOT NULL,
  `lastname` VARCHAR(100) NOT NULL,
  `password` VARCHAR(200) NOT NULL,
  `email` VARCHAR(250) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

INSERT INTO user (firstname, lastname, password, email ) VALUES ("Kenneth", "Phang", "password@123", "bunnyppl@hotmail.com");
INSERT INTO user (firstname, lastname, password, email ) VALUES ("Kenneth", "Phang", "password@123", "bunnyppl@gmail.com");
INSERT INTO user (firstname, lastname, password, email ) VALUES ("Kenneth", "Phang", "password@123", "bunnyppl@yahoo.com");


An identifier may be exported to permit access to it from another package. An identifier is exported if both: