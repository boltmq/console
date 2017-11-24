package main

import (
	"flag"
	"fmt"
	"git.oschina.net/cloudzone/cloudcommon-go/web"
	"git.oschina.net/cloudzone/smartgo/stgweb/web/route"
	"github.com/boltmq/console/web/g"
	"os"
	"strings"
	//"git.oschina.net/cloudzone/smartgo/stgcommon"
)

const (
	_version = "v1.0.0"
)

func main() {

	//os.Setenv(stgcommon.NAMESRV_ADDR_ENV, "10.112.68.190:9876;10.112.68.192:9876")
	//os.Setenv(stgcommon.BLOTMQ_WEB_CONFIG_ENV, "E:/source/src/github.com/boltmq/console/conf/cfg.toml")

	v := flag.Bool("v", false, "boltmq console version")
	c := flag.String("c", "", "console logger config file")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *v {
		fmt.Println(_version)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	configPath := strings.TrimSpace(*c)
	g.InitLogger(configPath)

	web.New(_version).Config(&g.GetConfig().Web).Call(func(ctx *web.Context) {
		ctx.Super().Action = route.Route
	}).End().Run()
}
