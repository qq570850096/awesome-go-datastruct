package BehavioralType

import "fmt"

// TVState is the state role. Each state decides how remote-control operations
// behave and when to transition the machine to another state.
type TVState interface {
	PowerOn(r *RemoteControlMachine)

	PowerOff(r *RemoteControlMachine)

	Play(r *RemoteControlMachine)

	Standby(r *RemoteControlMachine)
}

// StandByState handles operations while the TV is powered on but idle.
type StandByState struct {
	r *RemoteControlMachine
}

func (s *StandByState) PowerOn(r *RemoteControlMachine) {}

func (s *StandByState) PowerOff(r *RemoteControlMachine) {
	fmt.Println("power off")

	s.r = r
	s.r.SetCurrentSate(&PowerOffState{})

	s.r.PowerOff()
}

func (s *StandByState) Play(r *RemoteControlMachine) {
	fmt.Println("play")

	s.r = r
	s.r.SetCurrentSate(&PlayState{})

	s.r.Play()
}

func (s *StandByState) Standby(r *RemoteControlMachine) {

}

type PowerOffState struct {
	r *RemoteControlMachine
}

func (s *PowerOffState) PowerOn(r *RemoteControlMachine) {
	fmt.Println("power on")

	s.r = r
	s.r.SetCurrentSate(&StandByState{})

	s.r.Standby()
}

func (s *PowerOffState) PowerOff(r *RemoteControlMachine) {
}

func (s *PowerOffState) Play(r *RemoteControlMachine) {
}

func (s PowerOffState) Standby(r *RemoteControlMachine) {
}

type PlayState struct {
	r *RemoteControlMachine
}

func (s *PlayState) PowerOn(r *RemoteControlMachine) {}

func (s *PlayState) PowerOff(r *RemoteControlMachine) {
	fmt.Println("power off")

	s.r = r
	s.r.SetCurrentSate(&PowerOffState{})

	s.r.PowerOff()
}

func (s *PlayState) Play(r *RemoteControlMachine) {
}

func (s *PlayState) Standby(r *RemoteControlMachine) {
	fmt.Println("power on")

	s.r = r
	s.r.SetCurrentSate(&StandByState{})

	s.r.Standby()
}

type RemoteControlMachine struct {
	currentSate TVState
}

// PowerOn delegates behavior to the current state object.
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
