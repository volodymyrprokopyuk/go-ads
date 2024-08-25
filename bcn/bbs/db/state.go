package db

import (
	"bufio"
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
)

type State struct {
  Balances map[Account]uint
  txMempool []Tx
  dbFile *os.File
}

// Loads the state of TXs from the initial genesis.json and the historical tx.db
func LoadState() (*State, error) {
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
  return state, nil
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

// Saves TXs from the memory pool to the DB file
func (s *State) Save() error {
  for _, tx := range slices.Clone(s.txMempool) {
    txJson, err := json.Marshal(tx)
    if err != nil {
      return err
    }
    _, err = s.dbFile.Write(append(txJson, '\n'))
    if err != nil {
      return err
    }
    s.txMempool = s.txMempool[1:]
  }
  return nil
}

func (s *State) Close() {
  s.dbFile.Close()
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
