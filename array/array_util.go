package array

import (
	"fmt"
	"reflect"
)

//func DeleteItems(arr interface{}, item ...interface{}) interface{} {
//	m := make(map[interface{}][]int)
//	indexs := make([]int, 0)
//
//	av := reflect.ValueOf(arr)
//	switch reflect.TypeOf(arr).Kind() {
//	case reflect.Slice, reflect.Array:
//		for j := 0; j < av.Len(); j++ {
//			m[av.Index(j)] = append(m[av.Index(j)], j)
//		}
//	default:
//		return nil
//	}
//
//	for _, k := range item {
//		if len(m[k]) > 0 {
//			indexs = append(indexs, m[k]...)
//		}
//	}
//
//	sort.Ints(indexs)
//	l := len(indexs)
//	i := 1
//	for j := l - 1; j >= 0; j-- {
//		av.
//	}
//
//	return nil
//}
//
//func StringToInt64(a []string) []int64 {
//
//}

func IsExist(arr interface{}, item interface{}) bool {
	av := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < av.Len(); i++ {
			if av.Index(i).Interface() == item {
				return true
			}
		}
	default:
		return false
	}

	return false
}

// TODO:去重
type Interface interface {
	Index() int
	Map() map[interface{}]bool
	IsExist(o map[interface{}]bool, x interface{}) bool
	Remove()
}

func DuplicateRemoval(x []int64) {
	m := make(map[int64]bool)
	l := len(x)
	y := make([]int64, 0)
	for i := 0; i < l; i++ {
		if !m[x[i]] {
			m[x[i]] = true
			y = append(y, x[i])
		}
	}
	fmt.Println(l)
	x = make([]int64, len(y))
	x = y
}

// 数组翻转
func OverTurn[T any](a []T) []T {
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
	return a
}
