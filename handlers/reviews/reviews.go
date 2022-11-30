package reviews

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

func (s Server) GetReviews(w http.ResponseWriter, r *http.Request) {
	var ratings []model.Rating
	var reviews []string
	params := mux.Vars(r)["id"]
	s.Db.Where("product_id = ?", params).Find(&ratings)
	for _, rev := range ratings {
		reviews = append(reviews, rev.Review)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratings)
}

func (s Server) CreateReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rating model.Rating
	json.NewDecoder(r.Body).Decode(&rating)
	s.Db.Save(&rating)
	json.NewEncoder(w).Encode(&rating)
}

func (s Server) DeleteReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ratings []model.Rating
	productid := mux.Vars(r)["product_id"]
	ratingid := mux.Vars(r)["rating_id"]
	s.Db.Where("rating_id = ? AND product_id = ?", ratingid, productid).Delete(&ratings)
	json.NewEncoder(w).Encode(&ratings)
}

func (s Server) UpdateReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rating model.Rating
	json.NewDecoder(r.Body).Decode(&rating)
	productid := mux.Vars(r)["product_id"]
	ratingid := mux.Vars(r)["rating_id"]
	ratings := []model.Rating{}
	s.Db.Where("rating_id = ? and product_id = ?", ratingid, productid).Find(&ratings).Update(&rating)
}
