package main

import (
	//"math/rand"
	//"time"
	"sort"
)

func main() {
	ints := []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)
	indexOf4 := sort.SearchInts(sortedInts, 4)
	indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v (present: %v)", indexOf4,
		sortedInts[indexOf4] == 4)
	Printfln("Index of 3: %v (present: %v)", indexOf3,
		sortedInts[indexOf3] == 3)
}
