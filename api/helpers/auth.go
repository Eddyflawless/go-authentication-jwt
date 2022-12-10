package helpers

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, password string) (bool, string) {
	check := true
	msg := ""

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		msg = fmt.Sprintf("Email of password is incorrect")
		check = false
	}

	return check, msg

}

func CheckUserType(c *gin.Context, role string) (err error) {

	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("Unauthorized access to resource")
	}

	return err

}

func MatchUserTypeToUid(c *gin.Context, userId string, expectedUserType string) (err error) {

	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == expectedUserType && uid != userId {
		err = errors.New("Unauthorized access to resource")
	}

	return err

}
