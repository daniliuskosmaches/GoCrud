package Restaurant

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) MenuRepository(name string) {
	var menu string

	if name == "" {
		log.Println("name is nil")
	}

	query := "SELECT menu FROM restaurants WHERE LOWER(name) = LOWER(?) AND deleted_at IS NULL LIMIT 1"

	err := r.db.Raw(query, name).Scan(&menu).Error
	if err != nil {
		log.Println("didnt found, or nil")
	}
}
