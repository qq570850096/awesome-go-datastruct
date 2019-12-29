package DesignPatterns

import (
	"fmt"
	"testing"
)

func TestBenzCar_GetName(t *testing.T) {
	var (
		list []ICar
	)
	list = []ICar{}
	list = append(list,&FinanceBenzCar{BenzCar{"迈巴赫",99}})
	list = append(list,&FinanceBenzCar{BenzCar{"AMG",200}})
	list = append(list,&FinanceBenzCar{BenzCar{"V",40}})
	for _,v := range list {
		fmt.Println("车名:",v.GetName(),"\t价格:",v.GetPrice())
	}

	person := Person{ani:&Rubbit{}}
	person.WalkAnimal()
}
