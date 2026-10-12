package order

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) Calc(items []Item, clientType string) float64 {
	var sum float64

	for _, item := range items {
		sum += item.Price * float64(item.Quantity)
	}

	if clientType == "VIP" {
		sum *= 0.9
	}

	if clientType == "NEW" {
		sum *= 0.95
	}

	if sum > 1000 {
		sum -= 50
	}

	return sum
}
