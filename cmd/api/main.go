package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/hawful70/example-social/internal/db"
	"github.com/hawful70/example-social/internal/store"
	"github.com/hawful70/example-social/utils"
	"go.uber.org/zap"
)

const version = "1.1.0"

func main() {
	utils.InitEnv()

	cfg := config{
		addr:   ":" + os.Getenv("ADDR"),
		apiURL: os.Getenv("EXTERNAL_URL"),
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: 10,
			maxIdleConns: 10,
			maxIdleTime:  "15m",
		},
		redisCfg: redisConfig{
			addr:    os.Getenv("REDIS_ADDR"),
			pw:      os.Getenv("REDIS_PW"),
			db:      0,
			enabled: false,
		},
		env: os.Getenv("ENV"),
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer func(logger *zap.SugaredLogger) {
		err := logger.Sync()
		if err != nil {
			logger.Fatal(err)
		}
	}(logger)

	// Main Database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(db)
	logger.Info("database connection pool established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
