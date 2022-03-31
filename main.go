package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"strconv"
)

type todo struct {
	ID        int `json:"id"`
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

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id int) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
func getTodo(context *gin.Context) {
	id := context.Param("id")
	i, err := strconv.Atoi(id)

	todo, err := getTodoById(i)
	if err != nil {
	   context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	   return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	err := context.ShouldBindJSON(&newTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)
	context.JSON(http.StatusOK, newTodo)
}

func updateTodo(context *gin.Context) {
	id := context.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oneTodo, err := getTodoById(i)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var updatedTodo todo
	err = context.ShouldBindJSON(&updatedTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oneTodo.Item = updatedTodo.Item
	oneTodo.Completed = updatedTodo.Completed
	context.JSON(http.StatusOK, oneTodo)
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := getTodoById(i)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos[:i], todos[i+1:]...)
	context.JSON(http.StatusOK, todo)
}

func main() {
	fmt.Println("Program started")

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run("localhost:9090")
}
