package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func users(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/godeliver?charset=utf8")

	//w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, email from userotp")

	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		Id     int    `json:"id"`
		Email  string `json:"email"`
		Mobile string `json:"MobileNumber"`
		Otp    string `json:"otpNumber"`
	}

	var users []User

	for rows.Next() {
		var id int
		var mobile string
		var email string
		var otp string

		rows.Scan(&id, &mobile, &email, &otp)
		users = append(users, User{id, mobile, email, otp})
	}

	usersBytes, _ := json.Marshal(&users)

	w.Write(usersBytes)
	db.Close()
}

func users_xml(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/godeliver?charset=utf8")

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, email from userotp")

	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		Id     int    `json:"id"`
		Email  string `json:"email"`
		Mobile string `json:"MobileNumber"`
		Otp    string `json:"otpNumber"`
	}

	var users []User

	for rows.Next() {
		var id int
		var mobile string
		var email string
		var otp string

		rows.Scan(&id, &mobile, &email, &otp)
		users = append(users, User{id, mobile, email, otp})
	}

	usersBytes, _ := xml.Marshal(&users)

	w.Write(usersBytes)
	db.Close()
}

func main() {
	fmt.Println("godeliver started")

	http.HandleFunc("/go/json/", users)
	http.HandleFunc("/go/xml/", users_xml)
	http.ListenAndServe(":8080", nil)

	fmt.Println("terminated")
}
