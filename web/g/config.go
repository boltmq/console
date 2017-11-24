package g

import (
	"fmt"
	"git.oschina.net/cloudzone/cloudcommon-go/logger"
	"git.oschina.net/cloudzone/cloudcommon-go/web"
	"git.oschina.net/cloudzone/smartgo/stgcommon"
	"github.com/BurntSushi/toml"
	"github.com/toolkits/file"
	"log"
	"os"
	"reflect"
	"strings"
)

const (
	cfgName = "cfg.toml" // 配置文件名称
	cfgDir  = "conf"     // 配置文件父目录
)

var (
	cfg Config
)

type Config struct {
	Web web.Config    `json:"web" toml:"web"`
	Log logger.Config `json:"log" toml:"log"`
}

// InitLogger 初始化日志
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/24
func InitLogger(configPath string) {
	cfgPath := configPath
	if cfgPath == "" {
		cfgPath = getLoggerConfigPath()
	}

	_, err := toml.DecodeFile(cfgPath, &cfg)
	if err != nil {
		errMsg := fmt.Errorf("parse toml err: %s", err.Error())
		panic(errMsg)
	}

	log.Printf("read config file %s success \n", cfgPath)
	logger.SetCustomConfig(cfg.Log)
}

// getLoggerConfigPath 获取日志配置文件路径
//
// 	eg. export BLOTMQ_WEB_CONFIG = "/home/boltmq/console/conf/cfg.toml"
//
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/24
func getLoggerConfigPath() (cfgPath string) {
	cfgPath = strings.TrimSpace(os.Getenv(stgcommon.BLOTMQ_WEB_CONFIG_ENV))
	if file.IsExist(cfgPath) {
		return cfgPath
	}

	// 默认寻找当前目录的cfg.toml日志配置文件
	cfgPath = file.SelfDir() + string(os.PathSeparator) + cfgName
	if file.IsExist(cfgPath) {
		return cfgPath
	}

	// 此处为了兼容能够直接在idea上面利用web/g/默认配置文件目录
	// 当前“Config”结构体位于“github.com/boltmq/console/web/g”包,而目标toml文件位于“github.com/boltmq/console/conf”
	oldVal := strings.Split(reflect.TypeOf(cfg).PkgPath(), "/")
	newVal := oldVal
	if len(oldVal) > 2 {
		end := len(oldVal) - 2
		newVal = oldVal[0:end]
	}
	parentDir := strings.Join(newVal, "/")

	cfgPath = getConfigFullPath(parentDir)
	fmt.Printf("idea special cnosole config path = %s \n", cfgPath)
	return cfgPath
}

// parentDir 获取日志配置文件的完整目录
//
// eg.  github.com/boltmq/console  -->  E:/source/src/github.com/boltmq/console/conf/cfg.toml
//
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/24
func getConfigFullPath(parentDir string) string {
	//TODO: 可能存在“GOPATH”环境变量需要使用分号切割，获取第一个值的情况
	gopath := strings.TrimSpace(os.Getenv("GOPATH"))
	src := "/src/"
	return gopath + src + parentDir + "/" + cfgDir + "/" + cfgName
}

// GetConfig 获取配置
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/24
func GetConfig() *Config {
	return &cfg
}
