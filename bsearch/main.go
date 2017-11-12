package main

import "fmt"

func bsearch(a []int, key int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if a[mid] == key {
		return mid
	} else if key < a[mid] {
		return bsearch(a, key, low, mid-1)
	} else {
		return bsearch(a, key, mid+1, high)
	}
	return -1
}

func howmany(a []int, key int) int {
	m := bsearch(a, key, 0, len(a)-1)
	if m == -1 {
		return 0
	}
	l := m
	for {
		t := bsearch(a, key, 0, l-1)
		if t == -1 {
			break
		}
		l = t
	}
	r := m
	for {
		t := bsearch(a, key, r+1, len(a)-1)
		if t == -1 {
			break
		}
		r = t
	}
	return r - l + 1
}

func main() {
	a := []int{1, 2, 3, 4, 6, 6, 6, 7, 8, 8, 9}
	ret := howmany(a, 6)
	fmt.Println(ret)
}
