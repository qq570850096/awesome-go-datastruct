package BehavioralType

import "fmt"

type TVState interface {
	// 开机
	PowerOn(r *RemoteControlMachine)
	// 关机
	PowerOff(r *RemoteControlMachine)
	// 播放
	Play(r *RemoteControlMachine)
	// 待机
	Standby(r *RemoteControlMachine)
}
// 待机状态
type StandByState struct {
	r *RemoteControlMachine
}

func (s *StandByState) PowerOn(r *RemoteControlMachine) {}

func (s *StandByState) PowerOff(r *RemoteControlMachine) {
	fmt.Println("关机")
	// 使用遥控器设置电视机状态为关机
	s.r = r
	s.r.SetCurrentSate(&PowerOffState{})
	// 执行关机
	s.r.PowerOff()
}

func (s *StandByState) Play(r *RemoteControlMachine) {
	fmt.Println("播放")
	// 使用遥控器设置电视机状态为播放
	s.r = r
	s.r.SetCurrentSate(&PlayState{})
	// 执行播放
	s.r.Play()
}

func (s *StandByState) Standby(r *RemoteControlMachine) {
	// do nothing
}
// 关机状态
type PowerOffState struct {
	r *RemoteControlMachine
}

func (s *PowerOffState) PowerOn(r *RemoteControlMachine) {
	fmt.Println("开机")
	// 使用遥控器设置电视机状态为开机
	s.r = r
	s.r.SetCurrentSate(&StandByState{})
	// 执行播放
	s.r.Standby()
}

func (s *PowerOffState) PowerOff(r *RemoteControlMachine) {
}

func (s *PowerOffState) Play(r *RemoteControlMachine) {
}

func (s PowerOffState) Standby(r *RemoteControlMachine) {
}

// 播放状态
type PlayState struct {
	r *RemoteControlMachine
}

func (s *PlayState) PowerOn(r *RemoteControlMachine) {}

func (s *PlayState) PowerOff(r *RemoteControlMachine) {
	fmt.Println("关机")
	// 使用遥控器设置电视机状态为关机
	s.r = r
	s.r.SetCurrentSate(&PowerOffState{})
	// 执行关机
	s.r.PowerOff()
}

func (s *PlayState) Play(r *RemoteControlMachine) {
}

func (s *PlayState) Standby(r *RemoteControlMachine) {
	fmt.Println("开机")
	// 使用遥控器设置电视机状态为开机
	s.r = r
	s.r.SetCurrentSate(&StandByState{})
	// 执行播放
	s.r.Standby()
}

// 引入控制器（上下文角色）
type RemoteControlMachine struct {
	currentSate TVState
}

func (r *RemoteControlMachine) PowerOn() {
	r.currentSate.PowerOn(r)
}

func (r *RemoteControlMachine) PowerOff() {
	r.currentSate.PowerOff(r)
}

func (r *RemoteControlMachine) Play() {
	r.currentSate.Play(r)
}

func (r *RemoteControlMachine) Standby() {
	r.currentSate.Standby(r)
}

func (r *RemoteControlMachine) CurrentSate() TVState {
	return r.currentSate
}

func (r *RemoteControlMachine) SetCurrentSate(currentSate TVState) {
	r.currentSate = currentSate
}



