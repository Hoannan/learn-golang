package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

//变量学习
var (
	aa = 123
	bb = "abc"
	cc = true
)

func variable() {
	var age int
	var name string
	fmt.Printf("%d %q\n", age, name)
}

func variableInitialValue() {
	var age1, age2 int = 13, 14
	var name1, name2 string = "KHN", "xsy"
	fmt.Println(age1, age2, name1, name2)
}

func variableNoType() {
	var age1, age2, isok, name1 = 12, 22, true, "KHN"
	fmt.Println(age1, age2, isok, name1)
}

//第一次定义变量 需要使用：
func variablshorter() {
	age1, age2, isok, name1 := 12, 22, true, "KHN"
	age1 = 88
	fmt.Println(age1, age2, isok, name1)
}

func euler() {
	a := 3 + 4i
	fmt.Println(cmplx.Abs(a))

	//欧拉公式
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	//强制类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const name1 string = "khn"
	const a, b = 4, 5
	fmt.Println(name1, a, b)
}

func emum() {
	//普通枚举
	const (
		cpp  = 1
		php  = 2
		java = 3
	)

	fmt.Println(cpp, php, java)

	const (
		a = iota + 1

		c
		d
		e
		f
	)
	fmt.Println(a, c, d, e, f)
	//b, kb, mb, gb, tb, pb
	//自增枚举
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variable()
	variableInitialValue()
	variableNoType()
	variablshorter()
	fmt.Println(aa, bb, cc)

	euler()
	triangle()
	consts()
	emum()
}
