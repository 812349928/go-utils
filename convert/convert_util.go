package convert

import (
	"errors"
	"fmt"
	"github.com/812349928/go-utils/array"
	"math"
	"strconv"
	"strings"
)

var Items = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v'}

// ConvertByTen
// 十进制转换成其他进制
func ConvertToXByTen(n, to int64) string {
	if to == 10 {
		return fmt.Sprintf("%d", n)
	}

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
	arr = array.OverTurn[string](arr)
	x := strings.Join(arr, "")

	return x
}

// ConvertToTen
// 其他进制转换为十进制
func ConvertToTenByX(s string, from int64) (int64, error) {
	if b := CheckX(s, from); !b {
		return 0, errors.New("data is not legal")
	}

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

// 从某一进制转换到某一进制
func ConvertToXByY(s string, from, to int64) (string, error) {
	if from == to {
		return s, nil
	}

	n, err := ConvertToTenByX(s, from)
	if err != nil {
		return "", err
	}

	return ConvertToXByTen(n, to), nil
}

// 判断是否是合法的数据
func CheckX(s string, from int64) bool {
	m := make(map[byte]int)
	for i, item := range Items {
		m[item] = i
	}
	m[0] = -1

	for i, _ := range s {
		if m[s[i]] == 0 || m[s[i]] >= int(from) {
			return false
		}
	}

	return true
}
