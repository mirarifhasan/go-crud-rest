package services

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"project/config/database"
	"project/src/todo/dtos"
	"project/src/todo/models"
)

type todo struct {
	ID        int    `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{
		ID:        1,
		Item:      "Clean room",
		Completed: false,
	},
	{
		ID:        2,
		Item:      "Read book",
		Completed: false,
	},
	{
		ID:        3,
		Item:      "Buy grocery",
		Completed: false,
	},
}

var collection *mongo.Collection

func init() {
	collection = database.Collection
}

func GetTodos(context *gin.Context) {
	cur, err := database.Collection.Find(context, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var todos []models.Todo

	for cur.Next(context) {
		var todo models.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}
	defer cur.Close(context)

	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(_id string) (*models.Todo, error) {
	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		log.Fatal("Not found")
	}

	filter := bson.M{"_id": id}

	var result models.Todo
	if err = database.Collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &result, nil
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, _ := getTodoById(id)
	context.IndentedJSON(http.StatusOK, todo)
}

func AddTodo(context *gin.Context) {
	var newTodo dtos.CreateTodoRequest

	err := context.ShouldBindJSON(&newTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.Collection.InsertOne(context, newTodo)
	if err != nil || result == nil {
		panic(err)
	}

	context.JSON(http.StatusOK, result)
}

func UpdateTodo(context *gin.Context) {
	var dto models.Todo
	err := context.ShouldBindJSON(&dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		log.Fatal("Not found")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": dto.Name, "completed": dto.Completed}}

	result, err := database.Collection.UpdateOne(context, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count:", result.ModifiedCount)
}

func DeleteTodo(context *gin.Context) {
	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		log.Fatal("Not found")
	}
	filter := bson.M{"_id": id}

	deleteCount, err := database.Collection.DeleteOne(context, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count:", deleteCount.DeletedCount)
}
