package controllers

import (
	"net/http"

	"todo_app/config"
)

func StartMainSercer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
