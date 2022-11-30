package routers

import (
	product "ecommerce/handlers/products"
	"ecommerce/handlers/reviews"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//  Routing

	router := mux.NewRouter()
	router.HandleFunc("/api/products", product.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", product.GetProductsById).Methods("GET")
	router.HandleFunc("/api/products/{id}/reviews", reviews.GetReviews).Methods("GET")
	router.HandleFunc("/api/products/create", product.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}/reviews/create", reviews.CreateReview).Methods("POST")
	router.HandleFunc("/api/products/{product_id}/reviews/{rating_id}/delete", reviews.DeleteReview).Methods("DELETE")
	router.HandleFunc("/api/products/{product_id}/reviews/{rating_id}/update", reviews.UpdateReview).Methods("PUT")
	fmt.Println("server at 8080")
	http.ListenAndServe(":8080", router) // port opened at 8080

}
