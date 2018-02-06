// Copyright 2017 luoji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/boltmq/console/admin/graphql"
	"github.com/boltmq/console/admin/server"
	"github.com/juju/errors"
	daemon "github.com/sevlyar/go-daemon"
)

const (
	version = "1.0.0"
)

func main() {
	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	f := flag.Bool("f", false, "run front terminal")
	port := flag.Int("p", 8000, "listen port")
	pid := flag.String("pid", "console.pid", "pid file")
	root := flag.String("root", "", "web root")
	prefix := flag.String("perfix", "/", "web root prefix url")
	index := flag.String("index", "index.html", "default home url")
	debug := flag.Bool("debug", false, "debug model")

	flag.Parse()
	if *h {
		flag.Usage()
		os.Exit(0)
	}

	if *v {
		fmt.Println("version:", version)
		os.Exit(0)
	}

	if !*f {
		dctx, err := runDaemon(*pid)
		if err != nil {
			os.Exit(0)
		}
		defer dctx.Release()
	}

	fmt.Printf("console is running on port %d.\n", *port)
	fmt.Printf("Begin with Get      : http://localhost:%d\n", *port)
	server.New().Root(*prefix, *root, *index).
		LoadGraphQL("/api", graphql.Schema, &graphql.Resolver{}).
		Debug(*debug).Listen(*port).Run()
}

func runDaemon(pidfile string) (*daemon.Context, error) {
	cntxt := &daemon.Context{
		PidFileName: pidfile,
		PidFilePerm: 0644,
		LogFileName: "",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        nil,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		return nil, err
	}
	if d != nil {
		return nil, errors.Errorf("child process not nil.")
	}

	return cntxt, nil
}
