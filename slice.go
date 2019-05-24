package main

import "fmt"

func main() {
	/*
	切片是围绕动态数组的概念构建的，可以按需自动增长和缩小。
	切片的动态增长是通过内置函数 append 来实现的。
	这个函数可以快速且高效地增长切片。还可以通过对切片再次切片来缩小一个切片的大小。

	切片有 3 个字段分别是指向底层数组的指针、切片访问的元素的个数（即长度）和切片允许增长到的元素个数（即容量）

	从切片 slice1 创建出来的切片 slice2，slice1 和 slice2 共享底层数组，一个修改了共享部分的元素，另一个也会感知
	*/

	//使用make函数创建一个字符串切片，长度和容量都是5
	slice1 := make([]string, 5)

	slice1[2] = "abc"
	//2、创建一个int切片，长度是3，容量是5
	slice2 := make([]int, 6)

	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	fmt.Println(slice2)

	fmt.Println(slice1)

	//使用字面量创建切片，长度是5，容量是5

	slice3 := []int{1, 2, 3, 5, 9}
	fmt.Println(slice3)
	fmt.Printf("slice 3 is a  slice look %d ", slice3)

	//创建 nil 切片，长度为0，容量为0
	var slice4 []int
	slice5 := make([]int, 0)
	slice6 := []int{}
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Printf("slice 4 is a nil slice look %d", slice4)

	//6、自定义底层数组，通过该底层数组创建切片
	arr := [5]int{1,2,3,9,5}

	slice7 := arr[2:4]
	slice8 := arr[3:]
	slice9 := arr[0:4]


	fmt.Printf("arr to slice7 %d\n slice8 %d\n slice9 %d", slice7,slice8,slice9)

}
