package main

import (
	"database/sql"
	x "employee/pdf"
	"fmt"

	_ "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type employee struct {
	Name         string
	Id           int
	Salary       string
	Destignation string
}

type data struct {
	Emp []employee
}

func main() {

	r := x.NewRequestPdf("")

	db, err := sql.Open("sqlite3", "./employee.db")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM employee;")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SQL Data", rows)
	}

	var emp = []employee{}
	for rows.Next() {
		e := employee{}
		err := rows.Scan(&e.Name, &e.Id, &e.Salary, &e.Destignation)
		if err != nil {
			fmt.Println(err)
			continue
		}
		emp = append(emp, e)
	}

	var x = data{}
	x.Emp = emp
	fmt.Println("emp struct", emp)
	if err := r.ParseTemplate("templates/index.html", x); err == nil {
		pdf, _ := r.GeneratePDF("./examples.pdf")
		fmt.Println(pdf, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
