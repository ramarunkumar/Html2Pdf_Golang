package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ID    int64
	err   error
	db    *sql.DB
	templ *template.Template
)

type employee struct {
	name         string
	id           int
	salary       int
	destignation string
}

func main() {
	db, err = sql.Open("sqlite3", "./employee.db")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(db)
	}

	templ, err = templ.ParseGlob("templates/index.html")
	checkErr(err)

	http.HandleFunc("/", index)
	// http.HandleFunc("/create", createUser)
	http.HandleFunc("/read", readUser)
	http.Handle("/assets/", http.FileServer(http.Dir("."))) //serve other files in assets dir
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("server running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	templ.ExecuteTemplate(res, "index.html", nil)
}

func checkErr(err error) {
	fmt.Println(err)
}

func readUser(res http.ResponseWriter, req *http.Request) {
	// query
	rows, err := db.Query("SELECT * FROM employee")
	checkErr(err)

	var id int
	var name string
	var salary int
	var destignation string
	var e []employee

	for rows.Next() {
		err = rows.Scan(&name, &id, &salary, &destignation)
		checkErr(err)
		e = append(e, employee{name: name, id: id, salary: salary, destignation: destignation})
		return
	}

	templ.ExecuteTemplate(res, "index.html", e)
}
