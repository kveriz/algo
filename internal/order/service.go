package order

// OrderService содержит логику расчёта заказов.
type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

// CalculateTotal рассчитывает итоговую стоимость всех товаров в корзине
// с учётом скидок, зависящих от типа клиента и объёма заказа.
//
// Правила расчёта (применяются последовательно):
//   - базовая стоимость — сумма произведений цены товара на его количество;
//   - для клиента типа ClientTypeVIP применяется скидка 10 %;
//   - для клиента типа ClientTypeNew применяется скидка 5 %;
//   - если суммарное количество единиц товара в заказе превышает 10,
//     применяется дополнительная bulk-скидка 1 %;
//   - если итоговая сумма после применения скидок превышает 1000,
//     из неё вычитается фиксированная скидка 50.
//
// @param items      список товаров в корзине (может быть пустым);
// @param clientType тип клиента, влияющий на применяемую скидку;
// @return итоговая стоимость заказа после применения всех скидок.
func (s *OrderService) CalculateTotal(items []Item, clientType ClientType) float64 {
	subtotal := calculateSubtotal(items)
	subtotal = applyClientDiscount(subtotal, clientType)
	subtotal = applyBulkDiscount(subtotal, items)
	return applyThresholdDiscount(subtotal)
}

// calculateSubtotal суммирует стоимость всех позиций корзины.
func calculateSubtotal(items []Item) float64 {
	var subtotal float64
	for _, item := range items {
		subtotal += item.Total()
	}
	return subtotal
}
