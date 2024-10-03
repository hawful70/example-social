package main

import (
	"log"
	"os"

	"github.com/hawful70/example-social/internal/store"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	Init()

	cfg := config{
		addr: ":" + os.Getenv("ADDR"),
	}

	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
