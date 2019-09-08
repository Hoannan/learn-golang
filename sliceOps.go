package main

import "fmt"

func printSliceLen(s []int) {
	fmt.Printf("%v ,s len = %d and s cap=%d \n", s, len(s), cap(s))
}

func main() {
	var s []int

	for i := 0; i < 100; i++ {
		printSliceLen(s)
		s = append(s, i)
	}

	fmt.Println(s)

	s1 := []int{1, 2, 3, 4, 5}

	s2 := make([]int, 20)

	s3 := make([]int, 10, 32)

	printSliceLen(s1)
	printSliceLen(s2)
	printSliceLen(s3)

	s4 := append(s2, 100)

	printSliceLen(s4)

	s5 := append(s3, 50)
	printSliceLen(s5)

	//copy
	copy(s2, s1)
	printSliceLen(s2)

	//delete
	//s2[:3] + s2[4:]
	s2 = append(s2[:3], s2[4:]...)
	printSliceLen(s2)

}
