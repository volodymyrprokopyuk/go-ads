package pattern

import "fmt"

func PatDecorator() {
  type fun func(val int) int // decorated function type
  inc := func(val int) int { // decorated function
    return val + 1
  }
  logAround := func(f fun) fun { // decorator
    return func(val int) int {
      fmt.Printf("before %d\n", val)
      res := f(val)
      fmt.Printf("after %d\n", res)
      return res
    }
  }
  logAroundInc := logAround(inc)
  res := logAroundInc(1)
  fmt.Println(res) // before 1, after 2, 2
}
