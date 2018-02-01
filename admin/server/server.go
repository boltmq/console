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
	"fmt"
	"net/http"

	graphql "github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

type Server struct {
	port int
	mux  *http.ServeMux
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

// Index set directory index from server
func (srv *Server) Index(url string) *Server {
	srv.mux.Handle("/", http.RedirectHandler(url, http.StatusFound))
	return srv
}

// Root set web root directory server, use store static sources file, etc .html .css .js
func (srv *Server) Root(pattern, webRoot string) *Server {
	srv.mux.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(webRoot))))
	return srv
}

// LoadGraphQL load graphql to pattern.
func (srv *Server) LoadGraphQL(pattern, schemaString string, resolver interface{}, opts ...graphql.SchemaOpt) *Server {
	schema := graphql.MustParseSchema(schemaString, resolver, opts...)
	srv.mux.Handle(pattern, &relay.Handler{Schema: schema})
	return srv
}

// Debug enable debug model，use development evn.
func (srv *Server) Debug(open bool) *Server {
	if open {
		srv.mux.Handle("/debug", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(graphiqlPage)
		}))
	}
	return srv
}

// Listen set listen prot
func (srv *Server) Listen(port int) *Server {
	srv.port = port
	return srv
}

// Run startup server
func (srv *Server) Run() {
	http.ListenAndServe(fmt.Sprintf(":%d", srv.port), srv.mux)
}
