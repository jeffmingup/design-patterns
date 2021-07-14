package decorator

import "log"

//装饰器模式，给已有的方法包装一些动作，且不影响原方法的属性
type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}

func Double(n int) int {
	return n * 2
}

func run() {
	f := LogDecorate(Double)

	f(5)
}
