package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{5, 8, 13, 15, 2, 7, 3, 11, 4, 18}
	key := 15
	isFind := false
	sort.Ints(array)
	mid, max, min := 0, 0, 0
	max = len(array)

	for max >= min {
		mid = (min + max) >> 1
		if array[mid] == key {
			fmt.Printf("index:%d value:%d\n", mid, key)
			isFind = true
			break
		} else if array[mid] > key {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	if !isFind {
		fmt.Println("Not Find!")
	}

}
