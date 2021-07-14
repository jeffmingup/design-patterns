package template

import (
	"fmt"
	"testing"
)

//模版模式， 定义一个抽象类，实现大部分固定好的方法，然后让子类实现剩下的特定的方法，然后执行
type Cooker interface {
	fire()
	cooke()
	outfire()
}

// 类似于一个抽象类
type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜，交给具体的子类实现 //这样好像不能强制子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("做西红柿")
}

type ChaoJiDan struct {
	CookMenu
}

func (ChaoJiDan) cooke() {
	fmt.Println("做炒鸡蛋")
}

func TestTemplate(t *testing.T) {
	// 做西红柿
	xihongshi := &XiHongShi{}
	doCook(xihongshi)

	fmt.Println("\n=====> 做另外一道菜")
	// 做炒鸡蛋
	chaojidan := &ChaoJiDan{}
	doCook(chaojidan)

}
