package can_jump

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			pass := false
			for j := i - 1; j >= 0 ; j-- {
				if nums[j] + j > i || nums[j] + j == len(nums) - 1 {
					pass = true
					break
				}
			}
			if !pass {
				return false
			}
		}
	}

	return true
}
