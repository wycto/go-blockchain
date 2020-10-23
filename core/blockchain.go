package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.ApendBlock(&genesisBlock)
	return &blockchain
}
func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}
func (bc *Blockchain) ApendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 || isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

func (bc *Blockchain) Print() {
	fmt.Println("开始打印")
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevBlockHash: %s\n", block.PrevBlockHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n\n", block.Data)
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}

	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
