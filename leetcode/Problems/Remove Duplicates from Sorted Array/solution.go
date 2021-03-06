func removeDuplicates(nums []int) int {
	i := 0
	k := 1

	for i+k < len(nums) {
		e := nums[i+k]
		if nums[i] == e {
			k++
		} else {
			i++
			nums[i] = e
		}
	}

	return i + 1
}