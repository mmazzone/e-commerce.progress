package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// global pointer to database
var db *sql.DB

//
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// product struct (model)
type Product struct {
	ProductId    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductDesc  string `json:"product_desc"`
	ProductPrice string `json:"product_price"`
	InfoLink     string `json:"info_link"`
	ImgSrc       string `json:"img_src"`
	Categories   string `json:"categories"`
}

// GET ALL PRODUCTS

func getProducts(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	// INIT PRODUCTS VAR AS A SLICE PRODUCT STRUCT
	products := []Product{}

	result, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	// creating a new variable and setting the value to existing struct
	for result.Next() {
		// creating a new struct for each row
		var product Product
		// checking for discrepencies
		err := result.Scan(&product.ProductId, &product.ProductName, &product.ProductDesc,
			&product.ProductPrice, &product.InfoLink, &product.ImgSrc, &product.Categories)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err.Error())
		}
		// append the return values to original products variable
		products = append(products, product)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

func main() {
	// local db connection
	//                  driver
	database, err := sql.Open("mysql", "root:rootpassword!@tcp(database:3306)/lunar")
	if err != nil {
		panic(err.Error())
	}

	db = database

	defer db.Close()

	// Init router
	fmt.Println("Init Router")
	r := mux.NewRouter()

	// Route Handlers / end points
	r.HandleFunc("/products", getProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
