package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// SplitToInt64Array split str with sep, return int64 array
func SplitToInt64Array(s, sep string) ([]int64, error) {
	if s == "" {
		return []int64{}, nil
	}
	
	splitList := strings.Split(s, sep)
	result := make([]int64, len(splitList))
	for i, v := range splitList {
		id, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = id
	}
	return result, nil
}