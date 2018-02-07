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
	"context"
	"net/http"
)

type Context struct {
	ctx context.Context
}

type httpChain interface {
	Chain(w http.ResponseWriter, r *http.Request, ctx *Context) bool
}

func join(next http.Handler, chains ...httpChain) http.Handler {
	if len(chains) == 0 {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{ctx: r.Context()}
		for _, chain := range chains {
			if ok := chain.Chain(w, r, ctx); !ok {
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx.ctx))
	})
}
