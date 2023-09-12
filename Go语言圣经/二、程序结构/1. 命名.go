package 二_程序结构

/*

1. Go语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则：
	一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。
	大写字母和小写字母是不同的：heapSort和Heapsort是两个不同的名字。

2. Go语言中类似if和switch的关键字有25:
	break      default       func     interface   select
	case       defer         go       map         struct
	chan       else          goto     package     switch
	const      fallthrough   if       range       type
	continue   for           import   return      var

还有大约30多个预定义的名字：
	内建常量: true false iota nil
	内建类型: int int8 int16 int32 int64
			  uint uint8 uint16 uint32 uint64 uintptr
			  float32 float64 complex128 complex64
			  bool byte rune string error
	内建函数: make len cap new append copy close delete
			  complex real imag
			  panic recover

3. 如果一个名字是在函数内部定义，那么它就只在函数内部有效。如果是在函数外部定义，那么将在当前包的所有文件中都可以访问。
名字的开头字母的大小写决定了名字在包外的可见性。
如果一个名字是大写字母开头的（译注：必须是在函数外部定义的包级名字；包级函数名本身也是包级名字），那么它将是导出的，也就是说可以被外部的包访问

4. Go语言程序员推荐使用 驼峰式 命名，当名字由几个单词组成时优先使用大小写分隔，而不是优先用下划线分隔


*/