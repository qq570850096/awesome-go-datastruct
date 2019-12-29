package CreativeType

import (
	"fmt"
	"sync"
)
// 单例模式
type ChocolateBoiler struct {
	empty bool
	boiled bool
}

var instance *ChocolateBoiler
var once sync.Once

func GetInstance() *ChocolateBoiler {
	// 使用Once保证创建实例的方法永远只能运行一次,就算在并发状态下也一定只执行一次
	once.Do(func() {
		instance = &ChocolateBoiler{true,false}
	})
	return instance
}

func (c *ChocolateBoiler) IsEmpty() bool {
	return c.empty
}

func (c *ChocolateBoiler) IsBoiled() bool {
	return c.boiled
}

func (c *ChocolateBoiler) Fill(){
	if c.empty{
		c.empty = false
		fmt.Println("容器装满了")
	}
}

func (c *ChocolateBoiler) Drain(){
	if c.empty==false && c.boiled {
		c.empty = true
		c.boiled = false
		fmt.Println("倒入模具了")
	}
}

func (c *ChocolateBoiler) Boil(){
	if c.empty==false && c.boiled == false {
		fmt.Println("巧克力煮开了")
		c.boiled = true
	}
}
