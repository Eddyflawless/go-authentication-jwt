package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go-jwt/api/database"
	"go-jwt/api/helpers"
	"go-jwt/api/helpers/validators"
	"go-jwt/api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginUser struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `json:"email" validate:"required,email"`
	Password string             `json:"password" validate:"required"`
}

func Login(c *gin.Context) {

	var user LoginUser
	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	validationErr := validate.Struct(user)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	db := database.OpenCollection("users")

	fmt.Printf("Email: %v and password is %v \n", user.Email, user.Password)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	err := db.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.UserId)

	foundUser.Token = &token
	foundUser.RefreshToken = &refreshToken

	c.JSON(http.StatusOK, gin.H{"msg": "login --", "token": foundUser.Token})

}

func Signup(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validators.SignUpValidator(user)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	fmt.Printf("Email: %v and password is %v \n", *user.Email, *user.Password)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	db := database.OpenCollection("users")

	// check email availability

	count, err := db.CountDocuments(context.TODO(), bson.D{{"email", user.Email}})
	defer cancel()

	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occured"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email or phone number is already in use. AU001"})
		return
	}

	// check phone availability
	count, err = db.CountDocuments(ctx, bson.M{"phone": *user.Phone})
	defer cancel()

	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occured"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email or phone number is already in use. AU002"})
		return
	}

	// generate password
	password, _ := helpers.HashPassword(*user.Password)
	log.Printf("Password generated %v\n", password)
	user.Password = &password

	user.CreatedAt, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC3339))
	user.Id = primitive.NewObjectID()
	user.UserId = user.Id.Hex()

	//
	token, refreshToken, _ := helpers.GenerateAllTokens(user.UserId)

	user.Token = &token
	user.RefreshToken = &refreshToken

	resultInsertionNumber, err := db.InsertOne(ctx, user)
	defer cancel()

	if err != nil {
		msg := fmt.Sprintf("User was not added")
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, resultInsertionNumber)

}

func Me(c *gin.Context) {

	x_userId, ok := c.Request.Context().Value("x_user_id").(string)
	if !ok {
		c.Abort()
		return
	}

	log.Printf(" x_userId is %v\n", x_userId)

	userId, err := database.ConvertObjectIDToHex(x_userId)

	if err != nil {
		// log warning
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.OpenCollection("users")

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	var user models.FoundUser

	err = db.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	defer cancel()

	if err != nil {
		// log error

		// spit out a human-readable error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
