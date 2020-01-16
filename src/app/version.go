package app

import "MyTools/src/utils"

//版本输出
func showVersion() {
	utils.OutputLn("MyTools " + Version + " Build " + Build + "。")
	utils.OutputLn("开发者：陶然  博客：https://blog.cocofan.cn")
	utils.OutputLn("项目地址：https://github.com/xiaka53/MyTools")
}
