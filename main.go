package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Post struct {
	Id int
	Title string
	Body string
}

var db, err = sql.Open("mysql", "root:12345678@/go_course?charset=utf8")

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/", HomeHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.ExecuteTemplate(writer, "index.html", ListPosts()); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func ListPosts() []Post {
	rows, error := db.Query("select * from posts")
	checkError(error)

	var items []Post

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}
	return items
}