package main

import (
	"fmt"

	"todo_app/app/controllers"
	"todo_app/app/models"
)

func TestConnection() {
}

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
