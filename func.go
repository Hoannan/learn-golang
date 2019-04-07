package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func option(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported option: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

//两个参数值互换
func swap(a, b *int) {
	*a, *b = *b, *a
}

func newSwap(a, b int) (int, int) {

	return b, a
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name();
	fmt.Printf("Call function name %s with params %d , %d\n", opName, a, b);
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

func main() {
	fmt.Println(
		option(1, 2, "*"),
	)

	fmt.Println(div(333, 2))
	a, b := 4, 5
	swap(&a, &b)
	fmt.Println(a, b)

	c, e := newSwap(3, 2)
	fmt.Println(c, e)

	//初次尝试函数式编程
	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a, b int) int {
		return int(math.Mod(float64(a), float64(b)))
	}, 11, 3))

	fmt.Println(sum(22,11,12,12313,2131))
}
