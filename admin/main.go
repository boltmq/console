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
	"net/http"
	"os"
)

const (
	version = "1.0.0"
)

func main() {
	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	port := flag.Int("p", 8000, "listen port")
	webRoot := flag.String("root", "./sources", "web file root")
	durl := flag.String("def_url", "/sources/index.html", "default url redirect path")

	flag.Parse()
	if *h {
		flag.Usage()
		os.Exit(0)
	}

	if *v {
		fmt.Println("version:", version)
		os.Exit(0)
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *durl, http.StatusFound)
	}))
	http.Handle("/sources/", http.StripPrefix("/sources/", http.FileServer(http.Dir(*webRoot))))
	fmt.Printf("console is running on port %d.\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
