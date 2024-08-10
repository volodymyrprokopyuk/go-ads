package ads

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
