package analytics

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

func AnalyzeSales(sales []Sale) Report {
	report := Report{
		RevenueByCategory: make(map[string]int64),
	}

	if len(sales) == 0 {
		return report
	}

	unitsByProduct := make(map[string]int64)
	maxQuantity := int64(0)

	for _, sale := range sales {
		if sale.Cancelled {
			continue
		}
		if sale.Product == "" || sale.Category == "" || sale.Quantity <= 0 || sale.UnitPriceKopecks <= 0 {
			report.InvalidSales++
			continue
		}
		report.UnitsSold += sale.Quantity
		revenue := int64(sale.Quantity) * sale.UnitPriceKopecks
		report.RevenueKopecks += revenue
		report.RevenueByCategory[sale.Category] += revenue

		unitsByProduct[sale.Product] += int64(sale.Quantity)
		quantity := unitsByProduct[sale.Product]

		if quantity > maxQuantity ||
			(quantity == maxQuantity && sale.Product < report.TopProduct) {
			maxQuantity = quantity
			report.TopProduct = sale.Product
		}
	}

	return report
}
