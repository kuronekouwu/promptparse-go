package generator

import "github.com/mrwan200/promptparse/lib"

func BOTBarcode(billerId string, ref1 string, ref2 string, amount float64) string {
	barocde := lib.BOTBarcode{
		BillerID: billerId,
		Ref1:     ref1,
		Ref2:     ref2,
		Amount:   amount,
	}

	return barocde.ToString()
}
