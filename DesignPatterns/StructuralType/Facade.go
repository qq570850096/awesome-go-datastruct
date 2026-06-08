package StructuralType

import (
	"fmt"
)

type GunSystem struct {
}

func (GunSystem) Fire() {
	fmt.Println("fire")
}

func (GunSystem) UseBullet() {
	fmt.Println("load bullet")
}

type UserSystem struct {
}

func (UserSystem) AddScore() {
	fmt.Println("score")
}

func (UserSystem) LoseBlood() {
	fmt.Println("lose health")
}

type Facade struct {
	fire *GunSystem
	user *UserSystem
}

func (f *Facade) shooting() {
	f.fire.Fire()
	f.fire.UseBullet()
	f.user.AddScore()
	f.user.LoseBlood()
}
