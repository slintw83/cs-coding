package main

import (
	"fmt"
	"sort"
)

type DiskSort [][]int

func (d DiskSort) Len() int {
	return len(d)
}

func (d DiskSort) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DiskSort) Less(i, j int) bool {
	return d[i][2] < d[j][2]
}

func DiskStacking(disks [][]int) [][]int {
	fmt.Println(disks)
	// Sort disks first
	sort.Sort(sort.Reverse(DiskSort(disks)))
	fmt.Println(disks)
	return Lis(disks)
}

func Lis(disks [][]int) [][]int {
	maxSum := make([]int, len(disks))
	sequences := make([]int, len(disks))
	for i := range maxSum {
		maxSum[i] = disks[i][2]
		sequences[i] = -1
	}

	for i := range maxSum {
		for j := 0; j < i; j++ {
			if disks[i][0] < disks[j][0] && disks[i][1] < disks[j][1] && disks[i][2] < disks[j][2] {
				newSum := disks[i][2] + maxSum[j]
				if newSum > maxSum[i] {
					maxSum[i] = newSum
					sequences[i] = j
				}
			}
		}
	}

	maxIdx := 0
	for i := range maxSum {
		if maxSum[i] > maxSum[maxIdx] {
			maxIdx = i
		}
	}

	result := make([][]int, 0)
	for maxIdx != -1 {
		result = append(result, disks[maxIdx])
		maxIdx = sequences[maxIdx]
	}
	return result
}
