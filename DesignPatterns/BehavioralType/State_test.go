package BehavioralType

import "testing"

func TestTelev_Play(t *testing.T) {
	context := RemoteControlMachine{}

	context.SetCurrentSate(&PowerOffState{})
	// 如果直接播放，因为电视处于关机状态，所以不会有输出
	context.Play()

	context.PowerOn()
	context.Play()
	context.Standby()
	context.PowerOff()
}
