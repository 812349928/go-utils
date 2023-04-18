package convert

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	Two = 2
	Ten = 10
)

var Items = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v'}

func ConvertToBinByTen(n int64) string {
	var x int64
	arr := make([]int64, 0)
	for {
		if n <= 1 {
			arr = append(arr, n)
			break
		} else {
			arr = append(arr, (n % Two))
			n /= Two
		}
	}
	for i, v := range arr {
		x += v * int64(math.Pow10(i))
	}

	return fmt.Sprintf("%08d", x)
}

// ConvertByTen
func ConvertByTen(n, to int64) string {
	arr := make([]string, 0)
	for {
		var i int64
		if n <= to-1 {
			i = n
			arr = append(arr, string(Items[i]))
			break
		} else {
			i = n % to
			n /= to
			arr = append(arr, string(Items[i]))
		}
	}
	x := strings.Join(arr, "")

	return x
}

// ConvertToTen
// 其他进制转换为十进制
func ConvertToTen(s string, from int64) (int64, error) {
	//如果本身就是十进制则不进行转换
	if from == 10 {
		return strconv.ParseInt(s, 10, 64)
	}

	arr := make([]int64, 0)
	var x int64
	for i, _ := range s {
		var n int64
		if s[i] >= '0' && s[i] <= '9' {
			n = int64(s[i] - '0')
		} else if s[i] >= 'a' && s[i] <= 'v' {
			n = int64(s[i] - 'a' + 10)
		}
		arr = append(arr, n)
	}

	l := len(arr)

	for i, v := range arr {
		x += v * int64(math.Pow(float64(from), float64(l-i-1)))
	}

	return x, nil
}

func ConvertToTenByBin(s string) int64 {
	arr := make([]int64, 0)
	var x int64
	for i, _ := range s {
		n, _ := strconv.ParseInt(s[i:i+1], Ten, 64)
		arr = append(arr, n)
	}

	l := len(arr)

	for i, v := range arr {
		x += v * int64(math.Pow(Two, float64(l-i-1)))
	}

	return x
}
