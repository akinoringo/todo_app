package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	// filesにfile pathを格納
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// Mustを用いると、template.ParseFilesでエラーが生じた際にpanicのエラーを返す
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func top(w http.ResponseWriter, r *http.Request) {
	_, err := Session(w, r)
	// ログインしていない場合のみtopにアクセス
	if err != nil {
		// layoutも再度読み込ませないといけないっぽい
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	session, err := Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}
