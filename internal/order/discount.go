package order

// Правила скидок вынесены в именованные константы,
// чтобы избавиться от «магических» чисел в коде.
const (
	vipDiscountRate   = 0.10
	newDiscountRate   = 0.05
	thresholdAmount   = 1000.0
	thresholdDiscount = 50.0
)

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
