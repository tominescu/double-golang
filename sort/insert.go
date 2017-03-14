package sort

func InsertSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
}
