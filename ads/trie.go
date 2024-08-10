package ads

type Trie struct {
  htb map[rune]*Trie
}

func NewTrie() *Trie {
  return &Trie{htb: make(map[rune]*Trie, 52)}
}

// O(word.length)
func (t *Trie) Set(words ...string) {
  for _, word := range words {
    trie := t
    for _, rne := range word {
      if _, exist := trie.htb[rne]; !exist {
        trie.htb[rne] = NewTrie()
      }
      trie = trie.htb[rne]
    }
  }
}

// O(word.length)
func (t *Trie) Get(word string) bool {
  trie := t
  for _, rne := range word {
    if _, exist := trie.htb[rne]; !exist {
      return false
    }
    trie = trie.htb[rne]
  }
  return true
}
