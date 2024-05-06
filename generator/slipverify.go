package generator

import "github.com/mrwan200/promptparse-go/lib"

// Generate Slip Verify QR Code
//
// This also called "Mini-QR" that embedded in slip used for verify transactions
func SlipVerify(sendingBank string, transRef string) string {
	payload := []lib.TLVTag{
		lib.Tag("00", lib.Encode([]lib.TLVTag{
			lib.Tag("00", "000001"),
			lib.Tag("01", sendingBank),
			lib.Tag("02", transRef),
		})),
		lib.Tag("51", "TH"),
	}

	tag, err := lib.WithCRCTag(lib.Encode(payload), "91")
	if err != nil {
		return ""
	}

	return tag
}
