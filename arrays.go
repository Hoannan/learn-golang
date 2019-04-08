package main

import "fmt"

func PrintArr(arr [3]int) {
	for i := range arr {
		fmt.Println(i)
	}
}

func loopArr(arr [3]int) {
	//通用方法
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//只获取数组的value
	for i := range arr {
		fmt.Println(arr[i])
	}
	//获取数组的key和value
	for k, v := range arr {
		fmt.Println(k, v)
	}
	//只获取值 使用_省略变量
	for _, v := range arr {
		fmt.Println(v)
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

	//定义指针数组
	arr5 := [2]*string{new(string), new(string)}

	*arr5[0] = "hello"
	*arr5[1] = "hello"

	fmt.Println(arr5)

	//数组复制
	arr6 := arr5

	fmt.Println(arr6)

	arr7 := [3]int{11, 22, 222}
	loopArr(arr7)
}
