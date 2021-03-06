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

type consumeProgressConnectionResolver struct {
	cpcs  []*consumeProgress
	start int
	end   int
}

func (r *consumeProgressConnectionResolver) Total(ctx context.Context) int32 {
	return int32(len(r.cpcs))
}

func (r *consumeProgressConnectionResolver) Edges(ctx context.Context) ([]*consumeProgressEdgeResolver, error) {
	if r.end <= r.start {
		return []*consumeProgressEdgeResolver{}, nil
	}

	l := make([]*consumeProgressEdgeResolver, r.end-r.start)
	for i := range l {
		l[i] = &consumeProgressEdgeResolver{
			cursor: encodeCursor(r.start + i),
			cp:     r.cpcs[r.start+i],
		}
	}

	return l, nil
}

func (r *consumeProgressConnectionResolver) PageInfo(ctx context.Context) (*pageInfoResolver, error) {
	return &pageInfoResolver{
		startCursor: encodeCursor(r.start),
		endCursor:   encodeCursor(r.end - 1),
		hasNextPage: r.end < len(r.cpcs),
	}, nil
}

type consumeProgressEdgeResolver struct {
	cursor graphql.ID
	cp     *consumeProgress
}

func (r *consumeProgressEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *consumeProgressEdgeResolver) Node() *consumeProgressResolver {
	return &consumeProgressResolver{cp: r.cp}
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
