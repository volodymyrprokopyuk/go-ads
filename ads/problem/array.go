package problem

import (
	"fmt"
	"slices"
	"strings"
)

// * Determine if a string has all unique characters. What if you cannot use
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
func UniqueCharsBitVec(str string) bool {
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

// * Given two strings, determine if one is a permutation of the other
// O(n) time, O(1) space
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

// O(n) time, O(1) space
func IsPermutationBitVec(a, b string) bool {
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

// O(n) time, O(1) space
func IsPermutationSort(a, b string) bool {
  if len(a) != len(b) {
    return false
  }
  aslc, bslc := strings.Split(a, ""), strings.Split(b, "")
  slices.Sort(aslc); slices.Sort(bslc)
  return slices.Equal(aslc, bslc)
}

// O(n) time, O(1) space
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

// * Replace all spaces in a string with %20
// O(n) time, O(n) space
func URLify(url string) string {
  slc := strings.Split(url, "")
  res := make([]string, len(slc))
  for _, c := range slc {
    if c == " " {
      res = append(res, "%20")
    } else {
      res = append(res, c)
    }
  }
  str := strings.Join(res, "")
  return str
}

// Check is a string is a permutation of a palindrome
// O(n) time, O(1) space
func IsPalindrPermut(str string) bool {
  cnts := make([]int, 128) // ASCII alphabet
  for _, c := range str {
    cnts[c]++
  }
  middle := false
  for _, cnt := range cnts {
    if cnt % 2 == 1 {
      if middle {
        return false // More than one middle odd chars
      }
      middle = true
    }
  }
  return true
}

// O(n) time, O(1) space
func IsPalindrPermutBitVec(str string) bool {
  var cnts int // Bit vector
  for _, c := range str {
    // Flips the char bit to track even or odd char occurrence
    cnts ^= (1 << (c - 'a')) // ASCII alphabet
  }
  return cnts & (cnts - 1) == 0 // As most one odd is allowed
}

// Check if two strings are one edit away: insert, delete, or replace
// O(n) time, O(1) space
func IsOneEditAway (a, b string) bool {
  if len(a) < len(b) {
    a, b = b, a // The a string is larger or equal to the b string
  }
  if len(a) - len(b) > 1 {
    return false
  }
  if len(a) == len(b) {
    replace := false
    for i := range len(a) {
      if a[i] != b[i] {
        if replace {
          return false // More than one replace
        }
        replace = true
      }
    }
    return true
  }
  delete := false
  for i, j := 0, 0; i < len(a) && j < len(b); {
    if a[i] != b[j] {
      if delete {
        return false // More than one delete
      }
      delete = true
      i++ // Advance only the deleted char from the larger string
      continue
    }
    i++; j++ // Advance both strings with equal chars
  }
  return true
}

// Compress a string using counts of repeated characters
// O(n) time, O(n) space
func StrCompress(str string) string {
  var bld strings.Builder
  var zero rune
  last, cnt := zero, 0
  for _, c := range str {
    if c == last {
      cnt++ // Increase the counter of the repeated char
      continue
    }
    if last != zero { // Write the last char with its counter
      bld.WriteRune(last)
      bld.WriteString(fmt.Sprintf("%d", cnt))
    }
    last, cnt = c, 1 // Reset the counter for a new char
  }
  if last != zero { // Write the last char with its counter
    bld.WriteRune(last)
    bld.WriteString(fmt.Sprintf("%d", cnt))
  }
  arc := bld.String()
  if len(arc) < len(str) {
    return arc
  }
  return str
}
