package StructuralType

import "testing"

func TestComposite_Add(t *testing.T) {
	root := Composite{
		name: "Peace Hotel",
		arr:  make([]MenuComponent, 0),
	}

	branchLevel21 := Composite{
		name: "Signature dishes",
		arr:  make([]MenuComponent, 0),
	}
	branchLevel21.Add(&Leaf{"braised pork", false, "pork belly", 20.0})
	branchLevel21.Add(&Leaf{"sour shredded potato", true, "fresh", 10.0})
	branchLevel21.Add(&Leaf{"shredded pork with sweet bean sauce", false, "fresh pork", 30.0})

	root.Add(&branchLevel21)

	branchLevel22 := Composite{
		name: "home-style dishes",
		arr:  make([]MenuComponent, 0),
	}
	branchLevel22.Add(&Leaf{"pepper pork", false, "1", 14.2})
	branchLevel22.Add(&Leaf{"mixed platter", false, "1234", 15})
	branchLevel22.Add(&Leaf{"twice-cooked pork", false, "2134", 30})

	branchLevel221 := Composite{
		name: "malatang",
		arr:  make([]MenuComponent, 0),
	}
	branchLevel221.Add(&Leaf{"cardamom", true, "1", 0.5})
	branchLevel221.Add(&Leaf{"bean curd stick", true, "1", 0.5})
	branchLevel22.Add(&branchLevel221)

	root.Add(&branchLevel22)

	root.Display(1)

	root.Remove(&branchLevel22)
	root.Display(1)
}
