package main

import (
	"fmt"
	"os"

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
      err = state.Save()
      if err != nil {
        return err
      }
      fmt.Printf("Add TX %v: success", tx)
      return nil
    },
  }
  cmd.Flags().StringVarP(&from, "from", "f", "", "Debit account")
  cmd.Flags().StringVarP(&to, "to", "t", "", "Credit account")
  cmd.Flags().UintVarP(&value, "value", "v", 0, "Transfer value")
  cmd.Flags().StringVarP(&data, "data", "d", "", "TX metadata e.g. reward")
  return cmd
}

func main() {
  bbs := bbsCmd()
  bbs.AddCommand(balCmd(), txCmd())
  err := bbs.Execute()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
