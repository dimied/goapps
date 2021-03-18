package main;

import (
	"fmt"
	"os"
	"database/sql"

	// Postgres Driver 
	// go get github.com/lib/pq
	_ "github.com/lib/pq"

)

type PgDb struct {
	DB *sql.DB
}

type InfoRow struct {
	id int
	name string
}

func newDatabase(dbname string, user string, password string) (*PgDb, error) {
	connectionString := "user="+user+" dbname="+dbname+" sslmode=disable"
	connectionString = connectionString+" password="+password
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	return &PgDb {
		DB: db,
	}, nil
}

func InsertIntoTable(db *PgDb, name string) (int, error) {
	var id int

	err := db.DB.QueryRow("INSERT INTO info (name) VALUES($1) RETURNING id", name).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}

func FindAllValues(db *PgDb) ([]InfoRow, error) {
	rows, err := db.DB.Query("SELECT * from info")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]InfoRow, 0, 4)

	for rows.Next() {
		var resultRow InfoRow

		err = rows.Scan(&resultRow.id, &resultRow.name)

		if err != nil {
			fmt.Println("Failed to read a row from result")
			return nil, err
		}

		result = append(result, resultRow)
	}

	
	return result, nil
}

func main() {
	// Install postgres
	// create user with root privileges and password
	// CREATE table info (
	//   id serial PRIMARY KEY,
	//   name varchar(64)	
	// );
	db, err := newDatabase("admin", "admin", "admin")

	if err != nil {
		fmt.Println("Failed to connect", err)
		os.Exit(1)
	}

	 names := []string{"Adam", "Eve", "Snake"}

	 for _, val := range names {
		id, insertError := InsertIntoTable(db, val)

		if insertError != nil {
			fmt.Println("Failed to insert an entry for name = ", val)
			fmt.Println("Error: ", insertError)
			continue
		} 
		fmt.Printf("Entry created. ID = %d, Name = %s\n", id, val)
	 }

	 entries, findError := FindAllValues(db)

	 if findError != nil {
		fmt.Println("Failed to find entries")
	 } else {
		 for _, val := range entries {
			 fmt.Println("Entry: ", val)
		 }
	 }
	
	
}