package utils

import (
	"fmt"
	"strings"
)

// Generate an `UCS-2`-like? Hex string for Tag 81
func EncodeTag81(message string) string {
	var msg []string
	for _, val := range message {
		result := append(msg, fmt.Sprintf("%04X", int(val)))
		msg = result
	}

	return strings.ToUpper(strings.Join(msg, ""))
}
