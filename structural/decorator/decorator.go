package main

import "log"

//装饰器模式，给已有的方法包装一些动作，且不影响原方法的属性
type Object func(int) int

//重点 装饰器必须满足可以保存原有对象，返回值也和原有对象属于同一接口
type decorator func(Object) Object

func LogDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}
func LogDecorate2(fn Object) Object {
	return func(n int) int {
		log.Println("111")

		result := fn(n)

		log.Println("222")

		return result
	}
}

func Double(n int) int {
	return n * 2
}

func main() {
	f := LogDecorate(Double)
	f = LogDecorate2(f)
	f(5)
}
