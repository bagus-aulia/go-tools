package number

import (
	"github.com/leekchan/accounting"
)

// GenerateRupiah function to generate rupiah currency
func GenerateRupiah(value int, decimalPrecision int) string {
	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: decimalPrecision,
		Thousand:  ".",
		Decimal:   ",",
	}

	currency := ac.FormatMoney(value)
	if decimalPrecision == 0 {
		currency += ",-"
	}

	return currency
}
