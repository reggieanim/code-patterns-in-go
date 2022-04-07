package main

import (
	"math"
)

// Given an array of positive numbers and a positive number ‘S,’
// find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘S’.
//  Return 0 if no such subarray exists.
func main() {

}

func SmallestSubArrWGreaterSum(ints []int, k int) int {
	windowStart := 0
	sum := 0
	minLength := int(len(ints) + 100)

	for windowEnd := 0; windowEnd < len(ints); windowEnd++ {
		sum += ints[windowEnd]
		for sum >= k {
			minLength = int(math.Min(float64(minLength), float64(windowEnd-windowStart+1)))
			sum -= ints[windowStart]
			windowStart += 1
		}
	}
	// set minLength to impossible value
	if minLength == int(len(ints)+100) {
		return 0
	}
	return minLength
}