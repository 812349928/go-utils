package sort

import (
	"sort"
	"time"
)

func Int64s(x []int64) {
	sort.Sort(Int64Slice(x))
}

func Times(x []time.Time) {
	sort.Sort(TimeSlice(x))
}

type Int64Slice []int64

func (x Int64Slice) Len() int           { return len(x) }
func (x Int64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Int64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type TimeSlice []time.Time

func (x TimeSlice) Len() int           { return len(x) }
func (x TimeSlice) Less(i, j int) bool { return x[i].UnixMilli() < x[j].UnixMilli() }
func (x TimeSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
