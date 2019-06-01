package main

import "fmt"

func main() {
	type user struct {
		name string
		sex  int
		age  int
	}

	type admin struct {
		//自定义类
		person user
		//内置类型
		level int
	}

	//实例化类
	var haonan user

	fmt.Println(haonan)

	shuyi := user{"shuyi", 1, 99}

	fmt.Println(shuyi)

	haonan = user{"haonan", 2, 23}
	fmt.Println(haonan)

	khn := admin{haonan, 99}

	xsy := admin{shuyi, 82}

	fmt.Println(khn)
	fmt.Println(xsy)
}
