package StructuralType

import "testing"

func TestBarrageRedVideoPlay_GetMediaContent(t *testing.T) {
	mediaplay1 := &VideoPlay{
		videoContent: "Legend of the Condor Heroes",
		videoName:    "Guo Jing vs Ouyang Feng",
	}
	mediaplay := &BarrageRedVideoPlay{BarrageVideoPlay{mediaplay1}}

	mediaplay1.play()

	mediaplay.play()
}
