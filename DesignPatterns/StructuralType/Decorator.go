package StructuralType

import (
	"fmt"
	"time"
)

type MediaPlayer interface {
	GetMediaName() string
	GetMediaSeconds() int
	GetMediaContent() string
	play()
}

type VideoPlay struct {
	MediaPlayer
	videoContent string
	videoName    string
}

func (v *VideoPlay) GetMediaName() string {
	return v.videoName
}

func (v *VideoPlay) GetMediaSeconds() int {
	return len(v.videoContent)
}

func (v *VideoPlay) GetMediaContent() string {
	return v.videoContent
}

func (m *VideoPlay) play() {
	fmt.Println("Media:", m.GetMediaName(),
		"(duration: ", m.GetMediaSeconds(), " seconds) is playing")
	for i := 0; i < m.GetMediaSeconds(); i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("playing second ", i, "s"+m.GetMediaContent())
	}
}

type BarrageVideoPlay struct {
	MediaPlayer
}

type BarrageRedVideoPlay struct {
	BarrageVideoPlay
}

func (b *BarrageRedVideoPlay) GetMediaName() string {
	return b.MediaPlayer.GetMediaName() + " enables bullet comments"
}

func (b *BarrageRedVideoPlay) GetMediaSeconds() int {
	return b.MediaPlayer.GetMediaSeconds()
}

func (b *BarrageRedVideoPlay) GetMediaContent() string {
	return "bullet comments---" + b.MediaPlayer.GetMediaContent()
}

func (m *BarrageRedVideoPlay) play() {
	fmt.Println("Media:", m.GetMediaName(),
		"(duration: ", m.GetMediaSeconds(), " seconds) is playing")
	for i := 0; i < m.GetMediaSeconds(); i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("playing second ", i, "s"+m.GetMediaContent())
	}
}
