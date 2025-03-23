package controllers

import (
	"net/http"
	"text/template"
)

// topのハンドラー
func top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/views/templetes/top.html")
	t.Execute(w, nil)
}
