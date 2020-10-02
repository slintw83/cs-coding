package main

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	return InnerSum(array, 1)
}

func InnerSum(array SpecialArray, multiplier int) int {
	sum := 0
	for _, v := range array {
		if cast, ok := v.(SpecialArray); ok {
			sum += InnerSum(cast, multiplier+1)
		} else if cast, ok := v.(int); ok {
			sum += cast
		}
	}
	return sum * multiplier
}
