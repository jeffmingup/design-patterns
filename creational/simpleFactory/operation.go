package main

import "log"

type operation interface {
	SetParamsA(int)
	SetParamsB(int)
	GetResult() int
}
type params struct {
	a int
	b int
}

func (p *params) SetParamsA(a int) {
	p.a = a
}
func (p *params) SetParamsB(b int) {
	p.b = b
}

//实现了加法
type operationAdd struct {
	params
}

func (o *operationAdd) GetResult() int {
	return o.a + o.b
}

//减法
type operationSub struct {
	params
}

func (o *operationSub) GetResult() int {
	return o.a - o.b
}

//乘法
type operationMul struct {
	params
}

func (o *operationMul) GetResult() int {
	return o.a * o.b
}

//减法
type operationDiv struct {
	params
}

func (o *operationDiv) GetResult() int {
	return o.a / o.b
}

//简单工厂方法，返回的是接口值，客户端只需要操作接口定义好的函数就可以了
func getOperation(str string) operation {
	switch str {
	case "+":
		return &operationAdd{}
	case "-":
		return &operationSub{}
	case "*":
		return &operationMul{}
	case "/":
		return &operationDiv{}
	default:
		panic("error str")
	}
}

func main() {
	operation := getOperation("+")
	operation.SetParamsA(2)
	operation.SetParamsB(8)
	log.Println(operation.GetResult())

}
