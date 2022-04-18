/*
面向对象的多态
interface
基本要素：有接口、有子类（实现了父类全部的接口）、父类类型的变量（指针）指向/引用子类具体的数据变量
*/
package main

import "fmt"

// AnimalInterface 定义动物接口类，本质是一个指针
type AnimalInterface interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// Cat 具体的类
type Cat struct {
	color string // 猫的颜色
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat is sleep...")
}

func (cat *Cat) GetColor() string {
	return cat.color
}

func (cat *Cat) GetType() string {
	return "Cat"
}

// Dog 具体的类
type Dog struct {
	color string // 狗的颜色
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog is sleep...")
}

func (dog *Dog) GetColor() string {
	return dog.color
}

func (dog *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalInterface) {
	animal.Sleep()
	fmt.Println("Animal color is", animal.GetColor())
	fmt.Println("Kind of animal is", animal.GetType())
}

func main() {
	/*
		var animal AnimalInterface // 接口的数据类型，父类指针

		animal = &Cat{color: "Black"}
		animal.Sleep() // 调用的就是 Cat 的 Sleep() 方法，多态的现象

		animal = &Dog{color: "Yellow"}
		animal.Sleep() // 调用的就是 Dog 的 Sleep() 方法多态的现象
	*/

	cat := Cat{color: "Black"}
	dog := Dog{color: "Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)
}
