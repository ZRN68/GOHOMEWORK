package main

import "fmt"

type CustomerStatus int // статус и тип

const (
	Regular CustomerStatus = iota
	Silver
	Gold
)

const (
	deliveryPrice    = 299.0  // цена доставки
	freeDeliveryFrom = 3000.0 // порог для бесплатной доставки
)

func main() {
	var sum float64 = 3000 // cумма заказа
	var textstatus string
	var status CustomerStatus = Gold // статус покупателя

	// рассчет скидки по статусу
	var discountPercent float64
	switch status {
	case Regular:
		textstatus = "Regular"
	case Silver:
		textstatus = "Silver"
		discountPercent = 5
	case Gold:
		textstatus = "Gold"
		discountPercent = 10
	default:
		textstatus = "неизвестный статус"
	}
	// скидка в рублях
	discount := sum * discountPercent / 100

	// сумма после скидки
	afterDiscount := sum - discount

	// доставка
	var delivery float64
	if afterDiscount >= freeDeliveryFrom {
		delivery = 0
	} else {
		delivery = deliveryPrice
	}

	// итого
	total := afterDiscount + delivery

	// вывод
	fmt.Printf("Сумма: %.2f, статус: %s\n", sum, textstatus)
	fmt.Printf("Скидка: %.2f\n", discount)
	fmt.Printf("Доставка: %.2f\n", delivery)
	fmt.Printf("Итого: %.2f\n", total)
}
