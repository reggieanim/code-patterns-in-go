package main

import "math"

//complete
func main() {

}

func Solution(str string) int {
	windowStart := 0
	maximumLength := 0
	distinctChars := map[string]int{}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := string(str[windowEnd])
		if _, in := distinctChars[rightChar]; in {
			windowStart = int(math.Max(float64(distinctChars[rightChar]), float64(windowStart)))
		}
		distinctChars[rightChar] = windowEnd

		maximumLength = int(math.Max(float64(maximumLength), float64(windowEnd-windowStart+1)))
	}
	return maximumLength
}
