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

type topicsConnectionResolver struct {
	ts    []*topic
	start int
	end   int
	name  string
}

func (r *topicsConnectionResolver) Total(ctx context.Context) int32 {
	return int32(len(r.ts))
}

func (r *topicsConnectionResolver) Edges(ctx context.Context) ([]*topicsEdgeResolver, error) {
	if r.end <= r.start {
		return []*topicsEdgeResolver{}, nil
	}

	l := make([]*topicsEdgeResolver, r.end-r.start)
	for i := range l {
		l[i] = &topicsEdgeResolver{
			cursor: encodeCursor(r.start + i),
			t:      r.ts[r.start+i],
			name:   r.name,
		}
	}

	return l, nil
}

func (r *topicsConnectionResolver) PageInfo(ctx context.Context) (*pageInfoResolver, error) {
	return &pageInfoResolver{
		startCursor: encodeCursor(r.start),
		endCursor:   encodeCursor(r.end - 1),
		hasNextPage: r.end < len(r.ts),
	}, nil
}

type topicsEdgeResolver struct {
	cursor graphql.ID
	t      *topic
	name   string
}

func (r *topicsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *topicsEdgeResolver) Node() *topicResolver {
	return &topicResolver{t: r.t, name: r.name}
}

type topicResolver struct {
	name string
	t    *topic
}

func (r *topicResolver) Topic(ctx context.Context) string {
	return r.t.topic
}

func (r *topicResolver) Type(ctx context.Context) int32 {
	return int32(r.t.typ)
}

func (r *topicResolver) IsSystem(ctx context.Context) bool {
	return r.t.isSystem
}

func (r *topicResolver) Store(ctx context.Context) (*topicStoreResolver, error) {
	ts := mockQueryTopicsStore(r.name, r.t.topic)
	return &topicStoreResolver{ts: ts}, nil
}

func (r *topicResolver) Route(ctx context.Context) (*topicRouteResolver, error) {
	tr := mockQueryTopicsRoute(r.name, r.t.topic)
	return &topicRouteResolver{tr: tr}, nil
}

func (r *topicResolver) Groups(ctx context.Context) []string {
	return mockQueryTopicsGroup(r.name, r.t.topic)
}

func (r *topicResolver) ConsumeClients(ctx context.Context, args struct {
	First *int32
	After *graphql.ID
}) (*consumeClientConnectionResolver, error) {
	ccs := mockQueryTopicsConsumeClientConns(r.name, r.t.topic)
	start, end, err := parsePageStart2End(len(ccs), args.First, args.After)
	if err != nil {
		return nil, err
	}

	return &consumeClientConnectionResolver{name: r.name, topic: r.t.topic, ccs: ccs, start: start, end: end}, nil
}

func (r *topicResolver) ConsumeProgresses(ctx context.Context, args struct {
	Group *string
	First *int32
	After *graphql.ID
}) (*consumeProgressConnectionResolver, error) {
	cpcs := mockQueryTopicsConsumeProgress(r.name, r.t.topic, args.Group)
	start, end, err := parsePageStart2End(len(cpcs), args.First, args.After)
	if err != nil {
		return nil, err
	}

	return &consumeProgressConnectionResolver{cpcs: cpcs, start: start, end: end}, nil
}

type topicStoreResolver struct {
	ts *topicStore
}

func (r *topicStoreResolver) BrokerName(ctx context.Context) string {
	return r.ts.brokerName
}

func (r *topicStoreResolver) QueueId(ctx context.Context) int32 {
	return r.ts.queueId
}

func (r *topicStoreResolver) LastUpdateTime(ctx context.Context) string {
	return r.ts.lastUpdateTime
}

func (r *topicStoreResolver) MaxOffset(ctx context.Context) int32 {
	return r.ts.maxOffset
}

func (r *topicStoreResolver) MinOffset(ctx context.Context) int32 {
	return r.ts.minOffset
}

type topicRouteResolver struct {
	tr *topicRoute
}

func (r *topicRouteResolver) Queues(ctx context.Context) ([]*queueDataResolver, error) {
	var qdrs []*queueDataResolver
	for _, d := range r.tr.queueDatas {
		qdrs = append(qdrs, &queueDataResolver{qd: d})
	}

	return qdrs, nil
}

func (r *topicRouteResolver) Brokers(ctx context.Context) ([]*brokerDataResolver, error) {
	var bdrs []*brokerDataResolver
	for _, d := range r.tr.brokerDatas {
		bdrs = append(bdrs, &brokerDataResolver{bd: d})
	}

	return bdrs, nil
}

type queueDataResolver struct {
	qd *queueData
}

func (r *queueDataResolver) BrokerName(ctx context.Context) string {
	return r.qd.brokerName
}

func (r *queueDataResolver) ReadQueueNums(ctx context.Context) int32 {
	return r.qd.readQueueNums
}

func (r *queueDataResolver) WriteQueueNums(ctx context.Context) int32 {
	return r.qd.writeQueueNums
}

func (r *queueDataResolver) Perm(ctx context.Context) int32 {
	return r.qd.perm
}

func (r *queueDataResolver) SysFlag(ctx context.Context) int32 {
	return r.qd.sysFlag
}

type brokerDataResolver struct {
	bd *brokerData
}

func (r *brokerDataResolver) BrokerName(ctx context.Context) string {
	return r.bd.brokerName
}

func (r *brokerDataResolver) BrokerAddrs(ctx context.Context) ([]*brokerAddrResolver, error) {
	var bars []*brokerAddrResolver
	for bi, ba := range r.bd.brokerAddrs {
		bars = append(bars, &brokerAddrResolver{bi: bi, ba: ba})
	}

	return bars, nil
}

type brokerAddrResolver struct {
	bi int32
	ba string
}

func (r *brokerAddrResolver) BrokerId(ctx context.Context) int32 {
	return r.bi
}

func (r *brokerAddrResolver) Addr(ctx context.Context) string {
	return r.ba
}

// mutation api
type topicInput struct {
	Topic          string
	WriteQueueNums int32
	ReadQueueNums  int32
	Unit           bool
	Order          bool
}

type responseResolver interface {
	Code(ctx context.Context) int32
	Desc(ctx context.Context) string
}

type topicResponseResolver struct {
	code int32
	desc string
}

func (r *topicResponseResolver) Code(ctx context.Context) int32 {
	return r.code
}

func (r *topicResponseResolver) Desc(ctx context.Context) string {
	return r.desc
}
