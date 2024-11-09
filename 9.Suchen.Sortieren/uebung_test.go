package suchensortieren

import "fmt"

func ExampleUebungSuch() {
	fmt.Println(UebungSuch([]byte{'b', 'A', 't', 'a', 'G', 'm'}, 'a'))
	fmt.Println(UebungSuch([]byte{'A', 'H', 'Z', 'a', 'h', 'z'}, 'Z'))
	fmt.Println(UebungSuch([]byte{}, 'a'))
	// Output:
	// 3
	// 2
	// -1

}

func ExampleUebungSort() {
	fmt.Println(UebungSort([]byte{'b', 'A', 't', 'a', 'G', 'm'}))
	fmt.Println(UebungSort([]byte{'A', 'H', 'Z', 'a', 'h', 'z'}))
	fmt.Println(UebungSort([]byte{}))
	// Output:
	// [65 71 97 98 109 116]
	// [65 72 90 97 104 122]
	// []
}

// Lösung:
/*
	1.
	if len(bl) == 0 {
		return -1
	}
	if !slices.IsSorted(bl) { // Wenn die Liste nicht sortiert ist, wird eine einfache Suche durchgeführt
		for i := 0; i < len(bl); i++ {
			if bl[i] == s {
				return i
			}
		}
		return -1
	} else { // Wenn die Liste sortiert ist, wird eine binäre Suche durchgeführt
		var left = 0
		var right = len(bl) - 1
		for left <= right {
			var middle = (left + right) / 2
			if bl[middle] == s {
				return middle
			} else if bl[middle] < s {
				left = middle
			} else {
				right = middle - 1
			}
		}
	}
	return -1

	2.
	for i := 0; i < len(bl); i++ {
		for j := i + 1; j < len(bl); j++ {
			if bl[i] > bl[j] {
				var temp = bl[i]
				bl[i] = bl[j]
				bl[j] = temp
			}
		}
	}
	return bl

*/
