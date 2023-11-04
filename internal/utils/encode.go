package utils

import (
	"net/url"
	"strings"
)

func CustomURLEncode(str string) string {
	var encodedParts []string
	parts := strings.Split(str, "/")

	for _, part := range parts {
		if isChinese(part) {
			part = url.QueryEscape(part)
		}
		encodedParts = append(encodedParts, part)
	}

	return strings.Join(encodedParts, "/")
}

func isChinese(s string) bool {
	for _, r := range s {
		if r >= 0x4E00 && r <= 0x9FFF {
			return true
		}
	}
	return false
}
