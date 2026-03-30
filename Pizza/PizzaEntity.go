package Pizza

import (
	"github.com/jmoiron/sqlx"
)

type Pizza struct {
	ID           uint     `db:"id" json:"id" `
	Name         string   `db:"name" json:"name" `
	CheeseName   string   `db:"name" json:"cheeseName" `
	RestaurantID uint     `db:"restaurant_id" json:"restaurant_id"`
	Ingredients  []string `db:"ingredients" json:"ingredients"`
	sqlx.DB
}
