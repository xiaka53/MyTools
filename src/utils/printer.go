package utils

import (
	"fmt"
)

//输出一行信息（无样式）
func OutputLn(a ...interface{}) {
	fmt.Println(a...)
}

//输出信息（不换行）
func Outputs(a ...interface{}) {
	fmt.Print(a...)
}

//绿色颜色输出信息（不换行）
func Info(a ...interface{}) {
	fmt.Print("\033[32m")
	Outputs(a...)
	noStyle()
}

//绿色颜色输出信息
func InfoLn(a ...interface{}) {
	fmt.Print("\033[32m")
	Outputs(a...)
	noStyle()
}

//红色颜色输出信息（不换行）
func ErrInfo(a ...interface{}) {
	fmt.Print("\033[31m")
	Outputs(a...)
	noStyle()
}

//红色颜色输出信息
func ErrInfoLn(a ...interface{}) {
	fmt.Print("\033[31m")
	Outputs(a...)
	noStyle()
}

//消除所有样式
func noStyle() {
	fmt.Print("\033[0m")
}
