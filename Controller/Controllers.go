package Controller

import (
	"PizzaApi/Entity"
	"PizzaApi/Service"
	"PizzaApi/utils"
	"encoding/json"
	"io"
	"log"
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
	ingridients := utils.ParseJson[[]Entity.Pizza]
	if ingridients == nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
		return
	}

	s.Service.IngredientService(ingridients)

}

func (s Controller) GetChefs(r *http.Request, w http.ResponseWriter) {

}

func (s Controller) GetReviews(r *http.Request, w http.ResponseWriter) {
	getReviews := utils.ParseJson[Entity.Review]
	if getReviews = nil {

		log.Println("getReviews is nil")
	}

}

func (s Controller) GetRestaurants(r *http.Request, w http.ResponseWriter) {

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
	return
}
func (s Controller) GetRestrauntsMenu(r *http.Request, w http.ResponseWriter) {
	Restaurant := utils.ParseJson[Entity.Restaurant]
	if Restaurant == nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
	}
	menu, err := s.Service.MenuService(Restaurant)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Restaurant not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{
		"restaurant": name,
		"menu":       menu,
	})
	if err != nil {
		return
	}

}
