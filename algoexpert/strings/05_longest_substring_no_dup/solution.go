package main

type substring struct {
	start int
	end   int
}

func (ss substring) length() int { return ss.end - ss.start }

func LongestSubstringWithoutDuplication(str string) string {
	cache := make(map[rune]int)
	start := 0
	longest := substring{0, 1}

	for i, char := range str {
		if lastSeen, found := cache[char]; found && start < lastSeen+1 {
			start = lastSeen + 1
		}
		if i+1-start > longest.length() {
			longest = substring{start, i + 1}
		}
		cache[char] = i
	}

	return str[longest.start:longest.end]
}
