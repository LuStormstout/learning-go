/**
面向对象类的继承
*/
package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (human *Human) Eat() {
	fmt.Println("human eating...")
}

func (human *Human) Walk() {
	fmt.Println("human walking...")
}

// ======== 分割线 ========

// Superman 定义一个 Superman 类继承自 Human 类
type Superman struct {
	Human // 表示继承 Human 类
	// level 是 Superman 独有的属性
	level int
}

// Eat 重新定义（重写）父类中的 Eat() 方法
func (superman *Superman) Eat() {
	fmt.Println("Superman eating more...")
}

// Fly 子类的新方法
func (superman *Superman) Fly() {
	fmt.Println("Superman is flying...")
}

func (superman *Superman) PrintSuperman() {
	fmt.Println("name is :", superman.name)
	fmt.Println("sex is :", superman.sex)
	fmt.Println("level is :", superman.level)
}

func main() {
	h := Human{name: "ZhangSan", sex: "female"}
	h.Eat()
	h.Walk()

	fmt.Println("====== ====== ======")

	// 定义一个子类对象
	//s := Superman{Human{"LiSi", "male"}, 90}
	// 上面的方法不易阅读的话，下面这种方法就很好理解了，因为 Superman 继承了 Human 所以直接可以用 s. 来调用 Human 类中的属性
	var s Superman
	s.name = "LiSi"
	s.sex = "male"
	s.level = 90

	s.Walk() // 调用的是父类的 Walk()，父类的方法
	s.Eat()  // 调用执行的是重写过的，子类的方法
	s.Fly()  // 调用的是子类的方法
	s.PrintSuperman()
}
