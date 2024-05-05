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

	return fmt.Sprintf("|%s\r%s\r%s\r%s", bc.BillerID, bc.Ref1, bc.Ref2, fmt.Sprintf("%.2f", bc.Amount))
}

func (bc *BOTBarcode) ToQRTag30() string {
	tag30 := []TLVTag{
		Tag("00", "A000000677010112"),
		Tag("01", bc.BillerID),
		Tag("02", bc.Ref1),
		Tag("03", bc.Ref2),
	}

	payload := []TLVTag{
		Tag("00", "01"),
		Tag("01", "12"),
		Tag("54", fmt.Sprintf("%.2f", float64(bc.Amount*100)/100)),
		Tag("30", Encode(tag30)),
		Tag("53", "764"),
		Tag("58", "TH"),
	}

	tag, err := WithCRCTag(Encode(payload), "63")
	if err != nil {
		return ""
	}

	return tag
}
