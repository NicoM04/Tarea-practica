package services

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}
}

func GetMongoURL() string {
	return os.Getenv("DB_URL")
}

func GetDatabaseName() string {
	return os.Getenv("DB_NAME")
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetCollectionName() string {
	return os.Getenv("DB_COLLECTION")
}
