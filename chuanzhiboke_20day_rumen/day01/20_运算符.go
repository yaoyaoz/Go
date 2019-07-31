package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-07-31 22:25.
 * @Author: yaoyao
 */
func main() {
	fmt.Println("4 > 3 结果：", 4 > 3)
	fmt.Println("4 != 3 结果", 4 != 3)

	fmt.Println("!(4 > 3) 结果：", !(4 > 3))
	fmt.Println("!(4 != 3) 结果：", !(4 != 3))

	fmt.Println("true && true 结果：", true && true)

	/*
		&& 与
		|| 或
		非0就是真，0就是假（bool类型和int类型不能转换）

		a++ 没有++a
		a-- 没有--a

		& 取地址运算符
		* 取值运算符 指针变量所指向内存的值
	*/

}
