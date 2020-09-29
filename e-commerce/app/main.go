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

var db *sql.DB
var port = "8000"

// product struct (model)
type Product struct {
	ProductId    string `json:"product_id"`
	Discontinued string `json:"discontinued"`
	ProductName  string `json:"product_name"`
	ProductDesc  string `json:"product_desc"`
	ProductPrice string `json:"product_price"`
	ImgLink      string `json:"img_link"`
	Categories   string `json:"categories"`
}

// CONTACT INFO struct
type ContactInfo struct {
	Firstname    string `json:"first_name"`
	Lastname     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
}

// INIT PRODUCTS VAR AS A SLICE PRODUCT STRUCT
var products []Product

// GET ALL PRODUCTS
func getProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	fmt.Println("product Endpoint hit")
	fmt.Println()

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(products)

	result, err := db.Query("SELECT * FROM products")
	if err != nil {
		fmt.Println(err)
	}

	defer result.Close()

	for result.Next() {
		var product Product
		err := result.Scan(&product.ProductId, &product.Discontinued, &product.ProductName, &product.ProductDesc, &product.ProductPrice, &product.ImgLink, &product.Categories)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)

}

// // GET SINGLE PRODUCT
// funct getProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // GET PARAMS
// 	// LOOP THROUGH PRODUCTS AND FIND WITH ID
// 	for _, item := range products {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Product{})
// }

// CREATE A NEW CONTACT
// funct createContact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(10000000)) // MOCK ID
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }

func main() {
	// local db connection
	//                  driver
	database, err := sql.Open("mysql", "root:rootpassword!@tcp(database:3306)/lunar")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println()
	fmt.Println("connected to database")
	fmt.Println()

	db = database

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// MOCK DATA - @TODO - IMPLEMENT DB

	// Route Handlers / end points
	r.HandleFunc("/products", getProducts).Methods("GET")
	//r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	// r.HandleFunc("/webclient/contact", createContact).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
