package main

import "fmt" //导入包后，必须要使用

func main() {
	//变量：程序运行期间，可以改变的量

	//1、声明格式		var 变量名 类型，变量声明了，必须要使用
	//2、只是声明没有初始化的变量，默认值为0
	//3、同一个{}里，声明的变量名是唯一的
	var a int
	fmt.Println("a=", a)

	//4、可以同时声明多个变量
	var b, c int
	b = 10
	fmt.Println("b=", b, "; c=", c)

	//变量的初始化，声明变量时，同时赋值
	var d int = 10 //初始化，声明变量时，同时赋值（一步到位）
	d = 20         //赋值
	fmt.Println("d = ", d)

	//自动推导类型，必须初始化，通过初始化的值确定类型
	e := 30
	fmt.Printf("c type is %T\n", e) //%T打印变量所属的类型
}
