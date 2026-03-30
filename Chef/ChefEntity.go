package Chef

import "gorm.io/gorm"

type Chef struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	RestaurantID uint   `json:"restaurant_id"`
}
