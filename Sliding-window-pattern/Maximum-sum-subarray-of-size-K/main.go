package main

import (
	"math"
)

// Given an array of positive numbers and a positive number ‘k,’
// find the maximum sum of any contiguous subarray of size ‘k’

func main() {
}

func MaxSumOfSizeK(ints []int, k int) int {
	var maxSum int
	sum := 0
	windowStart := 0
	for i := 0; i < len(ints); i++ {
		sum += ints[i]
		if i >= k-1 {
			maxSum = int(math.Max(float64(sum), float64(maxSum)))
			sum -= ints[windowStart]
			windowStart += 1
		}
	}
	return maxSum
}
