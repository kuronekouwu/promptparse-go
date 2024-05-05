package generator

import (
	"fmt"
	"strings"

	"github.com/mrwan200/promptparse-go/lib"
)

const (
	/** Mobile number */
	MSIDN = "01"
	/** National ID or Tax ID */
	NATID = "02"
	/** E-Wallet ID */
	EWALLETID = "03"
	/** Bank Account (Reserved) */
	BANKACC = "04"
)

// Generate PromptPay AnyID (Tag 29) QR Code
func AnyID(types string, target string, amount float64) string {
	if types == MSIDN {
		msidn := target
		if strings.HasPrefix(target, "0") {
			msidn = "66" + msidn[1:]
		}
		target = fmt.Sprintf("%013s", msidn)
	}

	tag29 := []lib.TLVTag{
		lib.Tag("00", "A000000677010111"),
		lib.Tag(types, target),
	}

	payload := []lib.TLVTag{
		lib.Tag("00", "01"),
		lib.Tag("01", "11"),
		lib.Tag("29", lib.Encode(tag29)),
		lib.Tag("53", "764"),
		lib.Tag("58", "TH"),
	}

	if amount != 0 {
		payload[1] = lib.Tag("01", "12")
		// Append data
		result := append(payload, lib.Tag("54", fmt.Sprintf("%.2f", float64(amount*100)/100)))
		payload = result
	}

	tag, err := lib.WithCRCTag(lib.Encode(payload), "63")
	if err != nil {
		return ""
	}

	return tag
}
