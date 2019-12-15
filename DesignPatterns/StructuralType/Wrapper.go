package StructuralType

import "fmt"

type Target interface {
	TargetMethod1()
	TargetMethod2()
}

type Adaptee struct {
}

func (Adaptee) MethodA ()  {
	fmt.Println("Adaptee methodA invoked")
}

func (Adaptee) MethodB ()  {
	fmt.Println("Adaptee methodB invoked")
}

type Adapter struct {
	Adaptee
}

func (Adapter) TargetMethod1()  {
	fmt.Println("Adapter targetMethod1 invoked")
	Adaptee{}.MethodA()
}

func (Adapter) TargetMethod2()  {
	fmt.Println("Adapter targetMethod1 invoked")
	Adapter{}.MethodB()
}