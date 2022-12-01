package main

import (
	product "ecommerce/controllers/products"
	"ecommerce/controllers/reviews"
	"flag"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {

	db_name := flag.String("db", "product", "give db name")

	db, err := gorm.Open("postgres", fmt.Sprintf("user=postgres password=MohanNeelima@01 dbname=%v sslmode=disable", *db_name))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var prod product.Server = product.Server{db}
	var rev reviews.Server = reviews.Server{db}

	//  Routing

	router := mux.NewRouter()
	router.HandleFunc("/api/products", prod.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", prod.GetProductsById).Methods("GET")
	router.HandleFunc("/api/products/create", prod.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}/reviews", rev.GetReviews).Methods("GET")
	router.HandleFunc("/api/products/{id}/reviews/create", rev.CreateReview).Methods("POST")
	router.HandleFunc("/api/products/{product_id}/reviews/{rating_id}/delete", rev.DeleteReview).Methods("DELETE")
	router.HandleFunc("/api/products/{product_id}/reviews/{rating_id}/update", rev.UpdateReview).Methods("PUT")
	fmt.Println("server at 8080")

}
