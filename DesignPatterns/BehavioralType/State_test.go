package BehavioralType

import "testing"

func TestTelev_Play(t *testing.T) {
	context := RemoteControlMachine{}

	context.SetCurrentSate(&PowerOffState{})

	context.Play()

	context.PowerOn()
	context.Play()
	context.Standby()
	context.PowerOff()
}
