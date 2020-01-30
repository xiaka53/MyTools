package utils

import (
	"unicode"
)

//计算字符串长度
func ZhLen(str string) (length int) {
	for _, c := range str {
		if unicode.Is(unicode.Scripts["Han"], c) {
			length += 2
		} else {
			length += 1
		}
	}
	return
}

//左右填充
func FormatSeparator(title, c string, maxlength int) string {
	charslen := (maxlength - ZhLen(title)) / 2
	chars := ""
	for i := 0; i < charslen; i++ {
		chars += c
	}
	return chars + title + chars
}
