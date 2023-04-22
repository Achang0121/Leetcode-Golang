package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	result := make([]bool, n)

	maxValue := -1
	for _, v := range candies {
		if v >= maxValue {
			maxValue = v
		}
	}

	for i, v := range candies {
		if v == maxValue {
			result[i] = true
		}
		tmp := candies[i] + extraCandies
		if tmp >= maxValue {
			result[i] = true
		}
	}
	return result
}
