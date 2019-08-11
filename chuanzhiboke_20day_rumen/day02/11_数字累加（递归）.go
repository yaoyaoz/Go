package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 12:23
 * @Author: yaoyao
 */
func main() {
	var sum int
	sum = funcSum1()
	fmt.Println("sum = ", sum)

	sum2 := funcSum2(100)
	fmt.Println("sum2 = ", sum2)
}

//实现1+2+3+...+100
func funcSum1() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func funcSum2(i int) (sum int) {
	if i == 1 {
		return 1
	}
	return i + funcSum2(i-1)
}
