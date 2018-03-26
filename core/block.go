package core

import (
  "bytes"
  "crypto/sha256"
  "encoding/hex"
  "time"
)

type Block struct {
  Index     int64     `json:"index"`
  Data      string    `json:"data"`
  PrevHash  string    `json:"prevHash"`
  Hash      string    `json:"hash"`
  Timestamp time.Time `json:"timestamp"`
}

func NewBlock(prevBlock *Block, data string) (*Block, error) {
  newBlock := &Block{
    Index:     prevBlock.Index + 1,
    Data:      data,
    PrevHash:  prevBlock.Hash,
    Timestamp: time.Now(),
  }
  newBlock.Hash = newBlock.calculateHash()
  return newBlock, nil
}

func GenesisBlock() (*Block, error) {
  genesisBlock := &Block{
    Index:     0,
    Data:      "",
    PrevHash:  "",
    Timestamp: time.Now(),
  }
  genesisBlock.Hash = genesisBlock.calculateHash()
  return genesisBlock, nil
}

func (block *Block) String() string {
  var buffer bytes.Buffer
  buffer.WriteString(string(block.Index))
  buffer.WriteString(block.Timestamp.String())
  buffer.WriteString(block.Data)
  buffer.WriteString(block.PrevHash)
  return buffer.String()
}

func (block *Block) calculateHash() string {
  hash := sha256.Sum256([]byte(block.String()))
  return hex.EncodeToString(hash[:])
}

func (block *Block) Valid(prevBlock *Block) bool {

  if block.Index != prevBlock.Index+1 {
    return false
  }

  if block.PrevHash != prevBlock.Hash {
    return false
  }

  if block.calculateHash() != block.Hash {
    return false
  }

  return true
}
