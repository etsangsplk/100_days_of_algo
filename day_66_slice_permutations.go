package main

import (
	"fmt"
)

func range_slice(start, stop int) []int {
	if start > stop {
		panic("Slice ends before it started")
	}
	xs := make([]int, stop-start)
	for i := 0; i < len(xs); i++ {
		xs[i] = i + 1 + start
	}
	return xs
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func main() {

	no := permutation(range_slice(0, 4))
	fmt.Println("Number of Permutation:", len(no))
	for i := 0; i < len(no); i++ {
		fmt.Println(no[i])
	}
}
