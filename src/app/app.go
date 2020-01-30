package app

import (
	"flag"
)

var (
	Version string
	Build   string

	c       string //配置地址
	v       bool   //版本号
	h       bool   //帮助信息
	cp      bool   //复制
	upgrade bool   //升级版本
)

func init() {
	c = "./config.json"

	flag.StringVar(&c, "c", c, "配置文件路径")
	flag.StringVar(&c, "config", c, "配置文件路径")

	flag.BoolVar(&v, "v", v, "版本信息")
	flag.BoolVar(&v, "version", v, "版本信息")

	flag.BoolVar(&h, "h", h, "帮助信息")
	flag.BoolVar(&h, "help", h, "帮助信息")

	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) > 0 {
		arg := flag.Arg(0)
		switch arg {
		case "upgrade":
			upgrade = true
		case "cp":
			cp = true
		default:

		}
	}
}

func Run() {
	if v {
		showVersion()
	} else if h {
		showHelp()
	} else if upgrade {
		showUpgrade()
	} else if cp {

	} else {
		showServers()
	}
}
