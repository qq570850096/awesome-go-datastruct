package BehavioralType

import "fmt"

const (
	DIFFICULTY_LEVEL_1 = 1
	DIFFICULTY_LEVEL_2 = 2
	DIFFICULTY_LEVEL_3 = 3
)

type HandleMessage func(hand Handler,request IRequest)


type IRequest interface {
	// 请求级别
	GetRequestLevel() int
	// 获取要请求的内容
	GetRequest() string
}

type Request struct {
	// 难度1--初级工程师
	// 难度2--中级工程师
	// 难度3--高级工程师
	level int
	request string
}

func InitRequset(level int, request string) *Request {
	r := &Request{
		level:   level,
		request: request,
	}
	switch r.level {
	case 1:
		r.request = "难度级别1的请求是："+ request
	case 2:
		r.request = "难度级别2的请求是："+ request
	case 3:
		r.request = "难度级别3的请求是："+ request
	}
	return r
}

func (r Request) GetRequestLevel() int {
	return r.level
}

func (r Request) GetRequest() string {
	return r.request
}

type Handler interface {
	HandleMessage(request IRequest,handler Handler,message HandleMessage)
	SetNextHandler(handler Handler)
	Response(request IRequest)
	GetLevel()int
	GetNext() Handler
}

// 初级工程师
type Primary struct {
	level int
	request string
	next Handler
}

func (p *Primary) GetNext() Handler {
	return p.next
}

func (p *Primary) GetLevel() int {
	return p.level
}

func (p *Primary) HandleMessage(request IRequest,handler Handler,message HandleMessage) {
	message(p,request)
}

func (p *Primary) SetNextHandler(handler Handler) {
	p.next = handler
}

func (p *Primary) Response(request IRequest) {
	fmt.Println("---难度级别1的请求---")
	fmt.Printf(request.GetRequest())
	fmt.Println("初级工程师已经处理完毕")
}

func InitPrimary() Handler {
	return &Primary{
		level:   DIFFICULTY_LEVEL_1,
		request: "",
	}
}

type Middle struct {
	level int
	request string
	next Handler
}
func HandleMess (hand Handler,request IRequest)  {
	// 如果请求级别小于可以处理的级别就直接处理
	if request.GetRequestLevel() <= hand.GetLevel() {
		hand.Response(request)
	} else {
		if hand.GetNext() != nil {
			HandleMess(hand.GetNext(),request)
		} else {
			fmt.Println("---难度级别为",request.GetRequestLevel(),"的请求无法处理")
		}
	}
}
func (p *Middle) HandleMessage(request IRequest,handler Handler,message HandleMessage) {
	handler = p
	message(handler,request)
}

func (p *Middle) SetNextHandler(handler Handler) {
	p.next = handler
}

func (p *Middle) Response(request IRequest) {
	fmt.Println("---难度级别2的请求---")
	fmt.Printf(request.GetRequest())
	fmt.Println("中级工程师已经处理完毕")
}

func (p *Middle) GetLevel() int {
	return p.level
}

func (p *Middle) GetNext() Handler {
	return p.next
}

type Senior struct {
	level int
	request string
	next Handler
}

func (p *Senior) HandleMessage(request IRequest,handler Handler,message HandleMessage) {
	handler = p
	message(handler,request)
}

func (p *Senior) SetNextHandler(handler Handler) {
	p.next = handler
}

func (p *Senior) Response(request IRequest) {
	fmt.Println("---难度级别3的请求---")
	fmt.Printf(request.GetRequest())
	fmt.Println("高级工程师已经处理完毕")
}

func (p *Senior) GetLevel() int {
	return p.level
}

func (p *Senior) GetNext() Handler {
	return p.next
}








