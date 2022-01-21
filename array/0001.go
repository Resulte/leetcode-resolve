package array

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums)) //key: nums[i], val: i

	for index, val := range nums {
		if value, ok := numsMap[target-val]; ok {
			return []int{value, index}
		}
		numsMap[val] = index
	}

	return nil
}
