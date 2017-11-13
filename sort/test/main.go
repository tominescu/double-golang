package main

import (
	"fmt"

	"github.com/tominescu/double-golang/sort"
)

func main() {
	//nums := []int{5, 4, 3, 2, 1}
	nums := []int{1, 2, 5, 7, 9, 2, 4, 6, 8, 10}
	//sort.BubbleSort(nums)
	//sort.InsertSort(nums)
	//sort.SelectSort(nums)
	//sort.QuickSort(nums)
	sort.HeapSort2(nums)
	/*
		nums1 := []int{1, 3, 5}
		nums2 := []int{2, 4, 6, 8, 10}
		result := sort.MergeSort(nums1, nums2)
	*/
	fmt.Println(nums)
}
