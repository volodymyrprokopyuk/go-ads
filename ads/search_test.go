package ads_test

import (
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func TestBinarySearch(t *testing.T) {
  cases := []struct{
    name string
    slc []int
    val int
    exp int
  }{
    {"empty slice", []int{}, 9, -1},
    {"singleton slice found", []int{1}, 1, 0},
    {"singleton slice not found", []int{1}, 9, -1},
    {"duplicate elements found", []int{1, 2, 2, 3}, 2, 1},
    {"larger slice found first", []int{1, 2, 3, 4}, 1, 0},
    {"larger slice found middle", []int{1, 2, 3, 4}, 2, 1},
    {"larger slice found middle center", []int{1, 2, 3, 4, 5}, 3, 2},
    {"larger slice found last", []int{1, 2, 3, 4}, 4, 3},
    {"larger slice not found", []int{1, 2, 3, 4}, 9, -1},
  }
  for _, c := range cases {
    got := ads.BinarySearch(c.slc, c.val, cm)
    if got != c.exp {
      t.Errorf("%v: expected %v, got %v", c.name, c.exp, got)
    }
  }
}
