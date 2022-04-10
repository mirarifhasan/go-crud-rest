package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	var connectionString = os.Getenv("DB_URL")
	var dbName = os.Getenv("DB_NAME")
	var colName = os.Getenv("COL_NAME")

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("error")
	}
	fmt.Println("MongoDB connection success")

	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("collection instance ready")
}
