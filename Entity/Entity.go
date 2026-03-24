package Entity

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Address string `gorm:"type:varchar(255);not null" json:"address"`

	Chef Chef `json:"chef"`

	Pizzas []Pizza `gorm:"foreignKey:RestaurantID" json:"pizzas,omitempty"`

	Reviews []Review `gorm:"foreignKey:RestaurantID" json:"reviews,omitempty"`
}

type Chef struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	RestaurantID uint   `json:"restaurant_id"`
}

type Pizza struct {
	gorm.Model
	Name             string `gorm:"type:varchar(255);not null" json:"name"`
	CheeseType       string `gorm:"type:varchar(100)" json:"cheese_type"`
	DoughThickness   string `gorm:"type:varchar(50)" json:"dough_thickness"`
	SecretIngredient string `gorm:"type:varchar(255)" json:"secret_ingredient"`

	RestaurantID uint `json:"restaurant_id"`

	Ingredients []Ingredient `gorm:"many2many:pizza_ingredients;" json:"ingredients"`
}

type Ingredient struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique;not null" json:"name"`
}

type Review struct {
	gorm.Model
	RestaurantID uint   `json:"restaurant_id"`
	Rating       int    `gorm:"check:rating >= 1 AND rating <= 5" json:"rating"`
	Text         string `gorm:"type:text" json:"text"`

	Restaurant Restaurant `gorm:"foreignKey:RestaurantID" json:"restaurant_info,omitempty"`
}
