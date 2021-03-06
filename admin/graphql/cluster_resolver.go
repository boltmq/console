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

	"github.com/go-errors/errors"
	graphql "github.com/neelance/graphql-go"
)

type clusterResolver struct {
	c *cluster
}

func (r *clusterResolver) Name(ctx context.Context) string {
	return r.c.name
}

func (r *clusterResolver) Stats(ctx context.Context) (*clusterStatsResolver, error) {
	cs := mockQueryClusterStats(r.c.name)
	if cs == nil {
		return nil, errors.Errorf("cluster=%s not found stats.", r.c.name)
	}

	return &clusterStatsResolver{cs: cs}, nil
}

func (r *clusterResolver) Nodes(ctx context.Context) (*clusterNodeResolver, error) {
	cn := mockQueryClusterNode(r.c.name)
	if cn == nil {
		return nil, errors.Errorf("cluster=%s not found node.", r.c.name)
	}

	return &clusterNodeResolver{cn: cn, name: r.c.name}, nil
}

func (r *clusterResolver) Topics(ctx context.Context, args struct {
	Like  *string
	First *int32
	After *graphql.ID
}) (*topicsConnectionResolver, error) {
	ts := mockQueryTopics(r.c.name, args.Like)
	// TODO: sort topics

	start, end, err := parsePageStart2End(len(ts), args.First, args.After)
	if err != nil {
		return nil, err
	}

	return &topicsConnectionResolver{name: r.c.name, ts: ts, start: start, end: end}, nil
}

type clusterStatsResolver struct {
	cs *clusterStats
}

func (r *clusterStatsResolver) ProducerNums(ctx context.Context) int32 {
	return r.cs.producerNums
}

func (r *clusterStatsResolver) ConsumerNums(ctx context.Context) int32 {
	return r.cs.consumerNums
}

func (r *clusterStatsResolver) BrokerNums(ctx context.Context) int32 {
	return r.cs.brokerNums
}

func (r *clusterStatsResolver) NamesrvNums(ctx context.Context) int32 {
	return r.cs.namesrvNums
}

func (r *clusterStatsResolver) TopicNums(ctx context.Context) int32 {
	return r.cs.topicNums
}

func (r *clusterStatsResolver) OutTotalTodayNums(ctx context.Context) int32 {
	return r.cs.outTotalTodayNums
}

func (r *clusterStatsResolver) OutTotalYestNums(ctx context.Context) int32 {
	return r.cs.outTotalYestNums
}

func (r *clusterStatsResolver) InTotalTodayNums(ctx context.Context) int32 {
	return r.cs.inTotalTodayNums
}

func (r *clusterStatsResolver) InTotalYestNums(ctx context.Context) int32 {
	return r.cs.inTotalYestNums
}

type clusterNodeResolver struct {
	name string
	cn   *clusterNode
}

func (r *clusterNodeResolver) NamesrvAddrs(ctx context.Context) []string {
	return r.cn.namesrvAddrs
}

func (r *clusterNodeResolver) BrokerNodes(ctx context.Context) ([]*brokerNodeResolver, error) {
	var bnr []*brokerNodeResolver
	bns := mockQueryBrokerNodes(r.name)
	for _, bn := range bns {
		bnr = append(bnr, &brokerNodeResolver{bn: bn})
	}

	return bnr, nil
}

type brokerNodeResolver struct {
	bn *brokerNode
}

func (r *brokerNodeResolver) Role(ctx context.Context) int32 {
	return r.bn.role
}

func (r *brokerNodeResolver) Addr(ctx context.Context) string {
	return r.bn.addr
}

func (r *brokerNodeResolver) Version(ctx context.Context) string {
	return r.bn.version
}

func (r *brokerNodeResolver) Desc(ctx context.Context) string {
	return r.bn.desc
}

func (r *brokerNodeResolver) OutTps(ctx context.Context) float64 {
	return r.bn.outTps
}

func (r *brokerNodeResolver) InTps(ctx context.Context) float64 {
	return r.bn.inTps
}

func (r *brokerNodeResolver) OutTotalTodayNums(ctx context.Context) int32 {
	return r.bn.outTotalTodayNums
}

func (r *brokerNodeResolver) OutTotalYestNums(ctx context.Context) int32 {
	return r.bn.outTotalYestNums
}

func (r *brokerNodeResolver) InTotalTodayNums(ctx context.Context) int32 {
	return r.bn.inTotalTodayNums
}

func (r *brokerNodeResolver) InTotalYestNums(ctx context.Context) int32 {
	return r.bn.inTotalYestNums
}
