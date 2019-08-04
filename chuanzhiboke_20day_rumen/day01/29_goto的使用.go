package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-04 22:44
 * @Author: yaoyao
 */
func main() {

	//break	//break is not in a loop, switch, or select
	//continue	//continue is not in a loop

	//goto可以用在任何地方，但是不能跨函数使用
	fmt.Println("1111111111")

	goto End //goto是关键字，End是用户起的名字，它叫标签

	fmt.Println("2222222222")

End:
	fmt.Println("3333333333")

}
