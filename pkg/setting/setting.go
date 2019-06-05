package setting

/**
本包主要是引入配置文件
引入具体的方法：
	1，定义对应配置文件中各个key的变量
	2，根据配置文件中的分组定义各个方法为变量赋值
*/

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// 对应了配置文件中各个key名的变量
var (
	// 加载配置文件  类型为 ini.File类型
	Cfg *ini.File

	//
	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init() {
	var err error
	// 加载配置文件
	Cfg, err = ini.Load("conf/app.ini")
	// 如果加载出现错误
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	// 对应了ini配置文件中的不同的分组
	LoadBase()	// 加载基础模块的配置部分
	LoadServer()  // 加载服务器配置
	LoadApp() 	// 加载app配置
}

func LoadBase() {
	//获取某个值的方法： 键对应的值 = Cfg.Section("分组名").Key("健名").MustString("键对应的值")
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	// 获取整个分组的方法 : 分组, err := Cfg.GetSection("分组名")
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	// 知道了分组后获取对应值的方法 值 = 组.Key("HTTP_PORT").转换函数(默认值)
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	// Time 包中，定义有一个名为 Duration 的类型和一些辅助的常量
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// 分组中: jwtsecret ， 分页
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
