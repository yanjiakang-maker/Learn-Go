package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 获取对应的url，并将其源文本打印出来
// 这个例子的灵感来源于curl工具（译注：unix下的一个用来发http请求的工具，具体可以man curl）

/*
	1. http.Get函数是创建HTTP请求的函数，如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果
	2. resp的Body字段包括一个可读的服务器响应流。
	3. ioutil.ReadAll函数从response中读取到全部内容；将其结果保存在变量b中
	4. resp.Body.Close关闭resp的Body流，防止资源泄露，Printf函数会将结果b写出到标准输出流中
*/
func fetch()  {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}


func main() {
	fetch()
}