package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id int
	Title string
	Body string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		post := Post{Id: 1, Title: "Unamed Post", Body: "No content"}

		if title := request.FormValue("title"); title != "" {
			post.Title = title
		} //127.0.0.1:8080/?title=My new Post

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(writer, "index.html", post); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
	//localhost:8080
	//127.0.0.1:8080
}
