package array

import (
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
