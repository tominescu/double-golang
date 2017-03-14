package sort

func MergeSort(nums1, nums2 []int) []int {
	result := []int{}
	n1 := len(nums1)
	n2 := len(nums2)
	i, j := 0, 0
	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	for ; i < n1; i++ {
		result = append(result, nums1[i])
	}
	for ; j < n2; j++ {
		result = append(result, nums2[j])
	}
	return result
}
