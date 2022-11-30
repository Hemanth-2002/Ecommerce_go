package model

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func CreateDB() *gorm.DB {

	db_name := flag.String("db", "product", "give db name")

	db, err := gorm.Open("postgres", fmt.Sprintf("user=postgres password=MohanNeelima@01 dbname=%v sslmode=disable", *db_name))
	checkError(err)

	// db.DropTableIfExists(&Product{})
	// db.DropTableIfExists(&Rating{})
	// db.DropTableIfExists(&Variant{})
	// db.CreateTable(&Product{})
	// db.CreateTable(&Rating{})
	// db.CreateTable(&Variant{})

	// db.Save(&Product{
	// 	Name:     "Water",
	// 	Price:    350,
	// 	Quantity: 10,
	// 	Variant: []Variant{
	// 		{Color: "red", Image: "red_bottle.jpg"},
	// 		{Color: "green", Image: "green_bottle.jpg"},
	// 	},
	// 	Rating: []Rating{
	// 		{Name: "KHK", Review: "Nice", Rating: 4},
	// 		{Name: "GST", Review: "Great", Rating: 5},
	// 		{Name: "AB", Review: "Avg", Rating: 3},
	// 	},
	// })

	// db.Save(&Product{
	// 	Name:     "Water-Bottle",
	// 	Price:    350,
	// 	Quantity: 10,
	// 	Variant: []Variant{
	// 		{Color: "red", Image: "red_bottle.jpg"},
	// 		{Color: "green", Image: "green_bottle.jpg"},
	// 	},
	// 	Rating: []Rating{
	// 		{Name: "KHK", Review: "Nice", Rating: 4},
	// 		{Name: "GST", Review: "Great", Rating: 5},
	// 		{Name: "AB", Review: "Avg", Rating: 3},
	// 	},
	// })
	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	Image       string `sql:"type:VARCHAR(255)" json:"image"`
	Variant     []Variant
	Rating      []Rating
}

type Rating struct {
	RatingID  uint   `gorm:"AUTO_INCREMENT" json:"id"`
	ProductID uint   `json:"productid"`
	Name      string `json:"name"`
	Review    string `json:"review"`
	Rating    uint   `gorm:"check:rating<6" json:"rating"`
}

type Variant struct {
	ProductID uint   `json:"productid"`
	Color     string `json:"color"`
	Image     string `json:"image"`
}
