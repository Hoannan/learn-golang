package main

import "fmt"

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
	s2 := s1[2:5]

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("s2 =%v len(s2) = %d cap(s2) = %d", s2, len(s2), cap(s2) )

}
