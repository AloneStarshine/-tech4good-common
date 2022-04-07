package utils

import (
	"github.com/google/uuid"
	"strings"
)

// GetUUID 获取guid
func GetUUID() string {
	struuid := uuid.New().String()
	return strings.Replace(struuid, "-", "", -1)
}

// Split 支持多sep的字符串分割，简单使用递归，所以seps的个数不能超过50个
func Split(s string, seps ...string) []string {
	if len(seps) == 0 || len(seps[0]) == 0 {
		return []string{s}
	}
	if len(seps) == 1 {
		return strings.Split(s, seps[0])
	}
	curSep := seps[0]
	nextSeps := seps[1:]
	tRes := strings.Split(s, curSep)
	res := []string{}
	for _, item := range tRes {
		r := Split(item, nextSeps...)
		res = append(res, r...)
	}
	return res
}
