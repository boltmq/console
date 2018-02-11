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
package server

import (
	"net/http"
)

type fileInspection struct {
	prefix string
	index  string
	root   http.FileSystem
}

func (fi *fileInspection) Chain(w http.ResponseWriter, r *http.Request, ctx *Context) bool {
	if r.URL.Path == "" || r.URL.Path == "/" {
		return true
	}

	f, e := fi.root.Open(r.URL.Path)
	if e == nil {
		f.Close()
		return true
	}

	f, e = fi.root.Open(fi.index)
	if e != nil {
		return true
	}
	defer f.Close()

	d, e := f.Stat()
	if e != nil {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return false
	}

	content := make([]byte, d.Size())
	f.Read(content)

	w.Write(content)
	return false
}
