package main

import (
	"log"
	"os"

	"github.com/hawful70/example-social/internal/db"
	"github.com/hawful70/example-social/internal/store"
	"github.com/hawful70/example-social/utils"
)

func main() {
	utils.Init()

	cfg := config{
		addr: ":" + os.Getenv("ADDR"),
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: 10,
			maxIdleConns: 10,
			maxIdleTime:  "15m",
		},
	}

	// Main database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
