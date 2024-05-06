package generator

import (
	"fmt"

	"github.com/mrwan200/promptparse-go/lib"
)

// Generate PromptPay Bill Payment (Tag 30) QR Code
func BillPayment(billerID string, amount float64, ref1 string, ref2 string, ref3 string) string {
	tag30 := []lib.TLVTag{
		lib.Tag("00", "A000000677010112"),
		lib.Tag("01", billerID),
		lib.Tag("02", ref1),
	}

	if ref2 != "" {
		result := append(tag30, lib.Tag("03", ref2))
		tag30 = result
	}

	payload := []lib.TLVTag{
		lib.Tag("00", "01"),
		lib.Tag("01", "11"),
		lib.Tag("30", lib.Encode(tag30)),
		lib.Tag("53", "764"),
		lib.Tag("58", "TH"),
	}

	if amount != 0 {
		payload[1] = lib.Tag("01", "12")
		// Append data
		result := append(payload, lib.Tag("54", fmt.Sprintf("%.2f", float64(amount*100)/100)))
		payload = result
	}

	if ref3 != "" {
		// Append data
		result := append(payload, lib.Tag("62", lib.Encode([]lib.TLVTag{
			lib.Tag("07", ref3),
		})))
		payload = result
	}

	tag, err := lib.WithCRCTag(lib.Encode(payload), "63")
	if err != nil {
		return ""
	}

	return tag
}
