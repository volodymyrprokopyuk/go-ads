package node

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/volodymyrprokopyuk/go-ads/bcn/bbs/db"
)

const addr = ":8080"

func writeError(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusInternalServerError)
  json.NewEncoder(w).Encode(
    struct{ Error string `json:"error"`}{Error: err.Error()},
  )
}

// func txAdd(state *db.State) http.HandlerFunc {
//   return func(w http.ResponseWriter, r *http.Request) {

//   }
// }

type BalList struct {
  Hash db.Hash `json:"hash"`
  Balances map[db.Account]uint `json:"balances"`
}

func balList(state *db.State) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    balLst := BalList{Hash: state.LastBlockHash(), Balances: state.Balances}
    jsnBalLst, err := json.Marshal(balLst)
    if err != nil {
      writeError(w, err)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsnBalLst)
  }
}

func Listen() error {
  fmt.Printf("listening on port %v\n", addr)
  state, err := db.LoadState()
  if err != nil {
    return nil
  }
  defer state.Close()
  mux := http.NewServeMux()
  // mux.HandleFunc("/tx/add", txAdd(state))
  mux.HandleFunc("/bal/list", balList(state))
  return http.ListenAndServe(addr, mux)
}
