package main

import (
	"log"
	"os"

	db "go-jwt/api/database"
	_log "go-jwt/api/helpers"
	_routes "go-jwt/api/routes"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	port string
)

func init() {

	LoadEnv() // default to .env file

	db.CreateDBConnection() //connect to database

	port = os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type DocUser struct {
	Id        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `json:"email"`
	Phone     string             `bson:"phone"`
}

func main() {

	// middleware

	_log.LoggerInit()

	router := _routes.SetUpRoutes() // return router object

	router.Run(":" + port) // listen and serve on 0.0.0.0:{PORT}
}
