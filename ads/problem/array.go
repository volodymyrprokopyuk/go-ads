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
