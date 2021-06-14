package cmd

import "flag"

//声明变量用于接收命令行传入的参数值
var (
	IsInit   bool   // 判断是否为选择init还是生成model
	InitPath string // 初始化基本配置文件的路径
	ConfPath string // 依据配置文件生成model
)

func init() {
	flag.BoolVar(&IsInit,
		"isInit",
		false,
		"生成初始配置文件")

	flag.StringVar(&InitPath,
		"initPath",
		"./",
		"初始化文件地址")

	flag.StringVar(&ConfPath,
		"confPath",
		"./",
		"配置文件地址")
}
