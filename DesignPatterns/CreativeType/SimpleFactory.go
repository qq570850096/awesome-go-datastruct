package CreativeType

import "fmt"

type gril interface {
	weight()
}

type FatGril struct {
}

func (FatGril) weight() {
	fmt.Println("100kg")
}

type ThinGirl struct {
}

func (ThinGirl) weight() {
	fmt.Println("45kg")
}

type GirlFactory struct {
}

func (*GirlFactory) CreateGirl(like string) gril {
	switch like {
	case "fat":
		return &FatGril{}
	case "thin":
		return &ThinGirl{}
	default:
		return nil
	}
}
