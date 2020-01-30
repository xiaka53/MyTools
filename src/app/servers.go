package app

import (
	"MyTools/src/utils"
)

func showServers() {
	_ = utils.Clear() //清屏
	show()
	for {
		scanInput()
	}
}

func show() {
	maxLen := 60
	utils.InfoLn(utils.FormatSeparator(" 欢迎使用 My Tools 我的小助手 ", "=", maxLen))

	showMenu()
	utils.Info("请输入操作方法或者账号序号： ")
}
