package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type BOTBarcode struct {
	BillerID string
	Ref1     string
	Ref2     string
	Amount   float64
}

func (bc *BOTBarcode) FromString(payload string) {
	if !strings.HasPrefix(payload, "|") {
		return
	}

	data := strings.Split(payload[1:], "\r")
	if len(data) != 4 {
		return
	}
	data = data[:4]

	bc.BillerID = data[0]
	bc.Ref1 = data[1]
	if len(data[2]) > 0 {
		bc.Ref2 = data[2]
	}
	if data[3] != "0" {
		amt, _ := strconv.Atoi(data[3])
		acualAmt := float64(amt) / 100
		bc.Amount = float64(acualAmt*100) / 100
	} else {
		bc.Amount = 0
	}
}

func (bc *BOTBarcode) ToString() string {
	if bc.Amount != 0 {
		acualAmt := float64(int(bc.Amount*100)) / 100
		bc.Amount = acualAmt * 100
	}

	result := fmt.Sprintf("|%s\r%s\r%s\r", bc.BillerID, bc.Ref1, bc.Ref2)
	if bc.Amount != 0 {
		result += fmt.Sprintf("%d", int64(bc.Amount))
	} else {
		result += "0"
	}

	return result
}

// Converts BOT Barcode to PromptPay QR Tag 30 (Bill Payment)
//
// This method works for some biller, depends on destination bank
func (bc *BOTBarcode) ToQRTag30() string {
	tag30 := []TLVTag{
		Tag("00", "A000000677010112"),
		Tag("01", bc.BillerID),
		Tag("02", bc.Ref1),
	}

	if bc.Ref2 != "" {
		result := append(tag30, Tag("03", bc.Ref2))
		tag30 = result
	}

	payload := []TLVTag{
		Tag("00", "01"),
		Tag("01", "11"),
		Tag("30", Encode(tag30)),
		Tag("53", "764"),
		Tag("58", "TH"),
	}

	if bc.Amount != 0 {
		payload[1] = Tag("01", "12")
		// Append data
		result := append(payload, Tag("54", fmt.Sprintf("%.2f", float64(bc.Amount*100)/100)))
		payload = result
	}

	tag, err := WithCRCTag(Encode(payload), "63")
	if err != nil {
		return ""
	}

	return tag
}
