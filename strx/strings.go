package strx

import (
	"math/rand"
	"regexp"
	"time"
)

var htmlRe = regexp.MustCompile(`</?[^>]+>`)

func RemoveHtml(s string) string {
	return htmlRe.ReplaceAllString(s, "")
}

// TruncateStr truncates a string to a specified length, supporting multibyte characters.
func TruncateStr(s string, length int) string {
	if length <= 0 {
		return ""
	}
	rs := []rune(s)
	if len(rs) <= length {
		return s
	}
	return string(rs[:length])
}

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 定义随机字符串的字符集
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 生成随机字符串
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charSet[r.Intn(len(charSet))]
	}

	return string(randomString)
}
