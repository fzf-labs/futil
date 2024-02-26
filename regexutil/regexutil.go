package regexutil

import "regexp"

// IsRegexMatch 检查字符串是否与正则表达式匹配
func IsRegexMatch(str, regex string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}
