package ads_test

import (
	"strconv"
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

type Int int

func (i Int) String() string {
  return strconv.Itoa(int(i))
}

func TestHTable(t *testing.T) {
  htb := ads.NewHTable[Int, string](101)
  htb.Set(Int(66), ">66")
  htb.Set(Int(98), ">98")
  htb.Set(Int(98), ">>98")
  gotLength, expLength := htb.Length(), 2
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
  exp := ">66"
  got, exist := htb.Get(Int(66))
  if !exist || got != exp {
    t.Errorf("invalid get: expected %v, got %v", exp, got)
  }
  exp = ">>98"
  got, exist = htb.Get(Int(98))
  if !exist || got != exp {
    t.Errorf("invalid get updated: expected %v, got %v", exp, got)
  }
  key := 99
  _, exist = htb.Get(Int(key))
  if exist {
    t.Errorf("exists non-existing key %v", key)
  }
  gotLength = 0
  for _, _ = range htb.Entries() {
    gotLength++
    if gotLength == 2 {
      break // test early exit from the iterator
    }
  }
  if gotLength != expLength {
    t.Errorf("invalid entries: expected %v, got %v", expLength, gotLength)
  }
  exp = ">66"
  got, deleted := htb.Delete(Int(66))
  if !deleted || got != exp {
    t.Errorf("invalid delete: expected %v, got %v", exp, got)
  }
  exp = ">>98"
  got, deleted = htb.Delete(Int(98))
  if !deleted || got != exp {
    t.Errorf("invalid delete updated: expected %v, got %v", exp, got)
  }
  _, deleted = htb.Delete(Int(key))
  if deleted {
    t.Errorf("deleted non-existing key %v", key)
  }
  gotLength, expLength = htb.Length(), 0
  if gotLength != expLength {
    t.Errorf("invalid length: expected %v, got %v", expLength, gotLength)
  }
}
