package Ingridient

import (
	"PizzaApi/pkg"
	"net/http"
)

type Controller struct {
	service Service

	request []byte
}

// Конструктор
func NewController(service Service) *Controller {
	return &Controller{service: service}
}

/*
конкретно в этом тестовом задание мы будем просто забирать данные в джсон формате
потому что юзеру не нужно брать и создавать новые рестораны юзеру только нужно просмотреть
список всех ресторонов и пицц с ингридиентами
*/

func (s Controller) GetIngredients(w http.ResponseWriter) {
	ingridient, err := pkg.ParseJson[Ingredient](s.request)

	if err != nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
	}

	err = s.service.IngredientService(ingridient)
	if err != nil {
		return
	}

}
