package strutil

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// StrToLower 转换成小写字母
func StrToLower(str string) string {
	runeArr := []rune(str)
	for i := range runeArr {
		if runeArr[i] >= 65 && runeArr[i] <= 90 {
			runeArr[i] += 32
		}
	}
	return string(runeArr)
}

// StrConcat 连接字符串,性能比fmt.Sprintf和+号要好
func StrConcat(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, i := range s {
		buffer.WriteString(i)
	}
	return buffer.String()
}

// SubStr 截取字符串，并返回实际截取的长度和子串
func SubStr(str string, start, end int64) (sub string, err error) {
	reader := strings.NewReader(str)
	// Calling NewSectionReader method with its parameters
	r := io.NewSectionReader(reader, start, end)
	// Calling Copy method with its parameters
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	sub = buf.String()
	return sub, err
}

// SubStrReturnLeft 截取并返回left部分
func SubStrReturnLeft(str, target string) string {
	pos := strings.Index(str, target)
	if pos == -1 {
		return ""
	}
	return str[:pos]
}

// SubStrReturnRight 截取并返回right部分
func SubStrReturnRight(str, target string) string {
	pos := strings.Index(str, target)
	if pos == -1 {
		return ""
	}
	return str[pos+len(target):]
}

// StringToUint64 字符串转uint64
func StringToUint64(str string) (uint64, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint64(valInt), nil
}

// StringToInt64 字符串转int64
func StringToInt64(str string) (int64, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return int64(valInt), nil
}

// StringToInt 字符串转int
func StringToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return valInt, nil
}

// StringToBytes 字符串转字节切片
func StringToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Utf8StringLen 获得字符串按照uft8编码的长度
func Utf8StringLen(str string) int {
	return utf8.RuneCountInString(str)
}

// Utf8StringCut 按照uft8编码截取字符串
func Utf8StringCut(str string, n int) string {
	var result string
	runes := []rune(str)
	for i := range runes {
		result = string(runes[:i])
		size := utf8.RuneCountInString(result)
		if size >= n {
			break
		}
	}
	return result
}

// Utf8Index 按照uft8编码匹配子串，返回开头的索引
func Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:index])
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	if str == "" {
		return str
	}

	rs := []rune(str)
	f := rs[0]

	if 'a' <= f && f <= 'z' {
		return string(unicode.ToUpper(f)) + string(rs[1:])
	}
	return str
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	if str == "" {
		return str
	}

	rs := []rune(str)
	f := rs[0]

	if 'A' <= f && f <= 'Z' {
		return string(unicode.ToLower(f)) + string(rs[1:])
	}
	return str
}

// CamelCase 将字符串转换为驼峰式字符串。
func CamelCase(str string) string {
	if str == "" {
		return ""
	}
	result := ""
	blankSpace := " "
	regex := regexp.MustCompile("[-_&]+")
	ss := regex.ReplaceAllString(str, blankSpace)
	for i, v := range strings.Split(ss, blankSpace) {
		vv := []rune(v)
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[0] += 32
			}
			result += string(vv)
		} else {
			result += Capitalize(v)
		}
	}

	return result
}

// Capitalize 将字符串的第一个字符转换为大写，其余字符转换为小写。
func Capitalize(str string) string {
	if str == "" {
		return ""
	}
	out := make([]rune, len(str))
	for i, v := range str {
		if i == 0 {
			out[i] = unicode.ToUpper(v)
		} else {
			out[i] = unicode.ToLower(v)
		}
	}
	return string(out)
}

// KebabCase 将字符串转换为 kebab-case
func KebabCase(str string) string {
	if str == "" {
		return ""
	}
	regex := regexp.MustCompile(`[\W_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(str, blankSpace)
	rs := strings.Split(match, blankSpace)
	var result []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			result = append(result, splitWords...)
		}
	}
	return strings.Join(result, "-")
}

// SnakeCase 将字符串转换为蛇形大小写
func SnakeCase(str string) string {
	if str == "" {
		return ""
	}
	regex := regexp.MustCompile(`[\W_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(str, blankSpace)
	rs := strings.Split(match, blankSpace)
	var result []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			result = append(result, splitWords...)
		}
	}
	return strings.Join(result, "_")
}

// splitWordsToLower 将将字符串拆分为多个字符串,并转小写
func splitWordsToLower(str string) []string {
	var result []string
	upperIndexes := upperIndex(str)
	l := len(upperIndexes)
	if upperIndexes == nil || l == 0 {
		if str != "" {
			result = append(result, str)
		}
		return result
	}
	for i := 0; i < l; i++ {
		if i < l-1 {
			result = append(result, strings.ToLower(str[upperIndexes[i]:upperIndexes[i+1]]))
		} else {
			result = append(result, strings.ToLower(str[upperIndexes[i]:]))
		}
	}
	return result
}

// upperIndex 得到一个 int 切片，其中元素都是字符串的大写 char 索引
func upperIndex(str string) []int {
	var result []int
	for i := 0; i < len(str); i++ {
		if 64 < str[i] && str[i] < 91 {
			result = append(result, i)
		}
	}
	if str != "" && result != nil && result[0] != 0 {
		result = append([]int{0}, result...)
	}
	return result
}

// Reverse 返回字符顺序与给定字符串相反的字符串
func Reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Quote 返回双引号的字符串
func Quote(str string) string {
	return strconv.Quote(str)
}

// AddSlashes 为字符串添加斜线。
func AddSlashes(str string) string {
	if ln := len(str); ln == 0 {
		return ""
	}
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// StripSlashes 去除字符串的斜杠。
func StripSlashes(str string) string {
	ln := len(str)
	if ln == 0 {
		return ""
	}
	var skip bool
	var buf bytes.Buffer
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < ln && str[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Trim  如果 cutSet 为空，将去除空格。
func Trim(str string, cutSet ...string) string {
	if ln := len(cutSet); ln > 0 && cutSet[0] != "" {
		if ln == 1 {
			return strings.Trim(str, cutSet[0])
		}
		return strings.Trim(str, strings.Join(cutSet, ""))
	}
	return strings.TrimSpace(str)
}

// LTrim 字符串中的字符。如果 cutSet 为空，将去除空格。
func LTrim(str string, cutSet ...string) string {
	if ln := len(cutSet); ln > 0 && cutSet[0] != "" {
		if ln == 1 {
			return strings.TrimLeft(str, cutSet[0])
		}

		return strings.TrimLeft(str, strings.Join(cutSet, ""))
	}
	return strings.TrimLeft(str, " ")
}

// RTrim 字符串中的字符。如果 cutSet 为空，将去除空格。
func RTrim(str string, cutSet ...string) string {
	if ln := len(cutSet); ln > 0 && cutSet[0] != "" {
		if ln == 1 {
			return strings.TrimRight(str, cutSet[0])
		}
		return strings.TrimRight(str, strings.Join(cutSet, ""))
	}
	return strings.TrimRight(str, " ")
}

// PosFlag type
type PosFlag uint8

// Position for padding/resize string
const (
	PosLeft PosFlag = iota
	PosRight
	PosMiddle
)

// Padding 填充字符串。
func Padding(str, pad string, length int, pos PosFlag) string {
	diff := len(str) - length
	if diff >= 0 { // do not need padding.
		return str
	}
	if pad == "" || pad == " " {
		mark := ""
		if pos == PosRight { // to right
			mark = "-"
		}
		// padding left: "%7s", padding right: "%-7s"
		tpl := fmt.Sprintf("%str%d", mark, length)
		return fmt.Sprintf(`%`+tpl+`str`, str)
	}
	if pos == PosRight { // to right
		return str + Repeat(pad, -diff)
	}
	return Repeat(pad, -diff) + str
}

// PadLeft 左边填充一个字符串
func PadLeft(str, pad string, length int) string {
	return Padding(str, pad, length, PosLeft)
}

// PadRight 右边填充一个字符串
func PadRight(str, pad string, length int) string {
	return Padding(str, pad, length, PosRight)
}

// Repeat 重复一个字符串
func Repeat(str string, times int) string {
	if times <= 0 {
		return ""
	}
	if times == 1 {
		return str
	}
	ss := make([]string, 0, times)
	for i := 0; i < times; i++ {
		ss = append(ss, str)
	}
	return strings.Join(ss, "")
}

// Resize 按给定的长度和对齐设置调整字符串的大小。填充空间。
func Resize(str string, length int, align PosFlag) string {
	diff := len(str) - length
	if diff >= 0 { // do not need padding.
		return str
	}
	if align == PosMiddle {
		strLn := len(str)
		padLn := (length - strLn) / 2
		padStr := string(make([]byte, padLn))
		if diff := length - padLn*2; diff > 0 {
			str += " "
		}
		return padStr + str + padStr
	}
	return Padding(str, " ", length, align)
}

// ChineseCount 中文字符统计
func ChineseCount(str string) int {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	return count
}

// GetFirstChineseChar 返回第一个中文字符
func GetFirstChineseChar(str string) string {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return fmt.Sprintf("%c", r)
		}
	}
	return ""
}

// GetChineseChar 过滤返回中文字符切片
func GetChineseChar(str string) []string {
	ss := make([]string, 0)
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			ss = append(ss, fmt.Sprintf("%c", r))
		}
	}
	return ss
}

// GetChineseString 过滤返回中文字符
func GetChineseString(str string) string {
	var ss string
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			ss += fmt.Sprintf("%c", r)
		}
	}
	return ss
}
