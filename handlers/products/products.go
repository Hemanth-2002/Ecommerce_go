package product

import (
	model "ecommerce/db_model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := model.CreateDB()
	var products []model.Product
	db.Find(&products)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(products)
	defer db.Close()
}

func GetProductsById(w http.ResponseWriter, r *http.Request) {
	db := model.CreateDB()
	var products []model.Product
	params := mux.Vars(r)["id"]
	db.Find(&products).First(&products, params)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db := model.CreateDB()
	w.Header().Set("Content-Type", "application/json")
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)
	db.Save(&product)
	json.NewEncoder(w).Encode(&product)
}
