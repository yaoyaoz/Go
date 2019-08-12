package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-12 22:47
 * @Author: yaoyao
 */
func main() {
	var p *int
	//p = nil
	fmt.Println("p = ", p)

	//*p = 666 //err,因为p没有合法指向

	var a int
	p = &a //p指向a
	*p = 666
	fmt.Printf("a = %v, p = %v, *p = %v\n", a, p, *p)
}
