package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// セッションが存在するかどうかでエラーを返すかを分岐する関数
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|save|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		id, _ := strconv.Atoi(q[2])
		fmt.Println(id)
		fn(w, r, id)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top) // top
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
