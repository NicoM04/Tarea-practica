package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var TaskCollection *mongo.Collection

func InitMongo() {
	// Establece un contexto para la conexión
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conecta a MongoDB en localhost:27017
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(GetMongoURL()))

	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Elige la base de datos y colección a usar desde .env -> config.go
	TaskCollection = Client.Database(GetDatabaseName()).Collection(GetCollectionName())

	fmt.Println("Connected to MongoDB")
}
