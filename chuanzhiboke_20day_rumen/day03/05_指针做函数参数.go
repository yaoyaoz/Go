package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-13 23:26
 * @Author: yaoyao
 */
func main() {
	a, b := 10, 20

	//通过一个函数交换a和b的内容
	swap2(&a, &b) //地址传递
	fmt.Printf("swap: a = %d, b = %d\n", a, b)
}

func swap2(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Printf("main: *p1 = %d, *p2 = %d\n", *p1, *p2)
}
