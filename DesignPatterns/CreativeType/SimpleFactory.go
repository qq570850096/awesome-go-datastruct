package CreativeType

import "fmt"

type gril interface {
	weight()
}
// 简单工厂模式
type FatGril struct {
}

func (FatGril) weight()  {
	fmt.Println("100kg")
}

type ThinGirl struct {
}

func (ThinGirl) weight ()  {
	fmt.Println("45kg")
}

type GirlFactory struct {
}

func (*GirlFactory) CreateGirl (like string) gril {
	switch like {
		case "fat":
			return &FatGril{}
		case "thin":
			return &ThinGirl{}
		default:
			return nil
	}
}
