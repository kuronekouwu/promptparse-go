package promptparsego_test

import (
	"testing"

	"github.com/mrwan200/promptparse-go/generator"
)

func TestGenerateAnyID(t *testing.T) {
	result := generator.AnyID(generator.MSISDN, "0812223333", 0)
	if result != "00020101021129370016A0000006770101110113006681222333353037645802TH63041DCF" {
		t.Fatalf("Incorrect test (GenerateAnyID)")
	}
}

func TestGenerateAnyIDIncludeAmount(t *testing.T) {
	result := generator.AnyID(generator.MSISDN, "0812223333", 30)
	if result != "00020101021229370016A0000006770101110113006681222333353037645802TH540530.0063043CAD" {
		t.Fatalf("Incorrect test (TestGenerateAnyIDIncludeAmount)")
	}
}

func TestGenerateSlipVerify(t *testing.T) {
	result := generator.SlipVerify("002", "0002123123121200011")
	if result != "004000060000010103002021900021231231212000115102TH91049C30" {
		t.Fatalf("Incorrect test (TestGenerateSlipVerify)")
	}
}

func TestGenerateTruemoneyQR(t *testing.T) {
	result := generator.Truemoney("0801111111", 0, "")
	if result != "00020101021129390016A000000677010111031514000080111111153037645802TH63047C0F" {
		t.Fatalf("Incorrect test (TestGenerateTruemoneyQR)")
	}
}

func TestGenerateTruemoneyQRWithAmountAndMessage(t *testing.T) {
	result := generator.Truemoney("0801111111", 10.05, "Hello World!")
	if result != "00020101021229390016A000000677010111031514000080111111153037645802TH540510.05814800480065006C006C006F00200057006F0072006C006400216304F5A2" {
		t.Fatalf("Incorrect test (TestGenerateTruemoneyQRWithAmountAndMessage)")
	}
}

func TestGenerateBillPaymentWithRef3(t *testing.T) {
	result := generator.BillPayment("0112233445566", 0, "CUSTOMER001", "INV001", "SCB")
	if result != "00020101021130620016A000000677010112011301122334455660211CUSTOMER0010306INV00153037645802TH62070703SCB6304780E" {
		t.Fatalf("Incorrect test (TestGenerateBillPaymentWithRef3)")
	}
}

func TestGenerateBotBarcode(t *testing.T) {
	result := generator.BOTBarcode("099999999999990", "111222333444", "", 0)
	if result != "|099999999999990\r111222333444\r\r0" {
		t.Fatalf("Incorrect test (TestGenerateBotBarcode)")
	}
}

func TestGenerateBOTBarcodeWithRef2AndAmount(t *testing.T) {
	result := generator.BOTBarcode("099400016550100", "123456789012", "670429", 3649.22)
	if result != "|099400016550100\r123456789012\r670429\r364922" {
		t.Fatalf("Incorrect test (TestGenerateBOTBarcodeWithRef2AndAmount)")
	}
}
