package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "laptop-vincent"
var database = "DW"
var connString = ""

var DB *sql.DB

func GetConnect() {
	var err error

	connString := fmt.Sprintf("server=%s;database=%s;trusted_connection=yes", server, database)

	fmt.Println(connString)

	// 	("server=%s;user id=%s;password=%s;database=%s;",
	// server, user, password, database)

	DB, err = sql.Open("sqlserver", connString)

	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}

	fmt.Println("test ping")
	err = DB.Ping()

	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	fmt.Println("test ping done")

	if err != nil {
		log.Fatal("Error creating connection pool : " + err.Error())
	}

	fmt.Println("test ping")

	ctx := context.Background()

	err = DB.PingContext(ctx)

	fmt.Println("connected")

	if err != nil {
		log.Fatal("Error creating connection pool : " + err.Error())
	}

	rows, err := DB.Query(`SELECT 'vincent' as Nama`)

	if err != nil {

		fmt.Println(err.Error())
		panic(err)
	}

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer DB.Close()

}
