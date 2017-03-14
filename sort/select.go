package sort

func SelectSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}
