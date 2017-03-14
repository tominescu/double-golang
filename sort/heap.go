package sort

type Heap struct {
	l []int
}

func (h *Heap) Insert(e int) {
	h.l = append(h.l, e)
	h.PerlocateUp()
	//fmt.Println(h.l)
}

func (h *Heap) Delete() int {
	n := len(h.l)
	if n == 0 {
		return -1
	}
	first := h.l[0]
	h.l[0] = h.l[n-1]
	h.l = h.l[:n-1]
	h.PerlocateDown()
	//fmt.Println(h.l)
	return first
}

func (h *Heap) PerlocateUp() {
	curr := len(h.l) - 1
	parent := curr / 2
	for curr > 0 {
		if h.l[curr] < h.l[parent] {
			h.l[curr], h.l[parent] = h.l[parent], h.l[curr]
			curr = parent
			parent = curr / 2
		} else {
			break
		}
	}
}

func (h *Heap) PerlocateDown() {
	n := len(h.l)
	curr := 0
	left := curr*2 + 1
	right := curr*2 + 2
	for left < n {
		minIdx := left
		if right < n && h.l[right] < h.l[left] {
			minIdx = right
		}
		if h.l[curr] > h.l[minIdx] {
			h.l[curr], h.l[minIdx] = h.l[minIdx], h.l[curr]
			curr = minIdx
			left = curr*2 + 1
			right = curr*2 + 2
		} else {
			break
		}
	}
}

func HeapSort(nums []int) []int {
	result := []int{}
	h := Heap{l: []int{}}
	for _, num := range nums {
		h.Insert(num)
	}
	for i := 0; i < len(nums); i++ {
		result = append(result, h.Delete())
	}
	return result
}
