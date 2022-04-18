/**
类、属性、方法，如果要对外开放（其他包中调用）的话首字母都需要大写
*/
package main

import "fmt"

// Hero 如果类名首字母大写，标识其他包也能够访问
type Hero struct {
	// 如果说类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	Level int
}

//func (hero Hero) show() {
//	fmt.Println("Name =", hero.Name)
//	fmt.Println("Ad =", hero.Ad)
//	fmt.Println("Level =", hero.Level)
//}
//
//func (hero Hero) getName() string {
//	return hero.Name
//}
//
//func (hero Hero) setName(newName string) {
//	// hero 是调用该方法的对象的一个副本（值拷贝）
//	hero.Name = newName
//}

// Show 当前方法的首字母大写，表示其他的模块/包中也可以访问，否则的话是能在当前包/模块中访问
func (hero *Hero) Show() {
	fmt.Println("Name =", hero.Name)
	fmt.Println("Ad =", hero.Ad)
	fmt.Println("Level =", hero.Level)
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) SetName(newName string) {
	// hero 是调用该方法的对象的一个副本（值拷贝）
	hero.Name = newName
}

func main() {
	// 创建一个对象
	hero := Hero{Name: "ZhangSan", Ad: 100, Level: 1}
	hero.Show()

	hero.SetName("LiSi")
	hero.Show()
}
