package array

func ContainsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool, len(nums))

	for _, num := range nums {
		if numsMap[num] {
			return true
		}

		numsMap[num] = true
	}

	return false
}
