package main

//方式1
//import "fmt"
//import "os"

//方式2
import (
	//"fmt"
	"os"
)

//.操作
//import . "fmt"	//调用函数，无需通过包名

//给包名起别名
import io "fmt"

/*
忽略此包
_操作其实是引入该包，而不是直接使用包里面的函数，而是调用了该包里面的init函数
*/
//import _ "fmt"

/**
 * @Description:
 * @Date: 2019-08-11 22:02
 * @Author: yaoyao
 */
func main() {
	io.Println("this is a test")
	io.Println("os.Args = ", os.Args)

}
