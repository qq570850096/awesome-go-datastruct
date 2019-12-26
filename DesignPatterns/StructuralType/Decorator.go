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
	videoName string
}

func (v *VideoPlay) GetMediaName() string  {
	return v.videoName
}

func (v *VideoPlay) GetMediaSeconds() int  {
	return len(v.videoContent)
}

func (v *VideoPlay) GetMediaContent() string  {
	return v.videoContent
}

// 模拟媒体播放
func (m *VideoPlay) play() {
	fmt.Println("Media:",m.GetMediaName(),
		"(累计时长：",m.GetMediaSeconds(),"秒)正在播放")
	for i:=0;i<m.GetMediaSeconds();i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("当前播放第",i,"秒"+m.GetMediaContent())
	}
}

type BarrageVideoPlay struct {
	MediaPlayer
}

// 实现弹幕播放的类，具体的修饰器的实现类
type BarrageRedVideoPlay struct {
	BarrageVideoPlay
}

func (b *BarrageRedVideoPlay) GetMediaName () string {
	return b.MediaPlayer.GetMediaName()+"开启弹幕"
}

func (b *BarrageRedVideoPlay)GetMediaSeconds() int  {
	return b.MediaPlayer.GetMediaSeconds()
}

func (b *BarrageRedVideoPlay)GetMediaContent() string {
	return "弹幕中---"+b.MediaPlayer.GetMediaContent()
}

// 模拟媒体播放
func (m *BarrageRedVideoPlay) play() {
	fmt.Println("Media:",m.GetMediaName(),
		"(累计时长：",m.GetMediaSeconds(),"秒)正在播放")
	for i:=0;i<m.GetMediaSeconds();i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("当前播放第",i,"秒"+m.GetMediaContent())
	}
}