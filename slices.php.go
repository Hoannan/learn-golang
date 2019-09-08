package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 123
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s := arr[2:6]

	fmt.Println("arr[2:6] = ", s)

	fmt.Println("arr[2:6] = ", arr[2:])
	fmt.Println("arr[2:6] = ", arr[:6])
	fmt.Println("arr[2:6] = ", arr[:])

	updateSlice(s)

	updateSlice(arr[:])

	fmt.Println("after update s", s)
	fmt.Println("after update arr", arr)

	s1 := arr[1:5]
	s2 := s1[2:7]

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("s2 =%v len(s2) = %d cap(s2) = %d", s2, len(s2), cap(s2))

	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	ss1 := arr3[2:6]//2345
	ss2 := ss1[3:5]//56
	ss3 := append(ss2, 10)//5610
	ss4 := append(ss3, 11)//561011
	ss5 := append(ss4, 12)//561012

	fmt.Printf("arr3 = %v ss1=%v ss2=%v ss3=%v ss4=%v ss5=%v", arr3, ss1, ss2, ss3, ss4, ss5)
}
