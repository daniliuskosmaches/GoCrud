package Repository

import (
	"PizzaApi/Entity"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) PizzaRepository(pizzas *[]Entity.Pizza) {
	result := r.db.Find(pizzas)
	if result != nil {
		log.Fatal("result is null")
	}
	log.Println(pizzas)

}

func (r *Repository) IngredientRepository(db *gorm.DB) {
	result := db.Find(&[]Entity.Ingredient{})
	if result != nil {
		log.Fatal("result is null")
	}
	log.Println(result)

}

func (r *Repository) ReviewFindRepository(reviews *Entity.Review) {
	result := r.db.Find(reviews)

	if result != nil {
		log.Fatal("result is null")
	}
	log.Println(result)
}
func (r *Repository) ReviewPostRepostory(review *Entity.Review) {
	r.db.Create(review)

	fmt.Println("Review created")

}
