package StructuralType

import (
	"fmt"
)

type GunSystem struct {
}

func (GunSystem)Fire()  {
	fmt.Println("开火")
}

func (GunSystem)UseBullet()  {
	fmt.Println("上子弹")
}

type UserSystem struct {
}

func (UserSystem)AddScore()  {
	fmt.Println("得分")
}

func (UserSystem)LoseBlood()  {
	fmt.Println("掉血")
}

type Facade struct {
	fire *GunSystem
	user *UserSystem
}

func (f *Facade) shooting()  {
	f.fire.Fire()
	f.fire.UseBullet()
	f.user.AddScore()
	f.user.LoseBlood()
}

