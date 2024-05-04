package generator

import (
	"fmt"
	"strings"

	"github.com/mrwan200/promptparse/lib"
)

const (
	MSIDN     = "01"
	NATID     = "02"
	EWALLETID = "03"
	BANKACC   = "04"
)

func SlipVerify(sendingBank string, transRef string) string {
	payload := []lib.TLVTag{
		lib.Tag("00", *lib.Encode([]lib.TLVTag{
			{
				ID:    "00",
				Value: "000001",
			},
			{
				ID:    "01",
				Value: sendingBank,
			},
			{
				ID:    "02",
				Value: transRef,
			},
		})),
		lib.Tag("51", "TH"),
	}

	tag, err := lib.WithCRCTag(*lib.Encode(payload), "91")
	if err != nil {
		return ""
	}

	return *tag
}

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
		lib.Tag("29", *lib.Encode(tag29)),
		lib.Tag("53", "764"),
		lib.Tag("58", "TH"),
	}

	if amount != 0 {
		payload[1] = lib.Tag("01", "12")
		// Append data
		result := append(payload, lib.Tag("54", fmt.Sprintf("%.2f", float64(amount*100)/100)))
		payload = result
	}

	tag, err := lib.WithCRCTag(*lib.Encode(payload), "63")
	if err != nil {
		return ""
	}

	return *tag
}

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
		lib.Tag("30", *lib.Encode(tag30)),
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
		result := append(payload, lib.Tag("62", ref3))
		payload = result
	}

	tag, err := lib.WithCRCTag(*lib.Encode(payload), "63")
	if err != nil {
		return ""
	}

	return *tag
}
