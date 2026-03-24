package Service

import (
	"PizzaApi/Repository"
)

type Service struct {
	repository *Repository.Repository
}

func NewService(repository *Repository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) PizzaService() {

}
func (s *Service) IngriService() {

}

func (s *Service) ReviewService() {

}
func (s *Service) RestaurantService() {

}
func (s *Service) ChefService() {

}
func (s *Service) OrderService() {

}
