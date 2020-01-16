package app

import (
	"MyTools/src/utils"
	"flag"
)

func showHelp() {
	flag.Usage()
}

func usage() {
	str :=
		`一个私人助手，加密文件、密码等数据。
Usage:
  MyTools [options] [commands]

Options:
  -c, -config string    指定配置文件(default: ./config.json)。
  -v, -version          显示版本信息。
  -h, -help             显示帮助信息。

Commands:
  cp [-r] source target    复制传输。
  ${ServerNum}             使用编号登录指定服务器。
  ${ServerAlias}           使用别名登录指定服务器。
  upgrade                  检测并更新到最新版本。
`
	utils.OutputLn(str)
}
