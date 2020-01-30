package utils

import (
	"math/rand"
	"time"
)

type strs string

var R *rand.Rand

var NumStr strs
var NumEnStr strs
var NumEngStr strs

func init() {
	R = rand.New(rand.NewSource(time.Now().UnixNano()))
	NumStr = strs("1234567890")
	NumEnStr = strs("0123456789abcdefghijklmnopqrstuvwxyz")
	NumEngStr = strs("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

//随机字符串
func (str strs) Run(l int) string {
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < l; i++ {
		result = append(result, bytes[R.Intn(len(bytes))])
	}
	return string(result)
}
