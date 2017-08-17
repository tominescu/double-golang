package sort

func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		changed := false
		for j := n - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				changed = true
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
		if !changed {
			break
		}
	}
}
