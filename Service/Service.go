package Service

import (
	"PizzaApi/Entity"
	"PizzaApi/Repository"
	"log"
	"strings"
)

type Service struct {
	repository *Repository.Repository
}

func NewService(repository *Repository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) PizzaService(pizzas *[]Entity.Pizza) {
	if pizzas == nil {
		log.Println("pizzas is nil")
	}
	s.repository.PizzaRepository(pizzas)

}
func (s *Service) RestaurantService(rest *Entity.Restaurant) {
	if rest == nil {
		log.Println("rest is nil")
	}
	s.repository.RestaurantRepository(rest)
}

func (s *Service) ReviewService(review *Entity.Review) {
	if review == nil {
		log.Println("review request is null")
	}
	s.repository.ReviewPostRepostory(review)
	s.repository.ReviewFindRepository(review)

}
func (s *Service) MenuService(menu string, name string) (string error) {
	if menu == "" {
		log.Println("menu is nil")
	}

	cleaname := strings.TrimSpace(name)
	s.repository.MenuRepository(cleaname)

	return
}
func (s *Service) IngredientService(ingridient *Entity.Ingredient) error {
	if ingridient == nil {
		log.Println("ingridient is nil")
	}
	s.repository.IngredientRepository(ingridient)

	return nil
}
func (s *Service) ChefService(chef *Entity.Chef) {
	if chef == nil {
		log.Println("chef is nil")
	}
	s.repository.ChefRepository(chef)

}
