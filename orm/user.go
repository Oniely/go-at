package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connection to database")
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users Endpoint Hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to db")
	}

	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to db")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	data := db.Create(&User{
		Name:  name,
		Email: email,
	})

	json.NewEncoder(w).Encode(data)

	fmt.Fprintf(w, "New User Created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to db")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Deleted User: %v", name)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to db")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email

	db.Save(&user)

	fmt.Fprintf(w, "Updated User: %v, New Email: %v", name, email)
}
