package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-07-28 16:30.
 * @Author: yaoyao
 */
func main() {
	//给int64起一个别名bigint
	type bigint int64

	var a bigint //等价于int64 a
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d, ch = %c\n", b, ch)
}
