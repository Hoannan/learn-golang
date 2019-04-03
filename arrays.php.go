package main

import "fmt"

func PrintArr(arr [3]int){
	for i:=range arr{
		fmt.Println(i)
	}
}

func main() {
	var arr1 [5] int

	arr2 := [3] int{1, 4, 6}

	arr3 := [...] int{1, 2, 3, 4, 333, 222, 1}

	//二维数组
	var arr4 [4][6] bool

	fmt.Println(arr1, arr2, arr3, arr4)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	PrintArr(arr2)
}
