package BehavioralType

import (
	"testing"
)

func TestConcreteCommand_DoExec(t *testing.T) {
	// 创建接收者
	rece := &TV{}
	// 创建命令对象
	openComm := &OpenTvCommand{rece}
	changeComm := &ChangeTvCommand{rece}
	closeComm := &CloseTvCommand{rece}

	// 创建请求者，把命令对象设置进去
	tvR := &TVRemote{
		open:   openComm,
		change: changeComm,
		close:  closeComm,
	}
	tvR.Open()
	tvR.Change()
	tvR.Close()
}
