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

type cluster struct {
	name string
}

type clusterStats struct {
	producerNums      int32
	consumerNums      int32
	brokerNums        int32
	namesrvNums       int32
	topicNums         int32
	outTotalTodayNums int32
	outTotalYestNums  int32
	inTotalTodayNums  int32
	inTotalYestNums   int32
}

type clusterNode struct {
	namesrvAddrs []string
}

type brokerNode struct {
	role              int32
	addr              string
	version           string
	desc              string
	outTps            float64
	inTps             float64
	outTotalTodayNums int32
	outTotalYestNums  int32
	inTotalTodayNums  int32
	inTotalYestNums   int32
}

type topic struct {
	topic    string
	typ      TopicType
	isSystem bool
}

type TopicType int32

const (
	NORMAL_TOPIC TopicType = iota
	RETRY_TOPIC
	DLQ_TOPIC
)

type topicStore struct {
	brokerAddr     string
	brokerId       int32
	brokerName     string
	writeQueueNums int32
	readQueueNums  int32
	unit           bool
	order          bool
	perm           int32
}
