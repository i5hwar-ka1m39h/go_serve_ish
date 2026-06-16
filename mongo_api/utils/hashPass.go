package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)




func HashPass(password string) string{
	hash , err :=bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil{
		log.Println("error converting password hash", err)
		return  password
	}

	return  string(hash)
}