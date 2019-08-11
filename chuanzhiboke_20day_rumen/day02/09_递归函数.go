package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 12:15
 * @Author: yaoyao
 */
func main() {
	funca1(3)
	fmt.Println("main")
}

func funca1(a int) {
	funcb1(a - 1)
	fmt.Println("a = ", a)
}

func funcb1(b int) {
	funcc1(b - 1)
	fmt.Println("b = ", b)
}

func funcc1(c int) {
	fmt.Println("c = ", c)
}
