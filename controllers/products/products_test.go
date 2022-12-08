package products

import (
	model "ecommerce/db_model"
	mocks "ecommerce/mocks"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func (*Server) CheckErr(err error) {
	if err != nil {
		fmt.Printf("Thers an error")
	}
}
func TestGetProducts(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbProducts(mockcntrl)
	testProd := Server{Db: mockProd}

	prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	mockProducts := []model.Product{prod1}

	mockProd.EXPECT().GetProds().Return(mockProducts, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testProd.GetProducts)
	handler.ServeHTTP(rr, req)

	// Checking status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checking body
	var got JsonResponse
	json.NewDecoder(rr.Body).Decode(&got)
	testProd.CheckErr(err)
	prod2 := model.Product{Name: "Asus", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	mockProducts2 := []model.Product{prod2}
	var mock = JsonResponse{Type: "success", Data: mockProducts2}

	if !reflect.DeepEqual(got, mock) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, mock)
	}
}
