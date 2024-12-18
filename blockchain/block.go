package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func CreateGenesisBlock() *Block {
	block := &Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         "Genesis Block",
		PreviousHash: "000000",
		Hash:         "",
		Nonce:        0,
		Transaction:  []*Transaction{},
	}
	block.Hash = block.calculateHash()
	return block
}

func (bc *BlockChain) NewBlock(index int, data string, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Data:         data,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
		Nonce:        0,
		Transaction:  []*Transaction{},
	}

	fmt.Println("Mining new block...")
	block.ProofOfWork()
	bc.Blocks = append(bc.Blocks, block)
	wallet := block.NewWallet(bc)
	wallet.Value = bc.Token.Balance[block.Hash]
	return block
}

func (b *Block) calculateHash() string {
	data := fmt.Sprintf("%d%s%d%s%d", b.Index, b.PreviousHash, b.Nonce, b.Timestamp)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (bc *BlockChain) GetBlockByHash(hash string) (*Block, error) {
	block, exists := bc.BlockIndex[hash]
	if !exists {
		return nil, fmt.Errorf("no block found")
	}
	return block, nil
}

func (bc *BlockChain) GetBlockBalance(hash string) float64 {
	balance := bc.Token.Balance[hash]
	return balance
}
