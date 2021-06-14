/*
	生成初始模板文件
*/
package generate

import (
	"modelgenerator/cmd"
	"modelgenerator/util"
)

var init_content = "#配置文件，按照要求修改输入数据，当前仅支持mysql数据库\n" +
	"Database:\n" +
	"  #数据库配置\n" +
	"  Host: 127.0.0.1\n" +
	"  Port: 3306\n" +
	"  User: test\n" +
	"  Password: 123321\n" +
	"  DBName: gotest\n" +
	"GModel:\n" +
	"  #存储路径，最后面也需要分割符\n" +
	"  StorePath: \"./dao/models/\"\n" +
	"  #是否覆盖原本的文件，false为不覆盖\n" +
	"  ModelCover: true\n" +
	"  #存储的package名\n" +
	"  PackageName: \"models\"\n" +
	"  #需要生成model的表名，支持多个，支持全库扫描（不输入任何参数即可）\n" +
	"  TableName:\n" +
	"    - \"sreenshot_ods\"\n"

func InitConf() {
	if cmd.InitPath[len(cmd.InitPath)-1] != '/' {
		cmd.InitPath += "/"
	}
	util.MakeFile(cmd.InitPath, cmd.InitPath+"conf.yaml", init_content, true)
}
