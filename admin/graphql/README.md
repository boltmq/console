# GraphQL API

**Query API**
```
query clusters($name: String, $like: String, $group: String, $msgId: String!, $first: Int, $after: ID) {
  clusters(name: $name) {
    name
    stats {
      producerNums
      consumerNums
      brokerNums
      namesrvNums
      topicNums
      outTotalTodayNums
      outTotalYestNums
      inTotalTodayNums
      inTotalYestNums
    }
    nodes {
      namesrvAddrs
      brokerNodes {
        role
        addr
        version
        desc
        outTps
        inTps
        outTotalTodayNums
        outTotalYestNums
        inTotalTodayNums
        inTotalYestNums
      }
    }
    topics(like: $like, first: $first, after: $after) {
      total
      pageInfo {
        startCursor
        endCursor
        hasNextPage
      }
      edges {
        cursor
        node {
          topic
          type
          isSystem
          store {
            brokerName
            queueId
            maxOffset
            minOffset
            lastUpdateTime
          }
          route {
            queues {
              brokerName
              writeQueueNums
              readQueueNums
              perm
              sysFlag
            }
            brokers {
              brokerName
              brokerAddrs {
                brokerId
                addr
              }
            }
          }
          groups
          consumeClients(first: $first, after: $after) {
            describe
            total
            edges {
              cursor
              node {
                consumeGroup
                clientId
                clientAddr
                language
                version
                consumeTps
                consumeFromWhere
                consumeType
                diff
                messageModel
              }
            }
            pageInfo {
              startCursor
              endCursor
              hasNextPage
            }
          }
          consumeProgresses(group: $group, first: $first, after: $after) {
            total
            edges {
              cursor
              node {
                consumeGroup
                tps
                diff
                total
                progress {
                  brokerOffset
                  consumeOffset
                  diff
                  brokerName
                  queueId
                }
              }
            }
            pageInfo {
              startCursor
              endCursor
              hasNextPage
            }
          }
        }
      }
    }
  }
  msg(msgId: $msgId) {
    info {
      msgId
      topic
      flag
      body
      queueId
      storeSize
      queueOffset
      sysFlag
      bornTimestamp
      bornHost
      storeTimestamp
      storeHost
      commitLogOffset
      bodyCRC
      reconsumeTimes
      preparedTransactionOffset
      properties {
        key
        val
      }
    }
    tracks {
      code
      type
      consumeGroup
      desc
    }
  }
}
```

**Mutation API**
```
# boltmq contole graphql schema
schema {
	query: Query
	mutation: Mutation
}

# The query type, represents all of the entry points into our object graph
type Query {
    clusters(name: String): [Cluster]!
	msg(name: String, msgId: String!): Message
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
	topics(like: String, first: Int, after: ID): TopicsConnection!
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

# A connection object for a cluster's topics
type TopicsConnection {
	# The total number of friends
	total: Int!
	# The edges for each of the cluster's topics.
	edges: [TopicsEdge]!
	# Information for paginating this connection
	pageInfo: PageInfo!
}

# A edge object for a cluster's topics
type TopicsEdge {
	# A cursor used for pagination
	cursor: ID!
	# The character represented by this edge
	node: Topic
}

# Information for paginating this connection
type PageInfo {
	startCursor: ID
	endCursor: ID
	hasNextPage: Boolean!
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
	# The consume group
    groups: [String!]!
	# The consume connection
    consumeClients(first: Int, after: ID): ConsumeClientConnection!
	consumeProgresses(group: String, first: Int, after: ID): ConsumeProgressConnection!
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

# consume client connection
type ConsumeClientConnection {
	# The total number of friends
	total: Int!
	# The describe
    describe: String!
	# The edges for each of the consume client's edge.
	edges: [ConsumeClientEdge]!
	# Information for paginating this connection
	pageInfo: PageInfo!
}

# A edge object for a client's edge
type ConsumeClientEdge {
	# A cursor used for pagination
	cursor: ID!
	# The character represented by this edge
	node: ConsumeClient
}

# consume client info
type ConsumeClient {
	# The consume group name
    consumeGroup: String!
	# The client id
    clientId: String!
	# The client addr
    clientAddr: String!
	# The language
    language: String!
	# The version
    version: String!
	# The consume tps
    consumeTps: Float!
	# The consume from where
    consumeFromWhere: String!
	# The consume type
    consumeType: Int!
	# The message diff total
    diff: Int!
	# The message model
    messageModel: Int!
}

# consume type
enum ConsumeType {
    # actively consume
	CONSUME_ACTIVELY
    # passively consume
	CONSUME_PASSIVELY
}

# message model
enum MessageModel {
    # broadcasting
	BROADCASTING 
    # clustering
	CLUSTERING
}


# consume progress connection
type ConsumeProgressConnection {
	# The total number of friends
	total: Int!
	# The edges for each of the consume progress's edge.
	edges: [ConsumeProgressEdge]!
	# Information for paginating this connection
	pageInfo: PageInfo!
}

# consume progress edge
type ConsumeProgressEdge {
	# A cursor used for pagination
	cursor: ID!
	# The character represented by this edge
	node: ConsumeProgress
}

# consume progress
type ConsumeProgress {
	# The consume group name
	consumeGroup: String!
	# The consume tps
	tps: Float!
	# The consume diff
	diff: Int!
	# The total
	total: Int!
	# The progress data list
	progress: [ConsumeProgressData]!
}

# consume progress data
type ConsumeProgressData {
	# The broker offset
	brokerOffset: Int!
	# The broker offset
	consumeOffset: Int!
	# The consume diff
	diff: Int!
	# The broker name
	brokerName: String!
	# The queue id
	queueId: Int!
}

# message
type Message {
	# The message base info
	info: MessageInfo!
	# The message track list
	tracks: [MessageTrack]!
}

# message info
type MessageInfo {
	# The message id
	msgId: String!
	# The topic name
	topic: String!
	# The message flag
	flag: Int!
	# The message body
	body: String!
	# The queue id
	queueId: Int!
	# The store size
	storeSize: Int!
	# The queue offset
	queueOffset: Int!
	# The message sys flag
	sysFlag: Int!
	# The born timestamp
	bornTimestamp: String!
	# The born host 
	bornHost: String!
	# The store timestamp
	storeTimestamp: String!
	# The store host 
	storeHost: String!
	# The commitlog offset
	commitLogOffset: Int!
	# The message body crc
	bodyCRC: Int!
	# The reconsume times
	reconsumeTimes: Int!
	# The reconsume times
	preparedTransactionOffset: Int!
	# The properties
	properties: [Property!]!
}

# property, replace map
type Property {
	key: String!
	val: String!
}

# message track
type MessageTrack {
	# The track code, 0: success, non-0: failed
	code: Int!
	# track type
	type: Int!
	# consume group name
	consumeGroup: String!
	# error describe
	desc: String!
}

# track type
enum TrackType {
	# subscribed and consumed
	SUBSCRIBEDANDCONSUMED
	# subscribed but filterd
	SUBSCRIBEDBUTFILTERD
	# subscribed but pull
	SUBSCRIBEDBUTPULL
	# subscribed and not consume yet
	SUBSCRIBEDBUTNOTCONSUMEYET
	# unknow exeption
	UNKNOWEXEPTION
	# not subscribed and not consumed
	NOTSUBSCRIBEDANDNOTCONSUMED
	# consume groupId not online
	CONSUMEGROUPIDNOTONLINE
}

# The mutation type, represents all updates we can make to our data
type Mutation {
	create2UpdateTopic(name: String!, topic: TopicInput!): TopicResponse
	deleteTopic(name: String!, topic: String!): TopicResponse
}

# The input object sent when cluster is creating a new topic
input TopicInput {
	# topic
	topic: String!
	# The read queue nums, optional
	readQueueNums: Int!
	# The write queue nums, optional
    writeQueueNums: Int!
	# The order topic, optional
	order: Boolean! 
	# The unit topic, optional
	unit: Boolean!
}

# Represents a topic for a cluster
interface Response {
	code: Int!
	desc: String!
}

type TopicResponse implements Response {
	code: Int!
	desc: String!
}
```
