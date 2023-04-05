package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		if idx, ok := m[target-value]; ok {
			return []int{idx, index}
		}
		m[value] = index
	}
	return nil
}
