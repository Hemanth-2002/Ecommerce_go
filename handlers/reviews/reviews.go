package reviews

import (
	model "ecommerce/db_model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	db := model.CreateDB()
	var ratings []model.Rating
	var reviews []string
	params := mux.Vars(r)["id"]
	db.Where("product_id = ?", params).Find(&ratings)
	for _, rev := range ratings {
		reviews = append(reviews, rev.Review)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratings)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	db := model.CreateDB()
	w.Header().Set("Content-Type", "application/json")
	var rating model.Rating
	json.NewDecoder(r.Body).Decode(&rating)
	db.Save(&rating)
	json.NewEncoder(w).Encode(&rating)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	db := model.CreateDB()
	w.Header().Set("Content-Type", "application/json")
	var ratings []model.Rating
	productid := mux.Vars(r)["product_id"]
	ratingid := mux.Vars(r)["rating_id"]
	db.Where("rating_id = ? AND product_id = ?", ratingid, productid).Delete(&ratings)
	json.NewEncoder(w).Encode(&ratings)
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	db := model.CreateDB()
	w.Header().Set("Content-Type", "application/json")
	var rating model.Rating
	json.NewDecoder(r.Body).Decode(&rating)
	productid := mux.Vars(r)["product_id"]
	ratingid := mux.Vars(r)["rating_id"]
	ratings := []model.Rating{}
	db.Where("rating_id = ? and product_id = ?", ratingid, productid).Find(&ratings).Update(&rating)
}
