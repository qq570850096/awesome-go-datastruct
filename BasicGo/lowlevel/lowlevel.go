package lowlevel

import "unsafe"

type Layout struct {
	Flag  bool
	Count int32
	Name  string
}

type LayoutInfo struct {
	Size        uintptr
	Align       uintptr
	CountOffset uintptr
	NameOffset  uintptr
}

func InspectLayout() LayoutInfo {
	var value Layout
	return LayoutInfo{
		Size:        unsafe.Sizeof(value),
		Align:       unsafe.Alignof(value),
		CountOffset: unsafe.Offsetof(value.Count),
		NameOffset:  unsafe.Offsetof(value.Name),
	}
}
