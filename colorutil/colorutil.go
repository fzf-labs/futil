package colorutil

import (
	"strconv"
	"strings"
)

// ColorHexToRGB 将十六进制颜色转换为 RGB 颜色
func ColorHexToRGB(colorHex string) (red, green, blue int, err error) {
	colorHex = strings.TrimPrefix(colorHex, "#")
	color64, err := strconv.ParseInt(colorHex, 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	color := int(color64)
	return color >> 16, (color & 0x00FF00) >> 8, color & 0x0000FF, nil
}

// ColorRGBToHex 将 RGB 颜色转换为十六进制颜色
func ColorRGBToHex(red, green, blue int) string {
	r := strconv.FormatInt(int64(red), 16)
	g := strconv.FormatInt(int64(green), 16)
	b := strconv.FormatInt(int64(blue), 16)
	if len(r) == 1 {
		r = "0" + r
	}
	if len(g) == 1 {
		g = "0" + g
	}
	if len(b) == 1 {
		b = "0" + b
	}
	return "#" + r + g + b
}
