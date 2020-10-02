package main

func GetNthFib(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return GetNthFib(n-2) + GetNthFib(n-1)
}
