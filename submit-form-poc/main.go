package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	// "path/filepath"
	u "net/http/httputil"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	SKU           string `json:"sku"`
	Price         string `json:"price"`
	PriceCurrency string `json:"priceCurrency"`
	Brand         string `json:"brand"`
	Category      string `json:"category"`
	Image         string `json:"image"`
}

var db *sql.DB

func initDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/productdb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database")
}

func main() {
	initDB()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/submit", handleFormSubmit)

	//http.HandleFunc("/submit", dumper)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFormSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var product Product
	decoder := json.NewDecoder(r.Body)
	//fmt.Println(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, "Error decoding JSON"+err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO products (name, description, sku, price, priceCurrency, brand, category, image) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	price, err := strconv.ParseFloat(product.Price, 64)
	if err != nil {
		http.Error(w, "Error convering price into the database"+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("image text len:", len(product.Image))
	_, err = db.Exec(query, product.Name, product.Description, product.SKU, price, product.PriceCurrency, product.Brand, product.Category, product.Image)
	if err != nil {
		http.Error(w, "Error inserting data into the database"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product created successfully")
}

func dumper(w http.ResponseWriter, r *http.Request) {
	b, err := u.DumpRequest(r, true)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		w.Write([]byte("\nRemote Address:[" + r.RemoteAddr + "]\n"))
		now := time.Now()
		tmj, err := now.MarshalJSON()
		if err == nil {
			w.Write(tmj)
		} else {
			w.Write([]byte(err.Error()))
		}
		w.Write([]byte("\n"))
		if r.Body != nil {
			var b []byte = make([]byte, r.ContentLength)
			if _, err := r.Body.Read(b); err == nil {
				w.Write([]byte("\n-----------------BODY--------"))
				w.Write(b)
				w.Write([]byte("\n-----------------BODY--------"))
			} else {
				w.Write([]byte(err.Error()))
			}
		} else {
			w.Write([]byte("\n no Body received"))
		}

	}
	return

}
