package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/volodymyrprokopyuk/go-ads/bcn/bbs/db"
)

func bbsCmd() *cobra.Command {
  return &cobra.Command{
    Use: "bbs",
    Short: "Build Blockchain from Scratch CLI",
    Long: "bbs is the entry point to manage the BBS blockchain",
    Version: "0.1.0",
    SilenceUsage: true,
    SilenceErrors: true,
  }
}

func balCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use: "bal",
    Short: "Manage balances on the BBS blockchain",
  }
  cmd.AddCommand(balListCmd())
  return cmd
}

func balListCmd() *cobra.Command {
  return &cobra.Command{
    Use: "list",
    Short: "List all balances",
    RunE: func(_ *cobra.Command, args []string) error {
      state, err := db.LoadState()
      if err != nil {
        return err
      }
      defer state.Close()
      // fmt.Printf("DB snapshot: %x\n", state.Snapshot())
      for acc, bal := range state.Balances {
        fmt.Printf("%v: %v\n", acc, bal)
      }
      return nil
    },
  }
}

func txCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use: "tx",
    Short: "Manage transactions on the BBS blockchain",
  }
  cmd.AddCommand(txAddCmd())
  return cmd
}

func txAddCmd() *cobra.Command {
  var from, to, data string
  var value uint
  cmd := &cobra.Command{
    Use: "add",
    Short: "Add a new TX to the DB",
    RunE: func(_ *cobra.Command, args []string) error {
      tx := db.Tx{
        From: db.Account(from), To: db.Account(to), Value: value, Data: data,
      }
      state, err := db.LoadState()
      if err != nil {
        return err
      }
      defer state.Close()
      err = state.Add(tx)
      if err != nil {
        return err
      }
      snapshot, err := state.SaveSnapshot()
      if err != nil {
        return err
      }
      fmt.Printf("Add TX %v: success\n", tx)
      fmt.Printf("DB snapshot: %x\n", snapshot)
      return nil
    },
  }
  cmd.Flags().StringVarP(&from, "from", "f", "", "Debit account")
  cmd.Flags().StringVarP(&to, "to", "t", "", "Credit account")
  cmd.Flags().UintVarP(&value, "value", "v", 0, "Transfer value")
  cmd.Flags().StringVarP(&data, "data", "d", "", "TX metadata e.g. reward")
  return cmd
}

func packTxsIntoBlocks() error {
  state, err := db.LoadState()
  if err != nil {
    return err
  }
  defer state.Close()
  blk0 := db.NewBlock(
    db.Hash{},
    uint64(time.Now().Unix()),
    []db.Tx{
      {"andrej", "andrej", 3, ""},
      {"andrej", "andrej", 700, "reward"},
    },
  )
  err = state.AddBlock(blk0)
  if err != nil {
    return err
  }
  err = state.Save()
  if err != nil {
    return err
  }
  blk1 := db.NewBlock(
    state.LastBlockHash(),
    uint64(time.Now().Unix()),
    []db.Tx{
      {"andrej", "babayaga", 2000, ""},
      {"andrej", "andrej", 100, "reward"},
      {"babayaga", "andrej", 1, ""},
      {"babayaga", "caesar", 1000, ""},
      {"babayaga", "andrej", 50, ""},
      {"andrej", "andrej", 600, "reward"},
    },
  )
  err = state.AddBlock(blk1)
  if err != nil {
    return err
  }
  err = state.Save()
  if err != nil {
    return err
  }
  return nil
}

func main() {
  err := packTxsIntoBlocks()

  // bbs := bbsCmd()
  // bbs.AddCommand(balCmd(), txCmd())
  // err := bbs.Execute()

  // err := node.Listen()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
