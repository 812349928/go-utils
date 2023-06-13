package http

import (
	"fmt"
	"testing"
)

func Test_GetIp(t *testing.T) {
	ip, _ := GetIp()
	fmt.Println(GetIpPosition(ip))
	fmt.Println(GetIp())
}
