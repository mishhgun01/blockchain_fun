package pkg

import (
	"crypto/sha1"
	"encoding/hex"
)

type Chain struct {
	Root *Block
}

type Block struct {
	Value           string
	TransactionHash string
	Next, Previous  *Block
}

func New() *Chain {
	var b Chain
	b.Root = &Block{}
	b.Root.Next = b.Root
	b.Root.Previous = b.Root
	return &b
}

func MakeBlock(val interface{}, elem *Block) *Block {
	var b Block
	h := sha1.New()
	switch val.(type) {
	case string:
		value := val.(string)
		b.Value = value + elem.TransactionHash
		h.Write([]byte(b.Value))
		b.TransactionHash = hex.EncodeToString(h.Sum(nil))
	}
	return &b
}

func (c *Chain) Add(b Block) *Block {
	b.Next = c.Root
	b.Previous = c.Root.Previous
	c.Root.Previous.Next = &b
	c.Root.Previous = &b
	b.Previous.Next = &b
	return &b
}
