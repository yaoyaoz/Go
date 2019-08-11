package main

import "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 11:06
 * @Author: yaoyao
 */
func main() {

	MyFunc01(4, 4)

	MyFunc02(1)
	MyFunc02(1, 3, 5)

	MyFunc03("3", 3, 6)

}

//固定参数
func MyFunc01(a int, b int) {

}

/*
...int; ...type	不定参数类型
注意：不定参数，只能放在形参中的最后一关参数
*/
func MyFunc02(args ...int) { //传递的实参可以是0或多个
	fmt.Println("len(args) = ", len(args)) //获取用户传递参数的个数
	for i := 0; i < len(args); i++ {
		fmt.Printf("arg[%d] = %d\n", i, args[i])
	}

	//返回2个值，第一个是下标，第二个是下标所对应的数
	for i, data := range args {
		fmt.Printf("arg[%d] = %d\n", i, data)
	}
}

func MyFunc03(a string, args ...int) {

}
