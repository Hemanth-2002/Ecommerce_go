package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "ecommerce/database"
	model "ecommerce/db_model"
)

type Server struct {
	Db database.DbProducts
}
type JsonResponse struct {
	Type string
	Data []model.Product
}

func (s Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	got, err := s.Db.GetProds()
	if err != nil {
		fmt.Printf("Thers an error")
	}
	resp := JsonResponse{Type: "success", Data: got}
	json.NewEncoder(w).Encode(resp)
}

// func (s Server) GetProductById(w http.ResponseWriter, r *http.Request) {
// 	var products []model.Product
// 	params := mux.Vars(r)["id"]
// 	s.Db.Find(&products).First(&products, params)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(products)
// }

// func (s Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product model.Product
// 	json.NewDecoder(r.Body).Decode(&product)
// 	s.Db.Save(&product)
// 	json.NewEncoder(w).Encode(&product)
// }
