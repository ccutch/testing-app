package main

import "net/http"
import "html/template"

var baseTemplate *template.Template

func main() {
	baseTemplate, _ = template.New("index.html").ParseFiles("templates/header.html")
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexPage)
	http.ListenAndServe(":5000", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, map[string]string{"name": "Connor"})
	if err != nil {
		panic(err)
	}
}
