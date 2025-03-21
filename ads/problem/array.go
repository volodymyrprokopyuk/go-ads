package problem

import (
	"slices"
	"strings"
)

// Determine if a string has all unique characters. What if you cannot use
// additional data structures?
// O(n) time, O(1) space
func UniqueChars(str string) bool {
  seen := make([]bool, 128) // ASCII alphabet
  // seen := make(map[rune]bool) // Unicode alphabet
  for _, c := range str {
    if seen[c] {
      return false
    }
    seen[c] = true
  }
  return true
}

// O(n) time, O(1) space
func UniqueCharsBitVector(str string) bool {
  var seen int // Bit vector
  for _, c := range str {
    pos := 1 << (c - 'a') // a-z alphabet
    if seen & pos > 0 {
      return false
    }
    seen |= pos
  }
  return true
}

// O(n*log(n)) time, O(1) space
func UniqueCharsSort(str string) bool {
  slc := strings.Split(str, "")
  slices.Sort(slc)
  var last string
  for _, c := range slc {
    if last == c {
      return false
    }
    last = c
  }
  return true
}

// Given two strings, determine if one is a permutation of the other
// O(m + n) time, O(1) space
func IsPermutation(a, b string) bool {
  if len(a) != len(b) {
    return false
  }
  aseen, bseen := make([]bool, 128), make([]bool, 128) // ASCII alphabet
  for _, c := range a {
    aseen[c] = true
  }
  for _, c := range b {
    bseen[c] = true
  }
  return slices.Equal(aseen, bseen)
}

// O(m + n) time, O(1) space
func IsPermutationBitVector(a, b string) bool {
  if len(a) != len(b) {
    return false
  }
  var avec, bvec int
  for _, c := range a {
    avec |= (1 << (c - 'a')) // ASCII alphabet
  }
  for _, c := range b {
    bvec |= (1 << (c - 'a'))
  }
  return avec == bvec
}

func IsPermutationSort(a, b string) bool {
  if len(a) != len(b) {
    return false
  }
  aslc, bslc := strings.Split(a, ""), strings.Split(b, "")
  slices.Sort(aslc); slices.Sort(bslc)
  return slices.Equal(aslc, bslc)
}

func IsPermutationCounts(a, b string) bool {
  if len(a) != len(b) {
    return false
  }
  acnt, bcnt := make([]int, 128), make([]int, 128) // ASCII alphabet
  for _, c := range a {
    acnt[c]++
  }
  for _, c := range b {
    bcnt[c]++
  }
  return slices.Equal(acnt, bcnt)
}
