package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mrzalr/eatery-hub/internal/server"
	"github.com/mrzalr/eatery-hub/pkg/database/mysql"
)

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.New()
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(db)
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
