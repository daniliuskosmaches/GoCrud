package Controller

import (
	"PizzaApi/Entity"
	"PizzaApi/Service"
	"PizzaApi/utils"
	"io"
	"net/http"
)

type Controller struct {
	Service *Service.Service

	request []byte
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

func (s Controller) GetPizzas(r *http.Request, w http.ResponseWriter) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
		return
	}
	pizzas, _ := utils.ParseJson[[]Entity.Pizza](body)
	s.Service.PizzaService(&pizzas)

}
func (s Controller) GetIngredients(r *http.Request, w http.ResponseWriter) {

	s.Service.IngridientService()
}
func (s Controller) GetRestaurants(r *http.Request, w http.ResponseWriter) {

	s.Service.RestaurantService()
}
func (s Controller) GetChefs(r *http.Request, w http.ResponseWriter) {
	s.Service.ChefService()
}
func (s Controller) GetOrders(r *http.Request, w http.ResponseWriter) {
	s.Service.OrderService()
}
func (s Controller) GetReviews(r *http.Request, w http.ResponseWriter) {
	s.Service.ReviewService()
}

// но здесь мы позволим юзеру оставить отзыв на сайте тобишь в базе данных
func (s Controller) PostReviews(r *http.Request, w http.ResponseWriter) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
		return
	}
	review, _ := utils.ParseJson[Entity.Review](body)
	s.Service.ReviewService(&review)
}
