# GraphQL API

**查询集群节点
```
query clustersInfo($name: String) {
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
        intTps
        outTotalTodayNums
        outTotalYestNums
        inTotalTodayNums
        inTotalYestNums
      }
    }
  }
}
```

# GraphQL Schema
```
```
