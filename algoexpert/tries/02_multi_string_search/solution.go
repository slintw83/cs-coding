package main

import (
	"fmt"
	"sort"
	"strings"
)

type SuffixSort []string

func (ss SuffixSort) Len() int {
	return len(ss)
}

func (ss SuffixSort) Less(i, j int) bool {
	return ss[i] < ss[j]
}

func (ss SuffixSort) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func MultiStringSearch(bigString string, smallStrings []string) []bool {
	suffixArray := make([]string, 0)
	for i := range bigString {
		suffixArray = append(suffixArray, bigString[i:])
	}
	sort.Sort(SuffixSort(suffixArray)) // NlogN
	fmt.Println(suffixArray[0])
	result := make([]bool, len(smallStrings))
	for i := range smallStrings { // NlogN to search every small string position
		idx := sort.SearchStrings(suffixArray, smallStrings[i])
		if idx < len(suffixArray) {
			if strings.HasPrefix(suffixArray[idx], smallStrings[i]) { // M * NlogN
				result[i] = true
			}
		}
	}
	return result
}
