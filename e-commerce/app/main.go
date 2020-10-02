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
	InfoLink     string `json:"info_link"`
	ImgSrc       string `json:"img_src"`
	Categories   string `json:"categories"`
}

// CONTACT INFO struct
type ContactInfo struct {
	Firstname    string `json:"first_name"`
	Lastname     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
}

// GET ALL PRODUCTS

func getProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	fmt.Println("product Endpoint hit")
	fmt.Println()

	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(products)

	// INIT PRODUCTS VAR AS A SLICE PRODUCT STRUCT
	products := []Product{}

	result, err := db.Query("SELECT * FROM products")
	if err != nil {
		fmt.Println(err)
		return
	}

	// defer result.Close()

	// creating a new variable and setting the value to existing struct
	for result.Next() {
		// creating a new struct for each row
		var product Product
		// checking for discrepencies
		err := result.Scan(&product.ProductId, &product.Discontinued, &product.ProductName, &product.ProductDesc,
			&product.ProductPrice, &product.InfoLink, &product.ImgSrc, &product.Categories)
		if err != nil {
			fmt.Println(err)
			return
		}
		// append the return values to original products variable
		products = append(products, product)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

//GET SINGLE PRODUCT
// func getProduct(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println()
// 	fmt.Println("single product")
// 	fmt.Println()

// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // GET PARAMS
// 	// LOOP THROUGH PRODUCTS AND FIND WITH ID
// 	for _, Product := range products {
// 		if Product.ProductId == params["id"] {
// 			json.NewEncoder(w).Encode(Product)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(products)

// }

//CREATE A NEW CONTACT
// func createContact(w http.ResponseWriter, r *http.Request) {
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
	db, err := sql.Open("mysql", "root:rootpassword!@tcp(database:3306)/lunar")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("connected to database")
	fmt.Println()

	// db = database

	defer db.Close()

	// Init router
	fmt.Println("Init Router")
	r := mux.NewRouter()

	// Route Handlers / end points
	fmt.Println("Prior to products function.")
	r.HandleFunc("/products", getProducts).Methods("GET")
	fmt.Println("Post products function.")
	// r.HandleFunc("/products", getProduct).Methods("GET")
	// fmt.Println("get one product.")
	// r.HandleFunc("/contact", createContact).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
