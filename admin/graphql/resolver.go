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
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	graphql "github.com/neelance/graphql-go"
)

type Resolver struct{}

func (r *Resolver) Clusters(ctx context.Context, args struct{ Name *string }) ([]*clusterResolver, error) {
	var crs []*clusterResolver

	clusters := mockQueryCluster(args.Name)
	for _, c := range clusters {
		crs = append(crs, &clusterResolver{c: c})
	}

	return crs, nil
}

func (r *Resolver) Msg(ctx context.Context, args struct {
	Name  *string
	MsgId string
}) (*messageResolver, error) {
	return &messageResolver{name: args.Name, msgId: args.MsgId}, nil
}

func (r *Resolver) Create2UpdateTopic(ctx context.Context, args struct {
	Name  string
	Topic topicInput
}) (*topicResponseResolver, error) {
	return &topicResponseResolver{}, nil
}

func (r *Resolver) DeleteTopic(ctx context.Context, args struct {
	Name  string
	Topic string
}) (*topicResponseResolver, error) {
	return &topicResponseResolver{}, nil
}

type pageInfoResolver struct {
	startCursor graphql.ID
	endCursor   graphql.ID
	hasNextPage bool
}

func (r *pageInfoResolver) StartCursor() *graphql.ID {
	return &r.startCursor
}

func (r *pageInfoResolver) EndCursor() *graphql.ID {
	return &r.endCursor
}

func (r *pageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}

func parsePageStart2End(size int, first *int32, after *graphql.ID) (start, end int, err error) {
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(string(*after))
		if err != nil {
			return 0, 0, err
		}

		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return 0, 0, err
		}

		start = i
		if start > size {
			start = 0
		}
	}

	end = size
	if first != nil {
		end = start + int(*first)
		if end > size {
			end = size
		}
	}

	return
}

func encodeCursor(i int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1))))
}
