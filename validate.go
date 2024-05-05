package promptparsego

type SlipVerifyStruct struct {
	sendingBank string
	transRef    string
}

func SlipVerify(payload string) *SlipVerifyStruct {
	ppqr := Parse(payload, true, true)

	apiType := ppqr.GetTagValue("00", "00")
	sendingBank := ppqr.GetTagValue("00", "01")
	transRef := ppqr.GetTagValue("00", "02")

	if apiType != "000001" || sendingBank == "" || transRef == "" {
		return nil
	}

	return &SlipVerifyStruct{
		sendingBank: sendingBank,
		transRef:    transRef,
	}
}
