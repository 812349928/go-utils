package time

import (
	"errors"
	"github.com/812349928/go-utils/sort"
	"time"
)

const (
	ArrayIsEmpty = "array is empty"
)

func GetMax(ts ...time.Time) time.Time {
	if arr, err := getArrToTimes(ts...); err != nil {
		return time.Time{}
	} else {
		sort.Int64s(arr)
		return time.UnixMilli(int64(arr[len(arr)-1]))
	}
}

func GetMin(ts ...time.Time) time.Time {
	if arr, err := getArrToTimes(ts...); err != nil {
		return time.Time{}
	} else {
		sort.Int64s(arr)
		return time.UnixMilli(int64(arr[0]))
	}
}

func getArrToTimes(ts ...time.Time) ([]int64, error) {
	if len(ts) == 0 {
		return nil, errors.New(ArrayIsEmpty)
	}
	arr := make([]int64, 0)
	for _, t := range ts {
		arr = append(arr, t.UnixMilli())
	}

	return arr, nil
}
