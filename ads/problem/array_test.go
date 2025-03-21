package problem_test

import (
	"fmt"
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

func TestIsPermutation(t *testing.T) {
  cases := []struct{ a, b string; permut bool }{
    {"", "", true}, {"ab", "ba", true}, {"abc", "cab", true},
    {"a", "b", false}, {"abc", "bcd", false},
  }
  for _, c := range cases {
    t.Run(fmt.Sprintf("%s %s", c.a, c.b), func(t *testing.T) {
      permut := problem.IsPermutation(c.a, c.b)
      if permut != c.permut {
        t.Errorf("IsPermutation: expected %t, got %t", c.permut, permut)
      }
      permut = problem.IsPermutationBitVector(c.a, c.b)
      if permut != c.permut {
        t.Errorf("IsPermutationBitVector: expected %t, got %t", c.permut, permut)
      }
      permut = problem.IsPermutationSort(c.a, c.b)
      if permut != c.permut {
        t.Errorf("IsPermutationSort: expected %t, got %t", c.permut, permut)
      }
      permut = problem.IsPermutationCounts(c.a, c.b)
      if permut != c.permut {
        t.Errorf("IsPermutationCounts: expected %t, got %t", c.permut, permut)
      }
    })
  }
}
