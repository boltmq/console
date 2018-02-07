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
	auth bool
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

// Root set web root directory server, use store static sources file, etc .html .css .js
func (srv *Server) Root(pattern, webRoot, index string) *Server {
	if webRoot != "" {
		srv.mux.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(webRoot))))
		index = fmt.Sprintf("%s%s", pattern, index)
		if pattern != "/" {
			srv.mux.Handle("/", http.RedirectHandler(index, http.StatusFound))
		}
	}
	return srv
}

// LoadGraphQL load graphql to pattern.
func (srv *Server) LoadGraphQL(pattern, schemaString string, resolver interface{}, opts ...graphql.SchemaOpt) *Server {
	schema := graphql.MustParseSchema(schemaString, resolver, opts...)
	if srv.auth {
		srv.mux.Handle(pattern, join(&relay.Handler{Schema: schema}, &authenticator{}))
	} else {
		srv.mux.Handle(pattern, join(&relay.Handler{Schema: schema}))
	}
	return srv
}

// SetAuth set auth and set login url pattern.
func (srv *Server) SetAuth(open bool, pattern string) *Server {
	srv.auth = open
	srv.mux.Handle(pattern, &loginHandler{})
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
