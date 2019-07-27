package main

import "fmt"

func main() {
	//不同类型变量的声明（定义）
	//var a int
	//var b float64

	//var (
	//	a int = 1
	//	b float64 =4
	//)

	var (
		a = 1
		b = 6.0
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//const i int = 10
	//const j float64 = 3.14

	//const (
	//	i int = 10
	//	j float64 = 3.14
	//)

	//可以自动推导类型
	const (
		i = 10
		j = 3.14
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
}
