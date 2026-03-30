package Ingridient

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) IngredientRepository(ingridient Ingredient) {
	result := r.db.Find(ingridient)
	if result != nil {
		log.Fatal("result is null")
	}
	log.Println(result)
}
