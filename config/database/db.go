package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""
const dbName = "arif_task"
const colName = "todos"

var Collection *mongo.Collection

func ConnectDB() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("error")
	}
	fmt.Println("MongoDB connection success")

	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("collection instance ready")
}
