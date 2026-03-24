package Entity

type Pizza struct {
	Id          int
	Name        string
	Price       float64
	Ingredients []string
}
type Ingredient struct {
	Id   int
	Name string
}
type Review struct {
	Id        int
	PizzaId   int
	ChefId    int
	Text      string
	Rating    float64
	CreatedAt string
}
type Chef struct {
	Id   int
	Name string
	Bio  string
}

// Геттеры
func (p *Pizza) GetIngredients() []string {
	return p.Ingredients
}
func (p *Pizza) GetPrice() float64 {
	return p.Price
}
func (p *Pizza) GetName() string {
	return p.Name
}
func (p *Pizza) GetId() int {
	return p.Id
}

// Сеттеры
func (p *Pizza) SetIngredients(ingredients []string) {
	p.Ingredients = ingredients
}
func (p *Pizza) SetPrice(price float64) {
	p.Price = price

}
func (p *Pizza) SetName(name string) {
	p.Name = name
}
func (p *Pizza) SetId(id int) {
	p.Id = id
}
