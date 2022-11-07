package main

func twoSum(nums []int, target int) []int {

	if len(nums) < 2 {
		return nil
	}

	res := []int{0, 0}

	// brute force approach
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				res[0] = i
				res[1] = j
				return res
			}

		}
	}
	return nil
}

func main() {
}
