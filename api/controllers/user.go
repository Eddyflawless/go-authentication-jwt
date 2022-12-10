package controllers

import (
	"context"
	"go-jwt/api/database"
	"go-jwt/api/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *gin.Context) {

	ctx := context.Background()

	db := database.OpenCollection("users")

	cursor, err := db.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var users []models.FoundUser

	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {

	// For reference
	xUserId, ok := c.Request.Context().Value("x_user_id").(string)
	if !ok {
		c.Abort()
		return
	}

	log.Printf("Authenticated user %v\n", xUserId)

	userId, err := database.ConvertObjectIDToHex(c.Param("userId"))

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
