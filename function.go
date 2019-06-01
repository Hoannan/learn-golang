package main

import "fmt"

/*方法的定义方法实际上也是函数，只是在声明时，在关键字 func 和方法名之间增加了一个参数

普通的函数定义 func 方法名(入参) 返回值

自定义类型的方法定义 func (接收者) 方法名(入参) 返回值

方法的值传递和指针传递

func (u user) notify() 拷贝一份 user

func (u *user) changeEmail(newEmail string) 传递指针(即地址)，内部改变会影响外部
*/

type user struct {
	name string
	age  int
}

func (u user) makePerson() {
	fmt.Println("name:", u.name, "age:", u.age)
}

func (u *user) changeAge(age int) {
	u.age = age
}

func (u user) changeName() {
	u.name = "sam"
	fmt.Println("name:", u.name, "age:", u.age)
}

func main() {
	// 1. user类型的值可以用来调用使用值接收者声明的方法
	khn := user{"khn", 24}

	khn.makePerson()

	khn.changeAge(25)

	khn.makePerson()

	khn.changeName()

	khn.makePerson()
}
