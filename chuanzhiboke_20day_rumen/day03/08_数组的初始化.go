package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-13 23:45
 * @Author: yaoyao
 */
func main() {
	//声明定义同时赋值，叫初始化
	//1、全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	//部分初始化，没有初始化的元素，自动赋值为0
	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)

	//指定某个元素初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Println("d = ", d)
}

/*
输出：
a =  [1 2 3 4 5]
b =  [1 2 3 4 5]
c =  [1 2 3 0 0]
d =  [0 0 10 0 20]
*/
