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

func (r *Repository) IngredientRepository(ingridient *Entity.Ingredient) {
	result := r.db.Find(ingridient)
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

func (r *Repository) RestaurantRepository(rest *Entity.Restaurant) {
	result := r.db.Create(rest)
	if result != nil {
		log.Fatal("result is null")
	}
	log.Println(result)

}
func (r *Repository) ChefRepository(chef *Entity.Chef) {
	result := r.db.Find(chef)
	if result != nil {
		log.Fatal("result is null")
	}
	log.Println(result)

}

func (r *Repository) MenuRepository(name string) {
	var menu string

	if name == "" {
		log.Println("name is nil")
	}

	query := "SELECT menu FROM restaurants WHERE LOWER(name) = LOWER(?) AND deleted_at IS NULL LIMIT 1"

	err := r.db.Raw(query, name).Scan(&menu).Error
	if err != nil {
		log.Println("didnt found, or nil")
	}
}
