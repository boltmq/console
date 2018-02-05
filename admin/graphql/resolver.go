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
