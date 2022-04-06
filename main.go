package main

import (
	"fmt"
	"project/config/database"
	"project/routers"
)

func main() {
	fmt.Println("Program started")

	database.ConnectDB()

	routers.MyRouters()
}
