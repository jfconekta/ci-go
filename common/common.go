package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Initialize() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, sslMode,
	)
}

func InsertDB() int {
	connStr := Initialize()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	var sqlStatement = `
	insert into demo_table (value) values (cast(random()*100+1 as integer) )
	`
	db.QueryRow(sqlStatement)

	return 1

}

func QueryDB() int {
	connStr := Initialize()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT value FROM demo_table ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	var value int
	err = rows.Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return value
}
