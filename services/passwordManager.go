package services

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func generateHashedPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePassword(userEnteredPassword string, databasePassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(userEnteredPassword))
	if err != nil {
		return err
	}

	return nil
}
