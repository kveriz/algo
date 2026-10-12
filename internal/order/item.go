package order

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

func NewItem(name string, price float64, quantity int) Item {
	return Item{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}
