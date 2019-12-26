package StructuralType

import "testing"

func TestComposite_Add(t *testing.T) {
	root := Composite{
		name: "和平饭店",
		arr:  make([]MenuComponent, 0),
	}


	branchLevel21 := Composite{
		name: "招牌菜",
		arr:  make([]MenuComponent, 0),
	}
	branchLevel21.Add(&Leaf{"红烧肉",false,"精五花",20.0})
	branchLevel21.Add(&Leaf{"醋溜土豆丝",true,"新鲜",10.0})
	branchLevel21.Add(&Leaf{"京酱肉丝",false,"鲜肉",30.0})

	root.Add(&branchLevel21)

	// 并列的二级节点
	branchLevel22 := Composite{
		name: "家常菜",
		arr:  make([]MenuComponent, 0),
	}
	branchLevel22.Add(&Leaf{"辣椒炒肉",false,"1",14.2})
	branchLevel22.Add(&Leaf{"杂拌",false,"1234",15})
	branchLevel22.Add(&Leaf{"回锅肉",false,"2134",30})

	branchLevel221 := Composite{
		name: "麻辣烫",
		arr:  make([]MenuComponent, 0),
	}
	branchLevel221.Add(&Leaf{"豆蔻",true,"1",0.5})
	branchLevel221.Add(&Leaf{"腐竹",true,"1",0.5})
	branchLevel22.Add(&branchLevel221)

	root.Add(&branchLevel22)

	root.Display(1)

	root.Remove(&branchLevel22)
	root.Display(1)}
