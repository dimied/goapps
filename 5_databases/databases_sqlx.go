package main

import (
	//"errors"
	//"database/sql"
	//"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	// Postgres Driver
	// go get github.com/lib/pq
	_ "github.com/lib/pq"

	db "./nullable"
)

type PgDb struct {
	DB *sqlx.DB
}

type InfoRow struct {
	Id int
	// For value which can be NULL, you need to use sql.NullString as type
	MyName db.OurNewNullString `db:"name"` // using annotation for mapping
}

func newDatabase(dbname string, user string, password string) (*PgDb, error) {
	connectionString := "user=" + user + " dbname=" + dbname + " sslmode=disable"
	connectionString = connectionString + " password=" + password
	db, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	return &PgDb{
		DB: db,
	}, nil
}

func DeleteByName(db *PgDb, name string) (bool, error) {
	res, err := db.DB.Exec("DELETE FROM info WHERE name = $1", name)
	if err != nil {
		return false, err
	}
	affected, _ := res.RowsAffected()

	return affected > 0, nil
}

func InsertIntoTable(db *PgDb, name string) (int, error) {
	var id int

	err := db.DB.QueryRow("INSERT INTO info (name) VALUES($1) RETURNING id", name).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}

func InsertNullNameIntoTable(db *PgDb) (int, error) {
	var id int

	err := db.DB.QueryRow("INSERT INTO info (name) VALUES($1) RETURNING id", nil).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}

func InsertIntoTableWithTransaction(db *PgDb, names ...string) (int, error) {
	var count int

	tx := db.DB.MustBegin()
	success := true
	var errInsert error

	for _, name := range names {
		_, err := tx.Exec("INSERT INTO info (name) VALUES($1)", name)
		if err != nil {
			success = false
			errInsert = err
			// Not really required, it will rollback on any error
			tx.Rollback()
			break
		}
		count++
	}

	if success {
		err := tx.Commit()
		if err != nil {
			return 0, err
		}
		return count, nil
	} else {
		return 0, errInsert
	}

}

func FindAllValues(db *PgDb) ([]InfoRow, error) {

	result := make([]InfoRow, 0, 4)
	err := db.DB.Select(&result, "SELECT * from info")

	if err != nil {
		return nil, err
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

	count, errTx := InsertIntoTableWithTransaction(db, "Anna", "Hanna")
	if errTx != nil {
		fmt.Println("Failed to insert in transaction", errTx)
	} else {
		fmt.Println("Inserted using transaction #count = ", count)
	}

	_, insertErrorNull := InsertNullNameIntoTable(db)

	if insertErrorNull != nil {
		fmt.Println("Failed to insert null name", insertErrorNull)
	}

	entries, findError := FindAllValues(db)

	if findError != nil {
		fmt.Println("Failed to find entries")
	} else {
		fmt.Println("Find ...")
		for _, val := range entries {
			asJson, _ := json.Marshal(val)
			fmt.Println("Entry: ", string(asJson))
		}
	}
}
