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
func (s Controller) GetIngredients(w http.ResponseWriter) {
	ingridient, err := utils.ParseJson[Entity.Ingredient](s.request)

	if err != nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
	}

	err = s.Service.IngredientService(&ingridient)
	if err != nil {
		return
	}

}

func (s Controller) GetRestaraunt(w http.ResponseWriter) error {
	rest, err := utils.ParseJson[Entity.Restaurant](s.request)
	if err != nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
	}
	s.Service.RestaurantService(&rest)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rest)
	if err != nil {
		return err
	}
	return nil

}

func (s Controller) GetChefs(r *http.Request, w http.ResponseWriter) {

}

func (s Controller) GetReviews(r *http.Request, w http.ResponseWriter) {
	getReviews, err := utils.ParseJson[Entity.Review](s.request)
	if err != nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
	}

	s.Service.ReviewService(&getReviews)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(getReviews)
	if err != nil {
		return
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
	Restaurant, err := utils.ParseJson[Entity.Restaurant](s.request)
	var name string
	var menu string
	if err != nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
	}
	name = Restaurant.Name
	menu = Restaurant.Menu
	err = s.Service.MenuService(name, menu)

	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Restaurant not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{
		"restaurant": name,
		"menu":       menu,
	})
	if err != nil {
		return
	}

}
