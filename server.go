package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed views/*
var views embed.FS

// メイン画面の表示
var templates = template.Must(template.ParseFS(views, "views/main.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "main.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"ok"}`))
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/healthcheck", healthcheck)

	http.Handle("/views/", http.FileServer(http.FS(views)))

	http.ListenAndServe(":8080", nil)
}
