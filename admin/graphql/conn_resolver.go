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
package graphql

import "context"

type consumeConnResolver struct {
	name  string
	topic string
}

func (r *consumeConnResolver) Describe(ctx context.Context) string {
	return ""
}

func (r *consumeConnResolver) Conns(ctx context.Context) ([]*connectionResolver, error) {
	var crs []*connectionResolver
	cs := mockQueryTopicsConsumeConns(r.name, r.topic)
	for _, c := range cs {
		crs = append(crs, &connectionResolver{c: c})
	}

	return crs, nil
}

type connectionResolver struct {
	c *connection
}

func (r *connectionResolver) ConsumeGroup(ctx context.Context) string {
	return r.c.consumeGroup
}

func (r *connectionResolver) ClientId(ctx context.Context) string {
	return r.c.clientId
}

func (r *connectionResolver) ClientAddr(ctx context.Context) string {
	return r.c.clientAddr
}

func (r *connectionResolver) Language(ctx context.Context) string {
	return r.c.language
}

func (r *connectionResolver) Version(ctx context.Context) string {
	return r.c.version
}

func (r *connectionResolver) ConsumeTps(ctx context.Context) float64 {
	return r.c.consumeTps
}

func (r *connectionResolver) ConsumeFromWhere(ctx context.Context) string {
	return r.c.consumeFromWhere
}

func (r *connectionResolver) ConsumeType(ctx context.Context) int32 {
	return int32(r.c.consumeType)
}

func (r *connectionResolver) Diff(ctx context.Context) int32 {
	return r.c.diff
}

func (r *connectionResolver) MessageModel(ctx context.Context) int32 {
	return int32(r.c.messageModel)
}
