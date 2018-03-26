package core

import "errors"

type Chain []*Block

func NewChain() (*Chain, error) {
  genBlock, _ := GenesisBlock()
  chain := &Chain{genBlock}
  return chain, nil
}

func (bc *Chain) AddBlock(data string) (*Block, error) {

  prevBlock := (*bc)[len(*bc)-1]
  newBlock, err := NewBlock(prevBlock, data)
  if err != nil {
    return nil, errors.New("failed to create new block")
  }
  if newBlock.Valid(prevBlock) == false {
    return nil, errors.New("invalid Block")
  }

  newChain := append(*bc, newBlock)
  bc.Replace(&newChain)

  return newBlock, nil
}

func (bc *Chain) Replace(newChain *Chain) {
  if len(*bc) < len(*newChain) {
    *bc = *newChain
  }
}
