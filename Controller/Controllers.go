package Controller

import (
	"PizzaApi/Service"
)

type Controller struct {
	Service *Service.Service
}

// Конструктор
func NewController(service *Service.Service) *Controller {
	return &Controller{Service: service}
}

/*
конкретно в этом тестовом задание мы будем просто забирать данные в джсон формате
потому что юзеру не нужно брать и создавать новые рестораны юзеру только нужно просмотреть
список всех ресторонов и пицц с ингридиентами
*/

func (s Controller) GetPizzas() {
	s.Service.PizzaService()

}
func (s Controller) GetIngredients() {
	s.Service.IngriService()
}
func (s Controller) GetRestaurants() {
	s.Service.RestaurantService()
}
func (s Controller) GetChefs() {
	s.Service.ChefService()
}
func (s Controller) GetOrders() {
	s.Service.OrderService()
}
func (s Controller) GetReviews() {
	s.Service.ReviewService()
}

// для того чтоб оставить коментарий
func (s Controller) PostReviews() {
	s.Service.ReviewService()
}
