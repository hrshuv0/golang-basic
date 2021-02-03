package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func users(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/demoapp?charset=utf8")

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, fname, lname, email from employees")

	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		Id    int    `json:"id"`
		Fname string `json:"firstname"`
		Lname string `json:"lastname"`
		Email string `json:"email"`
	}

	var users []User

	for rows.Next() {
		var id int
		var fname string
		var lname string
		var email string

		rows.Scan(&id, &fname, &lname, &email)
		users = append(users, User{id, fname, lname, email})
	}

	usersBytes, _ := json.Marshal(&users)

	w.Write(usersBytes)
	db.Close()
}

func main() {
	fmt.Println("started")

	http.HandleFunc("/users/", users)
	http.ListenAndServe(":8080", nil)

	fmt.Println("terminated")
}

