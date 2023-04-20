package util

import (
	"strconv"
)

func Str2Int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func Str2Int64(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func Int2Str(i int) string {
	return strconv.Itoa(i)
}
