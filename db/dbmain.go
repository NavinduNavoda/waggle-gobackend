package dbmain

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	// Ensure the connection is valid
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the SQLite database")

	// Create a table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		age INTEGER
	)`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'users' created successfully")

	// Insert some data into the table
	insertStmt := "INSERT INTO users(name, age) VALUES(?, ?)"
	_, err = db.Exec(insertStmt, "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(insertStmt, "Bob", 25)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully")

	// Query the data from the table
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Retrieving data from the table:")
	for rows.Next() {
		var id, age int
		var name string
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	
}