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

import (
	"context"

	graphql "github.com/neelance/graphql-go"
)

type consumeClientConnectionResolver struct {
	name  string
	topic string
	ccs   []*consumeClient
	start int
	end   int
}

func (r *consumeClientConnectionResolver) Total(ctx context.Context) int32 {
	return int32(len(r.ccs))
}

func (r *consumeClientConnectionResolver) Describe(ctx context.Context) string {
	return ""
}

func (r *consumeClientConnectionResolver) Edges(ctx context.Context) ([]*consumeClientEdgeResolver, error) {
	if r.end <= r.start {
		return []*consumeClientEdgeResolver{}, nil
	}

	l := make([]*consumeClientEdgeResolver, r.end-r.start)
	for i := range l {
		l[i] = &consumeClientEdgeResolver{
			cursor: encodeCursor(r.start + i),
			cc:     r.ccs[r.start+i],
		}
	}

	return l, nil
}

func (r *consumeClientConnectionResolver) PageInfo(ctx context.Context) (*pageInfoResolver, error) {
	return &pageInfoResolver{
		startCursor: encodeCursor(r.start),
		endCursor:   encodeCursor(r.end - 1),
		hasNextPage: r.end < len(r.ccs),
	}, nil
}

type consumeClientEdgeResolver struct {
	cursor graphql.ID
	cc     *consumeClient
}

func (r *consumeClientEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *consumeClientEdgeResolver) Node() *consumeClientResolver {
	return &consumeClientResolver{c: r.cc}
}

type consumeClientResolver struct {
	c *consumeClient
}

func (r *consumeClientResolver) ConsumeGroup(ctx context.Context) string {
	return r.c.consumeGroup
}

func (r *consumeClientResolver) ClientId(ctx context.Context) string {
	return r.c.clientId
}

func (r *consumeClientResolver) ClientAddr(ctx context.Context) string {
	return r.c.clientAddr
}

func (r *consumeClientResolver) Language(ctx context.Context) string {
	return r.c.language
}

func (r *consumeClientResolver) Version(ctx context.Context) string {
	return r.c.version
}

func (r *consumeClientResolver) ConsumeTps(ctx context.Context) float64 {
	return r.c.consumeTps
}

func (r *consumeClientResolver) ConsumeFromWhere(ctx context.Context) string {
	return r.c.consumeFromWhere
}

func (r *consumeClientResolver) ConsumeType(ctx context.Context) int32 {
	return int32(r.c.consumeType)
}

func (r *consumeClientResolver) Diff(ctx context.Context) int32 {
	return r.c.diff
}

func (r *consumeClientResolver) MessageModel(ctx context.Context) int32 {
	return int32(r.c.messageModel)
}
