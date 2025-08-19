package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

type Book struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var ctx = context.Background()

func main() {
	db, err := sql.Open("mysql", "root:yangi_parol@tcp(127.0.0.1:3306)/Book")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 20,
	})

	
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		bookID := r.URL.Query().Get("id")
		if bookID == "" {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}

		book, err := getBook(ctx, rdb, db, bookID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	})

	// POST → Write-through
	http.HandleFunc("/book/update", func(w http.ResponseWriter, r *http.Request) {
		var book Book
		bookID := r.URL.Query().Get("id")
		if bookID == "" {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		if err := updateBook(ctx, rdb, db, bookID, book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Book updated successfully")
	})

	fmt.Println("Server 8081 portda ishlayapti...")
	http.ListenAndServe(":8081", nil)
}


func getBook(ctx context.Context, rdb *redis.Client, db *sql.DB, id string) (Book, error) {
	var book Book

	val, err := rdb.Get(ctx, "book:"+id).Bytes()
	if err == redis.Nil {
		row := db.QueryRow("SELECT name, description FROM books WHERE id=?", id)
		if err := row.Scan(&book.Name, &book.Description); err != nil {
			return book, err
		}

		// Redis set async + Gob
		go func(b Book) {
			var buf bytes.Buffer
			gob.NewEncoder(&buf).Encode(b)
			rdb.Set(ctx, "book:"+id, buf.Bytes(), 2*time.Minute)
		}(book)

	} else if err != nil {
		return book, err
	} else {
		buf := bytes.NewBuffer(val)
		gob.NewDecoder(buf).Decode(&book)
	}

	return book, nil
}


func updateBook(ctx context.Context, rdb *redis.Client, db *sql.DB, id string, book Book) error {

	_, err := db.Exec("UPDATE books SET name=?, description=? WHERE id=?", book.Name, book.Description, id)
	if err != nil {
		return err
	}


	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(book); err != nil {
		return err
	}

	err = rdb.Set(ctx, "book:"+id, buf.Bytes(), 60*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}
