package main

import (
	"fmt"

	"github.com/kveriz/algo/14/internal/order"
)

func main() {
	service := order.NewOrderService()

	fmt.Println("=============== ДЕМО ================")

	// Обычный заказ, без скидка за опт
	cart := []order.Item{
		order.NewItem("Книга «Чистый код»", 1500, 1),
		order.NewItem("Ручка", 100, 3),
		order.NewItem("Блокнот", 250, 2),
	}

	printCart("Корзина №1 (6 единиц товара — оптовая скидка не применяется)", cart)
	for _, ct := range []order.ClientType{
		order.ClientTypeRegular,
		order.ClientTypeVIP,
		order.ClientTypeNew,
	} {
		total := service.CalculateTotal(cart, ct)
		fmt.Printf("  Тип клиента: %-8s → итоговая стоимость: %.2f\n", ct, total)
	}

	// Заказ с одной позицией, но большим количеством
	bulkCart := []order.Item{
		order.NewItem("Ручка", 100, 11),
	}

	printCart("Корзина №2 (11 единиц одного товара — оптовая скидка применяется)", bulkCart)
	fmt.Printf("  Тип клиента: %-8s → итоговая стоимость: %.2f\n",
		order.ClientTypeRegular,
		service.CalculateTotal(bulkCart, order.ClientTypeRegular))

	// Скидка за счёт суммы количеств разных позиций ---
	mixedCart := []order.Item{
		order.NewItem("Ручка", 100, 6),
		order.NewItem("Блокнот", 200, 5),
	}

	printCart("Корзина №3 (11 единиц суммарно по двум товарам — оптовая скидка применяется)", mixedCart)
	fmt.Printf("  Тип клиента: %-8s → итоговая стоимость: %.2f\n",
		order.ClientTypeRegular,
		service.CalculateTotal(mixedCart, order.ClientTypeRegular))
}

func printCart(header string, items []order.Item) {
	fmt.Println()
	fmt.Println(header)
	fmt.Println("  Состав корзины:")

	totalUnits := 0
	for _, item := range items {
		fmt.Printf("    — %-22s %d шт. × %.2f = %.2f\n",
			item.Name, item.Quantity, item.Price, item.Total())
		totalUnits += item.Quantity
	}

	fmt.Printf("  Всего единиц товара: %d\n", totalUnits)
}
