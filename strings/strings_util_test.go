package strings

import (
	"fmt"
	"testing"
)

func Test_RepeatLeft(t *testing.T) {
	s := "asdf"
	count := 10
	pad := 'v'
	fmt.Println(RepeatLeft(s, pad, count))
}

func Test_RepeatRight(t *testing.T) {
	s := "asdf"
	count := 10
	pad := 'v'
	fmt.Println(RepeatRight(s, pad, count))
}
