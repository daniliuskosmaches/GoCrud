package Service

import (
	"PizzaApi/Entity"
	"PizzaApi/Repository"
	"log"
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
func (s *Service) IngridientService() {

}

func (s *Service) ReviewService(review *Entity.Review) {
	if review == nil {
		log.Println("review request is null")
	}
	s.repository.ReviewPostRepostory(review)
	s.repository.ReviewFindRepository(review)

}
func (s *Service) RestaurantService(entity *Entity.Restaurant) {

}
func (s *Service) ChefService(entity *Entity.Chef) {

}
func (s *Service) OrderService(repository *Repository.Repository) {
	if repository == nil {
		return
	}

}
