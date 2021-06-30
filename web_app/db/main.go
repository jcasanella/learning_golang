package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

)

func main() {
	// http://www.devgi.com/2018/11/install-mysql-docker-windows.html
	// How to connect using docker: https://towardsdatascience.com/connect-to-mysql-running-in-docker-container-from-a-local-machine-6d996c574e55
	db, err := sql.Open("mysql", "root:my-secret-pw@(localhost:3306)/MYSQLTEST?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	query := `
	CREATE TABLE users (
	    id INT AUTO_INCREMENT,
	    username TEXT NOT NULL,
	    password TEXT NOT NULL,
	    created_at DATETIME,
	    PRIMARY KEY (id)
	);`
	_, err = db.Exec(query)
	if err != nil {
		panic(err.Error())
	}

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	userID, err := result.LastInsertId()
	fmt.Println(userID)

	defer db.Close()
}
