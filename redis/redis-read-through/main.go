package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
    _ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

type Book struct{
	Name string `json:"name"`
	Description string `json:"description"`
}


var (
	ctx = context.Background()
)

func main(){
db, err := sql.Open("mysql", "root:yangi_parol@tcp(127.0.0.1:3306)/Book")
if  err != nil {
	fmt.Println(err)
}

defer db.Close()


    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })


	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		bookId := r.URL.Query().Get("id")
		if bookId == "" {
			   http.Error(w, "id required", http.StatusBadRequest)
            return
		}

		book, err := getBook(ctx, rdb, db, bookId)
		  if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		    w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(book)

	
	})
		   fmt.Println("Server 8081 portda ishlayapti...")
    http.ListenAndServe(":8081", nil)
}

func getBook(ctx context.Context, rdb *redis.Client, db *sql.DB, id string) (Book, error) {
    var book Book

  
    val, err := rdb.Get(ctx, "book:"+id).Result()
    if err == redis.Nil {

        row := db.QueryRow("SELECT name, description FROM books WHERE id=?", id)
        if err := row.Scan(&book.Name, &book.Description); err != nil {
            return book, err
        }

   
        bookJSON, _ := json.Marshal(book)
        go func() {
            _ = rdb.Set(ctx, "book:"+id, bookJSON, 10*time.Second).Err()
        }()
    } else if err != nil {
        return book, err
    } else {
     
        _ = json.Unmarshal([]byte(val), &book)
    }

    return book, nil
}
