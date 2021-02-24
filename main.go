package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	//Insert register
	//stmt, err := db.Prepare("Insert into posts(title, body) values(?,?)")
	//checkError(err)
	//
	//_, err = stmt.Exec("My First Post", "My first content")
	//checkError(err)
	//db.Close()

	rows, error := db.Query("select * from posts")
	checkError(error)

	var items []Post

	for rows.Next() {
		//var id int
		//var title string
		//var body string
		//rows.Scan(&id, &title, &body)
		//fmt.Println(id, title, body)

		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}

	db.Close()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(writer, "index.html", items); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
	//localhost:8080
	//127.0.0.1:8080
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}