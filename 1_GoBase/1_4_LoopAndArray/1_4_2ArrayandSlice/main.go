package main

import "fmt"

func main() {
	//new Array1　不可變動類型
	var array1 [3]int
	array1[0] = 12
	array1[2] = 9
	for i, r := range array1 {
		fmt.Printf("%d:%d \n", i, r)
	}
	fmt.Println(array1[0:2])
	//new Array 可變動類型
	var array2 []int
	array2 = append(array2, 10, 89, 70)
	fmt.Println(array2)
	// 參數2　表示數量　參數3　表示容量
	array3 := make([]int, 3, 10)
	array3 = append(array3, 10, 89, 70)
	fmt.Println(array3)

}
