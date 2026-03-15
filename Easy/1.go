package main

func twoSum(nums []int, target int) []int {
	for left := 0; left < len(nums); left++ {
		for right := left + 1; right < len(nums); right++ {
			if nums[left]+nums[right] == target {
				return []int{left, right}
			}
		}
	}
	return nil
}

