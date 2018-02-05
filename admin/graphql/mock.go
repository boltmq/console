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
			brokerName:     "broker1",
			queueId:        1,
			maxOffset:      100,
			minOffset:      10,
			lastUpdateTime: "2018-1-20 10:10",
		},
		"topic2": &topicStore{
			brokerName:     "broker2",
			queueId:        2,
			maxOffset:      200,
			minOffset:      20,
			lastUpdateTime: "2018-1-20 10:10",
		},
	},
	"cluster2": map[string]*topicStore{
		"topic3": &topicStore{
			brokerName:     "broker3",
			queueId:        1,
			maxOffset:      300,
			minOffset:      30,
			lastUpdateTime: "2018-1-20 10:10",
		},
		"topic4": &topicStore{
			brokerName:     "broker4",
			queueId:        1,
			maxOffset:      400,
			minOffset:      40,
			lastUpdateTime: "2018-1-20 10:10",
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

var mockTopicsRoute = map[string]map[string]*topicRoute{
	"cluster1": map[string]*topicRoute{
		"topic1": &topicRoute{
			queueDatas: []*queueData{
				&queueData{
					brokerName:     "broker1",
					readQueueNums:  10,
					writeQueueNums: 20,
					perm:           1,
					sysFlag:        2,
				},
				&queueData{
					brokerName:     "broker2",
					readQueueNums:  10,
					writeQueueNums: 20,
					perm:           1,
					sysFlag:        2,
				},
			},
			brokerDatas: []*brokerData{
				&brokerData{
					brokerName: "broker1",
					brokerAddrs: map[int32]string{
						0: "10.1.0.1:11911",
						1: "10.1.0.2:11911",
					},
				},
				&brokerData{
					brokerName: "broker2",
					brokerAddrs: map[int32]string{
						0: "10.2.0.1:11911",
						1: "10.2.0.2:11911",
					},
				},
			},
		},
		"topic2": &topicRoute{
			queueDatas: []*queueData{
				&queueData{
					brokerName:     "broker1",
					readQueueNums:  100,
					writeQueueNums: 200,
					perm:           1,
					sysFlag:        2,
				},
				&queueData{
					brokerName:     "broker2",
					readQueueNums:  100,
					writeQueueNums: 200,
					perm:           1,
					sysFlag:        2,
				},
			},
			brokerDatas: []*brokerData{
				&brokerData{
					brokerName: "broker1",
					brokerAddrs: map[int32]string{
						0: "10.3.0.1:11911",
						1: "10.3.0.2:11911",
					},
				},
				&brokerData{
					brokerName: "broker2",
					brokerAddrs: map[int32]string{
						0: "10.4.0.1:11911",
						1: "10.4.0.2:11911",
					},
				},
			},
		},
	},
	"cluster2": map[string]*topicRoute{
		"topic3": &topicRoute{
			queueDatas: []*queueData{
				&queueData{
					brokerName:     "broker1",
					readQueueNums:  10,
					writeQueueNums: 20,
					perm:           1,
					sysFlag:        2,
				},
				&queueData{
					brokerName:     "broker2",
					readQueueNums:  10,
					writeQueueNums: 20,
					perm:           1,
					sysFlag:        2,
				},
			},
			brokerDatas: []*brokerData{
				&brokerData{
					brokerName: "broker3",
					brokerAddrs: map[int32]string{
						0: "10.4.0.1:11911",
						1: "10.4.0.2:11911",
					},
				},
				&brokerData{
					brokerName: "broker4",
					brokerAddrs: map[int32]string{
						0: "10.5.0.1:11911",
						1: "10.5.0.2:11911",
					},
				},
			},
		},
		"topic4": &topicRoute{
			queueDatas: []*queueData{
				&queueData{
					brokerName:     "broker1",
					readQueueNums:  100,
					writeQueueNums: 200,
					perm:           1,
					sysFlag:        2,
				},
				&queueData{
					brokerName:     "broker2",
					readQueueNums:  100,
					writeQueueNums: 200,
					perm:           1,
					sysFlag:        2,
				},
			},
			brokerDatas: []*brokerData{
				&brokerData{
					brokerName: "broker5",
					brokerAddrs: map[int32]string{
						0: "10.6.0.1:11911",
						1: "10.6.0.2:11911",
					},
				},
				&brokerData{
					brokerName: "broker4",
					brokerAddrs: map[int32]string{
						0: "10.7.0.1:11911",
						1: "10.7.0.2:11911",
					},
				},
			},
		},
	},
}

func mockQueryTopicsRoute(name, topic string) *topicRoute {
	for n, tr := range mockTopicsRoute {
		if name == n {
			for tp, r := range tr {
				if topic == tp {
					return r
				}
			}
		}
	}

	return &topicRoute{}
}

var mockTopicsGroup = map[string]map[string][]string{
	"cluster1": map[string][]string{
		"topic1": []string{
			"subscription-group",
			"subscription-group2",
		},
		"topic2": []string{
			"subscription-group3",
			"subscription-group4",
		},
	},
	"cluster2": map[string][]string{
		"topic1": []string{
			"subscription-group5",
			"subscription-group6",
		},
		"topic2": []string{
			"subscription-group7",
			"subscription-group8",
		},
	},
}

func mockQueryTopicsGroup(name, topic string) []string {
	for n, tg := range mockTopicsGroup {
		if name == n {
			for tp, g := range tg {
				if topic == tp {
					return g
				}
			}
		}
	}

	return []string{}
}

var mockTopicsConsumeConns = map[string]map[string][]*connection{
	"cluster1": map[string][]*connection{
		"topic1": []*connection{
			&connection{
				consumeGroup:     "consume-group-1",
				clientId:         "client1",
				clientAddr:       "10.1.100.1",
				language:         "Golang",
				version:          "1.0.0",
				consumeTps:       100,
				consumeFromWhere: "cd",
				consumeType:      CONSUME_ACTIVELY,
				diff:             200,
				messageModel:     BROADCASTING,
			},
		},
		"topic2": []*connection{
			&connection{
				consumeGroup:     "consume-group-2",
				clientId:         "client2",
				clientAddr:       "10.2.100.1",
				language:         "Golang",
				version:          "1.0.0",
				consumeTps:       200,
				consumeFromWhere: "cd",
				consumeType:      CONSUME_ACTIVELY,
				diff:             400,
				messageModel:     BROADCASTING,
			},
		},
	},
	"cluster2": map[string][]*connection{
		"topic3": []*connection{
			&connection{
				consumeGroup:     "consume-group-3",
				clientId:         "client1",
				clientAddr:       "10.3.100.1",
				language:         "Golang",
				version:          "1.0.0",
				consumeTps:       300,
				consumeFromWhere: "cd",
				consumeType:      CONSUME_ACTIVELY,
				diff:             600,
				messageModel:     BROADCASTING,
			},
		},
		"topic4": []*connection{
			&connection{
				consumeGroup:     "consume-group-4",
				clientId:         "client4",
				clientAddr:       "10.4.100.1",
				language:         "Golang",
				version:          "1.0.0",
				consumeTps:       400,
				consumeFromWhere: "cd",
				consumeType:      CONSUME_ACTIVELY,
				diff:             800,
				messageModel:     BROADCASTING,
			},
		},
	},
}

func mockQueryTopicsConsumeConns(name, topic string) []*connection {
	for n, tg := range mockTopicsConsumeConns {
		if name == n {
			for tp, g := range tg {
				if topic == tp {
					return g
				}
			}
		}
	}

	return []*connection{}
}

var mockTopicsConsumeProgress = map[string]map[string][]*consumeProgress{
	"cluster1": map[string][]*consumeProgress{
		"topic1": []*consumeProgress{
			&consumeProgress{
				group: "consume-group-1",
				tps:   100,
				diff:  10,
				total: 1,
				data: []consumeProgressData{
					consumeProgressData{
						brokerOffset:  100,
						consumeOffset: 100,
						diff:          1,
						brokerName:    "broker1",
						queueId:       1,
					},
				},
			},
			&consumeProgress{
				group: "consume-group-2",
				tps:   200,
				diff:  20,
				total: 1,
				data: []consumeProgressData{
					consumeProgressData{
						brokerOffset:  200,
						consumeOffset: 200,
						diff:          2,
						brokerName:    "broker2",
						queueId:       2,
					},
				},
			},
		},
	},
	"cluster2": map[string][]*consumeProgress{
		"topic2": []*consumeProgress{
			&consumeProgress{
				group: "consume-group-3",
				tps:   300,
				diff:  30,
				total: 1,
				data: []consumeProgressData{
					consumeProgressData{
						brokerOffset:  300,
						consumeOffset: 300,
						diff:          1,
						brokerName:    "broker3",
						queueId:       3,
					},
				},
			},
			&consumeProgress{
				group: "consume-group-4",
				tps:   400,
				diff:  40,
				total: 1,
				data: []consumeProgressData{
					consumeProgressData{
						brokerOffset:  400,
						consumeOffset: 400,
						diff:          1,
						brokerName:    "broker4",
						queueId:       4,
					},
				},
			},
		},
	},
}

func mockQueryTopicsConsumeProgress(name, topic string, group *string) []*consumeProgress {
	cps := []*consumeProgress{}

	for n, tcps := range mockTopicsConsumeProgress {
		if name == n {
			for tp, scps := range tcps {
				if topic == tp {
					if group == nil || *group == "" {
						return scps
					}

					for _, cp := range scps {
						if cp.group == *group {
							cps = append(cps, cp)
						}
					}

					break
				}
			}
		}
	}

	return cps
}

var mockMsgsInfo = map[string]map[string]*messageInfo{
	"cluster1": map[string]*messageInfo{
		"msgid1abcdefghajklmnopqrstuvwxyz": &messageInfo{
			topic:                     "topic1",
			flag:                      1,
			body:                      "content body 1",
			queueId:                   1,
			storeSize:                 10,
			queueOffset:               10,
			sysFlag:                   1,
			bornTimestamp:             "2018-2-4 10:30:10",
			bornHost:                  "192.168.0.10",
			storeTimestamp:            "2018-2-4 10:30:20",
			storeHost:                 "10.1.0.2:11911",
			msgId:                     "msgid1abcdefghajklmnopqrstuvwxyz",
			commitLogOffset:           10,
			bodyCRC:                   1,
			reconsumeTimes:            0,
			preparedTransactionOffset: 0,
			properties:                map[string]string{"p1": "v1", "p2": "v2"},
		},
		"msgid2abcdefghajklmnopqrstuvwxyz": &messageInfo{
			topic:                     "topic2",
			flag:                      2,
			body:                      "content body 2",
			queueId:                   2,
			storeSize:                 20,
			queueOffset:               20,
			sysFlag:                   2,
			bornTimestamp:             "2018-2-4 10:30:10",
			bornHost:                  "192.168.0.11",
			storeTimestamp:            "2018-2-4 10:30:20",
			storeHost:                 "10.1.0.3:11911",
			msgId:                     "msgid1abcdefghajklmnopqrstuvwxyz",
			commitLogOffset:           20,
			bodyCRC:                   2,
			reconsumeTimes:            0,
			preparedTransactionOffset: 0,
			properties:                map[string]string{"pp1": "vv1", "pp2": "vv2"},
		},
	},
	"cluster2": map[string]*messageInfo{
		"msgid3abcdefghajklmnopqrstuvwxyz": &messageInfo{
			topic:                     "topic3",
			flag:                      3,
			body:                      "content body 3",
			queueId:                   3,
			storeSize:                 30,
			queueOffset:               30,
			sysFlag:                   3,
			bornTimestamp:             "2018-2-4 10:30:10",
			bornHost:                  "192.168.0.13",
			storeTimestamp:            "2018-2-4 10:30:20",
			storeHost:                 "10.1.0.4:11911",
			msgId:                     "msgid3abcdefghajklmnopqrstuvwxyz",
			commitLogOffset:           30,
			bodyCRC:                   3,
			reconsumeTimes:            0,
			preparedTransactionOffset: 0,
			properties:                map[string]string{"p3": "v3", "p4": "v4"},
		},
		"msgid4abcdefghajklmnopqrstuvwxyz": &messageInfo{
			topic:                     "topic4",
			flag:                      4,
			body:                      "content body 4",
			queueId:                   4,
			storeSize:                 40,
			queueOffset:               40,
			sysFlag:                   4,
			bornTimestamp:             "2018-2-4 10:30:10",
			bornHost:                  "192.168.0.14",
			storeTimestamp:            "2018-2-4 10:30:20",
			storeHost:                 "10.1.0.4:11911",
			msgId:                     "msgid4abcdefghajklmnopqrstuvwxyz",
			commitLogOffset:           40,
			bodyCRC:                   2,
			reconsumeTimes:            0,
			preparedTransactionOffset: 0,
			properties:                map[string]string{"pp3": "vv3", "pp4": "vv4"},
		},
	},
}

func mockQueryMsgsInfo(name *string, msgId string) *messageInfo {
	for n, amsgs := range mockMsgsInfo {
		if name == nil || *name == "" {
			for mid, msg := range amsgs {
				if mid == msgId {
					return msg
				}
			}
		} else {
			if *name == n {
				for mid, msg := range amsgs {
					if mid == msgId {
						return msg
					}
				}
			}
		}
	}

	return &messageInfo{}
}

var mockMsgTracks = map[string]map[string][]*messageTrack{
	"cluster1": map[string][]*messageTrack{
		"msgid1abcdefghajklmnopqrstuvwxyz": []*messageTrack{
			&messageTrack{
				code:         1,
				consumeGroup: "consume group 1",
				trackType:    SubscribedAndConsumed,
				desc:         "",
			},
			&messageTrack{
				code:         1,
				consumeGroup: "consume group 2",
				trackType:    SubscribedAndConsumed,
				desc:         "",
			},
		},
		"msgid2abcdefghajklmnopqrstuvwxyz": []*messageTrack{
			&messageTrack{
				code:         1,
				consumeGroup: "consume group 3",
				trackType:    SubscribedAndConsumed,
				desc:         "",
			},
			&messageTrack{
				code:         1,
				consumeGroup: "consume group 4",
				trackType:    SubscribedAndConsumed,
				desc:         "",
			},
		},
	},
	"cluster2": map[string][]*messageTrack{
		"msgid3abcdefghajklmnopqrstuvwxyz": []*messageTrack{
			&messageTrack{
				code:         1,
				consumeGroup: "consume group 5",
				trackType:    SubscribedAndConsumed,
				desc:         "",
			},
			&messageTrack{
				code:         1,
				consumeGroup: "consume group 6",
				trackType:    SubscribedAndConsumed,
				desc:         "",
			},
		},
		"msgid4abcdefghajklmnopqrstuvwxyz": []*messageTrack{
			&messageTrack{
				code:         1,
				consumeGroup: "consume group 7",
				trackType:    SubscribedAndConsumed,
				desc:         "",
			},
			&messageTrack{
				code:         1,
				consumeGroup: "consume group 8",
				trackType:    SubscribedAndConsumed,
				desc:         "",
			},
		},
	},
}

func mockQueryMsgTracks(name *string, msgId string) []*messageTrack {
	for _, tracks := range mockMsgTracks {
		for mid, ts := range tracks {
			if mid == msgId {
				return ts
			}
		}
	}

	return []*messageTrack{}
}
