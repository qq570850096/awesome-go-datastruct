package StructuralType

import (
	"fmt"
	"time"
)

// MediaPlayer is the component interface shared by the base object and
// decorators.
type MediaPlayer interface {
	GetMediaName() string
	GetMediaSeconds() int
	GetMediaContent() string
	play()
}

// VideoPlay is the concrete component being decorated.
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

// BarrageVideoPlay is the decorator base. It embeds a MediaPlayer so concrete
// decorators can forward unchanged behavior.
type BarrageVideoPlay struct {
	MediaPlayer
}

// BarrageRedVideoPlay is a concrete decorator that adds bullet-comment text to
// the media name and content.
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
