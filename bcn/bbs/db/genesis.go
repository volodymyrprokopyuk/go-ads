package db

import (
	"encoding/json"
	"os"
)

type genesis struct {
  Balances map[Account]uint `json:"balances"`
}

func loadGenesis(path string) (genesis, error) {
  cont, err := os.ReadFile(path)
  if err != nil {
    return genesis{}, err
  }
  var gen genesis
  err = json.Unmarshal(cont, &gen)
  if err != nil {
    return genesis{}, err
  }
  return gen, nil
}
