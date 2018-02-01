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

type QueryResolver struct{}

func (r *QueryResolver) Clusters(ctx context.Context, args struct{ Name *string }) []*clusterResolver {
	var crs []*clusterResolver

	clusters := mockQueryCluster(args.Name)
	for _, c := range clusters {
		crs = append(crs, &clusterResolver{c: c})
	}

	return crs
}

type clusterResolver struct {
	c *cluster
}

func (r *clusterResolver) Name(ctx context.Context) string {
	return r.c.name
}

func (r *clusterResolver) Stats(ctx context.Context) *clusterStatsResolver {
	//r.c.name
	return &clusterStatsResolver{}
}

type clusterStatsResolver struct {
	cs *clusterStats
}

func (r *clusterStatsResolver) ProducerNums(ctx context.Context) int32 {
	return 0
}

func (r *clusterStatsResolver) ConsumerNums(ctx context.Context) int32 {
	return 0
}

func (r *clusterStatsResolver) BrokerNums(ctx context.Context) int32 {
	return 0
}

func (r *clusterStatsResolver) NamesrvNums(ctx context.Context) int32 {
	return 0
}

func (r *clusterStatsResolver) TopicNums(ctx context.Context) int32 {
	return 0
}

func (r *clusterStatsResolver) OutTotalTodayNums(ctx context.Context) int32 {
	return 0
}

func (r *clusterStatsResolver) OutTotalYestNums(ctx context.Context) int32 {
	return 0
}

func (r *clusterStatsResolver) InTotalTodayNums(ctx context.Context) int32 {
	return 0
}

func (r *clusterStatsResolver) InTotalYestNums(ctx context.Context) int32 {
	return 0
}

/*
func (r *clusterResolver) Nodes() []*clusterNodeResolver {
	return nil
}

type clusterNodeResolver struct {
}
*/
