package migration

import (
	"log"
	"os"
	// _ "github.com/eminetto/clean-archecture-go/migrations"
	// migrate "github.com/eminetto/mongo-migrate"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("Missing options: up or down")
		return
	}

	option := os.Args[1]

	//session, err := mgo.Dial(os.Getenv("MONGODB_HOST"))

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	switch option {
	case "new":
		log.Printf("New migration created: %s\n", "")
	case "up":
		log.Printf("Migration up: %s\n", "")
	case "down":
		log.Printf("Migration down: %s\n", "")
	}
}
