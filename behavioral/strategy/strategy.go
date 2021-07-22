package strategy

import (
	"fmt"
	"testing"
)

// 策略模式
// 可以随意添加修改策略，而不用修改策略执行类的代码

type strategy func(int, int) int

func strategyAdd(a, b int) int {
	return a + b
}
func strategySub(a, b int) int {
	return a - b
}

// 具体策略的执行者
type Operator struct {
	strategy strategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy strategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy(a, b)
}

func TestStrategy(t *testing.T) {
	operator := Operator{}

	operator.setStrategy(strategyAdd)
	result := operator.calculate(1, 2)
	fmt.Println("add:", result)

	operator.setStrategy(strategySub)
	result = operator.calculate(2, 1)
	fmt.Println("reduce:", result)
}
