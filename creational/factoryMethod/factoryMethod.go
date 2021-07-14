package factoryMethod

import "fmt"

//工厂方法模式，工厂创建出一些可以创建实例的方法

type Person struct {
	name string
	age  int
}

func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			name: name,
			age:  age,
		}
	}
}

func Run() {

	newBaby := NewPersonFactory(1)
	baby := newBaby("john")
	fmt.Println(baby)

	newTeenager := NewPersonFactory(16)
	teen := newTeenager("jill")
	fmt.Println(teen)
}
