package utils

import (
	"net/url"
	"strings"
	"unicode"
)

// 自定义URL编码
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

// 判断字符串是否包含中文
func isChinese(s string) bool {
	for _, r := range s {
		// 使用标准库的方法来判断是否包含中文
		// r >= 0x4E00 && r <= 0x9FFF 
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}
