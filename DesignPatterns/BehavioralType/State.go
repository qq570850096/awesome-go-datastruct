package BehavioralType

import "fmt"

const (
	STANDBY_STATE = 1
	POWER_OFF_STATE = 2
	PLAY_STATE = 3
)
type ITelevision interface {
	// 开机
	PowerOn()
	// 关机
	PowerOff()
	// 播放
	Play()
	// 待机
	Standby()
}

type Telev struct {
	ITelevision
	state int
}

func (telev *Telev) State() int {
	return telev.state
}

func (telev *Telev) SetState(state int) {
	telev.state = state
}

func (t *Telev) PowerOn() {
	switch t.state {
		case STANDBY_STATE:
		case POWER_OFF_STATE:
			fmt.Println("开机")
			t.SetState(STANDBY_STATE)
		case PLAY_STATE:
		default:
	}
}
func (t *Telev) PowerOff() {
	// 待机和播放状态都可以关机
	switch t.state {
	case STANDBY_STATE:
		fmt.Println("关机")
		t.SetState(POWER_OFF_STATE)
	case PLAY_STATE:
		fmt.Println("关机")
		t.SetState(POWER_OFF_STATE)
	case POWER_OFF_STATE:
	default:
	}
}
func (t *Telev) Play() {
	switch t.state {
	case STANDBY_STATE:
		fmt.Println("播放")
		t.SetState(PLAY_STATE)
	default:
	}
}
func (t *Telev) Standby() {
	switch t.state {
	case POWER_OFF_STATE:
		fmt.Println("关机")
		t.SetState(POWER_OFF_STATE)
	case PLAY_STATE:
		fmt.Println("待机")
		t.SetState(PLAY_STATE)
	default:
	}
}



