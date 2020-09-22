package main

import "sort"

type ByChar []rune

func (s ByChar) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s ByChar) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByChar) Len() int {
	return len(s)
}

func key(word string) string {
	r := []rune(word)
	sort.Sort(ByChar(r))
	return string(r)
}

func GroupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		key := key(word)
		if _, ok := groups[key]; !ok {
			groups[key] = make([]string, 0)
		}

		groups[key] = append(groups[key], word)
	}

	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
