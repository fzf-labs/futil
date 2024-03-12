package jsonutil

import (
	"bytes"
	"regexp"
	"strings"
	"text/scanner"
)

// StripComments 去除 JSON 字符串的注释
func StripComments(src string) string {
	// multi line comments
	if strings.Contains(src, "/*") {
		src = regexp.MustCompile(`(?s:/\*.*?\*/\s*)`).ReplaceAllString(src, "")
	}

	// single line comments
	if !strings.Contains(src, "//") {
		return strings.TrimSpace(src)
	}

	// strip inline comments
	var s scanner.Scanner

	s.Init(strings.NewReader(src))
	s.Filename = "comments"
	s.Mode ^= scanner.SkipComments // don't skip comments

	buf := new(bytes.Buffer)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if !strings.HasPrefix(txt, "//") && !strings.HasPrefix(txt, "/*") {
			buf.WriteString(txt)
			// } else {
			// fmt.Printf("%s: %s\n", s.Position, txt)
		}
	}

	return buf.String()
}
