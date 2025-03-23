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
      unique = problem.UniqueCharsBitVec(c.str)
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
      permut = problem.IsPermutationBitVec(c.a, c.b)
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

func TestURLify(t *testing.T) {
  cases := []struct{ url, escaped string }{
    {"abc", "abc"}, {"a b c", "a%20b%20c"},
  }
  for _, c := range cases {
    t.Run(c.url, func(t *testing.T) {
      t.Parallel()
      escaped := problem.URLify(c.url)
      if escaped != c.escaped {
        t.Errorf("URLify: %s expected, %s got", c.escaped, escaped)
      }
    })
  }
}

func TestIsPalindrPermut(t *testing.T) {
  cases := []struct{ str string; permut bool }{
    // Palindromes
    {"", true}, {"a", true}, {"aa", true}, {"aaa", true}, {"aabb", true},
    {"aba", true}, {"aabcc", true}, {"abccab", true},
    // Palindrome permutations
    {"abab", true}, {"aab", true}, {"abcac", true}, {"abcabc", true},
    {"abc", false}, {"aaab", false}, {"baaa", false},
    {"tactcoa", true}, {"tactcoapapa", true},
    {"tacxtcoa", false}, {"tactcxoapapa", false},
  }
  for _, c := range cases {
    t.Run(c.str, func(t *testing.T) {
      t.Parallel()
      permut := problem.IsPalindrPermut(c.str)
      if permut != c.permut {
        t.Errorf("IsPalindrPermut: expected %t, got %t", c.permut, permut)
      }
      permut = problem.IsPalindrPermutBitVec(c.str)
      if permut != c.permut {
        t.Errorf("IsPalindrPermutBitVec: expected %t, got %t", c.permut, permut)
      }
    })
  }
}
