package generator

import "github.com/mrwan200/promptparse-go/lib"

// Generate Slip Verify QR Code
//
// This also called "Mini-QR" that embedded in slip used for verify transactions
func SlipVerify(sendingBank string, transRef string) string {
	payload := []lib.TLVTag{
		lib.Tag("00", lib.Encode([]lib.TLVTag{
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

	tag, err := lib.WithCRCTag(lib.Encode(payload), "91")
	if err != nil {
		return ""
	}

	return tag
}
