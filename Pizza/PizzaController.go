package Pizza

import (
	"PizzaApi/pkg"
	"io"
	"net/http"
)

type Controller struct {
	Service *Service.Service

	request []byte
}

func NewController(service *Service.Service) *Controller {
	return &Controller{Service: service}
}

func (s Controller) GetPizzas(r *http.Request, w http.ResponseWriter) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid Json", http.StatusInternalServerError)
		return
	}
	pizzas, _ := pkg.ParseJson[[]Entity.Pizza](body)
	s.Service.PizzaService(&pizzas)

}
