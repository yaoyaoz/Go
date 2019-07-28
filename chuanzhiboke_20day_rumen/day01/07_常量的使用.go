package main

import "fmt"

func main() {
	//变量：程序运行期间，可以改变的量。变量声明需要关键字var
	//常量：程序运行期间，不可以改变的量。常亮声明需要关键字const

	const a int = 10
	//a = 20	//err, 常亮不允许修改
	fmt.Println("a = ", a)

	const b = 10.3                  //没有用使用:=
	fmt.Printf("b type is %T\n", b) //%T 类型
	fmt.Println("b = ", b)
}
