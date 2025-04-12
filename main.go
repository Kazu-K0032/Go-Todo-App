package main

import (
	"fmt"

	"todo_app/app/controllers"
	"todo_app/config"
)

func main() {
	fmt.Println("サーバーを起動します。ポート: "+config.Config.Port)
	controllers.StartMainServer()
}
