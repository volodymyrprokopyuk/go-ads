package ads

// O(n^2), in-place, stable, moving target
func BubbleSort[T any](slc []T, ord func(a, b T) bool) {
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
func InsertSort[T any](slc []T, ord func(a, b T) bool) {
  for i := 1; i < len(slc); i++ {
    for j := i; j > 0; j-- {
      if ord(slc[j], slc[j - 1]) {
        slc[j - 1], slc[j] = slc[j], slc[j - 1]
      }
    }
  }
}

// O(n*log(n)), in-place, non-stable, generalization of InsertSort
func ShellSort[T any](slc []T, ord func(a, b T) bool) {
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
func SelectSort[T any](slc []T, ord func(a, b T) bool) {
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
func QuickSort[T any](slc []T, ord func(a, b T) bool) {
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

func merge[T any](a, b []T, ord func(a, b T) bool) []T {
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
func MergeSort[T any](slc []T, ord func(a, b T) bool) []T {
  if len(slc) < 2 {
    return slc
  }
  m := len(slc) / 2
  return merge(MergeSort(slc[:m], ord), MergeSort(slc[m:], ord), ord)
}

// O(n*log(n)), copy, non-stable, no duplicates, O(n^2) on sorted array
func BSTSort[T any](slc []T, ord func(a, b T) bool) []T {
  tree := NewBSTree(func(val T) T { return val }, ord)
  tree.Set(slc...)
  res := make([]T, 0, len(slc))
  for _, nd := range tree.InOrder() {
    res = append(res, nd.Value())
  }
  return res
}

// O(n*log(n)), copy, non-stable
func HeapSort[T any](slc []T, ord func(a, b T) bool) []T {
  heap := NewHeap(len(slc), func(val T) T { return val }, ord)
  heap.Push(slc...)
  res := make([]T, 0, len(slc))
  for heap.Length() > 0 {
    val, _ := heap.Pop()
    res = append(res, val)
  }
  return res
}

// O(log(n)), binary search of an ordered slice
func BinarySearch[T any](slc []T, val T, cmp func(a, b T) int) int {
  a, b := 0, len(slc)
  for a < b {
    m := (a + b - 1) / 2
    switch cmp(val, slc[m]) {
    case -1:
      b = m
    case 1:
      a = m + 1
    default:
      return m
    }
  }
  return -1
}
