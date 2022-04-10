package main

import (
	"fmt"
	"project/config/database"
	"project/routers"
)

func main() {
	fmt.Println("Program started")

	database.ConnectDB()

	fmt.Println("Program started")

	routers.MyRouters()
}
