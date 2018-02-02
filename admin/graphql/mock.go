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

var mockClusters = []*cluster{
	&cluster{name: "cluster1"},
	&cluster{name: "cluster2"},
}

func mockQueryCluster(name *string) []*cluster {
	var cs []*cluster

	for _, c := range mockClusters {
		if name == nil || *name == "" || *name == c.name {
			cs = append(cs, c)
		}
	}

	return cs
}

var mockClusterStats = map[string]*clusterStats{
	"cluster1": &clusterStats{
		producerNums:      100,
		consumerNums:      101,
		brokerNums:        102,
		namesrvNums:       103,
		topicNums:         104,
		outTotalTodayNums: 105,
		outTotalYestNums:  106,
		inTotalTodayNums:  107,
		inTotalYestNums:   108,
	},
	"cluster2": &clusterStats{
		producerNums:      200,
		consumerNums:      201,
		brokerNums:        202,
		namesrvNums:       203,
		topicNums:         204,
		outTotalTodayNums: 205,
		outTotalYestNums:  206,
		inTotalTodayNums:  207,
		inTotalYestNums:   208,
	},
}

func mockQueryClusterStats(name string) *clusterStats {
	for n, cs := range mockClusterStats {
		if name == n {
			return cs
		}
	}

	return nil
}

var mockClusterNodes = map[string]*clusterNode{
	"cluster1": &clusterNode{
		namesrvAddrs: []string{"10.1.0.1:9876", "10.1.0.1:9876"},
	},
	"cluster2": &clusterNode{
		namesrvAddrs: []string{"10.2.0.1:9876", "10.2.0.1:9876"},
	},
}

func mockQueryClusterNode(name string) *clusterNode {
	for n, cn := range mockClusterNodes {
		if name == n {
			return cn
		}
	}

	return nil
}

var mockBrokerNodes = map[string][]*brokerNode{
	"cluster1": []*brokerNode{
		&brokerNode{
			role:              0,
			addr:              "10.1.0.3",
			version:           "v1.0.0",
			desc:              "broker server",
			outTps:            100.3,
			inTps:             100.4,
			outTotalTodayNums: 105,
			outTotalYestNums:  106,
			inTotalTodayNums:  107,
			inTotalYestNums:   108,
		},
		&brokerNode{
			role:              1,
			addr:              "10.1.0.4",
			version:           "v1.0.0",
			desc:              "broker server",
			outTps:            200.3,
			inTps:             200.4,
			outTotalTodayNums: 205,
			outTotalYestNums:  206,
			inTotalTodayNums:  207,
			inTotalYestNums:   208,
		},
	},
	"cluster2": []*brokerNode{
		&brokerNode{
			role:              0,
			addr:              "10.2.0.3",
			version:           "v1.0.0",
			desc:              "broker server",
			outTps:            1000.3,
			inTps:             1000.4,
			outTotalTodayNums: 1005,
			outTotalYestNums:  1006,
			inTotalTodayNums:  1007,
			inTotalYestNums:   1008,
		},
		&brokerNode{
			role:              1,
			addr:              "10.2.0.4",
			version:           "v1.0.0",
			desc:              "broker server",
			outTps:            2000.3,
			inTps:             2000.4,
			outTotalTodayNums: 2005,
			outTotalYestNums:  2006,
			inTotalTodayNums:  2007,
			inTotalYestNums:   2008,
		},
	},
}

func mockQueryBrokerNodes(name string) []*brokerNode {
	for n, bns := range mockBrokerNodes {
		if name == n {
			return bns
		}
	}

	return []*brokerNode{}
}

var mockTopics = map[string][]*topic{
	"cluster1": []*topic{
		&topic{
			topic:    "topic1",
			typ:      NORMAL_TOPIC,
			isSystem: false,
		},
		&topic{
			topic:    "topic2",
			typ:      RETRY_TOPIC,
			isSystem: false,
		},
	},
	"cluster2": []*topic{
		&topic{
			topic:    "topic3",
			typ:      NORMAL_TOPIC,
			isSystem: false,
		},
		&topic{
			topic:    "topic4",
			typ:      DLQ_TOPIC,
			isSystem: true,
		},
	},
}

func mockQueryTopics(name string, like *string) []*topic {
	var ts []*topic
	for n, nts := range mockTopics {
		if name == n {
			for _, t := range nts {
				if like == nil || *like == "" || *like == t.topic {
					ts = append(ts, t)
				}
			}
			break
		}
	}

	return ts
}

var mockTopicsStore = map[string]map[string]*topicStore{
	"cluster1": map[string]*topicStore{
		"topic1": &topicStore{
			brokerAddr:     "10.1.0.1:11911",
			brokerId:       0,
			brokerName:     "broker1",
			writeQueueNums: 100,
			readQueueNums:  200,
			unit:           false,
			order:          false,
			perm:           1,
		},
		"topic2": &topicStore{
			brokerAddr:     "10.1.0.2:11911",
			brokerId:       1,
			brokerName:     "broker2",
			writeQueueNums: 101,
			readQueueNums:  201,
			unit:           true,
			order:          true,
			perm:           3,
		},
	},
	"cluster2": map[string]*topicStore{
		"topic3": &topicStore{
			brokerAddr:     "10.2.0.1:11911",
			brokerId:       0,
			brokerName:     "broker3",
			writeQueueNums: 100,
			readQueueNums:  200,
			unit:           false,
			order:          false,
			perm:           1,
		},
		"topic4": &topicStore{
			brokerAddr:     "10.2.0.2:11911",
			brokerId:       1,
			brokerName:     "broker4",
			writeQueueNums: 101,
			readQueueNums:  201,
			unit:           true,
			order:          true,
			perm:           3,
		},
	},
}

func mockQueryTopicsStore(name, topic string) *topicStore {
	for n, ts := range mockTopicsStore {
		if name == n {
			for tp, s := range ts {
				if topic == tp {
					return s
				}
			}
		}
	}

	return &topicStore{}
}
