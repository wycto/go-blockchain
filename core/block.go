package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index         int64
	Timestamp     int64
	PrevBlockHash string
	Data          string
	Hash          string
}

//区块hash生成
func calculateHash(b Block) string {
	blockData := strconv.Itoa(int(b.Index)) + strconv.Itoa(int(b.Timestamp)) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

/**
创建新区块
*/
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

//初始化区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "起始块")
}
