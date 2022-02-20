package main

import (
	"database/sql" // Provides by GO
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Thrid party package
)

//provide credentials for database
const (
	host     = "localhost"
	port     = 5432
	user     = "library"
	password = "admin"
	dbname   = "library"
)

type Book struct {
	Book_id          int
	Name             string
	Author           string
	ISBN             string
	Description      string
	Publication_date time.Time
}

//Exercise 1-4 - Create a crud application
func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Establish a connection to the database
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() // Always close the database connection, even if there is an error

	//test connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//CRUD action {Create, Read, Update, Delete}

	//C - Create - Create a new record while sanitizing the input

	insertBook := `
	INSERT INTO books(name, author, isbn, description, publication_date)
	VALUES($1, $2, $3, $4, $5)
	RETURNING book_id, name`

	book_id := 0
	name := ""

	err = db.QueryRow(insertBook, "The Go Programming Language", "Alan A. A. Donovan", "0133502883", "Go is a new language.", "2006-01-01").Scan(&book_id, &name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The recent inserted book: ", name, " has the id of: ", book_id, "\n")

	//U - Update a record

	//update query
	updateBook := `
	UPDATE books
	SET author = $1
	WHERE book_id = $2
	RETURNING book_id, name`

	//update the author of the book with id = 4
	book_id = 0
	err = db.QueryRow(updateBook, "Alan A. Donovan", 4).Scan(&book_id, &name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The book: ", name, " with id of: ", book_id, " has been updated \n")

	//D - Delete a record
	//delete query
	deleteBook := `
	DELETE FROM books
	WHERE book_id = $1
	RETURNING book_id, name`

	//delete the book with id = 4
	book_id = 0
	err = db.QueryRow(deleteBook, 4).Scan(&book_id, &name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The book: ", name, " with id of: ", book_id, " has been deleted\n")

	//R - Read - Read a record

	//read query
	readBooks := `
	SELECT *
	FROM books`

	//read all the books
	rows, err := db.Query(readBooks)
	if err != nil {
		log.Fatal(err)
	}

	// Release resources
	defer rows.Close()

	// Iterate through the result set and store it in books
	var books []Book

	for rows.Next() {
		// create a book for the current book
		var b Book
		err = rows.Scan(&b.Book_id, &b.Name, &b.Author, &b.ISBN, &b.Description, &b.Publication_date)
		if err != nil {
			log.Fatal(err)
		}

		//append the book to the books slice
		books = append(books, b)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Print the books
	for _, b := range books {
		fmt.Println("The book with id of: ", b.Book_id, " is: ", b.Name, " by ", b.Author, " published on ", b.Publication_date, " isbn: ", b.ISBN, " and has the following description: ", b.Description, "\n")
	}

}
