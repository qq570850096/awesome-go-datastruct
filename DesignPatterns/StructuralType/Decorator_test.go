package StructuralType

import "testing"

func TestBarrageRedVideoPlay_GetMediaContent(t *testing.T) {
	mediaplay1 := &VideoPlay{
		videoContent: "射雕英雄传",
		videoName:    "郭靖战欧阳锋",
	}
	mediaplay := &BarrageRedVideoPlay{BarrageVideoPlay{mediaplay1}}
	// 不加修饰时候的播放
	mediaplay1.play()
	// 添加修饰后的播放
	mediaplay.play()
}
