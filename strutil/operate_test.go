package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToLower(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case	1",
			args: args{
				str: "ABC",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StrToLower(tt.args.str), "StrToLower(%v)", tt.args.str)
		})
	}
}

func TestStrConcat(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: []string{"a", "b", "c"},
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, StrConcat(tt.args.s...), "StrConcat(%v)")
		})
	}
}

func TestSubStr(t *testing.T) {
	type args struct {
		str   string
		start int64
		end   int64
	}
	tests := []struct {
		name    string
		args    args
		wantSub string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				str:   "abcdefgh",
				start: 1,
				end:   2,
			},
			wantSub: "bc",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSub, err := SubStr(tt.args.str, tt.args.start, tt.args.end)
			if !tt.wantErr(t, err, fmt.Sprintf("SubStr(%v, %v, %v)", tt.args.str, tt.args.start, tt.args.end)) {
				return
			}
			assert.Equalf(t, tt.wantSub, gotSub, "SubStr(%v, %v, %v)", tt.args.str, tt.args.start, tt.args.end)
		})
	}
}

func TestSubStrReturnLeft(t *testing.T) {
	type args struct {
		str    string
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:    "abcdefg",
				target: "c",
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubStrReturnLeft(tt.args.str, tt.args.target)
			assert.Equalf(t, tt.want, got, "SubStrReturnLeft(%v, %v)", tt.args.str, tt.args.target)
		})
	}
}

func TestSubStrReturnRight(t *testing.T) {
	type args struct {
		str    string
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:    "abcdefg",
				target: "c",
			},
			want: "defg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubStrReturnRight(tt.args.str, tt.args.target)
			assert.Equalf(t, tt.want, got, "SubStrReturnRight(%v, %v)", tt.args.str, tt.args.target)
		})
	}
}

func TestStringToUint64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				str: "8",
			},
			want:    uint64(8),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToUint64(tt.args.str)
			if !tt.wantErr(t, err, fmt.Sprintf("StringToUint64(%v)", tt.args.str)) {
				return
			}
			assert.Equalf(t, tt.want, got, "StringToUint64(%v)", tt.args.str)
		})
	}
}

func TestStringToInt64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				str: "8",
			},
			want:    int64(8),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt64(tt.args.str)
			if !tt.wantErr(t, err, fmt.Sprintf("StringToInt64(%v)", tt.args.str)) {
				return
			}
			assert.Equalf(t, tt.want, got, "StringToInt64(%v)", tt.args.str)
		})
	}
}

func TestStringToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "case 1",
			args: args{
				str: "8",
			},
			want:    int(8),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt(tt.args.str)
			if !tt.wantErr(t, err, fmt.Sprintf("StringToInt(%v)", tt.args.str)) {
				return
			}
			assert.Equalf(t, tt.want, got, "StringToInt(%v)", tt.args.str)
		})
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "case 1",
			args: args{
				str: "8",
			},
			want: []byte("8"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringToBytes(tt.args.str), "StringToBytes(%v)", tt.args.str)
		})
	}
}

func TestUtf8StringLen(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				str: "a",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Utf8StringLen(tt.args.str), "Utf8StringLen(%v)", tt.args.str)
		})
	}
}

func TestUtf8StringCut(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "abcd",
				n:   2,
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Utf8StringCut(tt.args.str, tt.args.n), "Utf8StringCut(%v, %v)", tt.args.str, tt.args.n)
		})
	}
}

func TestUtf8Index(t *testing.T) {
	type args struct {
		str    string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				str:    "abcd",
				substr: "c",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Utf8Index(tt.args.str, tt.args.substr), "Utf8Index(%v, %v)", tt.args.str, tt.args.substr)
		})
	}
}

func TestUcFirst(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "abc",
			},
			want: "Abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UcFirst(tt.args.str), "UcFirst(%v)", tt.args.str)
		})
	}
}

func TestLcFirst(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "ABC",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LcFirst(tt.args.str), "LcFirst(%v)", tt.args.str)
		})
	}
}

func TestCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "hello-world",
			},
			want: "helloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CamelCase(tt.args.str), "CamelCase(%v)", tt.args.str)
		})
	}
}

func TestCapitalize(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "abc",
			},
			want: "Abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Capitalize(tt.args.str), "Capitalize(%v)", tt.args.str)
		})
	}
}

func TestKebabCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "helloWorld",
			},
			want: "hello-world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, KebabCase(tt.args.s), "KebabCase(%v)", tt.args.s)
		})
	}
}

func TestSnakeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "hello-world",
			},
			want: "hello_world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SnakeCase(tt.args.str), "SnakeCase(%v)", tt.args.str)
		})
	}
}
func TestReverse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "abc",
			},
			want: "cba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Reverse(tt.args.str), "Reverse(%v)", tt.args.str)
		})
	}
}

func TestQuote(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "abc",
			},
			want: "\"abc\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Quote(tt.args.str), "Quote(%v)", tt.args.str)
		})
	}
}

func TestAddSlashes(t *testing.T) {
	assert.Equal(t, AddSlashes(`{"key": 123}`), `{\"key\": 123}`)
}

func TestStripSlashes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "{\\\"key\\\": 123}",
			},
			want: "{\"key\": 123}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StripSlashes(tt.args.str), "StripSlashes(%v)", tt.args.str)
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		str    string
		cutSet []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case	1",
			args: args{
				str:    "  abc  ",
				cutSet: nil,
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Trim(tt.args.str, tt.args.cutSet...), "Trim(%v, %v)", tt.args.str, tt.args.cutSet)
		})
	}
}

func TestLTrim(t *testing.T) {
	type args struct {
		str    string
		cutSet []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:    "  abc",
				cutSet: nil,
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LTrim(tt.args.str, tt.args.cutSet...), "LTrim(%v, %v)", tt.args.str, tt.args.cutSet)
		})
	}
}

func TestRTrim(t *testing.T) {
	type args struct {
		str    string
		cutSet []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:    "  abc",
				cutSet: nil,
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RTrim(tt.args.str, tt.args.cutSet...), "RTrim(%v, %v)", tt.args.str, tt.args.cutSet)
		})
	}
}

func TestPadding(t *testing.T) {
	type args struct {
		str    string
		pad    string
		length int
		pos    PosFlag
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:    "abc",
				pad:    "123",
				length: 9,
				pos:    0,
			},
			want: "123123123123123123abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Padding(tt.args.str, tt.args.pad, tt.args.length, tt.args.pos), "Padding(%v, %v, %v, %v)", tt.args.str, tt.args.pad, tt.args.length, tt.args.pos)
		})
	}
}

func TestPadLeft(t *testing.T) {
	type args struct {
		str    string
		pad    string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:    "abc",
				pad:    "1",
				length: 4,
			},
			want: "1abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PadLeft(tt.args.str, tt.args.pad, tt.args.length), "PadLeft(%v, %v, %v)", tt.args.str, tt.args.pad, tt.args.length)
		})
	}
}

func TestPadRight(t *testing.T) {
	type args struct {
		str    string
		pad    string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:    "abc",
				pad:    "1",
				length: 4,
			},
			want: "abc1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PadRight(tt.args.str, tt.args.pad, tt.args.length), "PadRight(%v, %v, %v)", tt.args.str, tt.args.pad, tt.args.length)
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		str   string
		times int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:   "abc",
				times: 2,
			},
			want: "abcabc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Repeat(tt.args.str, tt.args.times), "Repeat(%v, %v)", tt.args.str, tt.args.times)
		})
	}
}

func TestResize(t *testing.T) {
	type args struct {
		str    string
		length int
		align  PosFlag
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str:    "abc",
				length: 10,
				align:  0,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Resize(tt.args.str, tt.args.length, tt.args.align), "Resize(%v, %v, %v)", tt.args.str, tt.args.length, tt.args.align)
		})
	}
}

func TestChineseCount(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				str: "abc",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ChineseCount(tt.args.str), "ChineseCount(%v)", tt.args.str)
		})
	}
}

func TestGetFirstChineseChar(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "我是中国人",
			},
			want: "我",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetFirstChineseChar(tt.args.str), "GetFirstChineseChar(%v)", tt.args.str)
		})
	}
}

func TestGetChineseChar(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				str: "我1是a中国人",
			},
			want: []string{"我", "是", "中", "国", "人"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetChineseChar(tt.args.str), "GetChineseChar(%v)", tt.args.str)
		})
	}
}

func TestGetChineseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str: "我1是a中国人",
			},
			want: "我是中国人",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetChineseString(tt.args.str), "GetChineseString(%v)", tt.args.str)
		})
	}
}
