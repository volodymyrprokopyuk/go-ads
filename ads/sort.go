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

// O(n*log(n)), in-place, non-stable
func QuickSort[T cmp.Ordered](slc []T, ord func(a, b T) bool) {
  partition := func(a, b int) int {
    i, p := a, b - 1 // pivot is the last element
    for j := i; j < p; j++ {
      if ord(slc[j], slc[p]) {
        slc[i], slc[j] = slc[j], slc[i]
        i++
      }
    }
    slc[i], slc[p] = slc[p], slc[i]
    return i
  }
  var sort func(a, b int) // declaration for recursive function expression
  sort = func(a, b int) {
    if b - a < 2 {
      return
    }
    p := partition(a, b)
    sort(a, p); sort(p + 1, b)
  }
  sort(0, len(slc))
}

func merge[T cmp.Ordered](a, b []T, ord func(a, b T) bool) []T {
  res := make([]T, 0, len(a) + len(b))
  i, j := 0, 0
  for i < len(a) && j < len(b) {
    if ord(a[i], b[j]) {
      res = append(res, a[i])
      i++
    } else {
      res = append(res, b[j])
      j++
    }
  }
  for i < len(a) {
    res = append(res, a[i])
    i++
  }
  for j < len(b) {
    res = append(res, b[j])
    j++
  }
  return res
}

// O(n*log(n)), copy, stable, external sorting in files
func MergeSort[T cmp.Ordered](slc []T, ord func(a, b T) bool) []T {
  if len(slc) < 2 {
    return slc
  }
  m := len(slc) / 2
  return merge(MergeSort(slc[:m], ord), MergeSort(slc[m:], ord), ord)
}
