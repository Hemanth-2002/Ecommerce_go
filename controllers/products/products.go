package product

import (
	model "ecommerce/db_model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Db *gorm.DB
}

func (s Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	s.Db.Find(&products)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(products)
	// defer s.Db.Close()
}

func (s Server) GetProductById(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	params := mux.Vars(r)["id"]
	s.Db.Find(&products).First(&products, params)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (s Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)
	s.Db.Save(&product)
	json.NewEncoder(w).Encode(&product)
}
