package strutil

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsEmpty 是否是空字符串
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return strings.TrimSpace(s) == ""
}

// IsNumeric 如果给定的字符串是数字，则返回 true，否则返回 false。
func IsNumeric(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

// IsAlpha 检查字符串是否只包含字母 (a-zA-Z)
func IsAlpha(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(s)
}

// IsAlphaNumeric 判断字符串是否为字母数字
func IsAlphaNumeric(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(str)
}

// IsAllUpperAlpha 检查字符串是否都是大写字母 A-Z
func IsAllUpperAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return s != ""
}

// IsAllLowerAlpha 检查字符串是否都是小写字母 a-z
func IsAllLowerAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return s != ""
}

// IsContainAlpha 检查字符串是否至少包含一个字母
func IsContainAlpha(str string) bool {
	return regexp.MustCompile(`[a-zA-Z]`).MatchString(str)
}

// IsContainUpperAlpha 检查字符串是否至少包含一个大写字母 A-Z
func IsContainUpperAlpha(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// IsContainLowerAlpha 检查字符串是否至少包含一个小写字母 a-z
func IsContainLowerAlpha(str string) bool {
	for _, r := range str {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// IsFloatStr 检查字符串是否可以转换为浮点数。
func IsFloatStr(str string) bool {
	_, e := strconv.ParseFloat(str, 64)
	return e == nil
}

// IsIntStr check if the string can convert to a integer.
func IsIntStr(str string) bool {
	return regexp.MustCompile(`^[\+-]?\d+$`).MatchString(str)
}

// IsChinese 检查字符串是中文字符
func IsChinese(s string) bool {
	return regexp.MustCompile("[\u4e00-\u9fa5]").MatchString(s)
}

// IsContainChinese 是否包含中文字符
func IsContainChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

// IsBase64 检查字符串是否为 base64 字符串。
func IsBase64(base64 string) bool {
	return regexp.MustCompile(`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`).MatchString(base64)
}

// IsValidUtf8 有效的 utf8 字符串检查
func IsValidUtf8(s string) bool {
	return utf8.ValidString(s)
}

// IsRuneWord char: a-zA-Z
func IsRuneWord(c rune) bool {
	return IsRuneLower(c) || IsRuneUpper(c)
}

// IsRuneLower char
func IsRuneLower(c rune) bool {
	return 'a' <= c && c <= 'z'
}

// IsRuneUpper char
func IsRuneUpper(c rune) bool {
	return 'A' <= c && c <= 'Z'
}

// NoCaseEq 检查两个字符串是否相等且不区分大小写
func NoCaseEq(s, t string) bool {
	return strings.EqualFold(s, t)
}

// HasOneSub 判断给定字符串是否有子字符串。
func HasOneSub(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// HasAllSubs 给定字符串中的所有子字符串。
func HasAllSubs(s string, subs []string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// HasOnePrefix 字符串以其中一个子项开头
func HasOnePrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}
