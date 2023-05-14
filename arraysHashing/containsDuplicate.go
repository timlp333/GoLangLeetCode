package arraysHashing

func ContainsDuplicate(nums []int) bool {
	checkMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		_, isExist := checkMap[nums[i]]
		if !isExist {
			checkMap[nums[i]] = nums[i]
		} else {
			return true
		}
	}
	return false
}
