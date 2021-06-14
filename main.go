package main

import (
	"flag"
	"modelgenerator/cmd"
	_ "modelgenerator/cmd"
	"modelgenerator/conf"
	"modelgenerator/generate"
)

func main() {

	/**
	使用方法：修改配置文件，参考conf.yaml文件注释修改
	*/

	// 指令初始化
	flag.Parse()

	// 生成初始配置文件
	if cmd.IsInit {
		generate.InitConf()
		return
	}

	// 读取初始化配置文件
	conf.InitFile(cmd.ConfPath + "/conf.yaml")
	// 生成目标表单
	generate.Genertate(conf.Tconf)
}
