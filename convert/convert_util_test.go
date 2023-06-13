package convert

import (
	"fmt"
	"github.com/812349928/go-utils/strings"
	"testing"
)

func Test_ConvertToXByY(t *testing.T) {
	s := "12"
	r, _ := ConvertToXByY(s, 10, 2)
	r = strings.RepeatLeft(r, '0', 8)
	fmt.Println(r)
}
