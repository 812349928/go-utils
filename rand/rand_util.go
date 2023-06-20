package rand

import (
	"math/rand"
	"time"
)

func Intn(max int) int {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数
	return rand.Intn(max)
}
