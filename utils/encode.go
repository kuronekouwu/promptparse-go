package utils

import (
	"fmt"
	"strings"
)

func EncodeTag81(message string) string {
	var msg []string
	for _, val := range message {
		// hex := strconv.FormatInt(int64(val), 16)

		result := append(msg, fmt.Sprintf("%04X", int(val)))
		msg = result
	}

	return strings.ToUpper(strings.Join(msg, ""))
}
