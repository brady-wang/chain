package blockChain

import (
	"fmt"
	"test3/block"
)

type Node struct {
	nextNode *Node
	Data     *block.Block
}

// 创建头结点 保存创世区块
func CreateHeaderNode(data *block.Block) *Node {
	headerNode := new(Node)
	headerNode.nextNode = nil
	headerNode.Data = data
	return headerNode
}

// 挖矿成功 添加节点
func AddNode(data *block.Block, preNode *Node) *Node {
	var newNode = new(Node)
	newNode.Data = data
	newNode.nextNode = nil
	preNode.nextNode = newNode
	return newNode
}

func ShowNodes(node *Node) {
	n := node
	for {
		if n.nextNode == nil {
			fmt.Printf("%#v\n", n.Data)
			break
		} else {
			fmt.Printf("%#v\n", n.Data)
			n = n.nextNode
		}
	}
}
