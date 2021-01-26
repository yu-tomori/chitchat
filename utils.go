package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/yugaraxy/chitchat/data"
)

func session(w http.ResponseWriter, r *http.Request) (session data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess := data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func generateHTML(writer http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
