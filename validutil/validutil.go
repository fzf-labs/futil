package validutil

import (
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// IsZero 检查是否是零值
func IsZero(any any) bool {
	if any == nil {
		return true
	}
	v := reflect.ValueOf(any)
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr, reflect.Func:
		return v.IsNil()
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	}
}

// IsIP 检查字符串是否为 IP 地址。
func IsIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil
}

// IsIPV4 检查字符串是否为 ipv4 地址。
func IsIPV4(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return strings.Contains(ipStr, ".")
}

// IsIPV6 检查字符串是否为 ipv6 地址。
func IsIPV6(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return strings.Contains(ipStr, ":")
}

// IsPort 检查字符串是否是有效的网络端口。
func IsPort(port int) bool {
	if port > 0 && port <= 65535 {
		return true
	}
	return false
}

// IsURL 检查字符串是否为 url。
func IsURL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return regexp.MustCompile(`^((ftp|http|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]+([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`).MatchString(str)
}

// IsEmail check if the string is a email address.
func IsEmail(email string) bool {
	return regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`).MatchString(email)
}

// IsPhoneTight 手机验证
// Phone format validation.
//
//  1. China Mobile:
//     134, 135, 136, 137, 138, 139, 150, 151, 152, 157, 158, 159, 182, 183, 184, 187, 188,
//     178(4G), 147(Net)；
//     172
//
//  2. China Unicom:
//     130, 131, 132, 155, 156, 185, 186 ,176(4G), 145(Net), 175
//
//  3. China Telecom:
//     133, 153, 180, 181, 189, 177(4G)
//
//  4. Sate lite:
//     1349
//
//  5. Virtual:
//     170, 173
//
//  6. 2018:
//     16x, 19x
func IsPhoneTight(phone string) bool {
	regular := `^13[\d]{9}$|^14[5,7]{1}\d{8}$|^15[^4]{1}\d{8}$|^16[\d]{9}$|^17[0,2,3,5,6,7,8]{1}\d{8}$|^18[\d]{9}$|^19[\d]{9}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

// IsPhoneLoose 宽松的手机号验证
// 13, 14, 15, 16, 17, 18, 19 can pass the verification (只要满足 13、14、15、16、17、18、19开头的11位数字都可以通过验证)
func IsPhoneLoose(phone string) bool {
	regular := `^1(3|4|5|6|7|8|9)\d{9}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

// IsTelephone 固话校验
func IsTelephone(telephone string) bool {
	regular := `^((\d{3,4})|\d{3,4}-)?\d{7,8}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(telephone)
}

// IsPostalCode 邮政编码
func IsPostalCode(postalCode string) bool {
	regular := `^\d{6}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(postalCode)
}

// IsResidentID 身份证ID校验
//
// xxxxxx yyyy MM dd 375 0  十八位
// xxxxxx   yy MM dd  75 0  十五位
//
// 地区：     [1-9]\d{5}
// 年的前两位：(18|19|([23]\d))  1800-2399
// 年的后两位：\d{2}
// 月份：     ((0[1-9])|(10|11|12))
// 天数：     (([0-2][1-9])|10|20|30|31) 闰年不能禁止29+
//
// 三位顺序码：\d{3}
// 两位顺序码：\d{2}
// 校验码：   [0-9Xx]
//
// 十八位：^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$
// 十五位：^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$
//
// 总：
// (^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)
func IsResidentID(id string) bool {
	id = strings.ToUpper(strings.TrimSpace(id))
	if len(id) != 18 {
		return false
	}
	var (
		weightFactor = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		checkCode    = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
		last         = id[17]
		num          = 0
	)
	for i := 0; i < 17; i++ {
		tmp, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		num += tmp * weightFactor[i]
	}
	if checkCode[num%11] != last {
		return false
	}
	regular := `(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(id)
}

// IsQQ 腾讯qq校验
func IsQQ(qq string) bool {
	regular := `^[1-9][0-9]{4,}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(qq)
}

// IsPassport 校验护照
// Universal passport format rule:
// Starting with letter, containing only numbers or underscores, length between 6 and 18.
func IsPassport(p string) bool {
	regular := `^[a-zA-Z]{1}\w{5,17}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(p)
}

// IsWeakPassword 检查字符串是否为弱密码
// Weak password: 只有字母或只有数字或字母+数字
func IsWeakPassword(password string, length int) bool {
	if len(password) < length {
		return false
	}
	var num, letter, special bool
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsLetter(r):
			letter = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}
	return (num || letter) && !special
}

// IsStrongPassword 检查字符串是否为强密码，如果 len(password) 小于长度参数，则返回 false
// Strong password: alpha(lower+upper) + number + special chars(!@#$%^&*()?><)
func IsStrongPassword(password string, length int) bool {
	if len(password) < length {
		return false
	}
	var num, lower, upper, special bool
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return num && lower && upper && special
}

// IsDomain 校验域名
func IsDomain(p string) bool {
	regular1 := `^([0-9a-zA-Z][0-9a-zA-Z\-]{0,62}\.)+([a-zA-Z]{0,62})$`
	return regexp.MustCompile(regular1).MatchString(p)
}

// IsMac mac地址校验
func IsMac(p string) bool {
	regular1 := `^([0-9A-Fa-f]{2}[\-:]){5}[0-9A-Fa-f]{2}$`
	return regexp.MustCompile(regular1).MatchString(p)
}
