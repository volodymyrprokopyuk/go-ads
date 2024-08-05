package ads

import "cmp"

// O(n^2), in-place, stable, moving target
func BubbleSort[T cmp.Ordered](slc []T, ord func(a, b T) bool) {
  for i := len(slc); i > 1; i-- {
    swp := false
    for j := 1; j < i; j++ {
      if ord(slc[j], slc[j - 1]) {
        slc[j - 1], slc[j] = slc[j], slc[j - 1]
        swp = true
      }
    }
    if !swp {
      break
    }
  }
}

// O(n^2), in-place, stable, fixed target, O(n) on sorted array
func InsertSort[T cmp.Ordered](slc []T, ord func(a, b T) bool) {
  for i := 1; i < len(slc); i++ {
    for j := i; j > 0; j-- {
      if ord(slc[j], slc[j - 1]) {
        slc[j - 1], slc[j] = slc[j], slc[j - 1]
      }
    }
  }
}

// O(n*log(n)), in-place, non-stable, generalization of InsertSort
func ShellSort[T cmp.Ordered](slc []T, ord func(a, b T) bool) {
  gaps := []int{31, 15, 7, 3, 1}
  for _, gap := range gaps {
    for i := gap; i < len(slc); i++ {
      for j := i; j > gap - 1; j-- {
        if ord(slc[j], slc[j - gap]) {
          slc[j - gap], slc[j] = slc[j], slc[j - gap]
        }
      }
    }
  }
}

// O(n^2), in-place, non-stable, fixed target, O(n^2) on sorted array
func SelectSort[T cmp.Ordered](slc []T, ord func(a, b T) bool) {
  for i := 0; i < len(slc) - 1; i++ {
    m, j := i, i + 1
    for ; j < len(slc); j++ {
      if ord(slc[j], slc[m]) {
        m = j
      }
    }
    if m != i {
      slc[i], slc[m] = slc[m], slc[i]
    }
  }
}
