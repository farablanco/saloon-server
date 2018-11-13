## Install dependecies

```
go get github.com/labstack/echo
 go get github.com/dgrijalva/jwt-go
```

https://github.com/labstack/echox/tree/master/cookbook/jwt
https://github.com/dgrijalva/jwt-go


curl -X POST -d "username=test&password=test" localhost:3000/login -H "Content-Type: application/x-www-form-urlencoded"

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQyMzUwMDY1LCJuYW1lIjoidGVzdCJ9.4Qt14esoxlzuR047ARipfwh2P_1n3JYJFTXMC80HjMI"}

curl -X POST localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQyMzUwMDY1LCJuYW1lIjoidGVzdCJ9.4Qt14esoxlzuR047ARipfwh2P_1n3JYJFTXMC80HjMI"

Welcome test!

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