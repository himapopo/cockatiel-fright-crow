package main

import (
	"embed"
	"html/template"
	"net/http"
)

// viewsディレクトリ下のファイルを全て変数に格納する
//
//go:embed views/*
var views embed.FS

// メイン画面の表示
var templates = template.Must(template.ParseFS(views, "views/main.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "main.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)

	// viewsに格納したファイルを全て公開
	http.Handle("/views/", http.FileServer(http.FS(views)))

	// httpサーバーを立ち上げ
	http.ListenAndServe(":8080", nil)
}
