package sort

func QuickSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	pivot := 1
	for i := 1; i < n; i++ {
		if nums[i] < nums[0] {
			nums[pivot], nums[i] = nums[i], nums[pivot]
			pivot++
		}
	}
	nums[0], nums[pivot-1] = nums[pivot-1], nums[0]
	QuickSort(nums[:pivot-1])
	QuickSort(nums[pivot:])
}
