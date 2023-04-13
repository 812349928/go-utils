package sort

import "sort"

func Int64s(x []int64) {
	sort.Sort(Int64Slice(x))
}

type Int64Slice []int64

func (x Int64Slice) Len() int           { return len(x) }
func (x Int64Slice) Less(i, j int) bool { return x[i] < x[j] || (isNaN(x[i]) && !isNaN(x[j])) }
func (x Int64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func isNaN(f int64) bool {
	return f != f
}

type InterfaceSlice []int64

func (x InterfaceSlice) Len() int           { return len(x) }
func (x InterfaceSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x InterfaceSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
