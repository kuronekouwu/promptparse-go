# Promptparse Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/mrwan200/promptparse-go)](https://goreportcard.com/report/github.com/mrwan200/promptparse-go)
[![Documentation](https://godoc.org/github.com/mrwan200/promptparse-go?status.svg)](http://godoc.org/github.com/mrwan200/promptparse-go)
[![GitHub release](https://img.shields.io/github/release/mrwan200/promptparse-go.svg)](https://github.com/mrwan200/promptparse-go/releases)

"All-in-one Golang for PromptPay & EMVCo QR Codes"

That No dependency required. So here we go!

## Original library from (Forked from)

[maythiwat/promptparse](https://github.com/maythiwat/promptparse)

## Insatallation
```sh
go get github.com/mrwan200/promptparse-go
```

## Features

- **Parse** &mdash; PromptPay & EMVCo QR Code data strings into object
- **Generate** &mdash; QR Code data from pre-made templates (for example: PromptPay AnyID, PromptPay Bill Payment, TrueMoney, etc.)
- **Manipulate** &mdash; any values from parsed QR Code data (for example: transfer amount, account number) and encodes back into QR Code data
- **Validate** &mdash; checksum and data structure for known QR Code formats (for example: Slip Verify API Mini QR)

## Usage

### Parsing data and get value from tag

```golang
package main

import (
    "log"

    promptparse "github.com/mrwan200/promptparse-go"
)

func main() {
    parsed := promptparse.Parse("000201010212293...", true, true)
    if parsed == nil {
        log.Println("Invalid payload.")
        return
    }

    // Now get tag transfer ref
    log.Println(parsed.GetTagValue("00", "")) // Result: YYYY/MM/DD HH:mm:ss 01
}
```

### Build QR data and append CRC tag

```golang
package main

import (
	"log"

	"github.com/mrwan200/promptparse-go/lib"
)

func main() {
	// Example data
	data := []lib.TLVTag{
		lib.Tag("00", "01"),
		lib.Tag("01", "11"),
	}

	// Set CRC Tag ID '63'
	tag, err := lib.WithCRCTag(lib.Encode(data), "63")
	if err != nil {
		log.Println("Error when create CRC tag")
        return
	}
	log.Println(tag) // Result: YYYY/MM/DD HH:mm:ss 00020101...
}
```

### Generate PromptPay Bill Payment QR

```golang
package main

import (
	"log"

	"github.com/mrwan200/promptparse-go/generator"
)

func main() {
	log.Println(generator.BillPayment("1xxxxxxxxxxxx", 300, "INV12345", "", ""))  // Result: YYYY/MM/DD HH:mm:ss 00020101021230490016...
}
```

### Validate & extract data from Slip Verify QR

```golang
package main

import (
	"log"

	promptparse "github.com/mrwan200/promptparse-go"
)

func main() {
	slip := promptparse.SlipVerify("0041000600.....")
	if slip == nil {
		log.Println("Invalid slip verify")
		return
	}

	log.Println(slip.SendingBank) // Result: YYYY/MM/DD HH:mm:ss 004
	log.Println(slip.TransRef) // Result: YYYY/MM/DD HH:mm:ss 0141012...
}
```

### Convert BOT Barcode to PromptPay QR Tag 30 (Bill Payment)

```golang
package main

import (
	"log"

	promptparse "github.com/mrwan200/promptparse-go"
)

func main() {
	bill := promptparse.ParseBarcode("|310109999999901\r...")
	if bill.BillerID == "" {
		log.Println("Invalid barcode.")
		return
	}

	log.Println(bill.ToQRTag30()) // Reseult: YYYY/MM/DD HH:mm:ss  00020101021....
}
```

## References
- [EMV QR Code](https://www.emvco.com/emv-technologies/qrcodes/)
- [Thai QR Payment Standard](https://www.bot.or.th/content/dam/bot/fipcs/documents/FPG/2562/ThaiPDF/25620084.pdf)
- [Slip Verify API Mini QR Data](https://developer.scb/assets/documents/documentation/qr-payment/extracting-data-from-mini-qr.pdf)
- [BOT Barcode Standard](https://www.bot.or.th/content/dam/bot/documents/th/our-roles/payment-systems/about-payment-systems/Std_Barcode.pdf)

## See also
- [phoomin2012/promptparse-php](https://github.com/phoomin2012/promptparse-php) PromptParse port for PHP


## License
This project is MIT licensed (see [LICENSE](LICENSE))