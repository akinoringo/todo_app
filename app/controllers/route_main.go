package controllers

import (
	"fmt"
	"html/template"
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
	// layoutも再度読み込ませないといけないっぽい
	generateHTML(w, "Hello", "layout", "top")
}
