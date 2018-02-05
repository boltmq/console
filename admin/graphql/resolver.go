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
)

type QueryResolver struct{}

func (r *QueryResolver) Clusters(ctx context.Context, args struct{ Name *string }) ([]*clusterResolver, error) {
	var crs []*clusterResolver

	clusters := mockQueryCluster(args.Name)
	for _, c := range clusters {
		crs = append(crs, &clusterResolver{c: c})
	}

	return crs, nil
}

func (r *QueryResolver) Msg(ctx context.Context, args struct {
	Name  *string
	MsgId string
}) (*messageResolver, error) {
	return &messageResolver{name: args.Name, msgId: args.MsgId}, nil
}

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

func (r *clusterResolver) Topics(ctx context.Context, args struct{ Like *string }) ([]*topicResolver, error) {
	var trs []*topicResolver

	ts := mockQueryTopics(r.c.name, args.Like)
	for _, t := range ts {
		trs = append(trs, &topicResolver{t: t, name: r.c.name})
	}

	return trs, nil
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

type consumeProgressResolver struct {
	cp *consumeProgress
}

func (r *consumeProgressResolver) ConsumeGroup(ctx context.Context) string {
	return r.cp.group
}

func (r *consumeProgressResolver) Tps(ctx context.Context) float64 {
	return r.cp.tps
}

func (r *consumeProgressResolver) Diff(ctx context.Context) int32 {
	return r.cp.diff
}

func (r *consumeProgressResolver) Total(ctx context.Context) int32 {
	return r.cp.total
}

func (r *consumeProgressResolver) Progress(ctx context.Context) ([]*consumeProgressDataResolver, error) {
	var cpdrs []*consumeProgressDataResolver
	for _, cpd := range r.cp.data {
		cpdrs = append(cpdrs, &consumeProgressDataResolver{cpd: cpd})
	}

	return cpdrs, nil
}

type consumeProgressDataResolver struct {
	cpd consumeProgressData
}

func (r *consumeProgressDataResolver) BrokerOffset(ctx context.Context) int32 {
	return int32(r.cpd.brokerOffset)
}

func (r *consumeProgressDataResolver) ConsumeOffset(ctx context.Context) int32 {
	return int32(r.cpd.consumeOffset)
}

func (r *consumeProgressDataResolver) Diff(ctx context.Context) int32 {
	return r.cpd.diff
}

func (r *consumeProgressDataResolver) BrokerName(ctx context.Context) string {
	return r.cpd.brokerName
}

func (r *consumeProgressDataResolver) QueueId(ctx context.Context) int32 {
	return r.cpd.queueId
}
