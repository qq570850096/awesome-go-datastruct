package BehavioralType

import "testing"

func TestPlatform_Attach(t *testing.T) {
	// 创建图书平台（发布者）
	platform := Platform{list: []IReader{}}
	// 创建读者A
	reader := Reader{name:"A"}
	// 读者A订阅图书通知
	platform.Attach(&reader)
	// 创建读者B
	reader2 := Reader{name:"B"}
	// 读者B订阅图书通知
	platform.Attach(&reader2)
	platform.Change("《go核心编程》")
	// 读者B取消订阅
	platform.Detach(&reader2)
	platform.Change("《go高级编程》")
}

