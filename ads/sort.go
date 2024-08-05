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
