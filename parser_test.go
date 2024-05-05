package promptparsego_test

import (
	"testing"

	promptparse "github.com/mrwan200/promptparse-go"
)

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
