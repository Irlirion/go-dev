package order_total

// CalculateOrder возвращает стоимость заказа до скидки, размер скидки и итог
// (всё в копейках). Скидка задаётся целым числом процентов от 0 до 100.
// При некорректном входе функция возвращает три нуля.
// Расчёт скидки выполняется целочисленно с округлением вниз до копейки.
// Можно считать, что суммы и произведения не переполняют int64.
func CalculateOrder(prices []int64, quantities []int, discount int) (subtotal, discountAmount, total int64) {
	if (len(prices) != len(quantities)) || (discount < 0) || (discount > 100) {
		return 0, 0, 0
	}
	for i := 0; i < len(prices); i++ {
		if (prices[i] < 0) || (quantities[i] < 0) {
			return 0, 0, 0
		}
	}
	for i, price := range prices {
		quantity := quantities[i]
		subtotal += price * int64(quantity)
	}
	discountAmount = subtotal * int64(discount) / 100
	total = subtotal - discountAmount

	return subtotal, discountAmount, total
}
