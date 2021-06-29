package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

func main() {
	db, err := sql.Open("mysql", "root:jenni135@(172.17.0.2:3306)/mysql?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	fmt.Println(err.Error())

	defer db.Close()
}
