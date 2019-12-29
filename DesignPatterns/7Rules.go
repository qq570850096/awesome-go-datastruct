package DesignPatterns

import "fmt"

type ICar interface {
	// 车名
	GetName() string
	// 价格
	GetPrice() int
}

type BenzCar struct {
	name string
	price int
}

func (b BenzCar) GetName() string {
	return b.name
}

func (b BenzCar) GetPrice() int {
	return b.price
}

type FinanceBenzCar struct {
	BenzCar
}


func (b FinanceBenzCar) GetPrice() int {
	// 获取原价
	selfPrice := b.price
	var finance int
	if selfPrice >= 100 {
		finance = selfPrice + selfPrice*5/100
	} else if selfPrice >= 50 {
		finance = selfPrice + selfPrice*2/100
	} else {
		finance = selfPrice
	}
	return finance
}

type Girl struct {

}

type GroupLeader struct {
	girls []Girl
}

func (g GroupLeader) CountGirls ()  {
	fmt.Println("The sum of girls is ", len(g.girls))
}

type Teacher struct {

}

func (t Teacher) Command(leader GroupLeader)  {
	leader.CountGirls()
}



