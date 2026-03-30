package Ingridient

type Ingredient struct {
	Name string `gorm:"type:varchar(100);unique;not null" json:"name"`
}
