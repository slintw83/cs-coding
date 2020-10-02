package main

func GetNthFib(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	a, b := 0, 1
	for i := 3; i <= n; i++ {
		c := b + a
		a = b
		b = c
	}
	return b
}
