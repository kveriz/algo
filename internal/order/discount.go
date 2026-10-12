package order

// Правила скидок
const (
	vipDiscountRate       = 0.10
	newDiscountRate       = 0.05
	bulkDiscountRate      = 0.01
	thresholdAmount       = 1000.0
	thresholdDiscount     = 50.0
	bulkQuantityThreshold = 10
)

// countUnits возвращает суммарное количество единиц товара в заказе.
func countUnits(items []Item) int {
	total := 0
	for _, item := range items {
		total += item.Quantity
	}
	return total
}

// applyBulkDiscount применяет дополнительную скидку 1 %,
// если заказ содержит больше bulkQuantityThreshold единиц товара.
func applyBulkDiscount(amount float64, items []Item) float64 {
	if countUnits(items) > bulkQuantityThreshold {
		return amount * (1 - bulkDiscountRate)
	}
	return amount
}

// applyClientDiscount применяет процентную скидку в зависимости от типа клиента.
func applyClientDiscount(amount float64, clientType ClientType) float64 {
	switch clientType {
	case ClientTypeVIP:
		return amount * (1 - vipDiscountRate)
	case ClientTypeNew:
		return amount * (1 - newDiscountRate)
	default:
		return amount
	}
}

// applyThresholdDiscount вычитает фиксированную скидку,
// если сумма превышает пороговое значение.
func applyThresholdDiscount(amount float64) float64 {
	if amount > thresholdAmount {
		return amount - thresholdDiscount
	}
	return amount
}
