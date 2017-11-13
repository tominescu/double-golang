package sort

func heapify(nums []int, size int, root int) {
	lchild := 2*root + 1
	rchild := 2*root + 2

	if lchild >= size {
		return
	}

	max := lchild
	if rchild < size && nums[rchild] > nums[lchild] {
		max = rchild
	}

	if nums[root] < nums[max] {
		nums[root], nums[max] = nums[max], nums[root]
		heapify(nums, size, max)
	}
}

func HeapSort2(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, len(nums), i)
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}
