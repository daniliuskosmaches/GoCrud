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

func (s *Service) ReviewService(review *Entity.Review) {
	if review == nil {
		log.Println("review request is null")
	}
	s.repository.ReviewPostRepostory(review)
	s.repository.ReviewFindRepository(review)

}
func (s *Service) MenuService(name func[T any](data []byte) (Entity.Restaurant, error)) (string error, err error) {
	var cleaname = strings.TrimSpace(name)
	s.repository.MenuRepository(cleaname)

	if cleaname == "" {
		return nil, err
	}

	return nil, err
}
func (s *Service) IngredientService(ingridient func[T any](data []byte) ([]Entity.Pizza, error)) {
	if ingridient == nil {
		log.Println("ingridient is nil")
	}
	s.repository.IngredientRepository(ingridient)

}
func (s *Service) ChefService(chef *Entity.Chef) {
	if chef == nil {
		log.Println("chef is nil")
	}
	s.repository.ChefRepository(chef)

}
