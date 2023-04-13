package time

import (
	"errors"
	"sort"
	"time"
)

const (
	ArrayIsEmpty = "array is empty"
)

func GetMax(ts ...time.Time) time.Time {
	if arr, err := getArrToTimes(ts...); err != nil {
		return time.Time{}
	} else {
		sort.Float64s(arr)
		return time.UnixMilli(int64(arr[len(arr)-1]))
	}
}

func GetMin(ts ...time.Time) time.Time {
	if arr, err := getArrToTimes(ts...); err != nil {
		return time.Time{}
	} else {
		sort.Float64s(arr)
		return time.UnixMilli(int64(arr[0]))
	}
}

func getArrToTimes(ts ...time.Time) ([]float64, error) {
	if len(ts) == 0 {
		return nil, errors.New(ArrayIsEmpty)
	}
	arr := make([]float64, 0)
	for _, t := range ts {
		arr = append(arr, float64(t.UnixMilli()))
	}

	return arr, nil
}
