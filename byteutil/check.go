package byteutil

// IsLetterUpper 检查给定字节 b 是否为大写。
func IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

// IsLetterLower 检查给定字节 b 是否为小写。
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// IsLetter 检查给定的字节 b 是否是一个字母。
func IsLetter(b byte) bool {
	return IsLetterUpper(b) || IsLetterLower(b)
}

// IsNumChar 检查给定的字节 b 是否是一个数字字符。
func IsNumChar(b byte) bool { return b >= '0' && b <= '9' }

// IsAlpha 判断字节是 字母
func IsAlpha(b byte) bool {
	// A 65 -> Z 90
	if b >= 'A' && b <= 'Z' {
		return true
	}
	// a 97 -> z 122
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}
