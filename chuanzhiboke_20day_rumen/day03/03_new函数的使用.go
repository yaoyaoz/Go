package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-13 23:12
 * @Author: yaoyao
 */
func main() {
	//a := 10	//整型变量a

	var p *int
	//指向一个合法内存
	//p = &a

	//p是*int,指向int类型
	p = new(int)
	*p = 666
	fmt.Println("*p = ", *p)

	q := new(int) //自动推导类型
	*q = 555
	fmt.Printf("q = %v, *q = %d\n", q, *q)

}
