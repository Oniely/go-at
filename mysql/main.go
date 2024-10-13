package main

import {
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()


}
