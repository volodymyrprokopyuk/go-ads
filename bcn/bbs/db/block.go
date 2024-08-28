package db

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Hash [32]byte

func (h Hash) MarshalText() ([]byte, error) {
  return []byte(hex.EncodeToString(h[:])), nil
}

func (h *Hash) UnmarshalText(str []byte) error {
  _, err := hex.Decode(h[:], str)
  return err
}

type BlockHeader struct {
  Parent Hash `json:"parent"`
  Time uint64 `json:"time"`
}

type Block struct {
  Header BlockHeader `json:"header"`
  Txs []Tx `json:"payload"`
}

func NewBlock(parent Hash, time uint64, txs []Tx) Block {
  return Block{Header: BlockHeader{Parent: parent, Time: time}, Txs: txs}
}

func (b *Block) Hash() (Hash, error) {
  jsonBlock, err := json.Marshal(b)
  if err != nil {
    return Hash{}, err
  }
  return sha256.Sum256(jsonBlock), nil
}

type BlockFS struct {
  Key Hash `json:"hash"`
  Value Block `json:"block"`
}
