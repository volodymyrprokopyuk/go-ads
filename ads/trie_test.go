package ads_test

import (
	"testing"

	"github.com/volodymyrprokopyuk/go-ads/ads"
)

func TestTrie(t *testing.T) {
  cases := []struct{
    word string
    exp bool
  }{{"go", true}, {"goal", true}, {"goals", false}}
  trie := ads.NewTrie()
  trie.Set("go", "goal")
  for _, c := range cases {
    if trie.Get(c.word) != c.exp {
      t.Errorf("invalid get: word %v, expected %v, got %v", c.word, c.exp, !c.exp)
    }
  }
}
