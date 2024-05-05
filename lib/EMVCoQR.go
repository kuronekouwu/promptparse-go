package lib

type EMVCoQRStruct struct {
	Payload string
	Tags    []TLVTag
}

func (emv *EMVCoQRStruct) GetTag(tagId string, subTagId string) TLVTag {
	return Get(emv.Tags, tagId, subTagId)
}

func (emv *EMVCoQRStruct) GetTagValue(tagId string, subTagId string) string {
	return Get(emv.Tags, tagId, subTagId).Value
}

func (emv *EMVCoQRStruct) GetTags(tagId string, subTagId string) []TLVTag {
	return emv.Tags
}

func (emv *EMVCoQRStruct) GetPayload() string {
	return emv.Payload
}

func (emv *EMVCoQRStruct) Validate(crcTagId string) (string, error) {
	var tlvTags []TLVTag

	for _, tags := range emv.Tags {
		if tags.ID != crcTagId {
			tlvTags = append(tlvTags, tags)
		}
	}

	// Encode tag
	encoded := Encode(tlvTags)

	expected, err := WithCRCTag(encoded, crcTagId)
	if err != nil {
		return "", err
	}

	return expected, nil
}
