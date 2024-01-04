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
			// 使用url.PathEscape而不是url.QueryEscape，处理特殊字符时不转义，如&被转义成%26
			//part = url.QueryEscape(part)
			part = url.PathEscape(part)
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
