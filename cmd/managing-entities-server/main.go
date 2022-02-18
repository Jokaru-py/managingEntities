package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// r := router.New()

	// v1 := r.Group("/api/estate-agency")

	// d := db.New()
	// db.AutoMigrate(d)

	// connStore := .NewProfileStore(d)
	// h := handlers.NewHandler(scs, sms)

	// h.Register(v1)

	// r.Logger.Fatal(r.Start(":8585"))
}
