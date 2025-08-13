package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	Port             = 0
)

func ConfigureEnvironment() {

	var err error

	if err = godotenv.Load(); err != nil {

		log.Fatal("Error to load environment configurations")
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {

		Port = 9000
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
