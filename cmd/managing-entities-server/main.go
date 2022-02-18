package main

import (
	"Jokaru-py/managingEntities/internal/handlers"
	"Jokaru-py/managingEntities/internal/store"
	"Jokaru-py/managingEntities/pkg/db"
	"Jokaru-py/managingEntities/pkg/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := router.New()

	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	connStore := store.NewConnStore(d)
	h := handlers.NewHandler(*connStore)

	h.Register(v1)

	r.Logger.Fatal(r.Start(":8585"))
}
