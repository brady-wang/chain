package block

import (
	"fmt"
	"github.com/brady-wang/go-tools/hashx"
	"strconv"
	"strings"
	"time"
)

const DIFF = 4

type Block struct {
	PreHash   string // 前一个的 hash 值
	MyHash    string // 我的 hash
	TimeStamp string // 时间戳
	Diff      int    // 控制 hash 有几个前导 0
	Data      string // 交易信息
	Index     int    // 区块高度
	Nonce     int    // 随机值
}

// 创世区块
func GenerateFirstBlock(data string) Block {
	var firstBlock Block
	firstBlock.PreHash = "0"
	firstBlock.TimeStamp = time.Now().String()
	firstBlock.Diff = DIFF
	firstBlock.Data = data
	firstBlock.Index = 1
	firstBlock.Nonce = 0
	firstBlock.MyHash = GenerateHash(firstBlock)
	return firstBlock
}

func GenerateNextBlock(data string, preBlock Block) Block {
	var newBlock Block
	newBlock.PreHash = preBlock.MyHash
	newBlock.TimeStamp = time.Now().String()
	newBlock.Diff = DIFF
	newBlock.Data = data
	newBlock.Index = preBlock.Index + 1
	newBlock.Nonce = 0
	newBlock.MyHash = Pow(newBlock.Diff, &newBlock)
	return newBlock
}

func Pow(diff int, block *Block) string {
	for {
		hash := GenerateHash(*block)
		fmt.Println("挖矿 " + hash)
		if strings.HasPrefix(hash, strings.Repeat("0", diff)) {
			fmt.Println("挖矿成功:" + hash)
			return hash
		} else {
			block.Nonce++
		}
	}
}

func GenerateHash(block Block) string {
	var hashData = strconv.Itoa(block.Nonce) +  strconv.Itoa(block.Index) + strconv.Itoa(block.Diff) + block.Data + block.TimeStamp + block.PreHash
	hash := hashx.Sha256(hashData)
	return hash
}