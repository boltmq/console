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

type messageResolver struct {
	name  *string
	msgId string
}

func (r *messageResolver) Info(ctx context.Context) (*messageInfoResolver, error) {
	msgInfo := mockQueryMsgsInfo(r.name, r.msgId)
	return &messageInfoResolver{msg: msgInfo}, nil
}

func (r *messageResolver) Tracks(ctx context.Context) ([]*messageTrackResolver, error) {
	var tracks []*messageTrackResolver
	mtrs := mockQueryMsgTracks(r.name, r.msgId)
	for _, track := range mtrs {
		tracks = append(tracks, &messageTrackResolver{track: track})
	}

	return tracks, nil
}

type messageInfoResolver struct {
	msg *messageInfo
}

func (r *messageInfoResolver) Topic(ctx context.Context) string {
	return r.msg.topic
}

func (r *messageInfoResolver) Flag(ctx context.Context) int32 {
	return r.msg.flag
}

func (r *messageInfoResolver) Body(ctx context.Context) string {
	return r.msg.body
}

func (r *messageInfoResolver) QueueId(ctx context.Context) int32 {
	return r.msg.queueId
}

func (r *messageInfoResolver) StoreSize(ctx context.Context) int32 {
	return r.msg.storeSize
}

func (r *messageInfoResolver) QueueOffset(ctx context.Context) int32 {
	return int32(r.msg.queueOffset)
}

func (r *messageInfoResolver) SysFlag(ctx context.Context) int32 {
	return r.msg.sysFlag
}

func (r *messageInfoResolver) BornTimestamp(ctx context.Context) string {
	return r.msg.bornTimestamp
}

func (r *messageInfoResolver) BornHost(ctx context.Context) string {
	return r.msg.bornHost
}

func (r *messageInfoResolver) StoreTimestamp(ctx context.Context) string {
	return r.msg.storeTimestamp
}

func (r *messageInfoResolver) StoreHost(ctx context.Context) string {
	return r.msg.storeHost
}

func (r *messageInfoResolver) MsgId(ctx context.Context) string {
	return r.msg.msgId
}

func (r *messageInfoResolver) CommitLogOffset(ctx context.Context) int32 {
	return int32(r.msg.commitLogOffset)
}

func (r *messageInfoResolver) BodyCRC(ctx context.Context) int32 {
	return r.msg.bodyCRC
}

func (r *messageInfoResolver) ReconsumeTimes(ctx context.Context) int32 {
	return r.msg.reconsumeTimes
}

func (r *messageInfoResolver) PreparedTransactionOffset(ctx context.Context) int32 {
	return int32(r.msg.preparedTransactionOffset)
}

func (r *messageInfoResolver) Properties(ctx context.Context) ([]*propertyResolver, error) {
	var pyr []*propertyResolver

	for k, v := range r.msg.properties {
		pyr = append(pyr, &propertyResolver{key: k, val: v})
	}

	return pyr, nil
}

type propertyResolver struct {
	key string
	val string
}

func (r *propertyResolver) Key(ctx context.Context) string {
	return r.key
}

func (r *propertyResolver) Val(ctx context.Context) string {
	return r.val
}

type messageTrackResolver struct {
	track *messageTrack
}

func (r *messageTrackResolver) Code(ctx context.Context) int32 {
	return r.track.code
}

func (r *messageTrackResolver) ConsumeGroup(ctx context.Context) string {
	return r.track.consumeGroup
}

func (r *messageTrackResolver) Type(ctx context.Context) int32 {
	return int32(r.track.trackType)
}

func (r *messageTrackResolver) Desc(ctx context.Context) string {
	return r.track.desc
}
