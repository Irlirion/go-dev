package order_status

// CountStatuses возвращает количество new, paid, shipped, cancelled,
// unknown и active (new + paid + shipped) в указанном порядке.
// Любой статус, не равный одному из четырёх известных, считается unknown.
func CountStatuses(statuses []string) (newCount, paidCount, shippedCount, cancelledCount, unknownCount, activeCount int) {
	if len(statuses) == 0 {
		return 0, 0, 0, 0, 0, 0
	}
	for _, status := range statuses {
		switch status {
			case "new":
				newCount++
				activeCount++
			case "paid":
				paidCount++
				activeCount++
			case "shipped":
				shippedCount++
				activeCount++
			case "cancelled":
				cancelledCount++
			default:
				unknownCount++
		}
	}
	return newCount, paidCount, shippedCount, cancelledCount, unknownCount, activeCount
}
