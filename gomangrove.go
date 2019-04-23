package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var connection string

func init() {

	dbname := os.Getenv("DBNAME")
	if len(dbname) < 1 {
		dbname = "mangrove"
	}

	pass := os.Getenv("DBPASS")
	if len(pass) < 1 {
		pass = "swid"
	}

	user := os.Getenv("DBUSER")
	if len(user) < 1 {
		user = "root"
	}

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "3306"
	}

	host := os.Getenv("HOST")
	if len(host) < 1 {
		host = "127.0.0.1"
	}

	connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port)
}

func main() {

	//init()
	fmt.Println("MySQL Details", connection)

	// Open up our database connection.
	db, err := sql.Open("mysql", connection)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
		fmt.Println("error with db")
	}

	var name string
	stmt, err := db.Prepare(dropDb)
	stmt.Exec()

	stmt, err = db.Prepare(createDb)
	stmt.Exec()

	stmt, err = db.Prepare(createTable)
	stmt.Exec()

	stmt, err = db.Prepare(insertPerson)
	stmt.Exec()

	err = db.QueryRow("select name from mangrove.people where id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

// Simple embedded SQL migrations
var createDb = `
CREATE database mangrove;
`
var dropDb = `
DROP database mangrove;
`
var dropTable = `
DROP table mangrove.people;
`
var createTable = `
CREATE TABLE IF NOT EXISTS mangrove.people (
     id      INTEGER PRIMARY KEY AUTO_INCREMENT
    ,name     VARCHAR(32)
    ,city        VARCHAR(32)
);
`
var insertPerson = `
INSERT INTO mangrove.people (name,city) VALUES ('ssx','tlv');
`
