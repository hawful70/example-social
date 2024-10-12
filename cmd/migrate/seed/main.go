package main

import (
	"log"
	"os"

	"github.com/hawful70/example-social/internal/db"
	"github.com/hawful70/example-social/internal/store"
	"github.com/hawful70/example-social/utils"
)

func main() {
	utils.InitEnv()
	addr := os.Getenv("DB_ADDR")
	conn, err := db.New(addr, 10, 10, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store, conn)
}
