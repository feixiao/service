package main

import (
	"strings"
	"errors"
)

// ErrEmpty 提示输入为空
var ErrEmpty = errors.New("empty string")

// 服务定义

// StringService 提供了针对string的操作
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int 
}

type stringService struct{}

func (stringService) Uppercase(s string )(string,error) {
	if s == "" {
		return "",ErrEmpty
	}
	return strings.ToUpper(s),nil
}

func (stringService) Count(s string) int {
	return len(s)
}


type ServiceMiddleware func(StringService) StringService