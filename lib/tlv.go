package lib

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrwan200/promptparse-go/utils"
)

type TLVTag struct {
	ID      string
	Value   string
	SubTags []TLVTag
	Length  int
}

// Decode TLV string into array of TLV Tags
func Decode(payload string) (*[]TLVTag, error) {
	var tags []TLVTag

	idx := 0
	for {
		if idx >= len(payload) {
			break
		}

		data := payload[idx:]
		id := data[0:2]
		length, err := strconv.Atoi(data[2:4])
		if err != nil {
			return nil, err
		}

		value := data[4 : 4+length]
		result := append(tags, TLVTag{
			ID:     id,
			Length: length,
			Value:  value,
		})
		tags = result
		idx += length + 4
	}

	return &tags, nil
}

// Encode TLV Tags array into TLV string
func Encode(tags []TLVTag) string {
	var payload string

	for _, tag := range tags {
		payload += tag.ID
		// Get length
		length := ("00" + fmt.Sprintf("%d", tag.Length))
		payload += length[len(length)-2:]
		if len(tag.SubTags) > 0 {
			result := Encode(tag.SubTags)
			// Append it
			payload += result
		}
		payload += tag.Value
	}

	return payload
}

// Generate CRC Checksum for provided string
func Checksum(payload string) (string, error) {
	sum, err := utils.CRC16XModem(payload, 0xffff)
	if err != nil {
		return "", err
	}
	result := strings.ToUpper(fmt.Sprintf("%04x", sum))
	return result, nil
}

// Get TLV string combined with CRC Tag
func WithCRCTag(payload string, crcTagId string) (string, error) {
	payload += fmt.Sprintf("%02s", crcTagId)
	payload += "04"
	// Checksum
	crc, err := Checksum(payload)
	if err != nil {
		return "", err
	}
	payload += crc

	return payload, nil

}

// Get Tag or Sub-tag by Tag ID in array of TLV Tags
func Get(tlvTags []TLVTag, tagId string, subTagId string) TLVTag {
	var tag TLVTag

	for _, t := range tlvTags {
		if t.ID == tagId {
			tag = t
			break
		}
	}

	// Check if subtag
	if subTagId != "" {
		for _, s := range tag.SubTags {
			if s.ID == subTagId {
				return s
			}
		}
	}

	return tag
}

// Create new TLV Tag
func Tag(tagId string, value string) TLVTag {
	return TLVTag{
		ID:     tagId,
		Length: len(value),
		Value:  value,
	}
}
