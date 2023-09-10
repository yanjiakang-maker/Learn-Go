package main

import (
	"fmt"
	"os"
	"strings"
)

// 模拟实现unix的echo命令，echo把它的命令行参数打印成一行
/*
			1. 程序的命令行参数可从os包的Args变量获取； os.Args
				os.Args变量是一个字符串（string）的切片（slice）
				s[i]访问单个元素，用s[m:n]获取子序列
				Go言里也采用左闭右开形式：比如a = [1, 2, 3, 4, 5], a[0:3] = [1, 2, 3]
			2. var声明定义了两个string类型的变量s和sep
				变量会在声明时直接初始化。如果变量没有显式初始化，则被隐式地赋予其类型的零值（zero value），数值类型是0，字符串类型是空字符串""。
			3. 循环索引变量i在for循环的第一部分中定义
				符号:= 是短变量声明（short variable declaration）的一部分，这是定义一个或多个变量并根据它们的初始值为这些变量赋予适当类型的语句。
			4. 自增语句i++给i加1；这和i += 1以及i = i + 1都是等价的
				它们是语句，而不像C系的其它语言那样是表达式。所以j = i++非法，而且++和--都只能放在变量名后面，因此--i也非法。
			5. Go语言只有for循环这一种循环语句
				for initialization; condition; post {
		    		// zero or more statements
				}
			a. initialization语句是可选的，在循环开始前执行
				initialization 如果存在，必须是一条简单语句（simple statement），即，短变量声明、自增语句、赋值语句或函数调用
			b. ion是一个布尔表达式（boolean expression）
				其值在每次循环迭代开始时计算。如果为true则执行循环体语句
			c. post语句在循环体执行结束后执行，之后再次对condition求值。condition值为false时,循环结束。
			d. for循环的这三个部分每个都可以省略
			省略前后，只保留条件
				for condition {
	    			// ...
				}

			无限循环，可用break或者return终止
				for {
					// ...
				}
*/

func echo1() {
	//	第一版
	var s, sep string
	for i := 1; 1 < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	// for循环的另一种形式，在某种数据类型区间(range)遍历（比如字符串、切片）
	// 我们这里不用os.Args了，用一个其他数组玩一下
	var s, sep string
	var args = []string{"arg1", "arg2", "arg3"}
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// range产生一对值；索引以及在该索引处的元素值。这个例子不需要索引，但range的语法要求，要处理元素，必须处理索引
	//Go语言中这种情况的解决方法是用空标识符（blank identifier），即_（也就是下划线）。
	//空标识符可用于在任何语法需要变量名但程序逻辑不需要的时候（如：在循环里）丢弃不需要的循环索引，
	/*
	   声明变量方式：（一般使用前两种）
	     s := "" （短变量声明，最简洁，只能用在函数内部，不能用于包变量）
	     var s string （默认初始化零值）
	     var s = ""
	     var s string = "" （显式标明变量类型）

	*/
}

func echo3() {
	// 第三版，使用strings包的Join函数进行拼接
	var args = []string{"arg1", "arg2", "arg3"}

	fmt.Println(strings.Join(args[0:], " "))
}

func homework1() {
	fmt.Println(os.Args[0])
}

func homework2() {
	var s string
	var args = []string{os.Args[0], "arg1", "arg2", "arg3"}
	num := 1
	for _, arg := range args {
		s = "[" + string(num) + "] " + arg
		fmt.Println(s)
		num += 1
	}
}

func main() {
	//echo1()
	//echo2()
	//echo3()
	//homework1()
	homework2()
}
