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
	brokerName     string
	queueId        int32
	maxOffset      int32
	minOffset      int32
	lastUpdateTime string
}

type topicRoute struct {
	queueDatas  []*queueData
	brokerDatas []*brokerData
}

type queueData struct {
	brokerName     string
	readQueueNums  int32
	writeQueueNums int32
	perm           int32
	sysFlag        int32
}

type brokerData struct {
	brokerName  string
	brokerAddrs map[int32]string
}

type consumeClient struct {
	consumeGroup     string
	clientId         string
	clientAddr       string
	language         string
	version          string
	consumeTps       float64
	consumeFromWhere string
	consumeType      ConsumeType
	diff             int32
	messageModel     MessageModel
}

type ConsumeType int32

const (
	CONSUME_ACTIVELY ConsumeType = iota
	CONSUME_PASSIVELY
)

type MessageModel int32

const (
	BROADCASTING MessageModel = iota
	CLUSTERING
)

type consumeProgress struct {
	group string
	tps   float64
	diff  int32
	total int32
	data  []consumeProgressData
}

type consumeProgressData struct {
	brokerOffset  int64
	consumeOffset int64
	diff          int32
	brokerName    string
	queueId       int32
}

type messageInfo struct {
	topic                     string
	flag                      int32
	body                      string
	queueId                   int32
	storeSize                 int32
	queueOffset               int64
	sysFlag                   int32
	bornTimestamp             string
	bornHost                  string
	storeTimestamp            string
	storeHost                 string
	msgId                     string
	commitLogOffset           int64
	bodyCRC                   int32
	reconsumeTimes            int32
	preparedTransactionOffset int64
	properties                map[string]string
}

type messageTrack struct {
	code         int32
	consumeGroup string
	trackType    TrackType
	desc         string
}

type TrackType int

const (
	SubscribedAndConsumed TrackType = iota
	SubscribedButFilterd
	SubscribedButPull
	SubscribedAndNotConsumeYet
	UnknowExeption
	NotSubscribedAndNotConsumed
	ConsumeGroupIdNotOnline
)
