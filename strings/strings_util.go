package strings

import (
	"fmt"
	"strings"
)

// RepeatLeft
// @description:字符串左边补位
func RepeatLeft(s string, pad rune, count int) string {
	return fmt.Sprintf("%s%s", strings.Repeat(string(pad), count-len(s)), s)
}

// RepeatRight
func RepeatRight(s string, pad rune, count int) string {
	return fmt.Sprintf("%s%s", s, strings.Repeat(string(pad), count-len(s)))
}
