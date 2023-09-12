package main

import "fmt"

/*

1. Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。

2. 一个Go语言编写的程序对应一个或多个以.go为文件后缀名的源文件。
每个源文件中以包的声明语句开始，说明该源文件是属于哪个包。
包声明语句之后是import语句导入依赖的其它包，
然后是包一级的类型、变量、常量、函数的声明语句，包一级的各种类型的声明语句的顺序无关紧要



 */

// 包一级范围声明
const biolingF = 222.22


func main() {
	// main函数内部声明
	var f = biolingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\\n", f, c)

}