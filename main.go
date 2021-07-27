package main

import (
	"github.com/brady-wang/gee/gee"
	"net/http"
	"test3/block"
	"test3/blockChain"
)

var preNode *blockChain.Node
var header *blockChain.Node

var first block.Block
var preBlock block.Block

func main() {

	first = block.GenerateFirstBlock("创世区块")
	header = blockChain.CreateHeaderNode(&first)
	preBlock = first
	preNode = header
	r := gee.New()
	r.GET("/next", nextBlock)
	r.Run(":8080")
}

func nextBlock(c *gee.Context) {
	block := block.GenerateNextBlock("hello", preBlock)
	node := blockChain.AddNode(&block, preNode)
	preBlock = block
	preNode = node
	blockChain.ShowNodes(header)
	c.JSON(http.StatusOK, node)
}
