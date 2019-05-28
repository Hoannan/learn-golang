package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Println(slice1)

	slice1[2] = 200

	fmt.Println(slice1)

	//根据切片创建切片，和根据自定义数组创建切片方式相同，长度是2=3-1，容量是4=5-1
	//但是需要格外注意，新生成的切片 slice2 和原始切片 slice1 的指针元素指向了相同的底层数组，所以修改元素要注意

	slice2 := slice1

	slice1[1] = 1000

	fmt.Println(slice2)

	//append 增加切片长度

	slice3 := []int{10,20,30,40,50}

	newSlice := slice3[1:3]

	fmt.Println(slice3)
	fmt.Println(newSlice)
	//由于底层数组还有容量，可以直接追加元素而容量不变
	newSlice = append(newSlice, 1000)

	fmt.Println(newSlice)
	fmt.Println(slice3)

	/*当切片容量（而非数组长度，默认切片容量等于数组长度，也可以显示指定）用完了，再追加需要扩容，此处会新建数组，长度为原数组的2倍，然后将旧数组元素拷贝到新数组，newSlice 的底层数组是新数组，newSlice 容量为8；而 slice 的底层数组是旧数组，二者互不影响
	slice 扩容机制：在切片的容量小于 1000 个元素时，总是会成倍地增加容量。一旦元素个数超过 1000，容量的增长因子会设为 1.25，也就是会每次增加 25% 的容量

	作者：原水寒
	链接：https://www.jianshu.com/p/78975d33a018
	来源：简书
	简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。*/
	slice4 := append(slice3, 3000)

	fmt.Println(slice4)

	/*在没有显示指定容量的情况下，切片容量就是其底层数组的长度
	如 2.3 所示，如果在创建切片时设置切片的容量和长度一样，就可以强制让新切片的第一个 append 操作创建新的底层数组，与原有的底层数组分离。
	新切片与原有的底层数组分离后，可以安全地进行后续修改*/

	slice5 := []int{1,2,3,4,5,6}
	//长度为1=3-2，容量为1=3-2  source[i:j:k] 长度=j-i 容量=k-i
	slice6 := slice5[2:3:3]

	fmt.Println(slice5)
	fmt.Println(slice6)

	slice6 = append(slice6, 100)

	fmt.Println(slice6)
	fmt.Println(slice5)

}
