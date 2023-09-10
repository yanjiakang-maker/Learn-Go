package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 模拟实现Unix的uniq命令，寻找相邻的重复行

/*
	1. map: kv键值对
	2. 打印占位符
		%d          十进制整数
		%x, %o, %b  十六进制，八进制，二进制整数。
		%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
		%t          布尔：true或false
		%c          字符（rune） (Unicode码点)
		%s          字符串
		%q          带双引号的字符串"abc"或带单引号的字符'c'
		%v          变量的自然形式（natural format）
		%T          变量的类型
		%%          字面上的百分号标志（无操作数）

*/

func dup1()  {
	// 第一版本：打印输入中多次出现的行，以重复次数开头

	// bufio包，它使处理输入和输出方便又高效
	// Scanner类型是该包最有用的特性之一，它读取输入并将其拆成行或单词；通常是处理行形式的输入最简单的方法
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// 每次dup读取一行输入，该行被当做键存入map，其对应的值递增
	// 每次调用input.Scan()，即读入下一行，并移除行末的换行符；读取的内容可以调用input.Text()得到
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts{
		if n > 1 {
			// 类似于C或其它语言里的printf函数，fmt.Printf函数对一些表达式产生格式化输出。
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}



/*
	1. os.Open函数返回两个值。第一个值是被打开的文件(*os.File）, 返回的第二个值是内置error类型的值
	2. map是一个由make函数创建的数据结构的引用。map作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），
		被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到
		（译注：类似于C++里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）
*/

func countLines(f* os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}


func dup2()  {
	// 第二版，读取标准输入或是使用os.Open打开各个具名文件，并操作它们

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}


// dup的前两个版本以"流”模式读取输入，并根据需要拆分成多个行。理论上，这些程序可以处理任意数量的输入数据。
//还有另一个方法，就是一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。下面这个版本，dup3，就是这么操作的。
// 引入了ReadFile函数（来自于io/ioutil包），其读取指定文件的全部内容，strings.Split函数把字符串分割成子串的切片。
func dup3()  {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		//ReadFile函数返回一个字节切片（byte slice），必须把它转换为string，才能用strings.Split分割
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}


func main() {
	dup1()
}