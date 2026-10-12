package order

// OrderService содержит логику расчёта заказов.
type OrderService struct{}

// NewOrderService создаёт сервис обработки заказов.
func NewOrderService() *OrderService {
	return &OrderService{}
}

// CalculateTotal рассчитывает итоговую стоимость всех товаров в корзине
// с учётом скидки, зависящей от типа клиента.
//
// Правила расчёта:
//   - базовая стоимость — сумма произведений цены товара на его количество;
//   - для клиента типа ClientTypeVIP применяется скидка 10 %;
//   - для клиента типа ClientTypeNew применяется скидка 5 %;
//   - если итоговая сумма после применения скидки превышает 1000,
//     из неё вычитается фиксированная скидка 50.
//
// Порядок применения скидок: сначала процентная скидка по типу клиента,
// затем — пороговая скидка при превышении суммы в 1000.
//
// @param items      список товаров в корзине (может быть пустым);
// @param clientType тип клиента, влияющий на применяемую скидку;
// @return итоговая стоимость заказа после применения всех скидок.
func (s *OrderService) CalculateTotal(items []Item, clientType ClientType) float64 {
	subtotal := calculateSubtotal(items)
	subtotal = applyClientDiscount(subtotal, clientType)
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
