package two_number_sum

func TwoNumberSum(array []int, target int) []int {
	cache := make(map[int]int)

	for _, v := range array {
		t, ok := cache[target-v]
		if ok {
			return []int{v, t}
		} else {
			cache[v] = v
		}
	}

	return []int{}
}
