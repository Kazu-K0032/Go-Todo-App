package main

import (
	// "fmt"

	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainSercer()
}
