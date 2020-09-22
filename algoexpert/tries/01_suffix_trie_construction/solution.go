package main

// Do not edit the class below except for the
// PopulateSuffixTrieFrom and Contains methods.
// Feel free to add new properties and methods
// to the class.
type SuffixTrie map[byte]SuffixTrie

// Feel free to add to this function.
func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {
	for i := range str {
		node := trie
		for j := i; j < len(str); j++ {
			c := str[j]
			if _, found := node[c]; !found {
				node[c] = NewSuffixTrie()
			}
			node = node[c]
		}
		node['*'] = nil
	}
}

func (trie SuffixTrie) Contains(str string) bool {
	node := trie
	for i := range str {
		c := str[i]
		if nextNode, found := node[c]; found {
			node = nextNode
		} else {
			return false
		}
	}
	_, found := node['*']
	return found
}
