package StructuralType

import "testing"

func TestGunSystem_Fire(t *testing.T) {
	facade := &Facade{
		fire: &GunSystem{},
		user: &UserSystem{},
	}
	facade.shooting()
}
