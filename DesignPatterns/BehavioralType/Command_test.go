package BehavioralType

import (
	"testing"
)

func TestConcreteCommand_DoExec(t *testing.T) {

	rece := &TV{}

	openComm := &OpenTvCommand{rece}
	changeComm := &ChangeTvCommand{rece}
	closeComm := &CloseTvCommand{rece}

	tvR := &TVRemote{
		open:   openComm,
		change: changeComm,
		close:  closeComm,
	}
	tvR.Open()
	tvR.Change()
	tvR.Close()
}
