package main

import (
	// "fmt"

	"todo_app/app/models"
)

func main() {
	// fmt.Println(models.Db)
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "testtest"
	// u.CreateUser()

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}
