package BehavioralType

import "testing"

func TestInitPrimary(t *testing.T) {
	var (
		pri Handler
		mid Handler
		sen Handler
		list []IRequest
	)
	list = make([]IRequest,0)
	list = append(list,&Request{
		level:   DIFFICULTY_LEVEL_1,
		request: "1+1=?",
	})
	list = append(list,&Request{
		level:   DIFFICULTY_LEVEL_2,
		request: "4*3",
	})
	list = append(list,&Request{
		level:   DIFFICULTY_LEVEL_3,
		request: "99*99",
	})
	list = append(list,&Request{
		level:   4,
		request: "aaaaaaaaaaa",
	})
	pri = InitPrimary()
	mid = &Middle{
		level:   DIFFICULTY_LEVEL_2,
		request: "",
		next:    nil,
	}
	sen = &Senior{
		level:   DIFFICULTY_LEVEL_3,
		request: "",
		next:    nil,
	}

	// 设置链的顺序
	pri.SetNextHandler(mid)
	mid.SetNextHandler(sen)
	for _,v := range list {
		// 责任链中处理该请求
		pri.HandleMessage(v,pri,HandleMess)
	}
}
