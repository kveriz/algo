package order

// ClientType описывает тип клиента, влияющий на применяемую скидку.
type ClientType string

const (
	// ClientTypeRegular — обычный клиент без процентной скидки.
	ClientTypeRegular ClientType = "REGULAR"

	// ClientTypeVIP — VIP-клиент со скидкой 10 %.
	ClientTypeVIP ClientType = "VIP"

	// ClientTypeNew — новый клиент со скидкой 5 %.
	ClientTypeNew ClientType = "NEW"
)
