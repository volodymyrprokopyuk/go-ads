package db

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"time"
)

type Snapshot [32]byte

type State struct {
  Balances map[Account]uint
  txMempool []Tx
  dbFile *os.File
  snapshot Snapshot
  lastBlockHash Hash
}

// Loads the state of TXs from the initial genesis.json and the historical tx.db
func LoadStateSnapshot() (*State, error) {
  path := filepath.Join("db", "genesis.json")
  gen, err := loadGenesis(path)
  if err != nil {
    return nil, err
  }
  path = filepath.Join("db", "tx.db")
  dbFile, err := os.OpenFile(path, os.O_APPEND | os.O_RDWR, 600)
  if err != nil {
    return nil, err
  }
  state := &State{
    Balances: maps.Clone(gen.Balances), txMempool: make([]Tx, 0), dbFile: dbFile,
  }
  scn := bufio.NewScanner(dbFile)
  for scn.Scan() {
    err := scn.Err()
    if err != nil {
      return nil, err
    }
    var tx Tx
    err = json.Unmarshal(scn.Bytes(), &tx)
    if err != nil {
      return nil, err
    }
    err = state.apply(tx)
    if err != nil {
      return nil, err
    }
  }
  err = state.takeSnapshot()
  if err != nil {
    return nil, err
  }
  return state, nil
}

// Loads the state of TXs from the initial genesis.json and the historical tx.db
func LoadState() (*State, error) {
  gen, err := loadGenesis(genPath())
  if err != nil {
    return nil, err
  }
  dbFile, err := os.OpenFile(blkPath(), os.O_APPEND | os.O_RDWR, 600)
  if err != nil {
    return nil, err
  }
  state := &State{
    Balances: maps.Clone(gen.Balances),
    txMempool: make([]Tx, 0),
    dbFile: dbFile,
    lastBlockHash: Hash{},
  }
  scn := bufio.NewScanner(dbFile)
  for scn.Scan() {
    err := scn.Err()
    if err != nil {
      return nil, err
    }
    var blockFS BlockFS
    err = json.Unmarshal(scn.Bytes(), &blockFS)
    if err != nil {
      return nil, err
    }
    err = state.applyBlock(blockFS.Value)
    if err != nil {
      return nil, err
    }
    state.lastBlockHash = blockFS.Key
  }
  return state, nil
}

func (s *State) Close() {
  s.dbFile.Close()
}

func (s *State) Snapshot() Snapshot {
  return s.snapshot
}

func (s *State) LastBlockHash() Hash {
  return s.lastBlockHash
}

// Adds a TX to the state balances and the memory pool
func (s *State) Add(tx Tx) error {
  err := s.apply(tx)
  if err != nil {
    return err
  }
  s.txMempool = append(s.txMempool, tx)
  return nil
}

// Adds all TXs in a block to the state balances and the memory pool
func (s *State) AddBlock(blk Block) error {
  for _, tx := range blk.Txs {
    err := s.Add(tx)
    if err != nil {
      return err
    }
  }
  return nil
}

// Saves TXs from the memory pool to the DB file
func (s *State) SaveSnapshot() (Snapshot, error) {
  for _, tx := range slices.Clone(s.txMempool) {
    jsonTx, err := json.Marshal(tx)
    if err != nil {
      return Snapshot{}, err
    }
    _, err = s.dbFile.Write(append(jsonTx, '\n'))
    if err != nil {
      return Snapshot{}, err
    }
    err = s.takeSnapshot()
    if err != nil {
      return Snapshot{}, err
    }
    s.txMempool = s.txMempool[1:]
  }
  return s.snapshot, nil
}

// Saves a block of TXs from the memory pool to the DB file
func (s *State) Save() error {
  blk := NewBlock(s.lastBlockHash, uint64(time.Now().Unix()), s.txMempool)
  blkHash, err := blk.Hash()
  if err != nil {
    return err
  }
  blkFS := BlockFS{Key: blkHash, Value: blk}
  jsonBlkFS, err := json.Marshal(blkFS)
  if err != nil {
    return err
  }
  _, err = s.dbFile.Write(append(jsonBlkFS, '\n'))
  if err != nil {
    return err
  }
  fmt.Printf("Save block: %x\n", blkHash)
  s.lastBlockHash = blkHash
  s.txMempool = make([]Tx, 0)
  return nil
}

// Validates a TX. Handles the reward. Transfers the TX value between balances
func (s *State) apply(tx Tx) error {
  if tx.IsReward() {
    s.Balances[tx.To] += tx.Value
    return nil
  }
  if s.Balances[tx.From] < tx.Value {
    return fmt.Errorf("account %v: insufficient funds < %v", tx.From, tx.Value)
  }
  s.Balances[tx.From] -= tx.Value
  s.Balances[tx.To] += tx.Value
  return nil
}

// Applies all TXs from a block
func (s *State) applyBlock(blk Block) error {
  for _, tx := range blk.Txs {
    err := s.apply(tx)
    if err != nil {
      return err
    }
  }
  return nil
}

// takes a snapshot of the whole DB file
func (s *State) takeSnapshot() error {
  _, err := s.dbFile.Seek(0, 0)
  if err != nil {
    return err
  }
  txCont, err := io.ReadAll(s.dbFile)
  if err != nil {
    return err
  }
  s.snapshot = sha256.Sum256(txCont)
  return nil
}
