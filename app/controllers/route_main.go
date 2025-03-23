package controllers

import (
	"net/http"
)

// topのハンドラー
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "top")
}
