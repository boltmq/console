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

var Schema = `
# boltmq contole graphql schema
schema {
	query: Query
}

# The query type, represents all of the entry points into our object graph
type Query {
    clusters(name: String): [Cluster]!
}

# A Cluster from the boltmq server
type Cluster {
	# The name of cluster
    name: String!
	# The stats info of cluster
	stats: ClusterStats!
	# The node info of cluster
    nodes: ClusterNode!
	# The topics of cluster
	topics(like: String): [Topic]!
}

# A ClusterStats info of boltmq cluster
type ClusterStats {
	# The producer nums of cluster
    producerNums: Int!
	# The consumer nums of cluster
    consumerNums: Int!
	# The broker nums of cluster
    brokerNums: Int!
	# The name server nums of cluster
    namesrvNums: Int!
	# The topic nums of cluster
    topicNums: Int!
	# The cluster consumer msg total number today
    outTotalTodayNums: Int!
	# The cluster consumer msg total number yest
    outTotalYestNums: Int!
	# The cluster producer msg total number yest
    inTotalTodayNums: Int!
	# The cluster producer msg total number today
    inTotalYestNums: Int!
}

# A Cluster node info of boltmq cluster
type ClusterNode {
	# The namesrv addr list fo cluster
    namesrvAddrs: [String!]!
	# The broker node list fo cluster
    brokerNodes: [BrokerNode]!
}

# A Boker node info of boltmq cluster
type BrokerNode {
	# The borker role
    role: Int!
	# The borker addr
    addr: String!
	# The borker server version
    version: String!
	# The borker server describe
    desc: String!
	# The borker server current out tps
    outTps: Float!
	# The borker server current in tps
    inTps: Float!
	# The cluster consumer msg total number today
    outTotalTodayNums: Int!
	# The cluster consumer msg total number yest
    outTotalYestNums: Int!
	# The cluster producer msg total number yest
    inTotalTodayNums: Int!
	# The cluster producer msg total number today
    inTotalYestNums: Int!
}

# A topic info of boltmq cluster
type Topic {
	# The topic name
    topic: String!
	# The topic type
    type: Int!
	# The topic type
    isSystem: Boolean!
	# The topic store
    store: TopicStore!
	# The topic route
    route: TopicRoute!
}

# topic type
enum TopicType {
    # normal topic
    NORMAL_TOPIC
    # retry topic
    RETRY_TOPIC
    # deadline queue topic
    DLQ_TOPIC
}

# A topic stroe info of boltmq cluster
type TopicStore {
	# The broker name
    brokerName: String!
	# The queue id
    queueId: Int!
	# The max offset
    maxOffset: Int!
	# The min offset
    minOffset: Int!
	# The last update time
	lastUpdateTime: String!
}

# A topic route info of boltmq cluster
type TopicRoute {
	# The route data of queue
    queues: [QueueData]!
	# The route data of broker
    brokers: [BrokerData]!
}

# A queue route data of topic
type QueueData {
	# The broker name
    brokerName: String!
	# The write queue nums
    writeQueueNums: Int!
	# The read queue nums
    readQueueNums: Int!
	# The permissions of topic on broker
    perm: Int!
	# The permissions of topic on broker
    sysFlag: Int!
}

# A broker route data of topic
type BrokerData {
	# The broker name
    brokerName: String!
	# The broker addrs
    brokerAddrs: [BrokerAddr]!
}

# A broker addr of topic route
type BrokerAddr {
	# The broker id
	brokerId: Int!
	# The broker addr
	addr: String!
}
`
