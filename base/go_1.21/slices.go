package go_1_21

import (
	"cmp"
	"fmt"
	"slices"
)

func slicesNew() {
	s := []int{-1, -2, 1154, 675, 786}
	fmt.Println(slices.Max(s))
	// clear(s)
	// fmt.Println(s, len(s), cap(s))

	fmt.Println("sort before")
	slices.SortFunc(s, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	fmt.Println("sort after: ", s)
	s1, s2 := 9.3, 4.3
	fmt.Println(max(s1, s2), min(s1, s2))

	fmt.Println("sort binarySearch")
	search, b := slices.BinarySearch(s, 675)
	fmt.Println(search, b)
	fmt.Println("sort binarySearch function")
}
