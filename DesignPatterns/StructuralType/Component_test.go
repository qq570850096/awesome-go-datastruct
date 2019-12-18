package StructuralType

import "testing"

func TestComposite_Add(t *testing.T) {
	root := Composite{
		name: "综合实验室",
		arr:  make([]Component, 0),
	}
	root.Add(&Leaf{"综合设备1"})
	root.Add(&Leaf{"综合设备2"})

	branchLevel21 := Composite{
		name: "化学实验室",
		arr:  make([]Component, 0),
	}
	branchLevel21.Add(&Leaf{"试管"})
	branchLevel21.Add(&Leaf{"烧杯"})
	branchLevel21.Add(&Leaf{"锥形瓶"})

	root.Add(&branchLevel21)

	// 并列的二级节点
	branchLevel22 := Composite{
		name: "物理实验室",
		arr:  make([]Component, 0),
	}
	branchLevel22.Add(&Leaf{"开关"})
	branchLevel22.Add(&Leaf{"电阻箱"})
	branchLevel22.Add(&Leaf{"锥形瓶"})

	branchLevel221 := Composite{
		name: "物理精密仪器组",
		arr:  make([]Component, 0),
	}
	branchLevel221.Add(&Leaf{"光学测量仪"})
	branchLevel221.Add(&Leaf{"精密机床"})
	branchLevel22.Add(&branchLevel221)

	root.Add(&branchLevel22)

	root.Display(1)

	root.Remove(&branchLevel22)
	root.Display(1)}
