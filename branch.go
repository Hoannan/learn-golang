package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func main() {
	const filename = "abc.txt"

	/*contents, err :=ioutil.ReadFile(filename)

	if err != nil{
		println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}*/
	//第二种写法
	if contents, err := ioutil.ReadFile(filename); err != nil {
		println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(100),
		grade(20),
		grade(22),
	)
}
