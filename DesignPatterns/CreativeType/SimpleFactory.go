package CreativeType

import "fmt"

// gril is the product interface for the simple factory example. The original
// spelling is preserved to avoid changing the public shape of the sample.
type gril interface {
	weight()
}

// FatGril is one concrete product variant.
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

// CreateGirl centralizes product selection so callers do not instantiate
// concrete variants directly.
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
