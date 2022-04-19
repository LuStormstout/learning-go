package main

import (
	"learning-go/init_function/lib1"
	//"learn_golang/init-function/lib1"
	// 导入包的时候起别名为 _ 为匿名导入，即使不使用 lib1 也不会报错（golang 中导入的包不使用会报错）
	// 匿名导入的包无法使用该包中的方法，但是会执行该包中的 init() 方法
	//_ "learn_golang/init-function/lib1"

	//"learn_golang/init-function/lib2"
	// 导入包的时候给导入的包起一个别名 mylib2
	mylib2 "learning-go/init-function/lib2"
	// 导入包的时候，起别名为 . 相当于把 lib2 中的方法直接导入到了当前的包中（虽然语法支持，但是不推荐使用，因为在导入多个包的时候会出现方法同名）
	//. "learn_golang/init-function/lib2"
)

func main() {
	lib1.TestLib1()

	//lib2.TestLib2()
	// 在包导入的时候给包起了别名，那么调用包里面的 API 的时候就使用别名来调用
	mylib2.TestLib2()
	// 在包导入的时候用 . 来当别名，调用 lib2 中的 TestLib2() 方法如下
	//TestLib2()
}

/**
main
⬇️
import lib1 ➡️ 	lib1
				⬇️
				import lib2	➡️	lib2
								⬇️
								const
								⬇️
								var
								⬇️
								init()
								⬅️
				const
				⬇️
				var
				⬇️
				init()
				⬅️
const
⬇️
var
⬇️
init()
⬇️
main()
*/
