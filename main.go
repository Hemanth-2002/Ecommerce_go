package main

import (
	product "ecommerce/controllers/products"
	"ecommerce/controllers/reviews"
	"flag"
	"fmt"
	"net/http"

	authorize "ecommerce/authorization"
	token "ecommerce/generate_token"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {

	token.GenerateJWt()

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
	router.Handle("/api/products", authorize.IsAuthorized(prod.GetProducts)).Methods("GET")
	router.Handle("/api/products/{id}", authorize.IsAuthorized(prod.GetProductById)).Methods("GET")
	router.Handle("/api/products/create", authorize.IsAuthorized(prod.CreateProduct)).Methods("POST")
	router.Handle("/api/products/{id}/reviews", authorize.IsAuthorized(rev.GetReviews)).Methods("GET")
	router.Handle("/api/products/{id}/reviews/create", authorize.IsAuthorized(rev.CreateReview)).Methods("POST")
	router.Handle("/api/products/{product_id}/reviews/{rating_id}/delete", authorize.IsAuthorized(rev.DeleteReview)).Methods("DELETE")
	router.Handle("/api/products/{product_id}/reviews/{rating_id}/update", authorize.IsAuthorized(rev.UpdateReview)).Methods("PUT")
	fmt.Println("server at 8080")
	http.ListenAndServe(":8080", router) // port opened at 8080
}
