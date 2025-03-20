package problem_test

import (
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads/problem"
)

func TestUniqueChars(t *testing.T) {
  cases := []struct { str string; unique bool }{
    {"", true}, {"abc", true}, {"abbc", false}, {"abcb", false},
  }
  for _, c := range cases {
    t.Run(c.str, func(t *testing.T) {
      t.Parallel()
      unique := problem.UniqueChars(c.str)
      if unique != c.unique {
        t.Errorf("UniqueChars: expected %t, got %t", c.unique, unique)
      }
      unique = problem.UniqueCharsBitVector(c.str)
      if unique != c.unique {
        t.Errorf("UniqueCharsBitVector: expected %t, got %t", c.unique, unique)
      }
      unique = problem.UniqueCharsSort(c.str)
      if unique != c.unique {
        t.Errorf("UniqueCharsSort: expected %t, got %t", c.unique, unique)
      }
    })
  }
}
