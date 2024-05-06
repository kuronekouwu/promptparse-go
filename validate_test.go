package promptparsego_test

import (
	"testing"

	promptparse "github.com/mrwan200/promptparse-go"
)

func TestValidateSlipVerifyValid(t *testing.T) {
	parsed := promptparse.SlipVerify("004100060000010103014022000111222233344ABCD125102TH910417DF")
	if parsed == nil {
		t.Fatalf("Invalid slip verify. (TestValidateChecksumTag)")
	}

	if parsed.SendingBank != "014" {
		t.Fatalf("Incorrect sending bank. (TestValidateSlipVerifyValid)")
	}

	if parsed.TransRef != "00111222233344ABCD12" {
		t.Fatalf("Incorrect transfer ref. (TestValidateSlipVerifyValid)")
	}
}

func TestValidateSlipVerifyInvalid(t *testing.T) {
	parsed := promptparse.SlipVerify("00020101021229370016A0000006770101110113006680111111153037645802TH540520.15630442BE")
	if parsed != nil {
		t.Fatalf("Incorrect test. (TestValidateSlipVerifyInvalid)")
	}
}
