// 基于redis
package uuid

import (
	"github.com/bwmarrin/snowflake"
)

var _node *snowflake.Node

// 唯一键
type UUID struct{}

// Init 初始化
func Init(dataCenterID DataCenterIDT, nodeID int) {
	var err error
	node := getNode(dataCenterID, nodeID)
	snowflake.NodeBits = 7          // 7位
	snowflake.Epoch = 1727712000000 // 2024-10-01 00:00:00
	_node, err = snowflake.NewNode(node)
	if err != nil {
		panic(err)
	}
}

// Int64 uuid整数格式
func (uuid UUID) Int64() int64 {
	return _node.Generate().Int64()
}

// String uuid字符串格式
func (uuid UUID) String() string {
	return _node.Generate().String()
}

// Time 当前的时间
func (uuid UUID) Time() int64 {
	return (_node.Generate().Int64()/4096/1024 + snowflake.Epoch) / 1000
}

// DataCenter 当前的数据中心
func (uuid UUID) DataCenter() int64 {
	return _node.Generate().Int64() % (4096 * 1024) >> 12
}

// DataCenterIDT 数据中心ID
type DataCenterIDT int

// getNode 获取节点
// dataCenterID 数据中心id
func getNode(dataCenterID DataCenterIDT, nodeID int) int64 {
	if dataCenterID >= 10 {
		panic("data center id can not be great than 100")
	}
	if nodeID >= 10 {
		panic("node id can not be great than 10")
	}
	return int64(int(dataCenterID)*10 + nodeID)
}
