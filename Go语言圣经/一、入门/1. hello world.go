// Go语言圣经：https://books.studygolang.com/gopl-zh/ch1/ch1-01.html

// Go语言的代码通过包（package）组织，每个源文件都要以package声明语句开始
// main包比较特殊，定义了一个独立可执行的程序，而不是一个库
package main

// import导入用到的包，必须跟在package声明之后
import "fmt"

// main包中的main函数也比较特殊，是整个走程序执行的入口（类似C/C++）
func main() {
	fmt.Println("Hello World !!")
}
