package main

import (
	"fmt"

	"todo_app/app/models"
)

func main() {
	u, _ := models.GetUser(1)
	fmt.Println(u)
}
