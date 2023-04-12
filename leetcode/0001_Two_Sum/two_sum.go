package twosum

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)

	for index, val := range nums {
		toSearch := target - val
		searchIndex, ok := mapping[toSearch]
		if ok {
			return []int{searchIndex, index}
		}
		mapping[val] = index
	}

	return []int{-1, -1}
}
