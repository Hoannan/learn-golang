package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(str int) string {
	result := ""

	for ; str > 0; str /= 2 {
		lsb := str % 2
		fmt.Println(str)
		fmt.Println(lsb)
		result = strconv.Itoa(lsb) + result
	}

	return result
}
//只有终止的for循环 相当于while
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//死循环
func forever(txt string){
	for {
		fmt.Println(txt)
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
	)
	printFile("abc.txt")
	forever("abc");
}
