package db

import (
	"path/filepath"
)

const dataDir = "."

func dbDir() string {
  return filepath.Join(dataDir, "db")
}

func genPath() string {
  return filepath.Join(dbDir(), "genesis.json")
}

func blkPath() string {
  return filepath.Join(dbDir(), "block.db")
}
