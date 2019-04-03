package main

import "fmt"

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
}
