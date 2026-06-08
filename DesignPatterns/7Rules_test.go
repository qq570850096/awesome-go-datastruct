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
	list = append(list, &FinanceBenzCar{BenzCar{"Maybach", 99}})
	list = append(list, &FinanceBenzCar{BenzCar{"AMG", 200}})
	list = append(list, &FinanceBenzCar{BenzCar{"V", 40}})
	for _, v := range list {
		fmt.Println("car:", v.GetName(), "\tprice:", v.GetPrice())
	}

}
