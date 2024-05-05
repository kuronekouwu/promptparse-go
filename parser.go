package promptparsego

import (
	"regexp"

	"github.com/mrwan200/promptparse-go/lib"
)

func Parse(payload string, strict bool, subTags bool) *lib.EMVCoQRStruct {
	reg, err := regexp.Compile(`^\d{4}.+`)
	if err != nil {
		return nil
	}

	match := reg.MatchString(payload)
	if !match {
		return nil
	}

	if strict {
		excepted := payload[len(payload)-4:]
		calculated, _ := lib.Checksum(payload[:len(payload)-4])
		if excepted != calculated {
			return nil
		}
	}

	// Decode
	tags, err := lib.Decode(payload)
	if err != nil {
		return nil
	}

	if subTags {

		for idx, tag := range *tags {
			// Check if invalid
			match := reg.MatchString(tag.Value)
			if !match {
				continue
			}

			sub, err := lib.Decode(tag.Value)
			if err != nil {
				return nil
			}

			for _, tag := range *sub {
				if tag.Length == 0 || tag.Length != len(tag.Value) {
					continue
				}
			}

			tag.SubTags = *sub
			(*tags)[idx] = tag
		}
	}

	emv := lib.EMVCoQRStruct{
		Payload: payload,
		Tags:    *tags,
	}
	return &emv
}

func ParseBarcode(payload string) lib.BOTBarcode {
	barcode := lib.BOTBarcode{}
	barcode.FromString(payload)

	return barcode
}
