package db

import (
	"encoding/json"
	"os"
)

type genesis struct {
  Balances map[Account]uint `json:"balances"`
}

func loadGenesis(path string) (genesis, error) {
  genCont, err := os.ReadFile(path)
  if err != nil {
    return genesis{}, err
  }
  var gen genesis
  err = json.Unmarshal(genCont, &gen)
  if err != nil {
    return genesis{}, err
  }
  return gen, nil
}
