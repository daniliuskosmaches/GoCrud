package Pizza

import "github.com/jmoiron/sqlx"

type Repository struct {
	db sqlx.DB
}

func NewRepository(db sqlx.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) GetPizza(pizza Pizza) {

}
