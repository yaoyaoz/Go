package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 21:46
 * @Author: yaoyao
 */
func main() {
	/*
		定义在{}里面的变量就是局部变量，只能在{}里面有效
		执行到定义变量那句话，才开始分配空间，离开作用域自动释放
		作用域，变量其作用的范围
	*/
	//a = 111
	{
		i := 13
		fmt.Println("i = ", i)
	}
	//i = 111

	if flag := 3; flag == 3 {
		fmt.Println("flag = ", flag)
	}
	//flag = 4
}

func test2() {
	a := 10
	fmt.Println("a = ", a)
}
