package generator

import (
	"fmt"

	"github.com/mrwan200/promptparse-go/lib"
	"github.com/mrwan200/promptparse-go/utils"
)

func Truemoney(mobileNo string, amount float64, message string) string {
	tag29 := []lib.TLVTag{
		lib.Tag("00", "A000000677010111"),
		lib.Tag("03", "14000"+mobileNo),
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

	if message != "" {
		// Append data
		result := append(payload, lib.Tag("81", utils.EncodeTag81(message)))
		payload = result
	}

	tag, err := lib.WithCRCTag(lib.Encode(payload), "63")
	if err != nil {
		return ""
	}

	return tag
}
