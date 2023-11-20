package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApiURL   = ""
	Port     = 0
	HashKey  []byte
	BlockKey []byte
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT")); 
	if err != nil {
		log.Fatal(err)
	}

	ApiURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}