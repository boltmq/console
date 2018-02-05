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

func (r *topicResolver) ConsumeConn(ctx context.Context) (*consumeConnResolver, error) {
	return &consumeConnResolver{name: r.name, topic: r.t.topic}, nil
}

func (r *topicResolver) ConsumeProgress(ctx context.Context, args struct{ Group *string }) ([]*consumeProgressResolver, error) {
	var cprs []*consumeProgressResolver

	cps := mockQueryTopicsConsumeProgress(r.name, r.t.topic, args.Group)
	for _, cp := range cps {
		cprs = append(cprs, &consumeProgressResolver{cp: cp})
	}

	return cprs, nil
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
