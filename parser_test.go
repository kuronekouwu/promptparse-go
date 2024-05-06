package promptparsego_test

import (
	"testing"

	promptparse "github.com/mrwan200/promptparse-go"
)

func TestInvalidStringPassedToParserWithRandomString(t *testing.T) {
	parsed := promptparse.Parse("811f6a6f6f0b546cca5bb5ad9818b6aea755a8cabb31561d8cfc018050b75df9", false, true)
	if parsed != nil {
		t.Fatalf("Incorrect test. (TestInvalidStringPassedToParser)")
	}
}

func TestInvalidStringPassedToParserWithRandomStringAndSpeicalCharacter(t *testing.T) {
	parsed := promptparse.Parse("1+LbROFK1ZIewyH2cKwXyZ8VGzTZjahsu5s6j52oL/k=", false, true)
	if parsed != nil {
		t.Fatalf("Incorrect test. (TestInvalidStringPassedToParser)")
	}
}

func TestInvalidStringPassedToParser(t *testing.T) {
	parsed := promptparse.Parse("AAAA0000", false, true)
	if parsed != nil {
		t.Fatalf("Incorrect test. (TestInvalidStringPassedToParser)")
	}
}

func TestParseTLVAndGetTagCount(t *testing.T) {
	parsed := promptparse.Parse("000411110104222202043333", false, true)
	if parsed == nil {
		t.Fatalf("Invalid payload. (TestParseTLVAndGetTagCount)")
	}

	// Get tag
	tag := parsed.GetTag("01", "")
	if len(tag.Value) < 3 {
		t.Fatalf("Incorrect test. (TestParseTLVAndGetTagCount)")
	}
}

func TestParseTLVAndGetOneTag(t *testing.T) {
	parsed := promptparse.Parse("000411110104222202043333", false, true)
	if parsed == nil {
		t.Fatalf("Invalid payload. (TestParseTLVAndGetOneTag)")
	}

	// Get tag
	tag := parsed.GetTag("01", "")
	if tag.Value != "2222" {
		t.Fatalf("Incorrect test. (TestParseTLVAndGetOneTag)")
	}
}

func TestParserPayloadStrictWithInvalidChecksum(t *testing.T) {
	parsed := promptparse.Parse("00020101021229370016A0000006770101110113006680111111153037645802TH540520.156304FFFF", true, true)
	if parsed != nil {
		t.Fatalf("Incorrect test. (TestParserPayloadStrictWithInvalidChecksum)")
	}
}

func TestParserPayloadStrictWithValidChecksumAndGetTagValue(t *testing.T) {
	parsed := promptparse.Parse("00020101021229370016A0000006770101110113006680111111153037645802TH540520.15630442BE", true, true)
	if parsed == nil {
		t.Fatalf("Invalid payload. (TestParserPayloadStrictWithValidChecksumAndGetTagValue)")
	}

	// Get tag
	value := parsed.GetTagValue("29", "01")
	if value != "0066801111111" {
		t.Fatalf("Incorrect test. (TestParserPayloadStrictWithValidChecksumAndGetTagValue)")
	}
}

func TestValidateChecksumTag(t *testing.T) {
	parsed := promptparse.Parse("00020101021229370016A0000006770101110113006680111111153037645802TH540520.15630442BE", false, true)
	if parsed == nil {
		t.Fatalf("Invalid payload. (TestValidateChecksumTag)")
	}

	// Get tag
	result := parsed.Validate("63")

	if !result {
		t.Fatalf("Incorrect test. (TestValidateChecksumTag)")
	}
}

func TestConvertBOTintoBillPaymentValid(t *testing.T) {
	parsed := promptparse.ParseBarcode("|099999999999990\r111222333444\r\r0")
	result := parsed.ToQRTag30()

	if result != "00020101021130550016A0000006770101120115099999999999990021211122233344453037645802TH63043EE7" {
		t.Fatalf("Incorrect test. (TestConvertBOTintoBillPaymentValid)")
	}
}

func TestConvertBOTintoBillPaymentWithRef2AndAmountValid(t *testing.T) {
	parsed := promptparse.ParseBarcode("|099400016550100\r123456789012\r670429\r364922")
	result := parsed.ToQRTag30()

	if result != "00020101021230650016A00000067701011201150994000165501000212123456789012030667042953037645802TH54073649.2263044534" {
		t.Fatalf("Incorrect test. (TestConvertBOTintoBillPaymentWithRef2AndAmountValid)")
	}
}

func TestConvertBOTintoBillPaymentInvalidWithIncorretPayload(t *testing.T) {
	parsed := promptparse.ParseBarcode("00020101021230650016A00000067701011201150994000165501000212123456789012030667042953037645802TH54073649.2263044534")

	if parsed.BillerID != "" {
		t.Fatalf("Incorrect test. (TestParserPayloadStrictWithInvalidChecksum)")
	}
}

func TestConvertBotBarcodeToBillPaymentInvalidWithDataLoss(t *testing.T) {
	parsed := promptparse.ParseBarcode("|099400016550100\r123456789012\r670429")

	if parsed.BillerID != "" {
		t.Fatalf("Incorrect test. (TestConvertBotBarcodeToBillPaymentInvalidWithDataLoss)")
	}
}
