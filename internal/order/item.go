package order

// Item описывает товар в корзине.
type Item struct {
	Name     string
	Price    float64
	Quantity int
}

// NewItem создаёт новый товар.
func NewItem(name string, price float64, quantity int) Item {
	return Item{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

// Total возвращает стоимость позиции: цена, умноженная на количество.
func (i Item) Total() float64 {
	return i.Price * float64(i.Quantity)
}
