package BehavioralType

import "testing"

func TestTelev_Play(t *testing.T) {
	tv := Telev{
		ITelevision: nil,
		state:       POWER_OFF_STATE,
	}
	// 这里因为电视还是关机状态，所以不会有任何的输出
	tv.Play()

	tv.PowerOn()
	tv.Play()
	tv.Standby()
	tv.PowerOff()
}
