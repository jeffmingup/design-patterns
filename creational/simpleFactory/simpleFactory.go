package simpleFactory

import "fmt"

type Door struct {
	width  int
	height int
}

func (door *Door) String() string {
	return fmt.Sprintf("door width:%v,height:%v", door.width, door.height)
}

//简单工程方法可以确保实例具有需要的参数，进而保证实例正常运行
func MakeDoor(width, height int) *Door {
	return &Door{width: width, height: height}
}
